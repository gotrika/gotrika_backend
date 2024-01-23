package core

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RawTrackerData struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	SiteID      primitive.ObjectID `bson:"site_id"`
	HashID      string             `bson:"hash_id"`
	Datetime    time.Time          `bson:"datetime"`
	Type        string             `bson:"type"`
	TrackerData []byte             `bson:"tracked_data"`
}

type Session struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	SiteID              primitive.ObjectID `bson:"site_id"`
	SessionID           string             `bson:"session_id"`
	ServerTimestamp     int                `bson:"server_timestamp"`
	ClientTimestamp     int                `bson:"client_timestamp"`
	ClientTimezone      string             `bson:"client_timezone"`
	Duration            int                `bson:"duration"`
	VisitorID           string             `bson:"visitor_id"`
	UTMSource           string             `bson:"utm_source"`
	UTMMedium           string             `bson:"utm_medium"`
	UTMCampaign         string             `bson:"utm_campaign"`
	UTMContent          string             `bson:"utm_content"`
	UTMTerm             string             `bson:"utm_term"`
	UTMReferrer         string             `bson:"utm_referrer"`
	Referrer            string             `bson:"referrer"`
	IPAdress            string             `bson:"ip_address"`
	CountryShort        string             `bson:"country_short"`
	CountryLong         string             `bson:"country_long"`
	RegionName          string             `bson:"region_name"`
	City                string             `bson:"city"`
	Longtitude          float64            `bson:"longtitude"`
	Latitude            float64            `bson:"latitude"`
	OSFamily            string             `bson:"os_family"`
	OSName              string             `bson:"os_name"`
	Browser             string             `bson:"browser"`
	BrowserFullVersion  string             `bson:"browser_full_version"`
	BrowserMinorVersion string             `bson:"browser_minor_version"`
	BrowserMajorVersion string             `bson:"browser_major_version"`
	ClientWindowWidth   int                `bson:"client_window_width"`
	ClientWindowHeight  int                `bson:"client_window_height"`
	Device              string             `bson:"device"`
}

type Event struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	SiteID          primitive.ObjectID `bson:"site_id"`
	SessionID       string             `bson:"session_id"`
	VisitorID       string             `bson:"visitor_id"`
	ClassName       string             `bson:"class_name"`
	Type            string             `bson:"Type"`
	ServerTimestamp int                `bson:"server_timestamp"`
	ClientTimestamp int                `bson:"client_timestamp"`
	Page            string             `bson:"page"`
	PageTitle       string             `bson:"page_title"`
	HitURL          string             `bson:"hit_url"`
	Referrer        string             `bson:"referrer"`
}
