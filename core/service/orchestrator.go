package service

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/port"
)

const (
	MaxRecoveryAttempts = 3
	RecoveryDelay       = 5 * time.Minute
	CheckInterval       = 24 * time.Hour
)

// Orchestrator manages job scheduling and execution
type Orchestrator struct {
	recurrencePatternRepo port.RecurrencePattern
	processor             port.TaskProcessor
	calculator            *ScheduleCalculator
	timeProvider          port.TimeProvider
	tickerProvider        port.TickerProvider
	stopChan              chan struct{}
	doneChan              chan struct{}
}

func NewOrchestrator(
	repository port.RecurrencePattern,
	processor port.TaskProcessor,
	timeProvider port.TimeProvider,
) *Orchestrator {
	return &Orchestrator{
		recurrencePatternRepo: repository,
		processor:             processor,
		calculator:            NewScheduleCalculator(timeProvider),
		timeProvider:          timeProvider,
		stopChan:              make(chan struct{}),
		doneChan:              make(chan struct{}),
	}
}

// Start begins the job scheduling loop
func (js *Orchestrator) Start(ctx context.Context) error {
	log.Info().Msg("Starting job orchestrator...")

	// Attempt recovery on startup
	js.recoverFromFailure(ctx)

	js.schedulingLoop(ctx)

	return nil
}

// Stop gracefully stops the orchestrator
func (js *Orchestrator) Stop() {
	log.Info().Msg("Stopping job orchestrator...")
	close(js.stopChan)
	<-js.doneChan
	log.Info().Msg("Job orchestrator stopped")
}

func (js *Orchestrator) schedulingLoop(ctx context.Context) {
	defer close(js.doneChan)

	ticker := js.tickerProvider.NewTicker(CheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-js.stopChan:
			return
		case <-ticker.C:
			go js.checkAndExecuteJobs(ctx)
		}
	}
}

func (js *Orchestrator) checkAndExecuteJobs(ctx context.Context) {
	// Get job configuration
	patterns, err := js.recurrencePatternRepo.GetTodayJobs(ctx)
	if err != nil {
		return
	}
	for _, pattern := range patterns {
		if !pattern.Active ||
			!js.calculator.ShouldExecuteNow(pattern.JobState.NextExecution) ||
			pattern.JobState.Status == coreentity.StatusRunning {
			continue
		}

		// Get current job state
		if pattern.JobState == nil {
			// Initialize state if not found
			pattern.JobState = &coreentity.JobState{
				Status:         coreentity.StatusIdle,
				NextExecution:  js.timeProvider.Now(),
				ExecutionCount: 0,
			}
		}
		errExec := js.executeJob(
			ctx,
			&pattern,
		)

		if errExec != nil {
			log.Err(errExec).Msgf(
				"Failed to execute job %s",
				pattern.ID,
			)
		}
	}
}

func (js *Orchestrator) executeJob(
	ctx context.Context,
	pattern *coreentity.RecurrencePattern,
) error {
	// Update state to running
	pattern.JobState.Status = coreentity.StatusRunning
	pattern.JobState.LastError = ""
	if err := js.recurrencePatternRepo.Update(
		ctx,
		*pattern,
	); err != nil {
		return fmt.Errorf(
			"failed to save running state: %w",
			err,
		)
	}

	// Execute the job
	err := js.processor.ProcessTask(
		ctx,
		*pattern,
	)

	// Update state based on result
	now := js.timeProvider.Now()
	if err != nil {
		pattern.JobState.Status = coreentity.StatusError
		pattern.JobState.LastError = err.Error()
		pattern.JobState.RecoveryAttempts++
		log.Printf(
			"Job execution failed: %v",
			err,
		)
	} else {
		pattern.JobState.Status = coreentity.StatusIdle
		pattern.JobState.LastExecution = now
		pattern.JobState.NextExecution = js.calculator.CalculateNextExecution(
			pattern,
			now,
		)
		pattern.JobState.ExecutionCount++
		pattern.JobState.RecoveryAttempts = 0
		log.Printf(
			"Job executed successfully. Next execution: %v",
			pattern.JobState.NextExecution,
		)
	}

	// Save final state
	if saveErr := js.recurrencePatternRepo.Update(
		ctx,
		*pattern,
	); saveErr != nil {
		log.Printf(
			"Failed to save job state: %v",
			saveErr,
		)
	}

	return err
}

func (js *Orchestrator) recoverFromFailure(ctx context.Context) {
	failedJobs, errGetJobState := js.recurrencePatternRepo.GetFailedJobs(ctx)
	if errGetJobState != nil || len(failedJobs) == 0 {
		return // No failedJobs to recover from
	}

	for _, failedJob := range failedJobs {
		go func() {
			// Check if we need recovery
			if failedJob.JobState.Status != coreentity.StatusError {
				return
			}

			// Check if we've exceeded max recovery attempts
			if failedJob.JobState.RecoveryAttempts >= MaxRecoveryAttempts {
				log.Printf(
					"Max recovery attempts reached (%d), manual intervention required",
					MaxRecoveryAttempts,
				)

				return
			}

			log.Printf(
				"Attempting recovery (attempt %d/%d)",
				failedJob.JobState.RecoveryAttempts+1,
				MaxRecoveryAttempts,
			)

			// Wait before recovery attempt
			js.timeProvider.Sleep(RecoveryDelay)

			// Mark as recovering
			failedJob.JobState.Status = coreentity.StatusRecovering
			if err := js.recurrencePatternRepo.Update(
				ctx,
				failedJob,
			); err != nil {
				log.Err(err).Msgf(
					"Failed to mark job ID %s as recovering",
					failedJob.ID,
				)
			}

			// Attempt to execute
			err := js.executeJob(
				ctx,
				&failedJob,
			)
			if err != nil {
				log.Err(err).Msgf(
					"Failed to recover job ID %s",
					failedJob.ID,
				)
			}
		}()
	}
}
