package service

import "github.com/gotrika/gotrika_backend/internal/repository"

type Dependencies struct {
	Repos *repository.Repositories
}

type Services struct{}

// NewServices: init services
func NewServices(deps Dependencies) *Services {
	return &Services{}
}
