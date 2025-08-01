package test

import (
	"github.com/stretchr/testify/assert"
	"proletariat-budget-core/core/domain/coreentity"
	"testing"
	"time"
)

func TestRecurrencePattern_Validate_ShouldReturnErrRecurrencyPatternInvalidIntervalValueWhenIntervalValueIsNegative(t *testing.T) {
	// Arrange
	now := time.Now()
	startDate := now.Add(24 * time.Hour)
	endDate := startDate.Add(48 * time.Hour)

	recurrencePattern := coreentity.RecurrencePattern{
		Frequency:     coreentity.NthDay,
		IntervalValue: -1,
		Amount:        100.0,
		Description:   "Test Description",
		StartDate:     &startDate,
		EndDate:       &endDate,
	}

	// Act
	err := recurrencePattern.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrRecurrencyPatternInvalidIntervalValue,
		err,
	)
}

func TestRecurrencePattern_Validate_ReturnsErrRecurrencyPatternInvalidIntervalValueWhenIntervalValueIsZero(t *testing.T) {
	// Arrange
	now := time.Now()
	startDate := now.Add(24 * time.Hour)
	endDate := startDate.Add(7 * 24 * time.Hour)

	recurrencePattern := &coreentity.RecurrencePattern{
		Frequency:     coreentity.NthDay,
		IntervalValue: 0,
		Amount:        100.0,
		Description:   "Test recurrence",
		StartDate:     &startDate,
		EndDate:       &endDate,
	}

	// Act
	err := recurrencePattern.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrRecurrencyPatternInvalidIntervalValue,
		err,
	)
}
func TestRecurrencePattern_Validate_ShouldReturnErrRecurrencyPatternEndDateInPast_WhenEndDateIsBeforeCurrentTime(t *testing.T) {
	// Arrange
	pastDate := time.Now().Add(-24 * time.Hour)
	futureDate := time.Now().Add(24 * time.Hour)

	recurrencePattern := &coreentity.RecurrencePattern{
		Frequency:     coreentity.NthDay,
		Amount:        100.0,
		EndDate:       &pastDate,
		Description:   "Test description",
		StartDate:     &futureDate,
		IntervalValue: 1,
	}

	// Act
	err := recurrencePattern.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrRecurrencyPatternEndDateInPast,
		err,
	)
}
func TestRecurrencePattern_Validate_ShouldReturnErrRecurrencyPatternInvalidFrequency_WhenFrequencyIsEmptyString(t *testing.T) {
	// Arrange
	now := time.Now()
	startDate := now.Add(24 * time.Hour)
	endDate := now.Add(48 * time.Hour)

	recurrencePattern := &coreentity.RecurrencePattern{
		Frequency:     "",
		IntervalValue: 1,
		Amount:        100.0,
		AccountID:     "account-123",
		Description:   "Test description",
		StartDate:     &startDate,
		EndDate:       &endDate,
		TaskType:      coreentity.TaskTypeRecurringExpenditure,
		Active:        true,
	}

	// Act
	err := recurrencePattern.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrRecurrencyPatternInvalidFrequency,
		err,
	)
}
func TestRecurrencePattern_Validate_ShouldReturnErrRecurrencyPatternAmountMustBePositive_WhenAmountIsNegative(t *testing.T) {
	// Arrange
	startDate := time.Now().Add(24 * time.Hour)
	endDate := time.Now().Add(48 * time.Hour)

	recurrencePattern := coreentity.RecurrencePattern{
		Frequency:     coreentity.NthDay,
		IntervalValue: 1,
		Amount:        -100.0,
		Description:   "Test description",
		StartDate:     &startDate,
		EndDate:       &endDate,
	}

	// Act
	err := recurrencePattern.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrRecurrencyPatternAmountMustBePositive,
		err,
	)
}
func TestRecurrencePattern_Validate_ShouldReturnErrRecurrencyPatternDescriptionEmpty_WhenDescriptionIsEmptyString(t *testing.T) {
	// Arrange
	startDate := time.Now().Add(24 * time.Hour)
	endDate := time.Now().Add(48 * time.Hour)
	recurrencePattern := &coreentity.RecurrencePattern{
		Frequency:     coreentity.NthDay,
		IntervalValue: 1,
		Amount:        100.0,
		Description:   "",
		StartDate:     &startDate,
		EndDate:       &endDate,
	}

	// Act
	err := recurrencePattern.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrRecurrencyPatternDescriptionEmpty,
		err,
	)
}
func TestRecurrencePattern_Validate_ShouldReturnErrRecurrencyPatternStartDateEmpty_WhenStartDateIsNil(t *testing.T) {
	// Arrange
	futureDate := time.Now().Add(24 * time.Hour)
	pattern := &coreentity.RecurrencePattern{
		Frequency:     coreentity.NthDay,
		IntervalValue: 1,
		Amount:        100.0,
		Description:   "Test pattern",
		StartDate:     nil,
		EndDate:       &futureDate,
	}

	// Act
	err := pattern.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrRecurrencyPatternStartDateEmpty,
		err,
	)
}
func TestRecurrencePattern_Validate_ReturnsErrRecurrencyPatternEndEmpty_WhenEndDateIsNil(t *testing.T) {
	// Arrange
	futureDate := time.Now().Add(24 * time.Hour)
	recurrencePattern := &coreentity.RecurrencePattern{
		Frequency:     coreentity.NthDay,
		Amount:        100.0,
		Description:   "Test description",
		StartDate:     &futureDate,
		EndDate:       nil,
		IntervalValue: 1,
	}

	// Act
	err := recurrencePattern.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrRecurrencyPatternEndEmpty,
		err,
	)
}
func TestRecurrencePattern_Validate_ReturnsErrRecurrencyPatternStartDateInPast(t *testing.T) {
	// Arrange
	pastDate := time.Now().Add(-24 * time.Hour)
	futureDate := time.Now().Add(24 * time.Hour)

	recurrencePattern := coreentity.RecurrencePattern{
		ID:            "pattern-123",
		Frequency:     coreentity.NthDay,
		IntervalValue: 1,
		Amount:        100.0,
		AccountID:     "account-123",
		Description:   "Test pattern",
		StartDate:     &pastDate,
		EndDate:       &futureDate,
		TaskType:      coreentity.TaskTypeRecurringExpenditure,
		Active:        true,
	}

	// Act
	err := recurrencePattern.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrRecurrencyPatternStartDateInPast,
		err,
	)
}
func TestRecurrencePattern_Validate_ShouldReturnErrRecurrencyPatternEndBeforeStartDate_WhenEndDateIsBeforeStartDate(t *testing.T) {
	// Arrange
	startDate := time.Now().Add(24 * time.Hour)
	endDate := startDate.Add(-12 * time.Hour) // End date before start date

	recurrencePattern := coreentity.RecurrencePattern{
		Frequency:     coreentity.NthDay,
		Amount:        100.0,
		Description:   "Test description",
		StartDate:     &startDate,
		EndDate:       &endDate,
		IntervalValue: 1,
	}

	// Act
	err := recurrencePattern.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrRecurrencyPatternEndBeforeStartDate,
		err,
	)
}
func TestRecurrencePattern_Validate_ReturnsErrRecurrencyPatternInvalidIntervalValueWhenFrequencyIsNthDayOfTheMonthAndIntervalValueIsBiggerThan31(t *testing.T) {
	// Arrange
	startDate := time.Now().Add(24 * time.Hour)
	endDate := time.Now().Add(48 * time.Hour)

	recurrencePattern := &coreentity.RecurrencePattern{
		Frequency:     coreentity.NthDayOfTheMonth,
		IntervalValue: 32,
		Amount:        100.0,
		Description:   "Test description",
		StartDate:     &startDate,
		EndDate:       &endDate,
	}

	// Act
	err := recurrencePattern.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrRecurrencyPatternInvalidIntervalValue,
		err,
	)
}

