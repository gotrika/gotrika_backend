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
}

func NewTrackerDataRepo(db *mongo.Database) *TrackerDataRepo {
	return &TrackerDataRepo{
		rawDataCollection: db.Collection(core.RawTrackerDataCollectioName),
	}
}

func (r *TrackerDataRepo) SaveRawTrackerData(ctx context.Context, td *dto.TrackerDataDTO) error {
	rawTrackerData := core.RawTrackerData{
		SiteID:      td.SiteID,
		HashID:      td.HashID,
		Datetime:    primitive.NewDateTimeFromTime(td.Timestamp),
		Type:        td.Type,
		TrackerData: td.TrackerData,
		Parsed:      false,
		InWork:      false,
	}
	_, err := r.rawDataCollection.InsertOne(ctx, &rawTrackerData)
	if err != nil {
		return err
	}
	return nil
}
