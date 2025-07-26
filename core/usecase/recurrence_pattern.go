package usecase

import (
	"context"
	"errors"

	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/port"
)

type RecurrencePattern struct {
	recurrencePatternRepo port.RecurrencePattern
}

func (rp *RecurrencePattern) Create(
	ctx context.Context,
	recurrencePattern coreentity.RecurrencePattern,
) (
	*coreentity.RecurrencePattern,
	error,
) {
	validationErr := recurrencePattern.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	recurrencePatternID, err := rp.recurrencePatternRepo.Create(
		ctx,
		recurrencePattern,
	)
	if err != nil {
		return nil, err
	}
	recurrencePattern.ID = recurrencePatternID

	return &recurrencePattern, nil
}

func (rp *RecurrencePattern) GetByID(
	ctx context.Context,
	id string,
) (
	*coreentity.RecurrencePattern,
	error,
) {
	recurrencePattern, err := rp.recurrencePatternRepo.GetByID(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return nil, coreentity.ErrRecurrencePatternNotFound
		}

		return nil, err
	}

	return recurrencePattern, nil
}

func (rp *RecurrencePattern) Update(
	ctx context.Context,
	recurrencePattern coreentity.RecurrencePattern,
) (
	*coreentity.RecurrencePattern,
	error,
) {
	validationErr := recurrencePattern.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	err := rp.recurrencePatternRepo.Update(
		ctx,
		recurrencePattern,
	)
	if err != nil {
		return nil, err
	}

	return &recurrencePattern, nil
}

func (rp *RecurrencePattern) Delete(ctx context.Context, id string) error {
	err := rp.recurrencePatternRepo.Delete(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return coreentity.ErrRecurrencePatternNotFound
		}

		return err
	}

	return nil
}
