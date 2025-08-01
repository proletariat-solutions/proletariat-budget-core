package coreentity

import (
	"errors"
	"time"
)

type SavingOperation struct {
	ID          string       `json:"id"`
	SavingsGoal *SavingsGoal `json:"savings_goal"`
	Transfer    *Transfer    `json:"transfer,omitempty"`
	Tags        *[]*Tag      `json:"tags,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
}

var (
	ErrSavingOperationEmptyID                          = errors.New("saving operation ID cannot be empty")
	ErrSavingOperationNilSavingsGoal                   = errors.New("savings goal cannot be nil")
	ErrSavingOperationNilTransfer                      = errors.New("transfer cannot be nil")
	ErrSavingOperationTransferAccountDifferentCurrency = errors.New("transfer account currency must match savings goal currency")
)

func (s SavingOperation) Clone() *SavingOperation {
	// TODO implement me
	panic("implement me")
}

func (s SavingOperation) Rollback(rollbackMessage string) *SavingOperation {
	// TODO implement me
	panic("implement me")
}

func (s SavingOperation) Validate() error {
	// TODO implement me
	panic("implement me")
}
