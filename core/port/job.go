package port

import (
	"context"

	domain "proletariat-budget-core/core/domain/coreentity"
)

// TaskProcessor defines how tasks are processed
type TaskProcessor interface {
	ProcessTasks(
		ctx context.Context,
		tasks []domain.RecurrencePattern,
	) error
}
