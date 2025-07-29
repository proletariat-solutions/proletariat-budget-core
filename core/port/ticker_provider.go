package port

import "time"

//go:generate mockgen -source=ticker_provider.go -destination=../../test/mocks/mock_ticker_provider.go -package mocks
type TickerProvider interface {
	NewTicker(duration time.Duration) *time.Ticker
}
