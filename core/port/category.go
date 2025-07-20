package port

import (
	"context"

	"proletariat-budget-core/core/domain/coreentity"
)

//go:generate mockgen -source=category.go -destination=../../test/mocks/mock_categoryrepo.go -package mocks
type Category interface {
	Create(
		ctx context.Context,
		category coreentity.Category,
	) (
		string,
		error,
	)
	Update(
		ctx context.Context,
		category coreentity.Category,
	) error
	Delete(
		ctx context.Context,
		id string,
	) error
	GetByID(
		ctx context.Context,
		id string,
	) (
		*coreentity.Category,
		error,
	)
	List(ctx context.Context) (
		[]coreentity.Category,
		error,
	)
	FindByType(
		ctx context.Context,
		categoryType coreentity.CategoryType,
	) (
		[]coreentity.Category,
		error,
	)
	FindByIDs(
		ctx context.Context,
		ids []string,
	) (
		[]coreentity.Category,
		error,
	)
}
