package usecase

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/port"
)

type RecurrencePattern struct {
	recurrencePatternRepo port.RecurrencePattern
	txManager             port.TransactionManager
}

func (rp *RecurrencePattern) Create(
	ctx context.Context,
	recurrencePattern coreentity.RecurrencePattern,
) (
	*coreentity.RecurrencePattern,
	error,
) {
	validationErr := recurrencePattern.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	switch recurrencePattern.TaskType {
	case coreentity.TaskTypeRecurringExpenditure:
		expenditureTemplate, err := rp.createExpenditureTemplate(
			ctx,
			recurrencePattern.Template.Expenditure,
		)
		if err != nil {
			return nil, err
		}
		recurrencePattern.Template.Expenditure = expenditureTemplate
	case coreentity.TaskTypeRecurringTransfer:
		transferTemplate, err := rp.createTransferTemplate(
			ctx,
			recurrencePattern.Template.Transfer,
		)
		if err != nil {
			return nil, err
		}
		recurrencePattern.Template.Transfer = transferTemplate
	case coreentity.TaskTypeRecurringSavingsContribution:
		savingsContributionTemplate, err := rp.createSavingsContributionTemplate(
			ctx,
			recurrencePattern.Template.SavingsContribution,
		)
		if err != nil {
			return nil, err
		}
		recurrencePattern.Template.SavingsContribution = savingsContributionTemplate
	case coreentity.TaskTypeRecurringIngress:
		ingressTemplate, err := rp.createIngressTemplate(
			ctx,
			recurrencePattern.Template.Ingress,
		)
		if err != nil {
			return nil, err
		}
		recurrencePattern.Template.Ingress = ingressTemplate
	default:
		return nil, coreentity.ErrRecurrencyPatternUnsupportedTaskType
	}

	recurrencePatternID, err := rp.recurrencePatternRepo.Create(
		ctx,
		recurrencePattern,
	)
	if err != nil {
		return nil, err
	}
	recurrencePattern.ID = recurrencePatternID

	return &recurrencePattern, nil
}

func (rp *RecurrencePattern) GetByID(
	ctx context.Context,
	id string,
) (
	*coreentity.RecurrencePattern,
	error,
) {
	recurrencePattern, err := rp.recurrencePatternRepo.GetByID(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return nil, coreentity.ErrRecurrencePatternNotFound
		}

		return nil, err
	}

	return recurrencePattern, nil
}

func (rp *RecurrencePattern) Update(
	ctx context.Context,
	recurrencePattern coreentity.RecurrencePattern,
) (
	*coreentity.RecurrencePattern,
	error,
) {
	validationErr := recurrencePattern.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	err := rp.recurrencePatternRepo.Update(
		ctx,
		recurrencePattern,
	)
	if err != nil {
		return nil, err
	}

	return &recurrencePattern, nil
}

func (rp *RecurrencePattern) Delete(
	ctx context.Context,
	id string,
) error {
	err := rp.recurrencePatternRepo.Delete(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return coreentity.ErrRecurrencePatternNotFound
		}

		return err
	}

	return nil
}

func (rp *RecurrencePattern) List(
	ctx context.Context,
) (
	[]coreentity.RecurrencePattern,
	error,
) {
	recurrencePatterns, err := rp.recurrencePatternRepo.List(
		ctx,
	)
	if err != nil {
		return nil, err
	}

	return recurrencePatterns, nil
}

func (rp *RecurrencePattern) createExpenditureTemplate(
	ctx context.Context,
	template *coreentity.Expenditure,
) (
	*coreentity.Expenditure,
	error,
) {
	errTx := rp.txManager.WithDatabaseTransaction(
		ctx,
		func(
			ctx context.Context,
			tx port.TransactionContext,
		) error {
			transactionRepo := tx.GetTransactionRepo()
			expenditureRepo := tx.GetExpenditureRepo()
			tagsRepo := tx.GetTagsRepo()

			// Process transaction
			statusCompleted := coreentity.TransactionStatusIgnored
			template.Transaction.Status = &statusCompleted

			txID, err := transactionRepo.Create(
				ctx,
				*template.Transaction,
			)
			if err != nil {
				return err
			}

			template.Transaction.ID = &txID

			var expID string
			var errCreate error
			// Create expenditure record
			expID, errCreate = expenditureRepo.Create(
				ctx,
				*template,
			)
			if errCreate != nil {
				return errCreate
			}

			// Link tags if present
			errLink := tagsRepo.LinkTagsToType(
				ctx,
				expID,
				template.Tags,
			)
			if errLink != nil {
				// Logging and notifying error to the user, since this is a recoverable error
				// TODO: notify user
				log.Err(errLink).Msgf(
					"Failed to link tags to expenditure %s",
					expID,
				)
			}

			return nil
		},
	)

	return template, errTx
}

