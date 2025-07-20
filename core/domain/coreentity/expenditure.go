package coreentity

import (
	"errors"
	"time"

	"proletariat-budget-core/core/domain/misc"
)

type Expenditure struct {
	ID                      string       `json:"id"`
	Category                *Category    `json:"category"`
	Declared                bool         `json:"declared"`
	Planned                 bool         `json:"planned"`
	Transaction             *Transaction `json:"transaction,omitempty"`
	Tags                    *[]*Tag      `json:"tags,omitempty"`
	Date                    time.Time    `json:"date"`
	FromRecurrencePatternID *string      `json:"from_recurrence_pattern,omitempty"`
}

var (
	ErrExpenditureNotFound = errors.New("expenditure not found")
)

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
