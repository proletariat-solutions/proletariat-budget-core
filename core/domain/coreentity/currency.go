package coreentity

import "errors"

type Currency struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

var (
	ErrCurrencyIDRequired     = errors.New("currency ID is required")
	ErrCurrencyNameRequired   = errors.New("currency name is required")
	ErrCurrencySymbolRequired = errors.New("currency symbol is required")
)

func (c *Currency) Validate() error {
	if c.ID == "" {
		return ErrCurrencyIDRequired
	}
	if c.Name == "" {
		return ErrCurrencyNameRequired
	}
	if c.Symbol == "" {
		return ErrCurrencySymbolRequired
	}

	return nil
}
