package bootstrap

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gotrika/gotrika_backend/internal/transport/amqp"
	"github.com/gotrika/gotrika_backend/pkg/logger"
)

func runWorker(workerType string) {
	if workerType != "session" && workerType != "event" {
		logger.Error("invalid worker type")
		return
	}
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
	defer amqpClient.Close()
	ctx := context.Background()
	workerManager, err := amqpClient.CreateWorkerManager(ctx)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	if workerType == "session" {
		err = workerManager.StartSessionWorker(ctx)
		if err != nil {
			logger.Error(err)
			return
		}
	}
	if workerType == "event" {
		err = workerManager.StartEventWorker(ctx)
		if err != nil {
			logger.Error(err)
			return
		}
	}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Wait for OS exit signal
	<-exit
	if workerType == "session" {
		workerManager.StopSessionWorker()
		return
	}
	if workerType == "event" {
		workerManager.StopEventWorker()
		return
	}

}

func RunSessionWorker() {
	runWorker("session")
}

func RunEventWorker() {
	runWorker("event")
}
