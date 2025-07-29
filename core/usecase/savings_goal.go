package usecase

import "proletariat-budget-core/core/port"

type SavingsGoal struct {
	accountRepo     port.Account
	savingsGoalRepo port.SavingsGoal
}
