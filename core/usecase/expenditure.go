package usecase

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/port"
)

type Expenditure struct {
	expenditureRepo port.Expenditure
	accountRepo     port.Account
	tagsRepo        port.Tags
	categoryRepo    port.Category
	transactionRepo port.Transaction
	txManager       port.TransactionManager
}

func NewExpenditureUseCase(
	expenditureRepo port.Expenditure,
	accountRepo port.Account,
	tagsRepo port.Tags,
	categoryRepo port.Category,
	transactionRepo port.Transaction,
	txManager port.TransactionManager,
) *Expenditure {
	return &Expenditure{
		expenditureRepo: expenditureRepo,
		accountRepo:     accountRepo,
		tagsRepo:        tagsRepo,
		categoryRepo:    categoryRepo,
		transactionRepo: transactionRepo,
		txManager:       txManager,
	}
}

func (u *Expenditure) Create(
	ctx context.Context,
	expenditure coreentity.Expenditure,
	recurrencePattern *coreentity.RecurrencePattern, // Optional, used only if expenditure is result of a recurring transaction
) (
	*coreentity.Expenditure,
	error,
) {
	// Validate expenditure
	if err := expenditure.Validate(); err != nil {
		return nil, err
	}

	// Validate account
	account, err := u.validateAccount(
		ctx,
		expenditure.Transaction.AccountID,
	)
	if err != nil {
		return nil, err
	}

	// Validate category
	category, err := u.validateCategory(
		ctx,
		expenditure.Category.ID,
	)
	if err != nil {
		return nil, err
	}

	expenditure.Category = category

	var expID string

	errTx := u.txManager.WithDatabaseTransaction(
		ctx,
		func(
			ctx context.Context,
			tx port.TransactionContext,
		) error {
			accountRepo := tx.GetAccountRepo()
			transactionRepo := tx.GetTransactionRepo()
			expenditureRepo := tx.GetExpenditureRepo()
			tagsRepo := tx.GetTagsRepo()

			// Process transaction
			errProcess := u.processTransaction(
				ctx,
				account,
				&expenditure,
				accountRepo,
				transactionRepo,
			)
			if errProcess != nil {
				return errProcess
			}

			var errCreate error
			// Create expenditure record
			expID, errCreate = expenditureRepo.Create(
				ctx,
				expenditure,
			)
			if errCreate != nil {
				return errCreate
			}

			// Link tags if present
			errLink := u.linkTags(
				ctx,
				expID,
				expenditure.Tags,
				tagsRepo,
			)
			if errLink != nil {
				if recurrencePattern != nil {
					// Logging and notifying error to the user, since this is a recoverable error
					// TODO: notify user
					log.Err(errLink).Msgf(
						"Failed to link tags to expenditure %s",
						expID,
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

	return u.expenditureRepo.GetByID(
		ctx,
		expID,
	)
}

func (u *Expenditure) Get(
	ctx context.Context,
	id string,
) (
	*coreentity.Expenditure,
	error,
) {
	expenditure, err := u.expenditureRepo.GetByID(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return nil, coreentity.ErrExpenditureNotFound
		}

		return nil, err
	}

	return expenditure, nil
}

func (u *Expenditure) List(
	ctx context.Context,
	params coreentity.ExpenditureListParams,
) (
	*coreentity.ExpenditureList,
	error,
) {
	return u.expenditureRepo.FindExpenditures(
		ctx,
		params,
	)
}

func (u *Expenditure) validateAccount(
	ctx context.Context,
	accountID string,
) (
	*coreentity.Account,
	error,
) {
	account, err := u.accountRepo.GetByID(
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

func (u *Expenditure) validateCategory(
	ctx context.Context,
	categoryID string,
) (
	*coreentity.Category,
	error,
) {
	category, err := u.categoryRepo.GetByID(
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

func (u *Expenditure) processTransaction(
	ctx context.Context,
	account *coreentity.Account,
	expenditure *coreentity.Expenditure,
	accountRepo port.Account,
	transactionRepo port.Transaction,
) error {
	if errDebit := account.DebitBalance(expenditure.Transaction.Amount); errDebit != nil {
		return errDebit
	}

	expenditure.Transaction.BalanceAfter = &account.CurrentBalance
	statusCompleted := coreentity.TransactionStatusCompleted
	expenditure.Transaction.Status = &statusCompleted

	txID, err := transactionRepo.Create(
		ctx,
		*expenditure.Transaction,
	)
	if err != nil {
		return err
	}

	expenditure.Transaction.ID = &txID

	return accountRepo.Update(
		ctx,
		*account,
	)
}

func (u *Expenditure) linkTags(
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
	if err != nil {
		return err
	}

	return nil
}
