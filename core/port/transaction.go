package port

import (
	"context"

	"proletariat-budget-core/core/domain/coreentity"
)

//go:generate mockgen -source=transaction.go -destination=../../test/mocks/mock_transactionrepo.go -package mocks
type Transaction interface {
	Create(
		ctx context.Context,
		transaction coreentity.Transaction,
	) (
		string,
		error,
	)
	GetByID(
		ctx context.Context,
		id string,
	) (
		*coreentity.Transaction,
		error,
	)
	List(
		ctx context.Context,
		params coreentity.ListTransactionsParams,
	) (
		*coreentity.TransactionList,
		error,
	)
}
