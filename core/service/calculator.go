package service

import (
	"time"

	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/port"
)

const DAYSINWEEK = 7
const DAYSINYEAR = 365
const MONTHS = 12
const DAYSINLEAPYEAR = 366

// ScheduleCalculator calculates next execution times
type ScheduleCalculator struct {
	timeProvider port.TimeProvider
}

func NewScheduleCalculator(timeProvider port.TimeProvider) *ScheduleCalculator {
	return &ScheduleCalculator{
		timeProvider: timeProvider,
	}
}

// CalculateNextExecution calculates the next execution time based on config and last execution
func (sc *ScheduleCalculator) CalculateNextExecution(
	pattern *coreentity.RecurrencePattern,
	lastExecution time.Time,
) time.Time {
	if lastExecution.IsZero() {
		return sc.timeProvider.Now()
	}

	switch pattern.Frequency {
	case coreentity.Daily:
		return sc.calculateDailyNext(
			lastExecution,
			pattern.IntervalValue,
		)
	case coreentity.Weekly:
		return sc.calculateWeeklyNext(
			lastExecution,
			pattern.IntervalValue,
		)
	case coreentity.Monthly:
		return sc.calculateWeeklyNext(
			lastExecution,
			pattern.IntervalValue,
		)
	case coreentity.Yearly:
		return sc.calculateYearlyNext(
			lastExecution,
			pattern.IntervalValue,
		)
	case coreentity.NthDayOfTheMonth:
		return sc.calculateEveryNthDayOfMonthNext(
			lastExecution,
			pattern.IntervalValue,
		)
	case coreentity.NthDayOfTheYear:
		return sc.calculateEveryNthDayYearNext(
			lastExecution,
			pattern.IntervalValue,
		)
	default:
		return sc.timeProvider.Now()
	}
}

func (sc *ScheduleCalculator) calculateDailyNext(
	lastExecution time.Time,
	intervalValue int,
) time.Time {
	return lastExecution.AddDate(
		0,
		0,
		intervalValue,
	)
}

func (sc *ScheduleCalculator) calculateWeeklyNext(
	lastExecution time.Time,
	intervalValue int,
) time.Time {
	return lastExecution.AddDate(
		0,
		0,
		DAYSINWEEK*intervalValue,
	)
}

func (sc *ScheduleCalculator) calculateMonthlyNext(
	lastExecution time.Time,
	intervalValue int,
) time.Time {
	return lastExecution.AddDate(
		0,
		intervalValue,
		0,
	)
}

func (sc *ScheduleCalculator) calculateYearlyNext(
	lastExecution time.Time,
	intervalValue int,
) time.Time {
	return lastExecution.AddDate(
		intervalValue,
		0,
		0,
	)
}

func (sc *ScheduleCalculator) calculateEveryNthDayNext(
	lastExecution time.Time,
	intervalValue int,
) time.Time {
	return lastExecution.AddDate(
		0,
		0,
		intervalValue,
	)
}

func (sc *ScheduleCalculator) calculateEveryNthDayYearNext(
	lastExecution time.Time,
	intervalValue int,
) time.Time {
	if intervalValue < 1 {
		intervalValue = 1 // Ensure n is at least 1
	}

	year := lastExecution.Year()

	// Check if we can use the nth day in the current year
	targetDay := intervalValue
	daysInCurrentYear := daysInYear(year)

	// If n is greater than days in current year, use last day (Dec 31st)
	if targetDay > daysInCurrentYear {
		targetDay = daysInCurrentYear
	}

	// Create the target date for current year (January 1st + (targetDay - 1) days)
	targetDate := time.Date(
		year,
		1,
		1,
		0,
		0,
		0,
		0,
		lastExecution.Location(),
	).AddDate(
		0,
		0,
		targetDay-1,
	)

	// If target date is after current date, return it
	if targetDate.After(lastExecution) {
		return targetDate
	}

	// Otherwise, move to next year
	nextYear := year + 1

	// Calculate target day for next year
	targetDay = intervalValue
	daysInNextYear := daysInYear(nextYear)

	// If n is greater than days in next year, use last day
	if targetDay > daysInNextYear {
		targetDay = daysInNextYear
	}

	// Create the target date for next year (January 1st + (targetDay - 1) days)
	return time.Date(
		nextYear,
		1,
		1,
		0,
		0,
		0,
		0,
		lastExecution.Location(),
	).AddDate(
		0,
		0,
		targetDay-1,
	)
}

// daysInYear returns the number of days in a given year (365 or 366 for leap years)
func daysInYear(year int) int {
	// Check if it's a leap year by creating Feb 29th and seeing if it's valid
	feb29 := time.Date(
		year,
		2,
		29,
		0,
		0,
		0,
		0,
		time.UTC,
	)
	if feb29.Month() == 2 && feb29.Day() == 29 {
		return DAYSINLEAPYEAR
	}

	return DAYSINYEAR
}

func (sc *ScheduleCalculator) calculateEveryNthDayOfMonthNext(
	lastExecution time.Time,
	intervalValue int,
) time.Time {
	year := lastExecution.Year()
	month := lastExecution.Month()

	// Check if we can use the nth day in the current month
	targetDay := intervalValue
	daysInCurrentMonth := daysInMonth(
		year,
		month,
	)

	// If n is greater than days in current month, use last day
	if targetDay > daysInCurrentMonth {
		targetDay = daysInCurrentMonth
	}

	// Create the target date for current month
	targetDate := time.Date(
		year,
		month,
		targetDay,
		0,
		0,
		0,
		0,
		lastExecution.Location(),
	)

	// If target date is after current date, return it
	if targetDate.After(lastExecution) {
		return targetDate
	}

	// Otherwise, move to next month
	nextMonth := month + 1
	nextYear := year

	// Handle year rollover
	if nextMonth > MONTHS {
		nextMonth = 1
		nextYear++
	}

	// Calculate target day for next month
	targetDay = intervalValue
	daysInNextMonth := daysInMonth(
		nextYear,
		nextMonth,
	)

	// If n is greater than days in next month, use last day
	if targetDay > daysInNextMonth {
		targetDay = daysInNextMonth
	}

	return time.Date(
		nextYear,
		nextMonth,
		targetDay,
		0,
		0,
		0,
		0,
		lastExecution.Location(),
	)
}

// daysInMonth returns the number of days in a given month and year
func daysInMonth(
	year int,
	month time.Month,
) int {
	// Create first day of next month and subtract one day to get last day of current month
	firstOfNextMonth := time.Date(
		year,
		month+1,
		1,
		0,
		0,
		0,
		0,
		time.UTC,
	)
	lastOfCurrentMonth := firstOfNextMonth.AddDate(
		0,
		0,
		-1,
	)

	return lastOfCurrentMonth.Day()
}

// ShouldExecuteNow checks if job should execute based on current time and next execution
func (sc *ScheduleCalculator) ShouldExecuteNow(nextExecution time.Time) bool {
	return sc.timeProvider.Now().After(nextExecution) || sc.timeProvider.Now().Equal(nextExecution)
}
