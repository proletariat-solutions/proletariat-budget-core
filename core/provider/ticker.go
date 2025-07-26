package provider

import (
	"time"
)

type TickerProvider struct {
}

func (t TickerProvider) NewTicker(duration time.Duration) *time.Ticker {
	return time.NewTicker(duration)
}
