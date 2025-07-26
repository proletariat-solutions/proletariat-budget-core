package port

import "time"

type TickerProvider interface {
	NewTicker(duration time.Duration) *time.Ticker
}
