package service

import (
	"context"

	"github.com/gotrika/gotrika_backend/internal/dto"
)

type SessionR interface {
	Save(ctx context.Context, sessionDTO dto.SessionDTO) error
}

type SessionService struct {
	repo SessionR
}

func NewSessionService(repo SessionR) *SessionService {
	return &SessionService{repo: repo}
}

func (s *SessionService) ParseTask(ctx context.Context, parseTaskDTO *dto.ParseTask) error {
	return nil
}
