package coreentity

type SavingsContribution struct {
	SavingOperation
	FromRecurrencePatternID *string `json:"from_recurrence_pattern,omitempty"`
}
