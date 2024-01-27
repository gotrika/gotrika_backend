package amqp

import (
	"context"

	"github.com/bzdvdn/cabbage/cabbage"
	"github.com/gotrika/gotrika_backend/internal/config"
	"github.com/gotrika/gotrika_backend/internal/service"
)

type TaskManager struct {
	services *service.Services
	cfg      *config.Config
}

func newTaskManager(services *service.Services, cfg *config.Config) *TaskManager {
	return &TaskManager{
		services: services,
		cfg:      cfg,
	}
}

func (m *TaskManager) CreateSchedulerTasks(ctx context.Context) []*cabbage.ScheduleTask {
	every5min, _ := cabbage.EntryEveryMinute(5)
	eventScheduleTaskFunc, _ := m.services.TrackerService.ScheduleEventFunc(ctx)
	sessionScheduleTaskFunc, _ := m.services.TrackerService.ScheduleSessionFunc(ctx)
	tasks := []*cabbage.ScheduleTask{
		{
			Name:      "parseSessions",
			QueueName: m.cfg.CabbageConfig.SessionQueueName,
			Func:      sessionScheduleTaskFunc,
			Entries:   cabbage.Entries{&cabbage.Entry{Schedule: "* * * * *"}},
		}, {
			Name:      "parseEvents",
			QueueName: m.cfg.CabbageConfig.EventQueueName,
			Func:      eventScheduleTaskFunc,
			Entries:   cabbage.Entries{every5min},
		}}
	return tasks
}

func (m *TaskManager) CreateWorkerTasks() []*cabbage.Task {
	return nil
}
