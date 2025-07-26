package provider

import "time"

// RealTimeProvider implements TimeProvider using real time
type RealTimeProvider struct{}

func NewRealTimeProvider() *RealTimeProvider {
	return &RealTimeProvider{}
}

func (r *RealTimeProvider) Now() time.Time {
	return time.Now()
}

func (r *RealTimeProvider) Sleep(duration time.Duration) {
	time.Sleep(duration)
}

func (r *RealTimeProvider) After(duration time.Duration) <-chan time.Time {
	return time.After(duration)
}
