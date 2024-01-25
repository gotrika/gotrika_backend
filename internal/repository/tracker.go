package repository

import (
	"context"

	"github.com/gotrika/gotrika_backend/internal/core"
	"github.com/gotrika/gotrika_backend/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TrackerDataRepo struct {
	rawDataCollection *mongo.Collection
	sessionCollection *mongo.Collection
	eventCollection   *mongo.Collection
}

func NewTrackerDataRepo(db *mongo.Database) *TrackerDataRepo {
	return &TrackerDataRepo{
		rawDataCollection: db.Collection(core.RawTrackerDataCollectioName),
		sessionCollection: db.Collection(core.SessionCollectionName),
		eventCollection:   db.Collection(core.EventCollectionName),
	}
}

func (r *TrackerDataRepo) SaveRawTrackerData(ctx context.Context, td *dto.AddRawTrackerDataDTO) error {
	rawTrackerData := core.RawTrackerData{
		SiteID:      td.SiteID,
		HashID:      td.HashID,
		Datetime:    primitive.NewDateTimeFromTime(td.Timestamp),
		Type:        td.Type,
		TrackerData: td.TrackerData,
	}
	_, err := r.rawDataCollection.InsertOne(ctx, &rawTrackerData)
	if err != nil {
		return err
	}
	return nil
}
