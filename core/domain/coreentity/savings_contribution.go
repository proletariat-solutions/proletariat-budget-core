package coreentity

type SavingsContribution struct {
	SavingOperation
	RecurrenceTransactionInfo *RecurrentTransactionInfo `json:"recurrence_transaction_info,omitempty"`
	AuditData
}

func (s SavingsContribution) Clone() *SavingsContribution {
	//TODO implement me
	panic("implement me")
}

func (s SavingsContribution) Rollback(rollbackMessage string) *SavingsContribution {
	//TODO implement me
	panic("implement me")
}

func (s SavingsContribution) Validate() error {
	//TODO implement me
	panic("implement me")
}
