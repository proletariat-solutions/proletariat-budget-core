package service

import (
	"context"

	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/usecase"
)

type TaskProcessorService struct {
	usecase *usecase.UseCases
}

func (t *TaskProcessorService) ProcessTask(
	ctx context.Context,
	task *coreentity.RecurrencePattern,
) error {
	switch task.TaskType {
	case coreentity.TaskTypeRecurringIngress:
		err := t.createIngressFromTemplate(
			ctx,
			task.Template.Ingress,
			task,
		)

		return err
	case coreentity.TaskTypeRecurringExpenditure:
		err := t.createExpenditureFromTemplate(
			ctx,
			task.Template.Expenditure,
			task,
		)

		return err
	case coreentity.TaskTypeRecurringSavingsContribution:
		err := t.createSavingsContributionFromTemplate(
			ctx,
			task.Template.SavingsContribution,
			task,
		)

		return err
	case coreentity.TaskTypeRecurringTransfer:
		err := t.createTransferFromTemplate(
			ctx,
			task.Template.Transfer,
			task,
		)

		return err
	default:
	}

	return nil
}

func (t *TaskProcessorService) createIngressFromTemplate(
	ctx context.Context,
	template *coreentity.Ingress,
	task *coreentity.RecurrencePattern,
) error {
	_, err := t.usecase.Ingress.Create(
		ctx,
		*template,
		task,
	)

	return err
}

func (t *TaskProcessorService) createExpenditureFromTemplate(
	ctx context.Context,
	template *coreentity.Expenditure,
	task *coreentity.RecurrencePattern,
) error {
	_, err := t.usecase.Expenditure.Create(
		ctx,
		*template,
		task,
	)

	return err
}

func (t *TaskProcessorService) createSavingsContributionFromTemplate(
	ctx context.Context,
	template *coreentity.SavingsContribution,
	task *coreentity.RecurrencePattern,
) error {
	_, err := t.usecase.SavingOperation.CreateContribution(
		ctx,
		*template,
		task,
	)

	return err
}

func (t *TaskProcessorService) createTransferFromTemplate(
	ctx context.Context,
	template *coreentity.Transfer,
	task *coreentity.RecurrencePattern,
) error {
	_, err := t.usecase.Transfer.Create(
		ctx,
		*template,
		task,
	)

	return err
}
