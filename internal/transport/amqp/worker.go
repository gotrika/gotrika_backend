package amqp

import (
	"context"

	"github.com/bzdvdn/cabbage/cabbage"
	"github.com/gotrika/gotrika_backend/internal/config"
)

type WorkerManager struct {
	cabbageClient *cabbage.CabbageClient
	sessionWorker *cabbage.CabbageWorker
	eventWorker   *cabbage.CabbageWorker
}

func NewWorkerManager(cl *cabbage.CabbageClient, cfg *config.Config) (*WorkerManager, error) {
	sessionWorker, err := cl.CreateWorker(cfg.CabbageConfig.SessionQueueName, 10)
	if err != nil {
		return nil, err
	}
	eventWorker, err := cl.CreateWorker(cfg.CabbageConfig.EventQueueName, 10)
	if err != nil {
		return nil, err
	}
	return &WorkerManager{
		cabbageClient: cl,
		sessionWorker: sessionWorker,
		eventWorker:   eventWorker,
	}, nil
}

func (m *WorkerManager) Close() {
	m.sessionWorker.StopWorker()
	m.eventWorker.StartWorker()
	m.cabbageClient.Close()
}

func (m *WorkerManager) StartSessionWorker(ctx context.Context) error {
	return m.sessionWorker.StartWorkerWithContext(ctx)
}
func (m *WorkerManager) StartEventWorker(ctx context.Context) error {
	return m.eventWorker.StartWorkerWithContext(ctx)
}

func (m *WorkerManager) RegisterTask(task *cabbage.Task) error {
	return m.cabbageClient.RegisterTask(task)
}

func (m *WorkerManager) StopSessionWorker() {
	m.sessionWorker.StopWorker()
}

func (m *WorkerManager) StopEventWorker() {
	m.eventWorker.StopWorker()
}
