package bootstrap

import (
	"fmt"

	"github.com/gotrika/gotrika_backend/internal/config"
	"github.com/gotrika/gotrika_backend/internal/repository"
	"github.com/gotrika/gotrika_backend/internal/service"
	"github.com/gotrika/gotrika_backend/pkg/database"
	"github.com/gotrika/gotrika_backend/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type Dependencies struct {
	cfg         *config.Config
	mongoClient *mongo.Client
	services    *service.Services
}

func (d *Dependencies) Config() *config.Config {
	return d.cfg
}

func (d *Dependencies) MongoClient() *mongo.Client {
	return d.mongoClient
}

func (d *Dependencies) Services() *service.Services {
	return d.services
}

func InitDependencies() (*Dependencies, error) {
	logger.Info("Init config")
	cfg := config.GetConfig()
	// Dependencies
	logger.Info("Init db connection")
	mongodbURI := fmt.Sprintf("mongodb://%s:%s/", cfg.MongoConfig.HOST, cfg.MongoConfig.PORT)
	mongoClient, err := database.NewMongoClient(mongodbURI, cfg.MongoConfig.User, cfg.MongoConfig.Password)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	db := mongoClient.Database(cfg.MongoConfig.DBName)
	// Services & Repos
	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Dependencies{
		Repos: repos,
	})
	return &Dependencies{
		cfg:         cfg,
		mongoClient: mongoClient,
		services:    services,
	}, nil
}
