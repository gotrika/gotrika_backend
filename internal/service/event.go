package service

import (
	"context"
	"encoding/json"

	"github.com/gotrika/gotrika_backend/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventR interface {
	Save(ctx context.Context, eventDTO dto.EventTaskDTO) error
	InsertManyEvents(ctx context.Context, eventDTOs []dto.EventTaskDTO) error
}

type EventService struct {
	repo        EventR
	trackerRepo TrackerR
}

func NewEventService(repo EventR, trackerRepo TrackerR) *EventService {
	return &EventService{
		repo:        repo,
		trackerRepo: trackerRepo,
	}
}

func (s *EventService) ParseTask(ctx context.Context, parseTaskDTO *dto.ParseTask) error {
	ids := make([]primitive.ObjectID, len(parseTaskDTO.IDS))
	for index, id := range parseTaskDTO.IDS {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}
		ids[index] = objID
	}
	rawEvents, err := s.trackerRepo.GetTrackerDataByIDs(ctx, ids)
	if err != nil {
		return err
	}
	if err := s.trackerRepo.ToWorkTrackerData(ctx, ids); err != nil {
		return err
	}
	dtos := make([]dto.EventTaskDTO, len(rawEvents))
	for index, rawEvent := range rawEvents {
		var eventDTO dto.EventTaskDTO
		err := json.Unmarshal(rawEvent.TrackerData, &eventDTO)
		if err != nil {
			return err
		}
		eventDTO.ServerTimestamp = int(rawEvent.Datetime.Time().Unix())
		eventDTO.SiteID = rawEvent.SiteID
		dtos[index] = eventDTO
	}
	if err := s.repo.InsertManyEvents(ctx, dtos); err != nil {
		return err
	}
	if err := s.trackerRepo.ToParsedTrackerData(ctx, ids); err != nil {
		return err
	}
	return nil
}
