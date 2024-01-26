package repository

import (
	"context"

	"github.com/gotrika/gotrika_backend/internal/core"
	"github.com/gotrika/gotrika_backend/internal/dto"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionRepo struct {
	collection *mongo.Collection
}

func NewSessionRepo(db *mongo.Database) *SessionRepo {
	return &SessionRepo{
		collection: db.Collection(core.SessionCollectionName),
	}
}

func (r *SessionRepo) Save(ctx context.Context, sessionDTO dto.SessionTaskDTO) error {
	return nil
}
