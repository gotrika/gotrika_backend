package repository

import (
	"context"

	"github.com/gotrika/gotrika_backend/internal/core"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionRepo struct {
	rawDataCollection *mongo.Collection
	collection        *mongo.Collection
}

func NewSessionRepo(db *mongo.Database) *SessionRepo {
	return &SessionRepo{
		rawDataCollection: db.Collection(core.RawTrackerDataCollectioName),
		collection:        db.Collection(core.SessionCollectionName),
	}
}

func (r *SessionRepo) Save(ctx context.Context, session core.Session) error {
	return nil
}

func (r *SessionRepo) InsertManySessions(ctx context.Context, sessions []core.Session) error {
	coreSessions := make([]interface{}, len(sessions))
	for index, session := range sessions {
		coreSessions[index] = session
	}
	_, err := r.collection.InsertMany(ctx, coreSessions)
	return err
}
