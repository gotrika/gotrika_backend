package repository

import (
	"context"

	"github.com/gotrika/gotrika_backend/internal/core"
	"github.com/gotrika/gotrika_backend/internal/dto"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventRepo struct {
	collection *mongo.Collection
}

func NewEventRepo(db *mongo.Database) *EventRepo {
	return &EventRepo{
		collection: db.Collection(core.EventCollectionName),
	}
}

func (r *EventRepo) Save(ctx context.Context, eventDTO dto.EventTaskDTO) error {
	return nil
}
