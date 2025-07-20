package coreentity

import (
	"errors"
	"time"
)

type ExchangeRate struct {
	ID                 string    `json:"id"`
	FromCurrency       *Currency `json:"from_currency"`
	ToCurrency         *Currency `json:"to_currency"`
	SellRate           float32   `json:"sell_rate"`
	BuyRate            float32   `json:"buy_rate"`
	Date               time.Time `json:"date"`
	FromManualExchange bool      `json:"from_manual_exchange"`
}

var (
	ErrExchangeRateNotFound               = errors.New("exchange rate not found")
	ErrExchangeRateSellRateMustBePositive = errors.New("sell rate must be positive")
	ErrExchangeRateBuyRateMustBePositive  = errors.New("buy rate must be positive")
	ErrExchangeRateFromCurrencyNil        = errors.New("from currency cannot be nil")
	ErrExchangeRateToCurrencyNil          = errors.New("to currency cannot be nil")
	ErrExchangeRateSameCurrency           = errors.New("from and to currencies cannot be the same")
)

func (e ExchangeRate) Validate() error {
	if e.FromCurrency == nil {
		return ErrExchangeRateFromCurrencyNil
	}
	if e.ToCurrency == nil {
		return ErrExchangeRateToCurrencyNil
	}
	if e.FromCurrency.ID == e.ToCurrency.ID {
		return ErrExchangeRateSameCurrency
	}
	if e.SellRate <= 0 {
		return ErrExchangeRateSellRateMustBePositive
	}
	if e.BuyRate <= 0 {
		return ErrExchangeRateBuyRateMustBePositive
	}

	return nil
}
