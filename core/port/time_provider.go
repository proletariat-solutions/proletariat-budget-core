package port

import "time"

// TimeProvider allows mocking time for tests
//
//go:generate mockgen -source=time_provider.go -destination=../../test/mocks/mock_time_provider.go -package mocks
type TimeProvider interface {
	Now() time.Time
	Sleep(duration time.Duration)
	After(duration time.Duration) <-chan time.Time
}
