package port

import (
	"context"

	"proletariat-budget-core/core/domain/coreentity"
)

type Transfer interface {
	Create(
		ctx context.Context,
		transfer coreentity.Transfer,
	) (
		string,
		error,
	)
	GetByID(
		ctx context.Context,
		id string,
	) (
		*coreentity.Transfer,
		error,
	)
	List(
		ctx context.Context,
		params coreentity.ListTransfersParams,
	) (
		*coreentity.TransferList,
		error,
	)
	UpdateTemplate(
		ctx context.Context,
		template coreentity.Transfer,
	) error
	DeleteTemplate(
		ctx context.Context,
		template coreentity.Transfer,
	) error
}
