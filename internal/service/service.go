package service

import (
	"context"

	"github.com/gotrika/gotrika_backend/internal/dto"
	"github.com/gotrika/gotrika_backend/internal/repository"
	"github.com/gotrika/gotrika_backend/pkg/auth"
	"github.com/gotrika/gotrika_backend/pkg/hash"
)

type Users interface {
	SignUp(ctx context.Context, input dto.RegisterUserDTO) (string, error)
	SignIn(ctx context.Context, input dto.AuthUserDTO) (*dto.AuthResponse, error)
	GetUserByID(ctx context.Context, userID string) (*dto.UserRetrieveDTO, string, error)
	TokenManager() auth.TokenManager
}

type Dependencies struct {
	Repos        *repository.Repositories
	Hasher       hash.Hasher
	TokenManager auth.TokenManager
}

type Services struct {
	Users Users
}

// NewServices: init services
func NewServices(deps Dependencies) *Services {
	return &Services{
		Users: NewUserService(deps.Repos.Users, deps.Hasher, deps.TokenManager),
	}
}
