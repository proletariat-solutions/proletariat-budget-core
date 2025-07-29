package usecase

import (
	"context"

	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/port"
)

type Transfer struct {
	txManager port.TransactionManager
}

func (t *Transfer) Create(
	ctx context.Context,
	transfer coreentity.Transfer,
	recurrencePattern *coreentity.RecurrencePattern,
) (
	*coreentity.Transfer,
	error,
) {
	panic("implement me")
}
