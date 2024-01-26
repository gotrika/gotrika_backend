package service

import (
	"context"

	"github.com/gotrika/gotrika_backend/internal/dto"
)

type EventR interface {
	Save(ctx context.Context, eventDTO dto.EventTaskDTO) error
}

type EventService struct {
	repo EventR
}

func NewEventService(repo EventR) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) ParseTask(ctx context.Context, parseTaskDTO dto.ParseTask) error {
	return nil
}
