package bootstrap

import (
	"context"

	"github.com/gotrika/gotrika_backend/internal/transport/amqp"
	"github.com/gotrika/gotrika_backend/pkg/logger"
)

func RunScheduler() {
	deps, err := InitDependencies()
	if err != nil {
		logger.Error("Invalid get dependenices")
		return
	}
	amqpClient, err := amqp.NewAMQPHandler(deps.Services(), deps.Config())
	if err != nil {
		logger.Error("Failed init amqp")
		return
	}
	ctx := context.Background()
	scheduler := amqpClient.CreateScheduler(ctx)
	<-scheduler.Start()
}
