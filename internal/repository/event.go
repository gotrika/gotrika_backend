package repository

import (
	"context"

	"github.com/gotrika/gotrika_backend/internal/core"
	"github.com/gotrika/gotrika_backend/internal/dto"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventRepo struct {
	rawDataCollection *mongo.Collection
	collection        *mongo.Collection
}

func NewEventRepo(db *mongo.Database) *EventRepo {
	return &EventRepo{
		rawDataCollection: db.Collection(core.RawTrackerDataCollectioName),
		collection:        db.Collection(core.EventCollectionName),
	}
}

func (r *EventRepo) Save(ctx context.Context, eventDTO dto.EventTaskDTO) error {
	return nil
}

func (r *EventRepo) InserManyEvents(ctx context.Context, eventDTOs []dto.EventTaskDTO) error {
	return nil
}
