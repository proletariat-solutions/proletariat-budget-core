package usecase

import (
	"context"

	"proletariat-budget-core/core/domain/coreentity"
)

type SavingOperation struct{}

func (s *SavingOperation) CreateContribution(
	ctx context.Context,
	operation coreentity.SavingsContribution,
	recurrencePattern *coreentity.RecurrencePattern, // Optional, used only if ingress is result of a recurring transaction
) (
	*coreentity.SavingsContribution,
	error,
) {
	panic("Savings contribution creation is not implemented yet")
}
