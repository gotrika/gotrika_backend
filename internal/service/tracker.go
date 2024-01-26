package service

import (
	"context"

	"github.com/gotrika/gotrika_backend/internal/dto"
)

type TrackerR interface {
	SaveRawTrackerData(ctx context.Context, td *dto.TrackerDataDTO) error
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
