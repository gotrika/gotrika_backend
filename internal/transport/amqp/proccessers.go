package amqp

import (
	"context"

	"github.com/gotrika/gotrika_backend/internal/dto"
)

type Events interface {
	ParseTask(ctx context.Context, parseTaskDTO dto.ParseTask) error
}
type Sessions interface {
	ParseTask(ctx context.Context, parseTaskDTO dto.ParseTask) error
}

type EventTaskProccesser struct {
	eventService Events
}

func NewEventTaskPoccesser(es Events) *EventTaskProccesser {
	return &EventTaskProccesser{eventService: es}
}

type SessionTaskProccesser struct {
	sessionService Sessions
}

func NewSessionTaskProccesser(ss Sessions) *SessionTaskProccesser {
	return &SessionTaskProccesser{sessionService: ss}
}

func (tp *SessionTaskProccesser) ProccessTask(ctx context.Context, body []byte, ID string) error {
	return nil
}

func (tp *EventTaskProccesser) ProccessTask(ctx context.Context, body []byte, ID string) error {
	return nil
}
