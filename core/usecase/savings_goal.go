package usecase

import "proletariat-budget-core/core/port"

type SavingGoal struct {
	accountRepo     port.Account
	savingsGoalRepo port.SavingsGoal
}
