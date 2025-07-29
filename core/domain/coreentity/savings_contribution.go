package coreentity

type SavingsContribution struct {
	SavingOperation
	RecurrenceTransactionInfo *RecurrentTransactionInfo `json:"recurrence_transaction_info,omitempty"`
	AuditData
}
