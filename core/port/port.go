package port

type Ports struct {
	Account            *Account
	Auth               *Auth
	Category           *Category
	Expenditure        *Expenditure
	HouseholdMembers   *HouseholdMember
	Ingress            *Ingress
	Notification       *Notification
	RecurrencePattern  *RecurrencePattern
	SavingGoal         *SavingsGoal
	Tags               *Tags
	TimeProvider       *TimeProvider
	TickerProvider     *TickerProvider
	Transaction        *Transaction
	TransactionManager *TransactionManager
	Transfer           *Transfer
}
