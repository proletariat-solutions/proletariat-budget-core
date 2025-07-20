package coreentity

type SavingsWithdrawal struct {
	SavingOperation
	Reason string `json:"reason"`
}
