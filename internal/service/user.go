package service

import (
	"context"
	"errors"

	"github.com/gotrika/gotrika_backend/internal/core"
	"github.com/gotrika/gotrika_backend/internal/dto"
	"github.com/gotrika/gotrika_backend/pkg/auth"
	"github.com/gotrika/gotrika_backend/pkg/hash"
)

type UserR interface {
	CreateUser(ctx context.Context, userDTO *dto.CreateUserDTO) (string, error)
	GetUserById(ctx context.Context, userID string) (*core.User, error)
	GetUserByUsername(ctx context.Context, username string) (*core.User, error)
}

type UserService struct {
	repo         UserR
	hasher       hash.Hasher
	tokenManager auth.TokenManager
}

func NewUserService(repo UserR, hasher hash.Hasher, tokenManager auth.TokenManager) *UserService {
	service := &UserService{
		repo:         repo,
		hasher:       hasher,
		tokenManager: tokenManager,
	}
	return service
}
func (s *UserService) TokenManager() auth.TokenManager {
	return s.tokenManager
}

func (s *UserService) SignUp(ctx context.Context, input dto.RegisterUserDTO) (string, error) {
	if input.Password1 != input.Password2 {
		return "", errors.New("passwords missmatch")
	}
	hashPassword, err := s.hasher.Hash(input.Password1)
	if err != nil {
		return "", err
	}
	userDTO := dto.CreateUserDTO{
		Username:       input.Username,
		HashedPassword: hashPassword,
		Name:           input.Name,
		IsActive:       true,
	}
	userID, err := s.repo.CreateUser(ctx, &userDTO)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func (s *UserService) SignIn(ctx context.Context, input dto.AuthUserDTO) (*dto.AuthResponse, error) {
	user, err := s.repo.GetUserByUsername(ctx, input.Username)
	if err != nil {
		return nil, core.ErrUserNotFound
	}
	hashPassword, err := s.hasher.Hash(input.Password)
	if err != nil {
		return nil, err
	}
	if hashPassword != user.Password {
		return nil, errors.New("invalid password")
	}
	tokens, err := s.tokenManager.GenerateTokens(user.ID.Hex(), user.Sign)
	if err != nil {
		return nil, err
	}
	resp := dto.AuthResponse{
		ID:           tokens.UserID,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		Expires:      tokens.Expires,
		ExpiresAt:    tokens.ExpiresAt,
	}
	return &resp, nil
}

func (s *UserService) GetUserByID(ctx context.Context, userID string) (*dto.UserRetrieveDTO, string, error) {
	user, err := s.repo.GetUserById(ctx, userID)
	if err != nil {
		return nil, "", err
	}
	userDTO := dto.UserRetrieveDTO{
		ID:       user.ID.Hex(),
		Username: user.Username,
		Name:     user.Name,
		IsActive: user.IsActive,
		IsAdmin:  user.IsAdmin,
	}
	return &userDTO, user.Sign, nil
}

func (s *UserService) UpdateTokens(ctx context.Context, refreshToken string) (*dto.AuthResponse, error) {
	userID, sign, err := s.tokenManager.Parse(refreshToken)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.GetUserById(ctx, userID)
	if err != nil {
		return nil, err
	}
	if !user.IsActive {
		return nil, errors.New("user inactive")
	}
	if sign != user.Sign {
		return nil, errors.New("invalid signature")
	}
	tokens, err := s.tokenManager.GenerateTokens(user.ID.Hex(), user.Sign)
	if err != nil {
		return nil, err
	}
	resp := dto.AuthResponse{
		ID:           tokens.UserID,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		Expires:      tokens.Expires,
		ExpiresAt:    tokens.ExpiresAt,
	}
	return &resp, nil
}
