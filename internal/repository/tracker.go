package repository

import (
	"context"

	"github.com/gotrika/gotrika_backend/internal/core"
	"github.com/gotrika/gotrika_backend/internal/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *TrackerDataRepo) GetUnparsedTrackerData(ctx context.Context, dtype string) ([]*core.RawTrackerData, error) {
	var data []*core.RawTrackerData
	limit := int64(250)
	opts := &options.FindOptions{
		Limit: &limit,
	}
	query := bson.M{"in_work": false, "type": dtype}
	cur, err := r.rawDataCollection.Find(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *TrackerDataRepo) ToWorkTrackerData(ctx context.Context, ids []primitive.ObjectID) error {
	query := bson.M{"_id": bson.M{"$in": ids}}
	update := bson.M{"$set": bson.M{"in_work": true}}
	_, err := r.rawDataCollection.UpdateMany(ctx, query, update)
	return err
}

func (r *TrackerDataRepo) ToParsedTrackerData(ctx context.Context, ids []primitive.ObjectID) error {
	query := bson.M{"_id": bson.M{"$in": ids}}
	update := bson.M{"$set": bson.M{"in_work": false, "parsed": true}}
	_, err := r.rawDataCollection.UpdateMany(ctx, query, update)
	return err
}
