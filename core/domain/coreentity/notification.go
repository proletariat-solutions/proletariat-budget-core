package coreentity

import (
	"time"

	"proletariat-budget-core/core/domain/misc"
)

type Notification struct {
	ID    string           `json:"id"`
	Date  *time.Time       `json:"date"`
	Title string           `json:"message"`
	Text  string           `json:"text"`
	Type  NotificationType `json:"type"`
	Read  bool             `json:"read"`
}

type NotificationType string

const (
	InfoNotificationType    NotificationType = "info"
	WarningNotificationType NotificationType = "warning"
	ErrorNotificationType   NotificationType = "error"
)

type NotificationList struct {
	Items    []Notification    `json:"items"`
	Metadata misc.ListMetadata `json:"metadata"`
}

type NotificationListParams struct {
	Type     *NotificationType `json:"type"`
	DateFrom *time.Time        `json:"date_from"`
	DateTo   *time.Time        `json:"date_to"`
	Read     *bool             `json:"read"`
	misc.ListParams
}
