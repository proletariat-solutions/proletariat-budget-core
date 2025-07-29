package coreentity

import (
	"errors"
	"time"

	"proletariat-budget-core/core/domain/misc"
)

type SavingsGoal struct {
	ID                      string            `json:"id"`
	Name                    string            `json:"name"`
	Category                *Category         `json:"category"`
	Description             string            `json:"description"`
	TargetAmount            float32           `json:"target_amount"`
	TargetDate              time.Time         `json:"target_date"`
	InitialAmount           float32           `json:"initial_amount"`
	CurrentAmount           float32           `json:"current_amount"`
	PercentCompleted        float32           `json:"percent_completed"`
	Account                 *Account          `json:"account"`
	Priority                uint              `json:"priority"`
	AutoContribute          bool              `json:"auto_contribute"`
	AutoContributeAmount    *float32          `json:"auto_contribute_amount"`
	AutoContributeFrequency *Frequency        `json:"auto_contribute_frequency"`
	AutoContributeInterval  *uint             `json:"auto_contribute_interval"`
	Tags                    *[]*Tag           `json:"tags,omitempty"`
	Status                  SavingsGoalStatus `json:"status"`
	ProjectedCompletionDate time.Time         `json:"projected_completion_date"`
	AuditData
}

type SavingsGoalStatus string

const (
	SavingsGoalStatusActive    SavingsGoalStatus = "active"
	SavingsGoalStatusCompleted SavingsGoalStatus = "completed"
	SavingsGoalStatusAbandoned SavingsGoalStatus = "abandoned"
	SavingsGoalStatusInactive  SavingsGoalStatus = "inactive"
)

// Savings domain errors
var (
	ErrSavingsGoalNotFound                       = errors.New("savings goal not found")
	ErrSavingsGoalHasActiveWithdrawals           = errors.New("savings goal has active withdrawals and cannot be deleted")
	ErrSavingsGoalNameEmpty                      = errors.New("savings goal name cannot be empty")
	ErrSavingsGoalDescriptionEmpty               = errors.New("savings goal description cannot be empty")
	ErrSavingsGoalTargetAmountMustBePositive     = errors.New("savings goal target amount must be positive")
	ErrSavingsGoalTargetDateMustBeInFuture       = errors.New("savings goal target date must be in the future")
	ErrSavingsGoalInitialAmountMustBeNonNegative = errors.New("savings goal initial amount must be non-negative")
	ErrSavingsGoalCurrentAmountMustBeNonNegative = errors.New("savings goal current amount must be non-negative")
)

func (sg *SavingsGoal) Validate() error {
	if sg.Name == "" {
		return ErrSavingsGoalNameEmpty
	}
	if sg.Description == "" {
		return ErrSavingsGoalDescriptionEmpty
	}
	if sg.TargetAmount <= 0 {
		return ErrSavingsGoalTargetAmountMustBePositive
	}
	if sg.TargetDate.Before(time.Now()) {
		return ErrSavingsGoalTargetDateMustBeInFuture
	}
	if sg.InitialAmount < 0 {
		return ErrSavingsGoalInitialAmountMustBeNonNegative
	}
	if sg.CurrentAmount < 0 {
		return ErrSavingsGoalCurrentAmountMustBeNonNegative
	}

	return nil
}

func (sg *SavingsGoal) CalculatePercentCompleted() {
	sg.PercentCompleted = (sg.CurrentAmount / sg.TargetAmount) * 100
}

type SavingOperation struct {
	ID          string       `json:"id"`
	SavingsGoal *SavingsGoal `json:"savings_goal"`
	Transfer    *Transfer    `json:"transfer,omitempty"`
	Tags        *[]*Tag      `json:"tags,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
}

type ListSavingsTransactionsParams struct {
}

type SavingsTransactionList struct {
	Metadata misc.ListMetadata `json:"metadata"`
}
