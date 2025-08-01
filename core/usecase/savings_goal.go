package usecase

import (
	"context"
	"errors"

	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/port"
)

type SavingsGoal struct {
	accountRepo     port.Account
	savingsGoalRepo port.SavingsGoal
}

func NewSavingsGoalUseCase(
	accountRepo port.Account,
	savingsGoalRepo port.SavingsGoal,
) *SavingsGoal {
	return &SavingsGoal{accountRepo: accountRepo, savingsGoalRepo: savingsGoalRepo}
}

func (sg *SavingsGoal) Create(
	ctx context.Context,
	savingsGoal coreentity.SavingsGoal,
) (
	*coreentity.SavingsGoal,
	error,
) {
	err := savingsGoal.Validate()
	if err != nil {
		return nil, err
	}
	id, err := sg.savingsGoalRepo.Create(
		ctx,
		savingsGoal,
	)
	if err != nil {
		return nil, err
	}
	savingsGoal.ID = id

	return &savingsGoal, nil
}

func (sg *SavingsGoal) Update(
	ctx context.Context,
	savingsGoal coreentity.SavingsGoal,
) error {
	err := savingsGoal.Validate()
	if err != nil {
		return err
	}
	return sg.savingsGoalRepo.Update(
		ctx,
		savingsGoal,
	)
}

func (sg *SavingsGoal) Delete(
	ctx context.Context,
	id string,
) error {
	err := sg.savingsGoalRepo.Delete(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return coreentity.ErrSavingsGoalNotFound
		}
		return err
	}
	return nil
}

func (sg *SavingsGoal) GetByID(
	ctx context.Context,
	id string,
) (
	*coreentity.SavingsGoal,
	error,
) {
	savingsGoal, err := sg.savingsGoalRepo.GetByID(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return nil, coreentity.ErrSavingsGoalNotFound
		}
		return nil, err
	}
	return savingsGoal, nil
}

func (sg *SavingsGoal) List(
	ctx context.Context,
	params coreentity.ListSavingsGoalsParams,
) (
	*coreentity.SavingsGoalsList,
	error,
) {
	savingsGoals, err := sg.savingsGoalRepo.List(
		ctx,
		params,
	)
	if err != nil {
		return nil, err
	}
	return &savingsGoals, nil
}

func (sg *SavingsGoal) IsActive(
	ctx context.Context,
	id string,
) (
	bool,
	error,
) {
	isActive, err := sg.savingsGoalRepo.IsActive(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return false, coreentity.ErrSavingsGoalNotFound
		}
		return false, err
	}
	return isActive, nil
}

func (sg *SavingsGoal) MarkAsAbandoned(
	ctx context.Context,
	id string,
) error {
	err := sg.savingsGoalRepo.MarkAsAbandoned(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return coreentity.ErrSavingsGoalNotFound
		}
		return err
	}
	return nil
}

func (sg *SavingsGoal) MarkAsCompleted(
	ctx context.Context,
	id string,
) error {
	err := sg.savingsGoalRepo.MarkAsCompleted(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return coreentity.ErrSavingsGoalNotFound
		}
		return err
	}
	return nil
}
