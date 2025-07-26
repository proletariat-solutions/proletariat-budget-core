package port

import (
	"context"

	"proletariat-budget-core/core/domain/coreentity"
)

//go:generate mockgen -source=recurrence_pattern.go -destination=../../test/mocks/mock_recurrence_pattern_repo.go -package mocks
type RecurrencePattern interface {
	Create(
		ctx context.Context,
		recurrencePattern coreentity.RecurrencePattern,
	) (
		string,
		error,
	)
	Update(
		ctx context.Context,
		recurrencePattern coreentity.RecurrencePattern,
	) error
	Delete(
		ctx context.Context,
		id string,
	) error
	GetByID(
		ctx context.Context,
		id string,
	) (
		*coreentity.RecurrencePattern,
		error,
	)
	GetFailedJobs(
		ctx context.Context,
	) (
		[]coreentity.RecurrencePattern,
		error,
	)
	GetTodayJobs(
		ctx context.Context,
	) (
		[]coreentity.RecurrencePattern,
		error,
	)
	List(
		ctx context.Context,
	) (
		[]coreentity.RecurrencePattern,
		error,
	)
}
