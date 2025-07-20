package config

import "proletariat-budget-core/core/domain/coreentity"

type AppConfig struct {
	DisplayConfig
	WorkersConfig
	FirstStart bool `json:"first_start"` // App initial launch, to determine if we show a wizard or not
}

type DisplayConfig struct {
	MainCurrency      string `json:"main_currency"`
	CurrencyFormat    string `json:"currency_format"`
	DateFormat        string `json:"date_format"`
	TimeFormat        string `json:"time_format"`
	ShowRelativeDates bool   `json:"show_relative_dates"`
}

/*
WorkersConfig defines config for worker tasks that are scheduled to run.
Ex: Recurring ingresses, automatic contributions to transfers, etc.
*/
type WorkersConfig struct {
	WorkerFrequency      coreentity.Frequency `json:"worker_frequency"`
	WorkerInterval       uint                 `json:"worker_interval"`
	WorkerTimeOfDayStart string               `json:"worker_time_of_day_start"`
}

func DefaultAppConfig() AppConfig {
	return AppConfig{
		DisplayConfig: DisplayConfig{
			MainCurrency:      "USD",
			CurrencyFormat:    "${{amount}}",
			DateFormat:        "MM/DD/YYYY",
			TimeFormat:        "hh:mm:ss a",
			ShowRelativeDates: true,
		},
		WorkersConfig: WorkersConfig{ // Running every day, midnight
			WorkerFrequency:      coreentity.Daily,
			WorkerInterval:       1,
			WorkerTimeOfDayStart: "00:00:00",
		},
		FirstStart: true,
	}
}
