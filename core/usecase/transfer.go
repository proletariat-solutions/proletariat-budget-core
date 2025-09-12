package usecase

import (
	"context"
	"errors"

	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/port"
)

type Transfer struct {
	accountRepo port.Account
	txManager   port.TransactionManager
}

func (t *Transfer) Create(
	ctx context.Context,
	transfer coreentity.Transfer,
	recurrencePattern *coreentity.RecurrencePattern,
) (
	*coreentity.Transfer,
	error,
) {
	sourceAccount, errGetSrcAcct := t.getAccount(
		ctx,
		*transfer.SourceAccount.ID,
	)
	if errGetSrcAcct != nil {
		return nil, errGetSrcAcct
	}
	destinationAccount, errGetTgtAcct := t.getAccount(
		ctx,
		*transfer.DestinationAccount.ID,
	)
	if errGetTgtAcct != nil {
		return nil, errGetTgtAcct
	}
	transfer.SourceAccount = sourceAccount
	transfer.DestinationAccount = destinationAccount

	if errValidate := transfer.Validate(); errValidate != nil {
		return nil, errValidate
	}

	errTransfer := transfer.DoTransfer()
	if errTransfer != nil {
		return nil, errTransfer
	}

	errTx := u.txManager.WithDatabaseTransaction(
		ctx,
		func(
			ctx context.Context,
			tx port.TransactionContext,
		) error {

		}
	)

	if recurrencePattern != nil {
		transfer.RecurrenceTransactionInfo = &coreentity.RecurrentTransactionInfo{
			FromRecurrencePatternID: &recurrencePattern.ID,
			IsTemplate:              false,
		}
	}
	return nil, nil
}

func (t *Transfer) getAccount(
	ctx context.Context,
	accountID string,
) (
	*coreentity.Account,
	error,
) {
	account, errGetAcct := t.accountRepo.GetByID(
		ctx,
		accountID,
	)
	if errGetAcct != nil {
		if errors.Is(
			errGetAcct,
			port.ErrRecordNotFound,
		) {
			return nil, coreentity.ErrAccountNotFound
		}
		return nil, errGetAcct
	}
	return account, nil
}
