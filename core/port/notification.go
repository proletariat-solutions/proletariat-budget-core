package port

import (
	"context"

	domain "proletariat-budget-core/core/domain/coreentity"
)

type Notification interface {
	Create(
		ctx context.Context,
		notification domain.Notification,
	) (
		*domain.Notification,
		error,
	)
	GetByID(
		ctx context.Context,
		id string,
	) (
		*domain.Notification,
		error,
	)
	List(
		ctx context.Context,
		params domain.NotificationListParams,
	) (
		*domain.NotificationList,
		error,
	)
	Delete(
		ctx context.Context,
		id string,
	) error
	DeleteAll(
		ctx context.Context,
	) error
	MarkAsRead(
		ctx context.Context,
		ids []string,
	) error
	MarkAsUnread(
		ctx context.Context,
		ids []string,
	) error
	MarkAllAsRead(
		ctx context.Context,
	) error
}
