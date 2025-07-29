package port

import (
	"context"

	"proletariat-budget-core/core/domain/coreentity"
)

//go:generate mockgen -source=expenditure.go -destination=../../test/mocks/mock_expenditurerepo.go -package mocks
type Expenditure interface {
	// Expenditure operations
	Create(
		ctx context.Context,
		expenditure coreentity.Expenditure,
	) (
		string,
		error,
	)
	GetByID(
		ctx context.Context,
		id string,
	) (
		*coreentity.Expenditure,
		error,
	)
	FindExpenditures(
		ctx context.Context,
		queryParams coreentity.ExpenditureListParams,
	) (
		*coreentity.ExpenditureList,
		error,
	)
	UpdateTemplate(
		ctx context.Context,
		template coreentity.Expenditure,
	) error
	DeleteTemplate(
		ctx context.Context,
		template coreentity.Expenditure,
	) error
}
