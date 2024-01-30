package service

import (
	"context"

	"github.com/bzdvdn/cabbage/cabbage"
	"github.com/gotrika/gotrika_backend/internal/dto"
	"github.com/gotrika/gotrika_backend/internal/repository"
	"github.com/gotrika/gotrika_backend/pkg/auth"
	"github.com/gotrika/gotrika_backend/pkg/hash"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users interface {
	SignUp(ctx context.Context, input dto.RegisterUserDTO) (string, error)
	SignIn(ctx context.Context, input dto.AuthUserDTO) (*dto.AuthResponse, error)
	GetUserByID(ctx context.Context, userID string) (*dto.UserRetrieveDTO, string, error)
	UpdateTokens(ctx context.Context, refreshToken string) (*dto.AuthResponse, error)
	TokenManager() auth.TokenManager
}

type Sites interface {
	CreateSite(ctx context.Context, url string, userID primitive.ObjectID, siteDTO dto.CreateSiteDTO) (*dto.SiteRetrieveDTO, error)
	GetSiteByID(ctx context.Context, isAdmin bool, userID, siteID primitive.ObjectID) (*dto.SiteRetrieveDTO, error)
	DeleteSite(ctx context.Context, isAdmin bool, userID, siteID primitive.ObjectID) error
	ListSites(ctx context.Context, search string, isAdmin bool, userID primitive.ObjectID, limit int, offset int) ([]*dto.SiteRetrieveDTO, int, error)
	UpdateSite(ctx context.Context, isAdmin bool, userID primitive.ObjectID, siteID primitive.ObjectID, siteDTO *dto.UpdateSiteDTO) (*dto.SiteRetrieveDTO, error)
}

type TrackerSrv interface {
	SaveRawTrackerData(ctx context.Context, td *dto.TrackerDataDTO) error
	ScheduleEventFunc(ctx context.Context) (func() (tpublisher cabbage.TaskPublisher), error)
	ScheduleSessionFunc(ctx context.Context) (func() (tpublisher cabbage.TaskPublisher), error)
}

type Events interface {
	ParseTask(ctx context.Context, parseTaskDTO *dto.ParseTask) error
}
type Sessions interface {
	ParseTask(ctx context.Context, parseTaskDTO *dto.ParseTask) error
}

type Dependencies struct {
	Repos        *repository.Repositories
	Hasher       hash.Hasher
	TokenManager auth.TokenManager
}

type Services struct {
	Users          Users
	Sites          Sites
	TrackerService TrackerSrv
	Events         Events
	Sessions       Sessions
}

// NewServices: init services
func NewServices(deps Dependencies) *Services {
	return &Services{
		Users:          NewUserService(deps.Repos.Users, deps.Hasher, deps.TokenManager),
		Sites:          NewSiteService(deps.Repos.Sites),
		TrackerService: NewTrackerService(deps.Repos.TrackerRepo),
		Sessions:       NewSessionService(deps.Repos.Sessions, deps.Repos.TrackerRepo),
		Events:         NewEventService(deps.Repos.Events, deps.Repos.TrackerRepo),
	}
}
