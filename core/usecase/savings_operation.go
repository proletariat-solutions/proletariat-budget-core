package usecase

import (
	"context"
	"proletariat-budget-core/core/port"

	"proletariat-budget-core/core/domain/coreentity"
)

type SavingOperation struct {
	accountRepo     port.Account
	savingsGoalRepo port.SavingsGoal
	txManager       port.TransactionManager
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

	transfer, err := s.createSavingTransfer(operation)
	if err != nil {
		return nil, err
	}

}

func (s *SavingOperation) createSavingTransfer(operation coreentity.SavingsContribution) (
	*coreentity.Transfer,
	error,
) {

}
