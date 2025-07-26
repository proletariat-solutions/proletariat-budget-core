package coreentity

import (
	"errors"
	"time"
)

// JobState represents the current state of job execution
type JobState struct {
	ID               string    `json:"id"`
	LastExecution    time.Time `json:"last_execution"`
	NextExecution    time.Time `json:"next_execution"`
	ExecutionCount   int       `json:"execution_count"`
	LastError        string    `json:"last_error,omitempty"`
	RecoveryAttempts int       `json:"recovery_attempts"`
	Status           JobStatus `json:"status"` // "running", "idle", "error", "recovering"
}

type JobStatus string

const (
	StatusRunning    JobStatus = "running"
	StatusIdle       JobStatus = "idle"
	StatusError      JobStatus = "error"
	StatusRecovering JobStatus = "recovering"
)

var (
	ErrGetJobConfig = errors.New("failed to get job configuration")
	ErrGetTask      = errors.New("failed to get task")
)
