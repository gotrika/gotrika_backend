package amqp

import (
	"github.com/gotrika/gotrika_backend/internal/config"
	"github.com/gotrika/gotrika_backend/internal/service"
)

type AMQPClient struct {
	cfg      *config.Config
	services *service.Services
}

func NewAMQPHandler(services *service.Services, cfg *config.Config) *AMQPClient {
	return &AMQPClient{
		services: services,
		cfg:      cfg,
	}
}
