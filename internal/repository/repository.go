package repository

import (
	"context"

	"github.com/gotrika/gotrika_backend/internal/core"
	"github.com/gotrika/gotrika_backend/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UsersR interface {
	CreateUser(ctx context.Context, userDTO *dto.CreateUserDTO) (string, error)
	GetUserById(ctx context.Context, userID string) (*core.User, error)
	GetUserByUsername(ctx context.Context, username string) (*core.User, error)
}

type SitesR interface {
	CreateSite(ctx context.Context, userID primitive.ObjectID, siteDTO dto.CreateSiteDTO, scriptUrl string) (*core.Site, error)
	UpdateSite(ctx context.Context, siteID primitive.ObjectID, siteDTO *dto.UpdateSiteDTO) (*core.Site, error)
	DeleteSite(ctx context.Context, siteID primitive.ObjectID) error
	GetSiteByID(ctx context.Context, siteID primitive.ObjectID) (*core.Site, error)
	ListSites(ctx context.Context, listDTO *dto.ListSitesDTO) ([]core.Site, int, error)
}

type TrackerR interface {
	SaveRawTrackerData(ctx context.Context, td *dto.AddRawTrackerDataDTO) error
}

type Repositories struct {
	Users       UsersR
	Sites       SitesR
	TrackerRepo TrackerR
}

// NewRepositories: ini all repos
func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Users:       NewUserRepo(db),
		Sites:       NewSiteRepo(db),
		TrackerRepo: NewTrackerDataRepo(db),
	}
}

func getPaginationOpts(limit int, skip int) *options.FindOptions {
	if limit > 100 || limit < 0 {
		limit = 100
	}
	if skip > 0 {
		skip = 0
	}
	skipParam := int64(skip)
	limitParam := int64(limit)
	return &options.FindOptions{
		Skip:  &skipParam,
		Limit: &limitParam,
	}
}
