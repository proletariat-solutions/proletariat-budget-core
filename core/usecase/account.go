package usecase

import (
	"context"
	"errors"

	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/port"
)

type Account struct {
	accountRepo         port.Account
	householdMemberRepo port.HouseholdMember
}

func NewAccountUseCase(
	accountRepo port.Account,
	householdMemberRepo port.HouseholdMember,
) *Account {
	return &Account{
		accountRepo:         accountRepo,
		householdMemberRepo: householdMemberRepo,
	}
}

func (a *Account) Create(
	ctx context.Context,
	account coreentity.Account,
) (
	*string,
	error,
) {
	householdMember, err := a.householdMemberRepo.GetByID(
		ctx,
		*account.OwnerID,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return nil, coreentity.ErrMemberNotFound
		}

		return nil, err
	} else if !householdMember.Active {
		return nil, coreentity.ErrMemberInactive
	}
	account.Owner = householdMember
	ID, err := a.accountRepo.Create(
		ctx,
		account,
	)
	if err != nil {
		return nil, err
	}

	return ID, nil
}

func (a *Account) GetByID(
	ctx context.Context,
	id string,
) (
	*coreentity.Account,
	error,
) {
	account, err := a.accountRepo.GetByID(
		ctx,
		id,
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

	return account, nil
}

func (a *Account) Update(
	ctx context.Context,
	account coreentity.Account,
) (
	*coreentity.Account,
	error,
) {
	err := a.accountRepo.Update(
		ctx,
		account,
	)
	if err != nil {
		return nil, err
	}
	updatedAccount, errGet := a.accountRepo.GetByID(
		ctx,
		*account.ID,
	)
	if errGet != nil {
		if errors.Is(
			errGet,
			port.ErrRecordNotFound,
		) {
			return nil, coreentity.ErrAccountNotFound
		}

		return nil, errGet
	}

	return updatedAccount, nil
}

func (a *Account) Deactivate(
	ctx context.Context,
	id string,
) error {
	account, err := a.accountRepo.GetByID(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return coreentity.ErrAccountNotFound
		}

		return err
	}
	if errInactive := account.SetInactive(); errInactive != nil {
		return errInactive
	}
	errUpdate := a.accountRepo.Update(
		ctx,
		*account,
	)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (a *Account) Activate(
	ctx context.Context,
	id string,
) error {
	account, err := a.accountRepo.GetByID(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return coreentity.ErrAccountNotFound
		}

		return err
	}
	if errActive := account.SetActive(); errActive != nil {
		return errActive
	}
	errUpdate := a.accountRepo.Update(
		ctx,
		*account,
	)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (a *Account) Delete(
	ctx context.Context,
	id string,
) error {
	hasTransactions, errHasTransactions := a.accountRepo.HasTransactions(
		ctx,
		id,
	)
	if errHasTransactions != nil {
		return errHasTransactions
	}
	if hasTransactions {
		return coreentity.ErrAccountHasTransactions
	}
	err := a.accountRepo.Delete(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return coreentity.ErrAccountNotFound
		}

		return err
	}

	return nil
}

func (a *Account) List(
	ctx context.Context,
	params coreentity.AccountListParams,
) (
	*coreentity.AccountList,
	error,
) {
	accounts, err := a.accountRepo.List(
		ctx,
		params,
	)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (a *Account) HasTransactions(
	ctx context.Context,
	id string,
) (
	bool,
	error,
) {
	hasTransactions, err := a.accountRepo.HasTransactions(
		ctx,
		id,
	)
	if err != nil {
		return false, err
	}

	return hasTransactions, nil
}
