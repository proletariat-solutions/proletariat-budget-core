package coreentity

import (
	"errors"
	"time"

	"proletariat-budget-core/core/domain/misc"
)

type Expenditure struct {
	ID                        string                    `json:"id"`
	Category                  *Category                 `json:"category"`
	Declared                  bool                      `json:"declared"`
	Planned                   bool                      `json:"planned"`
	Transaction               *Transaction              `json:"transaction,omitempty"`
	Tags                      *[]*Tag                   `json:"tags,omitempty"`
	RecurrenceTransactionInfo *RecurrentTransactionInfo `json:"recurrence_transaction_info,omitempty"`
	AuditData
}

var (
	ErrExpenditureNotFound       = errors.New("expenditure not found")
	ErrExpenditureCategoryNil    = errors.New("category is required")
	ErrExpenditureTransactionNil = errors.New("transaction is required")
)

func (e *Expenditure) Validate() error {
	if e.Category == nil {
		return ErrExpenditureCategoryNil
	}
	if e.Transaction == nil {
		return ErrExpenditureTransactionNil
	}

	return e.Transaction.Validate()
}

type ExpenditureList struct {
	Expenditures []Expenditure     `json:"expenditures"`
	Metadata     misc.ListMetadata `json:"metadata"`
}

type ExpenditureListParams struct {
	CategoryID  *string    `json:"categoryid"`
	StartDate   *time.Time `json:"date_from"`
	EndDate     *time.Time `json:"date_to"`
	Declared    *bool      `json:"declared"`
	Planned     *bool      `json:"planned"`
	Currency    *string    `json:"currency"`
	Description *string    `json:"description"`
	AccountID   *string    `json:"account_id"`
	Tags        *[]string  `json:"tags"`
	misc.ListParams
}
