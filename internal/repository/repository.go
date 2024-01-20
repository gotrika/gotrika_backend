package repository

import (
	"context"

	"github.com/gotrika/gotrika_backend/internal/core"
	"github.com/gotrika/gotrika_backend/internal/dto"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersR interface {
	CreateUser(ctx context.Context, userDTO *dto.CreateUserDTO) (string, error)
	GetUserById(ctx context.Context, userID string) (*core.User, error)
	GetUserByUsername(ctx context.Context, username string) (*core.User, error)
}

type Repositories struct {
	Users UsersR
}

// NewRepositories: ini all repos
func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Users: NewUserRepo(db),
	}
}
