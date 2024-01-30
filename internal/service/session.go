package service

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/gotrika/gotrika_backend/internal/core"
	"github.com/gotrika/gotrika_backend/internal/dto"
	"github.com/mileusna/useragent"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/exp/slices"
)

var utmsParams = []string{
	"utm_source",
	"utm_medium",
	"utm_campaign",
	"utm_content",
	"utm_term",
	"utm_referrer",
}

type SessionR interface {
	Save(ctx context.Context, session core.Session) error
	InsertManySessions(ctx context.Context, sessions []core.Session) error
}

type SessionService struct {
	repo        SessionR
	trackerRepo TrackerR
}

func NewSessionService(repo SessionR, trackerRepo TrackerR) *SessionService {
	return &SessionService{
		repo:        repo,
		trackerRepo: trackerRepo,
	}
}

func (s *SessionService) parseUserAgent(ua string) map[string]string {
	result := make(map[string]string)
	parsedUA := useragent.Parse(ua)
	result["OSFamily"] = parsedUA.OS
	result["OSName"] = parsedUA.OSVersion
	result["Browser"] = parsedUA.Name
	result["BrowserFullVersion"] = parsedUA.VersionNoFull()
	result["BrowserMinorVersion"] = strconv.Itoa(parsedUA.OSVersionNo.Minor)
	result["Device"] = parsedUA.Device
	return result
}

func (s *SessionService) parseUTMS(enterURL string) (map[string]string, error) {
	utms := map[string]string{
		"utm_source":   "direct",
		"utm_medium":   "none",
		"utm_campaign": "",
		"utm_content":  "",
		"utm_term":     "",
		"utm_referrer": "",
	}
	parsedURL, err := url.Parse(enterURL)
	if err != nil {
		return utms, err
	}
	values := parsedURL.Query()
	for key, v := range values {
		if slices.Contains(utmsParams, key) {
			utms[key] = v[0]
		}
	}
	return utms, nil
}

func (s *SessionService) ParseTask(ctx context.Context, parseTaskDTO *dto.ParseTask) error {
	ids := make([]primitive.ObjectID, len(parseTaskDTO.IDS))
	for index, id := range parseTaskDTO.IDS {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}
		ids[index] = objID
	}
	rawSessions, err := s.trackerRepo.GetTrackerDataByIDs(ctx, ids)
	if err != nil {
		return err
	}
	if err := s.trackerRepo.ToWorkTrackerData(ctx, ids); err != nil {
		return err
	}
	sessions := make([]core.Session, len(rawSessions))
	for index, rawSession := range rawSessions {
		var sessionDTO dto.SessionTaskDTO
		err := json.Unmarshal(rawSession.TrackerData, &sessionDTO)
		if err != nil {
			return err
		}
		parsedUserAgent := s.parseUserAgent(sessionDTO.UserAgent)
		utms, err := s.parseUTMS(sessionDTO.EnterURL)
		if err != nil {
			return err
		}
		session := core.Session{
			SiteID:              rawSession.SiteID,
			SessionID:           sessionDTO.SessionID,
			ServerTimestamp:     int(rawSession.Datetime.Time().Unix()),
			ClientTimestamp:     sessionDTO.SessionTimestamp,
			ClientTimezone:      sessionDTO.ClientTimezone,
			OSFamily:            parsedUserAgent["OSFamily"],
			OSName:              parsedUserAgent["OSName"],
			Browser:             parsedUserAgent["Browser"],
			BrowserFullVersion:  parsedUserAgent["BrowserFullVersion"],
			BrowserMinorVersion: parsedUserAgent["BrowserMinorVersion"],
			BrowserMajorVersion: parsedUserAgent["BrowserMajorVersion"],
			Device:              parsedUserAgent["Device"],
			Location:            sessionDTO.Location,
			Referrer:            sessionDTO.Referrer,
			ClientWindowWidth:   sessionDTO.UserScreenWidth,
			ClientWindowHeight:  sessionDTO.UserScreenHeight,
			EnterURL:            sessionDTO.EnterURL,
			UTMSource:           utms["utm_source"],
			UTMMedium:           utms["utm_medium"],
			UTMCampaign:         utms["utm_campaign"],
			UTMContent:          utms["utm_content"],
			UTMReferrer:         utms["utm_referrer"],
		}
		sessions[index] = session
	}
	if err := s.repo.InsertManySessions(ctx, sessions); err != nil {
		return err
	}
	if err := s.trackerRepo.ToParsedTrackerData(ctx, ids); err != nil {
		return err
	}
	return nil
}