func (rp *RecurrencePattern) createIngressTemplate(
	ctx context.Context,
	template *coreentity.Ingress,
) (
	*coreentity.Ingress,
	error,
) {
	errTx := rp.txManager.WithDatabaseTransaction(
		ctx,
		func(
			ctx context.Context,
			tx port.TransactionContext,
		) error {
			transactionRepo := tx.GetTransactionRepo()
			ingressRepo := tx.GetIngressRepo()
			tagsRepo := tx.GetTagsRepo()

			// Process transaction
			statusCompleted := coreentity.TransactionStatusIgnored
			template.Transaction.Status = &statusCompleted

			txID, err := transactionRepo.Create(
				ctx,
				*template.Transaction,
			)
			if err != nil {
				return err
			}

			template.Transaction.ID = &txID

			var ingID string
			var errCreate error
			// Create expenditure record
			ingID, errCreate = ingressRepo.Create(
				ctx,
				*template,
			)
			if errCreate != nil {
				return errCreate
			}

			// Link tags if present
			errLink := tagsRepo.LinkTagsToType(
				ctx,
				ingID,
				template.Tags,
			)
			if errLink != nil {
				// Logging and notifying error to the user, since this is a recoverable error
				// TODO: notify user
				log.Err(errLink).Msgf(
					"Failed to link tags to expenditure %s",
					ingID,
				)
			}

			return nil
		},
	)

	return template, errTx
}

func (rp *RecurrencePattern) createTransferTemplate(
	ctx context.Context,
	template *coreentity.Transfer,
) (
	*coreentity.Transfer,
	error,
) {
	errTx := rp.txManager.WithDatabaseTransaction(
		ctx,
		func(
			ctx context.Context,
			tx port.TransactionContext,
		) error {
			transactionRepo := tx.GetTransactionRepo()
			transferRepo := tx.GetTransferRepo()
			tagsRepo := tx.GetTagsRepo()

			// Process transaction
			statusCompleted := coreentity.TransactionStatusIgnored
			template.IncomingTransaction.Status = &statusCompleted
			template.OutgoingTransaction.Status = &statusCompleted

			txIDIncoming, errIncomingTx := transactionRepo.Create(
				ctx,
				*template.IncomingTransaction,
			)
			if errIncomingTx != nil {
				return errIncomingTx
			}

			txIDOutgoing, errOutgoingTx := transactionRepo.Create(
				ctx,
				*template.OutgoingTransaction,
			)
			if errOutgoingTx != nil {
				return errOutgoingTx
			}

			template.IncomingTransaction.ID = &txIDIncoming
			template.OutgoingTransaction.ID = &txIDOutgoing

			var trID string
			var errCreate error
			// Create expenditure record
			trID, errCreate = transferRepo.Create(
				ctx,
				*template,
			)
			if errCreate != nil {
				return errCreate
			}

			// Link tags if present
			errLink := tagsRepo.LinkTagsToType(
				ctx,
				trID,
				template.Tags,
			)
			if errLink != nil {
				// Logging and notifying error to the user, since this is a recoverable error
				// TODO: notify user
				log.Err(errLink).Msgf(
					"Failed to link tags to transfer %s",
					trID,
				)
			}

			return nil
		},
	)

	return template, errTx
}

func (rp *RecurrencePattern) createSavingsContributionTemplate(
	ctx context.Context,
	template *coreentity.SavingsContribution,
) (
	*coreentity.SavingsContribution,
	error,
) {
	// TODO: Implement logic to create savings contribution template
	return template, nil
}
