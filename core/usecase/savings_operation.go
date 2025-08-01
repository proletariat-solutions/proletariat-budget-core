package usecase

import (
	"context"

	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/port"
)

type SavingOperation struct {
	accountRepo     port.Account
	savingsGoalRepo port.SavingsGoal
	txManager       port.TransactionManager
}

func NewSavingOperationUseCase(
	accountRepo port.Account,
	savingsGoalRepo port.SavingsGoal,
	txManager port.TransactionManager,
) *SavingOperation {
	return &SavingOperation{
		accountRepo:     accountRepo,
		savingsGoalRepo: savingsGoalRepo,
		txManager:       txManager,
	}
}

func (s *SavingOperation) CreateContribution(
	ctx context.Context,
	operation coreentity.SavingsContribution,
	recurrencePattern *coreentity.RecurrencePattern, // Optional, used only if ingress is result of a recurring transaction
) (
	*coreentity.SavingsContribution,
	error,
) {
	validationErr := operation.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	_, err := s.createSavingTransfer(operation)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *SavingOperation) createSavingTransfer(operation coreentity.SavingsContribution) (
	*coreentity.Transfer,
	error,
) {
	return nil, nil // Placeholder for actual implementation
}
