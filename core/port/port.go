package port

type Ports struct {
	Account            *Account
	Auth               *Auth
	Category           *Category
	Expenditure        *Expenditure
	HouseholdMembers   *HouseholdMember
	Ingress            *Ingress
	SavingGoal         *SavingsGoal
	Tags               *Tags
	Transaction        *Transaction
	TransactionManager *TransactionManager
	Transfer           *Transfer
	RecurrencePattern  *RecurrencePattern
}
