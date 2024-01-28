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

func (r *EventRepo) InsertManyEvents(ctx context.Context, eventDTOs []dto.EventTaskDTO) error {
	coreEvents := make([]interface{}, len(eventDTOs))
	for index, eventDTO := range eventDTOs {
		coreEvents[index] = core.Event{
			SiteID:          eventDTO.SiteID,
			SessionID:       eventDTO.SessionID,
			VisitorID:       eventDTO.VisitorID,
			ClassName:       eventDTO.ClassName,
			Page:            eventDTO.Page,
			PageTitle:       eventDTO.PageTitle,
			HitURL:          eventDTO.HitURL,
			Type:            eventDTO.Type,
			ServerTimestamp: eventDTO.ServerTimestamp,
			ClientTimestamp: eventDTO.ClientTimestamp,
			Referrer:        eventDTO.Referrer,
		}
	}
	_, err := r.collection.InsertMany(ctx, coreEvents)
	return err
}
