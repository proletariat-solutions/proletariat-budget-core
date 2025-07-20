package coreentity

import (
	"errors"
	"time"

	"proletariat-budget-core/core/domain/misc"
)

var (
	ErrTransferNotFound              = errors.New("transfer not found")
	ErrSourceAccountNotEnoughBalance = errors.New("source account does not have enough balance")
	ErrRateMultiplierMustBePositive  = errors.New("exchange rate multiplier must be positive")
	ErrFeesMustBePositive            = errors.New("fees must be positive")
	ErrSourceAccountNotFound         = errors.New("source account not found")
	ErrDestinationAccountNotFound    = errors.New("destination account not found")
	ErrSourceAccountNotActive        = errors.New("source account is not active")
	ErrDestinationAccountNotActive   = errors.New("destination account is not active")
)

type Transfer struct {
	ID                     string   `json:"id"`
	SourceAccount          *Account `json:"source_account"`
	DestinationAccount     *Account `json:"destination_account"`
	ExchangeRateMultiplier float32  `json:"exchange_rate_multiplier"`
	// Fees are calculated based on the amount and currency of the source account.
	Fees                float32      `json:"fees"`
	OutgoingTransaction *Transaction `json:"outgoing_transaction"`
	IncomingTransaction *Transaction `json:"incoming_transaction"`
	TransferType        TransferType `json:"transfer_type"`
	CreatedAt           time.Time    `json:"created_at"`
}

type TransferType string

const (
	// TransferTypeCurrencyConversion between accounts with different currencies
	TransferTypeCurrencyConversion TransferType = "currency_conversion"
	// TransferTypeDeposit cash to account
	TransferTypeDeposit TransferType = "deposit"
	// TransferTypeWithdrawal account to cash
	TransferTypeWithdrawal TransferType = "withdrawal"
	// TransferTypeTransfer between accounts
	TransferTypeTransfer TransferType = "transfer"
	// TransferTypeSavingOperation deposit or withdrawal to savings goal, source & destination accounts can be the same
	TransferTypeSavingOperation TransferType = "saving_operation"
)

func (t *Transfer) Validate() error {
	if !t.SourceAccount.Active {
		return ErrSourceAccountNotActive
	}

	if !t.DestinationAccount.Active {
		return ErrDestinationAccountNotActive
	}

	if t.ExchangeRateMultiplier <= 0 {
		return ErrRateMultiplierMustBePositive
	}

	if t.Fees < 0 {
		return ErrFeesMustBePositive
	}

	if t.SourceAccount.CurrentBalance < t.OutgoingTransaction.Amount {
		return ErrSourceAccountNotEnoughBalance
	}

	return nil
}

type TransferList struct {
	Transfers []*Transfer       `json:"transfers"`
	Metadata  misc.ListMetadata `json:"metadata"`
}

type ListTransfersParams struct {
	// SourceAccountId Filter by source account ID
	SourceAccountId *string `json:"sourceAccountId,omitempty"`

	// DestinationAccountId Filter by destination account ID
	DestinationAccountId *string `json:"destinationAccountId,omitempty"`

	// StartDate Filter by start date (inclusive)
	StartDate *time.Time `json:"startDate,omitempty"`

	// EndDate Filter by end date (inclusive)
	EndDate *time.Time `json:"endDate,omitempty"`

	misc.ListParams
}
