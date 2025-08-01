package port

import (
	"context"

	"proletariat-budget-core/core/domain/coreentity"
)

//go:generate mockgen -source=account.go -destination=../../test/mocks/mock_accountrepo.go -package mocks
type Account interface {
	Create(
		ctx context.Context,
		account coreentity.Account,
	) (
		*string,
		error,
	)
	GetByID(
		ctx context.Context,
		id string,
	) (
		*coreentity.Account,
		error,
	)
	Update(
		ctx context.Context,
		account coreentity.Account,
	) error
	Delete(
		ctx context.Context,
		id string,
	) error
	List(
		ctx context.Context,
		params coreentity.AccountListParams,
	) (
		*coreentity.AccountList,
		error,
	)
	HasTransactions(
		ctx context.Context,
		id string,
	) (
		bool,
		error,
	)
	Activate(
		ctx context.Context,
		id string,
	) error

	Deactivate(
		ctx context.Context,
		id string,
	) error
}
