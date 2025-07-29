package port

import (
	"context"

	domain "proletariat-budget-core/core/domain/coreentity"
)

// TaskProcessor defines how tasks are processed

//go:generate mockgen -source=task_processor.go -destination=../../test/mocks/mock_task_processor.go -package mocks
type TaskProcessor interface {
	ProcessTask(
		ctx context.Context,
		task domain.RecurrencePattern,
	) error
}
