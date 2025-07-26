package service

import (
	"context"
	"log"

	"proletariat-budget-core/core/port"
)

// SchedulerService is the application service that orchestrates the scheduler
type SchedulerService struct {
	scheduler *JobScheduler
}

func NewSchedulerService(
	repository port.RecurrencePattern,
	processor port.TaskProcessor,
	timeProvider port.TimeProvider,
) *SchedulerService {
	scheduler := NewJobScheduler(
		repository,
		processor,
		timeProvider,
	)

	return &SchedulerService{
		scheduler: scheduler,
	}
}

// Start starts the scheduler service
func (s *SchedulerService) Start(ctx context.Context) error {
	log.Println("Starting scheduler service...")

	return s.scheduler.Start(ctx)
}

// Stop stops the scheduler service
func (s *SchedulerService) Stop() {
	log.Println("Stopping scheduler service...")
	s.scheduler.Stop()
}
