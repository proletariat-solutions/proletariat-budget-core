package coreentity

import (
	"errors"
	"time"
)

var (
	ErrTransactionNotFound = errors.New("transaction not found")
)

type Transaction struct {
	ID              *string            `json:"id"`
	AccountID       string             `json:"account_id"`
	Amount          float32            `json:"amount"`
	Currency        string             `json:"currency"`
	TransactionDate time.Time          `json:"transaction_date"`
	Description     string             `json:"description"`
	TransactionType TransactionType    `json:"transaction_type"`
	BalanceAfter    *float32           `json:"balance_after"`
	Status          *TransactionStatus `json:"status"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
}

type ListTransactionsParams struct{}

type TransactionList struct{}

type TransactionType string

const (
	TransactionTypeExpenditure TransactionType = "expenditure"
	TransactionTypeIngress     TransactionType = "ingress"
	TransactionTypeTransfer    TransactionType = "transfer"
	TransactionTypeRollback    TransactionType = "rollback"
)

// String returns the string representation of the transaction type
func (t TransactionType) String() string {
	return string(t)
}

// TransactionStatus represents the status of a financial transaction
type TransactionStatus string

// Transaction status enum values
const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusCompleted TransactionStatus = "completed"
	TransactionStatusFailed    TransactionStatus = "failed"
	TransactionStatusCancelled TransactionStatus = "canceled"
	TransactionStatusReversed  TransactionStatus = "reversed"
)

// String returns the string representation of the transaction status
func (s TransactionStatus) String() string {
	return string(s)
}

func RollbackTransaction(
	t *Transaction,
	description string,
) *Transaction {
	statusCompleted := TransactionStatusCompleted
	balanceAfter := *t.BalanceAfter + t.Amount

	return &Transaction{
		ID:              t.ID,
		AccountID:       t.AccountID,
		Amount:          -t.Amount,
		Currency:        t.Currency,
		TransactionDate: time.Now(),
		Description:     description,
		TransactionType: TransactionTypeRollback,
		BalanceAfter:    &balanceAfter,
		Status:          &statusCompleted,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}
