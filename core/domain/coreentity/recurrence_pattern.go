package coreentity

import (
	"errors"
	"time"
)

type RecurrencePattern struct {
	ID            string     `json:"id"`
	Frequency     Frequency  `json:"frequency"`
	IntervalValue uint       `json:"interval"`
	Amount        float32    `json:"amount"`
	AccountID     string     `json:"account_id"`
	Account       *Account   `json:"account"`
	Description   string     `json:"description"`
	StartDate     *time.Time `json:"start_date"`
	EndDate       *time.Time `json:"end_date"`
	LastRunDate   *time.Time `json:"last_run_date"`
	TaskType      TaskType   `json:"task_type"`
	Active        bool       `json:"active"`
}

type Frequency string

const (
	Daily            Frequency = "daily"
	Weekly           Frequency = "weekly"
	Monthly          Frequency = "monthly"
	Yearly           Frequency = "yearly"
	NthDayOfTheMonth Frequency = "nth_day_of_the_month" // For these, interval will be the N value (e.g., 2nd day of the month)
	NthDayOfTheWeek  Frequency = "nth_day_of_the_week"
	NthDayOfTheYear  Frequency = "nth_day_of_the_year"
)

type TaskType string

const (
	TaskTypeRecurringIngress             TaskType = "ingress"
	TaskTypeRecurringExpenditure         TaskType = "expenditure"
	TaskTypeRecurringTransfer            TaskType = "transfer"
	TaskTypeRecurringSavingsContribution TaskType = "savings_contribution"
)

var (
	ErrRecurrencePatternNotFound             = errors.New("recurrence pattern not found")
	ErrRecurrencyPatternInvalidFrequency     = errors.New("invalid recurrence frequency")
	ErrRecurrencyPatternAmountMustBePositive = errors.New("amount must be positive")
	ErrRecurrencyPatternEndDateInPast        = errors.New("end date cannot be in the past")
	ErrRecurrencyPatternDescriptionEmpty     = errors.New("description cannot be empty")
	ErrRecurrencyPatternStartDateInPast      = errors.New("start date cannot be in the past")
	ErrRecurrencyPatternStartDateEmpty       = errors.New("start date cannot be empty")
	ErrRecurrencyPatternEndEmpty             = errors.New("end date cannot be empty")
	ErrRecurrencyPatternEndBeforeStartDate   = errors.New("end date cannot be before start date")
	ErrRecurrencyPatternInvalidIntervalValue = errors.New("invalid interval value")
)

func (r *RecurrencePattern) Validate() error {
	if r.Frequency == "" {
		return ErrRecurrencyPatternInvalidFrequency
	}
	if r.Amount < 0 {
		return ErrRecurrencyPatternAmountMustBePositive
	}
	if r.EndDate != nil && r.EndDate.Before(time.Now()) {
		return ErrRecurrencyPatternEndDateInPast
	}
	if r.Description == "" {
		return ErrRecurrencyPatternDescriptionEmpty
	}
	if r.StartDate == nil {
		return ErrRecurrencyPatternStartDateEmpty
	}
	if r.EndDate == nil {
		return ErrRecurrencyPatternEndEmpty
	}
	if r.StartDate.Before(time.Now()) {
		return ErrRecurrencyPatternStartDateInPast
	}
	if r.EndDate.Before(*r.StartDate) {
		return ErrRecurrencyPatternEndBeforeStartDate
	}
	if r.IntervalValue == 0 {
		return ErrRecurrencyPatternInvalidIntervalValue
	}
	if r.Frequency == NthDayOfTheMonth && r.IntervalValue > 31 {
		// For months with less than 31 days, days between 29 & 31 will default to the last day of the month
		return ErrRecurrencyPatternInvalidIntervalValue
	}
	if r.Frequency == NthDayOfTheWeek && r.IntervalValue > 7 {
		return ErrRecurrencyPatternInvalidIntervalValue
	}
	if r.Frequency == NthDayOfTheYear && r.IntervalValue > 365 {
		return ErrRecurrencyPatternInvalidIntervalValue
	}

	return nil
}

func IsRecurrencyValidationError(err error) bool {
	return errors.Is(
		err,
		ErrRecurrencyPatternInvalidFrequency,
	) || errors.Is(
		err,
		ErrRecurrencyPatternAmountMustBePositive,
	) || errors.Is(
		err,
		ErrRecurrencyPatternEndDateInPast,
	) || errors.Is(
		err,
		ErrRecurrencyPatternDescriptionEmpty,
	) || errors.Is(
		err,
		ErrRecurrencyPatternStartDateInPast,
	) || errors.Is(
		err,
		ErrRecurrencyPatternStartDateEmpty,
	) || errors.Is(
		err,
		ErrRecurrencyPatternEndEmpty,
	) || errors.Is(
		err,
		ErrRecurrencyPatternEndBeforeStartDate,
	) || errors.Is(
		err,
		ErrRecurrencyPatternInvalidIntervalValue,
	)
}
