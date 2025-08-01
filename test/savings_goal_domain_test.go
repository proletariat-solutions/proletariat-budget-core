package test

import (
	"github.com/stretchr/testify/assert"
	"proletariat-budget-core/core/domain/coreentity"
	"testing"
	"time"
)

func TestSavingsGoal_Validate_ShouldReturnNilWhenAllSavingsGoalFieldsAreValid(t *testing.T) {
	// Arrange
	futureDate := time.Now().Add(24 * time.Hour)
	savingsGoal := &coreentity.SavingsGoal{
		Name:          "Emergency Fund",
		Description:   "Fund for emergencies",
		TargetAmount:  10000.0,
		TargetDate:    futureDate,
		InitialAmount: 1000.0,
		CurrentAmount: 2500.0,
	}

	// Act
	err := savingsGoal.Validate()

	// Assert
	assert.NoError(
		t,
		err,
	)
}

func TestSavingsGoal_Validate_ShouldReturnErrSavingsGoalNameEmptyWhenNameIsAnEmptyString(t *testing.T) {
	// Arrange
	futureDate := time.Now().Add(24 * time.Hour)
	savingsGoal := &coreentity.SavingsGoal{
		Name:          "",
		Description:   "Fund for emergencies",
		TargetAmount:  10000.0,
		TargetDate:    futureDate,
		InitialAmount: 1000.0,
		CurrentAmount: 2500.0,
	}

	// Act
	err := savingsGoal.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrSavingsGoalNameEmpty,
		err,
	)
}

func TestSavingsGoal_Validate_ShouldReturnErrSavingsGoalDescriptionEmptyWhenDescriptionIsAnEmptyString(t *testing.T) {
	// Arrange
	futureDate := time.Now().Add(24 * time.Hour)
	savingsGoal := &coreentity.SavingsGoal{
		Name:          "Emergency Fund",
		Description:   "",
		TargetAmount:  10000.0,
		TargetDate:    futureDate,
		InitialAmount: 1000.0,
		CurrentAmount: 2500.0,
	}

	// Act
	err := savingsGoal.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrSavingsGoalDescriptionEmpty,
		err,
	)
}
func TestSavingsGoal_Validate_ShouldReturnErrSavingsGoalTargetAmountMustBePositiveWhenTargetAmountIsZero(t *testing.T) {
	// Arrange
	futureDate := time.Now().Add(24 * time.Hour)
	savingsGoal := &coreentity.SavingsGoal{
		Name:          "Emergency Fund",
		Description:   "Fund for emergencies",
		TargetAmount:  0.0,
		TargetDate:    futureDate,
		InitialAmount: 1000.0,
		CurrentAmount: 2500.0,
	}

	// Act
	err := savingsGoal.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrSavingsGoalTargetAmountMustBePositive,
		err,
	)
}

func TestSavingsGoal_Validate_ShouldReturnErrSavingsGoalTargetAmountMustBePositiveWhenTargetAmountIsNegative(t *testing.T) {
	// Arrange
	futureDate := time.Now().Add(24 * time.Hour)
	savingsGoal := &coreentity.SavingsGoal{
		Name:          "Emergency Fund",
		Description:   "Fund for emergencies",
		TargetAmount:  -1000.0,
		TargetDate:    futureDate,
		InitialAmount: 1000.0,
		CurrentAmount: 2500.0,
	}

	// Act
	err := savingsGoal.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrSavingsGoalTargetAmountMustBePositive,
		err,
	)
}

func TestSavingsGoal_Validate_ShouldReturnErrSavingsGoalTargetDateMustBeInFutureWhenTargetDateIsInThePast(t *testing.T) {
	// Arrange
	pastDate := time.Now().Add(-24 * time.Hour)
	savingsGoal := &coreentity.SavingsGoal{
		Name:          "Emergency Fund",
		Description:   "Fund for emergencies",
		TargetAmount:  10000.0,
		TargetDate:    pastDate,
		InitialAmount: 1000.0,
		CurrentAmount: 2500.0,
	}

	// Act
	err := savingsGoal.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrSavingsGoalTargetDateMustBeInFuture,
		err,
	)
}

func TestSavingsGoal_Validate_ShouldReturnErrSavingsGoalTargetDateMustBeInFutureWhenTargetDateIsExactlyCurrentTime(t *testing.T) {
	// Arrange
	currentTime := time.Now()
	savingsGoal := &coreentity.SavingsGoal{
		Name:          "Emergency Fund",
		Description:   "Fund for emergencies",
		TargetAmount:  10000.0,
		TargetDate:    currentTime,
		InitialAmount: 1000.0,
		CurrentAmount: 2500.0,
	}

	// Act
	err := savingsGoal.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrSavingsGoalTargetDateMustBeInFuture,
		err,
	)
}

func TestSavingsGoal_Validate_ShouldReturnErrSavingsGoalInitialAmountMustBeNonNegativeWhenInitialAmountIsNegative(t *testing.T) {
	// Arrange
	futureDate := time.Now().Add(24 * time.Hour)
	savingsGoal := &coreentity.SavingsGoal{
		Name:          "Emergency Fund",
		Description:   "Fund for emergencies",
		TargetAmount:  10000.0,
		TargetDate:    futureDate,
		InitialAmount: -500.0,
		CurrentAmount: 2500.0,
	}

	// Act
	err := savingsGoal.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrSavingsGoalInitialAmountMustBeNonNegative,
		err,
	)
}

func TestSavingsGoal_Validate_ShouldReturnNilWhenInitialAmountAndCurrentAmountAreZero(t *testing.T) {
	// Arrange
	futureDate := time.Now().Add(24 * time.Hour)
	savingsGoal := &coreentity.SavingsGoal{
		Name:          "Emergency Fund",
		Description:   "Fund for emergencies",
		TargetAmount:  10000.0,
		TargetDate:    futureDate,
		InitialAmount: 0.0,
		CurrentAmount: 0.0,
	}

	// Act
	err := savingsGoal.Validate()

	// Assert
	assert.NoError(
		t,
		err,
	)
}

func TestSavingsGoal_Validate_ShouldReturnErrSavingsGoalCurrentAmountMustBeNonNegativeWhenCurrentAmountIsNegative(t *testing.T) {
	// Arrange
	futureDate := time.Now().Add(24 * time.Hour)
	savingsGoal := &coreentity.SavingsGoal{
		Name:          "Emergency Fund",
		Description:   "Fund for emergencies",
		TargetAmount:  10000.0,
		TargetDate:    futureDate,
		InitialAmount: 1000.0,
		CurrentAmount: -500.0,
	}

	// Act
	err := savingsGoal.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrSavingsGoalCurrentAmountMustBeNonNegative,
		err,
	)
}
