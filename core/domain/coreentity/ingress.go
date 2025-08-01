package coreentity

import (
	"errors"
	"time"

	"proletariat-budget-core/core/domain/misc"
)

type Ingress struct {
	ID                        string                    `json:"id"`
	Category                  *Category                 `json:"category"`
	Transaction               *Transaction              `json:"transaction,omitempty"`
	Tags                      *[]*Tag                   `json:"tags,omitempty"`
	RecurrenceTransactionInfo *RecurrentTransactionInfo `json:"recurrence_transaction_info,omitempty"`
	AuditData
}

// Ingress domain errors
var (
	ErrIngressNotFound             = errors.New("ingress not found")
	ErrIngressAmountMustBePositive = errors.New("amount must be positive")
	ErrIngressCategoryRequired     = errors.New("category is required")
	ErrTransactionRequired         = errors.New("transaction is required")
	ErrIngressDateRequired         = errors.New("date is required")
	ErrIngressAccountRequired      = errors.New("account is required")
)

func (i *Ingress) Validate() error {
	if i.Transaction == nil {
		return ErrTransactionRequired
	}
	if i.Category == nil {
		return ErrIngressCategoryRequired
	}
	if i.Date == nil {
		return ErrIngressDateRequired
	}

	return i.Transaction.Validate()
}

func (i *Ingress) Clone() *Ingress {
	txClone := i.Transaction.Clone()

	return &Ingress{
		ID:          "",
		Category:    i.Category,
		Transaction: txClone,
		Tags:        i.Tags,
		RecurrenceTransactionInfo: &RecurrentTransactionInfo{
			FromRecurrencePatternID: i.RecurrenceTransactionInfo.FromRecurrencePatternID,
			IsTemplate:              false,
		},
		AuditData: i.AuditData,
	}
}

func (i *Ingress) Rollback(rollbackMessage string) *Ingress {
	txRollback := i.Transaction.Rollback(rollbackMessage)
	now := time.Now()

	return &Ingress{
		ID:          i.ID,
		Category:    i.Category,
		Transaction: txRollback,
		Tags:        i.Tags,
		RecurrenceTransactionInfo: &RecurrentTransactionInfo{
			FromRecurrencePatternID: i.RecurrenceTransactionInfo.FromRecurrencePatternID,
			IsTemplate:              false,
		},
		AuditData: AuditData{
			Date:      &now,
			CreatedBy: i.CreatedBy, // TODO: Take the actual user doing the rollback after auth implementation
		},
	}
}

type IngressList struct {
	Ingresses []Ingress         `json:"ingresses"`
	Metadata  misc.ListMetadata `json:"metadata"`
}

type IngressListParams struct {
	CategoryID  *string    `form:"category,omitempty" json:"category,omitempty"`
	Source      *string    `form:"source,omitempty" json:"source,omitempty"`
	TagIDs      *[]string  `form:"tags,omitempty" json:"tags,omitempty"`
	StartDate   *time.Time `form:"startDate,omitempty" json:"startDate,omitempty"`
	EndDate     *time.Time `form:"endDate,omitempty" json:"endDate,omitempty"`
	IsRecurring *bool      `form:"isRecurring,omitempty" json:"isRecurring,omitempty"`
	Currency    *string    `form:"currency,omitempty" json:"currency,omitempty"`
	misc.ListParams
}
