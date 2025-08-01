package usecase

import (
	"context"
	"errors"

	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/port"
)

type Category struct {
	categoryRepo port.Category
}

func NewCategoryUseCase(categoryRepo port.Category) *Category {
	return &Category{categoryRepo: categoryRepo}
}

func (uc *Category) ListCategories(
	ctx context.Context,
	categoryType *coreentity.CategoryType,
) (
	[]coreentity.Category,
	error,
) {
	var categories []coreentity.Category
	var err error
	if categoryType == nil {
		categories, err = uc.categoryRepo.List(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		categories, err = uc.categoryRepo.FindByType(
			ctx,
			*categoryType,
		)
		if err != nil {
			return nil, err
		}
	}

	return categories, nil
}

func (uc *Category) GetCategory(
	ctx context.Context,
	id string,
) (
	*coreentity.Category,
	error,
) {
	category, err := uc.categoryRepo.GetByID(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return nil, coreentity.ErrCategoryNotFound
		}

		return nil, err
	}

	return category, nil
}

func (uc *Category) CreateCategory(
	ctx context.Context,
	category coreentity.Category,
) (
	*coreentity.Category,
	error,
) {
	if err := category.Validate(); err != nil {
		return nil, err
	}
	id, err := uc.categoryRepo.Create(
		ctx,
		category,
	)
	if err != nil {
		return nil, err
	}

	return uc.GetCategory(
		ctx,
		id,
	)
}

func (uc *Category) UpdateCategory(
	ctx context.Context,
	category coreentity.Category,
) (
	*coreentity.Category,
	error,
) {
	if err := category.Validate(); err != nil {
		return nil, err
	}

	err := uc.categoryRepo.Update(ctx, category)
	if err != nil {
		if errors.Is(err, port.ErrRecordNotFound) {
			return nil, coreentity.ErrCategoryNotFound
		}

		return nil, err
	}

	// Single call to get updated entity
	return uc.GetCategory(ctx, category.ID)
}

func (uc *Category) DeleteCategory(
	ctx context.Context,
	id string,
) error {
	err := uc.categoryRepo.Delete(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return coreentity.ErrCategoryNotFound
		}

		return err
	}

	return nil
}

func (uc *Category) Activate(ctx context.Context, id string) error {
	// Single repository call for activation
	err := uc.categoryRepo.Activate(ctx, id)
	if err != nil {
		if errors.Is(err, port.ErrRecordNotFound) {
			return coreentity.ErrCategoryNotFound
		}

		return err
	}

	return nil
}

func (uc *Category) Deactivate(ctx context.Context, id string) error {
	// Single repository call for deactivation
	err := uc.categoryRepo.Deactivate(ctx, id)
	if err != nil {
		if errors.Is(err, port.ErrRecordNotFound) {
			return coreentity.ErrCategoryNotFound
		}

		return err
	}

	return nil
}
