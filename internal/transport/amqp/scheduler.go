package amqp

import (
	"github.com/bzdvdn/cabbage/cabbage"
)

type Scheduler struct {
	cabbageScheduler *cabbage.Scheduler
}

func newScheduler(cl *cabbage.CabbageClient) *Scheduler {
	cabbageScheduler := cl.CreateScheduler()
	return &Scheduler{
		cabbageScheduler: cabbageScheduler,
	}
}

func (s *Scheduler) Start() chan bool {
	return s.cabbageScheduler.Start()
}

func (s *Scheduler) AddTasks(shTasks []*cabbage.ScheduleTask) error {
	return s.cabbageScheduler.AddScheduleTasks(shTasks)
}
