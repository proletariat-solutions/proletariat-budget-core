package usecase

import (
	"context"
	"errors"

	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/port"
)

type Ingress struct {
	accountRepo port.Account
	ingressRepo port.Ingress
	txManager   port.TransactionManager
}

func NewIngressUseCase(
	accountRepo port.Account,
	txManager port.TransactionManager,
) *Ingress {
	return &Ingress{accountRepo: accountRepo, txManager: txManager}
}

func (u *Ingress) Create(
	ctx context.Context,
	ingress coreentity.Ingress,
	recurrencePattern *coreentity.RecurrencePattern, // Optional, used only if ingress is result of a recurring transaction
) (
	*coreentity.Ingress,
	error,
) {
	errIngressValidation := ingress.Validate(false)
	if errIngressValidation != nil {
		return nil, errIngressValidation
	}
	// Validate account
	account, err := u.accountRepo.GetByID(
		ctx,
		ingress.Transaction.AccountID,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return nil, coreentity.ErrAccountNotFound
		}

		return nil, err
	}
	if !account.Active {
		return nil, coreentity.ErrAccountInactive
	}
	// Create transaction
	errTx := u.txManager.WithDatabaseTransaction(
		ctx,
		func(
			ctx context.Context,
			tx port.TransactionContext,
		) error {
			// Set recurrence pattern id if provided
			ingress.FromRecurrencePatternID = &recurrencePattern.ID

			// Update account balance
			account.CreditBalance(ingress.Transaction.Amount)

			// prepare transaction
			ingress.Transaction.BalanceAfter = &account.CurrentBalance
			statusCompleted := coreentity.TransactionStatusCompleted
			ingress.Transaction.Status = &statusCompleted

			// Create transaction
			txID, errCreateTx := tx.GetTransactionRepo().Create(
				ctx,
				*ingress.Transaction,
			)
			if errCreateTx != nil {
				return errCreateTx
			}
			ingress.Transaction.ID = &txID

			// Update account
			errUpdateAccount := tx.GetAccountRepo().Update(
				ctx,
				*account,
			)
			if errUpdateAccount != nil {
				return errUpdateAccount
			}

			// Create ingress
			ingressID, errCreateIngress := tx.GetIngressRepo().Create(
				ctx,
				ingress,
			)
			if errCreateIngress != nil {
				return errCreateIngress
			}

			ingress.ID = ingressID

			return nil
		},
	)

	if errTx != nil {
		return nil, errTx
	}

	return &ingress, nil
}

func (u *Ingress) GetByID(
	ctx context.Context,
	id string,
) (
	*coreentity.Ingress,
	error,
) {
	ingress, err := u.ingressRepo.GetByID(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return nil, coreentity.ErrIngressNotFound
		}

		return nil, err
	}

	return &ingress, nil
}

func (u *Ingress) List(
	ctx context.Context,
	params coreentity.IngressListParams,
) (
	*coreentity.IngressList,
	error,
) {
	ingresses, err := u.ingressRepo.List(
		ctx,
		params,
	)
	if err != nil {
		return nil, err
	}

	return &ingresses, nil
}
