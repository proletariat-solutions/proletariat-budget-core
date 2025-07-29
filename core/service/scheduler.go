package service

import (
	"context"
	"log"

	"proletariat-budget-core/core/port"
)

// SchedulerService is the application service that manages the orchestrator
type SchedulerService struct {
	orchestrator *Orchestrator
}

func NewSchedulerService(
	repository port.RecurrencePattern,
	processor port.TaskProcessor,
	timeProvider port.TimeProvider,
) *SchedulerService {
	orchestrator := NewOrchestrator(
		repository,
		processor,
		timeProvider,
	)

	return &SchedulerService{
		orchestrator: orchestrator,
	}
}

// Start starts the orchestrator service
func (s *SchedulerService) Start(ctx context.Context) error {
	log.Println("Starting orchestrator service...")

	return s.orchestrator.Start(ctx)
}

// Stop stops the orchestrator service
func (s *SchedulerService) Stop() {
	log.Println("Stopping orchestrator service...")
	s.orchestrator.Stop()
}