func TestRecurrencePattern_Validate_ReturnsErrRecurrencyPatternInvalidIntervalValueWhenFrequencyIsNthWeekAndIntervalValueIsBiggerThan52(t *testing.T) {
	// Arrange
	startDate := time.Now().Add(24 * time.Hour)
	endDate := time.Now().Add(48 * time.Hour)

	recurrencePattern := &coreentity.RecurrencePattern{
		Frequency:     coreentity.NthWeek,
		IntervalValue: 53,
		Amount:        100.0,
		Description:   "Test description",
		StartDate:     &startDate,
		EndDate:       &endDate,
	}

	// Act
	err := recurrencePattern.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrRecurrencyPatternInvalidIntervalValue,
		err,
	)
}

func TestRecurrencePattern_Validate_ReturnsErrRecurrencyPatternInvalidIntervalValueWhenFrequencyIsNthDayOfTheYearAndIntervalValueIsBiggerThan366(t *testing.T) {
	// Arrange
	startDate := time.Now().Add(24 * time.Hour)
	endDate := time.Now().Add(48 * time.Hour)

	recurrencePattern := &coreentity.RecurrencePattern{
		Frequency:     coreentity.NthDayOfTheYear,
		IntervalValue: 367,
		Amount:        100.0,
		Description:   "Test description",
		StartDate:     &startDate,
		EndDate:       &endDate,
	}

	// Act
	err := recurrencePattern.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrRecurrencyPatternInvalidIntervalValue,
		err,
	)
}

func TestRecurrencePattern_Validate_ReturnsNilWhenRecurrencePatternPassesValidation(t *testing.T) {
	// Arrange
	startDate := time.Now().Add(24 * time.Hour)
	endDate := time.Now().Add(48 * time.Hour)

	recurrencePattern := &coreentity.RecurrencePattern{
		Frequency:     coreentity.NthDay,
		IntervalValue: 1,
		Amount:        100.0,
		Description:   "Test description",
		StartDate:     &startDate,
		EndDate:       &endDate,
	}

	// Act
	err := recurrencePattern.Validate()

	// Assert
	assert.NoError(
		t,
		err,
	)
}
