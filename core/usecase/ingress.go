package usecase

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/port"
)

type Ingress struct {
	accountRepo  port.Account
	categoryRepo port.Category
	ingressRepo  port.Ingress
	txManager    port.TransactionManager
}

func NewIngressUseCase(
	accountRepo port.Account,
	categoryRepo port.Category,
	ingressRepo port.Ingress,
	txManager port.TransactionManager,
) *Ingress {
	return &Ingress{
		accountRepo:  accountRepo,
		categoryRepo: categoryRepo,
		ingressRepo:  ingressRepo,
		txManager:    txManager,
	}
}

func (i *Ingress) Create(
	ctx context.Context,
	ingress coreentity.Ingress,
	recurrencePattern *coreentity.RecurrencePattern, // Optional, used only if ingress is result of a recurring transaction
) (
	*coreentity.Ingress,
	error,
) {
	errIngressValidation := ingress.Validate()

	if errIngressValidation != nil {
		return nil, errIngressValidation
	}

	account, errValidateAccount := i.validateAccount(
		ctx,
		ingress.Transaction.AccountID,
	)

	if errValidateAccount != nil {
		return nil, errValidateAccount
	}

	category, errValidateCategory := i.validateCategory(
		ctx,
		ingress.Category.ID,
	)
	if errValidateCategory != nil {
		return nil, errValidateCategory
	}

	ingress.Category = category

	// Create transaction
	errTx := i.txManager.WithDatabaseTransaction(
		ctx,
		func(
			ctx context.Context,
			tx port.TransactionContext,
		) error {
			if recurrencePattern != nil {
				// Set recurrence pattern id if provided
				recurrence := coreentity.RecurrentTransactionInfo{
					FromRecurrencePatternID: &recurrencePattern.ID,
					IsTemplate:              false,
				}

				ingress.RecurrenceTransactionInfo = &recurrence
			}

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
			// Link tags if present
			errLink := i.linkTags(
				ctx,
				ingressID,
				ingress.Tags,
				tx.GetTagsRepo(),
			)
			if errLink != nil {
				if recurrencePattern != nil {
					// Logging and notifying error to the user, since this is a recoverable error
					// TODO: notify user
					log.Err(errLink).Msgf(
						"Failed to link tags to ingress %s",
						ingressID,
					)
				} else {
					return errLink
				}
			}

			return nil
		},
	)

	if errTx != nil {
		return nil, errTx
	}

	return &ingress, nil
}

func (i *Ingress) GetByID(
	ctx context.Context,
	id string,
) (
	*coreentity.Ingress,
	error,
) {
	ingress, err := i.ingressRepo.GetByID(
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

func (i *Ingress) List(
	ctx context.Context,
	params coreentity.IngressListParams,
) (
	*coreentity.IngressList,
	error,
) {
	ingresses, err := i.ingressRepo.List(
		ctx,
		params,
	)
	if err != nil {
		return nil, err
	}

	return &ingresses, nil
}

func (i *Ingress) validateAccount(
	ctx context.Context,
	accountID string,
) (
	*coreentity.Account,
	error,
) {
	// Validate account
	account, err := i.accountRepo.GetByID(
		ctx,
		accountID,
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

	return account, nil
}

func (i *Ingress) validateCategory(
	ctx context.Context,
	categoryID string,
) (
	*coreentity.Category,
	error,
) {
	category, err := i.categoryRepo.GetByID(
		ctx,
		categoryID,
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
	if !category.Active {
		return nil, coreentity.ErrCategoryInactive
	}

	return category, nil
}

func (i *Ingress) linkTags(
	ctx context.Context,
	expID string,
	tags *[]*coreentity.Tag,
	repo port.Tags,
) error {
	if tags == nil || len(*tags) == 0 {
		return nil
	}

	err := repo.LinkTagsToType(
		ctx,
		expID,
		tags,
	)

	return err
}
