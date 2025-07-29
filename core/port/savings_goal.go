package port

import (
	"context"

	"proletariat-budget-core/core/domain/coreentity"
)

//go:generate mockgen -source=savings_goal.go -destination=../../test/mocks/mock_savings_goalrepo.go -package mocks
type SavingsGoal interface {
	Create(
		ctx context.Context,
		savingsGoal coreentity.SavingsGoal,
	) (
		string,
		error,
	)
	Update(
		ctx context.Context,
		id string,
		savingsGoal coreentity.SavingsGoal,
	) error
	Delete(
		ctx context.Context,
		id string,
	) error
	GetByID(
		ctx context.Context,
		id string,
	) (
		*coreentity.SavingsGoal,
		error,
	)
	List(ctx context.Context) (
		[]coreentity.SavingsGoal,
		error,
	)
	IsActive(
		ctx context.Context,
		id string,
	) (
		bool,
		error,
	)
	MarkAsCompleted(
		ctx context.Context,
		id string,
	) error
	MarkAsAbandoned(
		ctx context.Context,
		id string,
	) error

	CreateWithdrawal(
		ctx context.Context,
		withdrawal coreentity.SavingsWithdrawal,
	) (
		string,
		error,
	)
	DeleteWithdrawal(
		ctx context.Context,
		id string,
	) error
	GetWithdrawalByID(
		ctx context.Context,
		id string,
	) (
		*coreentity.SavingsWithdrawal,
		error,
	)

	CreateContribution(
		ctx context.Context,
		contribution coreentity.SavingsContribution,
	) (
		string,
		error,
	)
	DeleteContribution(
		ctx context.Context,
		id string,
	) error
	GetContributionByID(
		ctx context.Context,
		id string,
	) (
		*coreentity.SavingsContribution,
		error,
	)

	ListSavingsTransactions(
		ctx context.Context,
		params coreentity.ListSavingsTransactionsParams,
	) (
		*coreentity.SavingsTransactionList,
		error,
	)
	UpdateTemplate(
		ctx context.Context,
		template coreentity.SavingsContribution,
	) error
	DeleteTemplate(
		ctx context.Context,
		template coreentity.SavingsContribution,
	) error
}
