package service

import (
	"context"

	"github.com/bzdvdn/cabbage/cabbage"
	"github.com/gotrika/gotrika_backend/internal/core"
	"github.com/gotrika/gotrika_backend/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TrackerR interface {
	SaveRawTrackerData(ctx context.Context, td *dto.TrackerDataDTO) error
	GetUnparsedTrackerData(ctx context.Context, dtype string) ([]*core.RawTrackerData, error)
	ToWorkTrackerData(ctx context.Context, ids []primitive.ObjectID) error
	ToParsedTrackerData(ctx context.Context, ids []primitive.ObjectID) error
	GetTrackerDataByIDs(ctx context.Context, ids []primitive.ObjectID) ([]*core.RawTrackerData, error)
}

type TrackerService struct {
	repo TrackerR
}

func NewTrackerService(repo TrackerR) *TrackerService {
	service := &TrackerService{
		repo: repo,
	}
	return service
}

func (s *TrackerService) SaveRawTrackerData(ctx context.Context, td *dto.TrackerDataDTO) error {
	return s.repo.SaveRawTrackerData(ctx, td)
}

func (s *TrackerService) ScheduleEventFunc(ctx context.Context) (func() (tpublisher cabbage.TaskPublisher), error) {
	return s.createScheduleFunc(ctx, "event")
}

func (s *TrackerService) ScheduleSessionFunc(ctx context.Context) (func() (tpublisher cabbage.TaskPublisher), error) {
	return s.createScheduleFunc(ctx, "session")
}

func (s *TrackerService) createScheduleFunc(ctx context.Context, dtype string) (func() (tpublisher cabbage.TaskPublisher), error) {
	trackerData, err := s.repo.GetUnparsedTrackerData(ctx, dtype)
	if err != nil {
		return nil, err
	}
	return func() (tpublisher cabbage.TaskPublisher) {
		ids := make([]string, len(trackerData))
		for index, event := range trackerData {
			ids[index] = event.ID.Hex()
		}
		return &dto.ParseTask{IDS: ids, Type: dtype}
	}, nil
}
