package port

import "context"

// TransactionManager defines the interface for managing database transactions
//
//go:generate mockgen -source=transaction_manager.go -destination=../../test/mocks/mock_transaction_manager.go -package mocks
type TransactionManager interface {
	// WithTransaction executes the given function within a database transaction
	// If the function returns an error, the transaction is rolled back
	// Otherwise, the transaction is committed
	WithDatabaseTransaction(
		ctx context.Context,
		fn func(
			ctx context.Context,
			tx TransactionContext,
		) error,
	) error
}

// TransactionContext provides access to repositories within a transaction
type TransactionContext interface {
	GetAccountRepo() Account
	GetTransactionRepo() Transaction
	GetTransferRepo() Transfer
	GetIngressRepo() Ingress
	GetExpenditureRepo() Expenditure
	GetCategoryRepo() Category
	GetTagsRepo() Tags
}
