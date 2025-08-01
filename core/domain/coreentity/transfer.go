package coreentity

import (
	"errors"
	"time"

	"proletariat-budget-core/core/domain/misc"
)

var (
	ErrTransferNotFound                      = errors.New("transfer not found")
	ErrTransferSourceAccountNotEnoughBalance = errors.New("source account does not have enough balance")
	ErrTransferRateMultiplierMustBePositive  = errors.New("exchange rate multiplier must be positive and not zero")
	ErrTransferFeesMustBePositive            = errors.New("fees must be positive")
	ErrTransferSourceAccountNotFound         = errors.New("source account not found")
	ErrTransferDestinationAccountNotFound    = errors.New("destination account not found")
	ErrTransferSourceAccountNotActive        = errors.New("source account is not active")
	ErrTransferDestinationAccountNotActive   = errors.New("destination account is not active")
	ErrTransferNotATemplate                  = errors.New("transfer is not a template")
	ErrTransferExchangeRateInvalid           = errors.New("exchange rate must be 1 for same currency accounts")
)

type Transfer struct {
	ID                 string   `json:"id"`
	SourceAccount      *Account `json:"source_account"`
	DestinationAccount *Account `json:"destination_account"`
	// ExchangeRateMultiplier has 1 if same currency
	ExchangeRateMultiplier float32 `json:"exchange_rate_multiplier"`
	// Fees are calculated based on the amount and currency of the source account.
	Fees                      float32                   `json:"fees"`
	OutgoingTransaction       *Transaction              `json:"outgoing_transaction"`
	IncomingTransaction       *Transaction              `json:"incoming_transaction"`
	TransferType              TransferType              `json:"transfer_type"`
	Tags                      *[]*Tag                   `json:"tags,omitempty"`
	RecurrenceTransactionInfo *RecurrentTransactionInfo `json:"recurrence_transaction_info,omitempty"`
	AuditData
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
		return ErrTransferSourceAccountNotActive
	}

	if !t.DestinationAccount.Active {
		return ErrTransferDestinationAccountNotActive
	}

	if t.SourceAccount.Currency.ID == t.DestinationAccount.Currency.ID && t.ExchangeRateMultiplier != 1 {
		return ErrTransferExchangeRateInvalid
	}

	if t.ExchangeRateMultiplier <= 0 {
		return ErrTransferRateMultiplierMustBePositive
	}

	if t.Fees < 0 {
		return ErrTransferFeesMustBePositive
	}

	if t.SourceAccount.CurrentBalance < t.OutgoingTransaction.Amount {
		return ErrTransferSourceAccountNotEnoughBalance
	}

	if errIncoming := t.IncomingTransaction.Validate(); errIncoming != nil {
		return errIncoming
	}
	if errOutgoing := t.OutgoingTransaction.Validate(); errOutgoing != nil {
		return errOutgoing
	}

	return nil
}

func (t *Transfer) Clone() *Transfer {
	now := time.Now()

	return &Transfer{
		ID:                        "",
		SourceAccount:             t.SourceAccount,
		DestinationAccount:        t.DestinationAccount,
		ExchangeRateMultiplier:    t.ExchangeRateMultiplier,
		Fees:                      t.Fees,
		OutgoingTransaction:       t.OutgoingTransaction.Clone(),
		IncomingTransaction:       t.IncomingTransaction.Clone(),
		TransferType:              t.TransferType,
		Tags:                      t.Tags,
		RecurrenceTransactionInfo: t.RecurrenceTransactionInfo,
		AuditData: AuditData{
			Date:      &now,
			CreatedBy: nil, // TODO: Take the actual user doing the rollback after auth implementation

		},
	}
}

func (t *Transfer) Rollback(rollbackMessage string) *Transfer {
	now := time.Now()

	return &Transfer{
		ID:                        t.ID,
		SourceAccount:             t.SourceAccount,
		DestinationAccount:        t.DestinationAccount,
		ExchangeRateMultiplier:    t.ExchangeRateMultiplier,
		Fees:                      float32(0), // Fees handling is done on the use case
		OutgoingTransaction:       t.OutgoingTransaction.Rollback(rollbackMessage),
		IncomingTransaction:       t.IncomingTransaction.Rollback(rollbackMessage),
		TransferType:              t.TransferType,
		Tags:                      t.Tags,
		RecurrenceTransactionInfo: t.RecurrenceTransactionInfo,
		AuditData: AuditData{
			Date:      &now,
			CreatedBy: nil, // TODO: Take the actual user doing the rollback after auth implementation
		},
	}
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
