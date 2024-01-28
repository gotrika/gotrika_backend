package amqp

import (
	"context"

	"github.com/bzdvdn/cabbage/cabbage"
	"github.com/gotrika/gotrika_backend/internal/config"
	"github.com/gotrika/gotrika_backend/internal/service"
)

type AMQPClient struct {
	cfg           *config.Config
	services      *service.Services
	cabbageClient *cabbage.CabbageClient
	taskManager   *TaskManager
}

func NewAMQPHandler(services *service.Services, cfg *config.Config) (*AMQPClient, error) {
	broker, err := cabbage.NewRabbitMQBroker(cfg.CabbageConfig.BrokerURI, 5)
	if err != nil {
		return nil, err
	}
	cabbageClient := cabbage.NewCabbageClient(broker)
	return &AMQPClient{
		services:      services,
		cfg:           cfg,
		cabbageClient: cabbageClient,
		taskManager:   newTaskManager(services, cfg),
	}, nil
}

func (c *AMQPClient) Close() {
	c.cabbageClient.Close()
}

func (c *AMQPClient) CreateScheduler(ctx context.Context) *Scheduler {
	scheduler := newScheduler(c.cabbageClient)
	scheduleTasks := c.taskManager.CreateSchedulerTasks(ctx)
	scheduler.AddTasks(scheduleTasks)
	return scheduler
}

func (c *AMQPClient) CreateWorkerManager(ctx context.Context) (*WorkerManager, error) {
	wm, err := NewWorkerManager(c.cabbageClient, c.cfg)
	if err != nil {
		return nil, err
	}
	sessionTask := c.taskManager.CreateSessionTask()
	eventTask := c.taskManager.CreateEventTask()
	if err := wm.RegisterTask(sessionTask); err != nil {
		return nil, err
	}
	if err := wm.RegisterTask(eventTask); err != nil {
		return nil, err
	}
	return wm, nil
}
