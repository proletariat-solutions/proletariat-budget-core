package coreentity

import "time"

type AuditData struct {
	Date      *time.Time       `json:"date"`
	CreatedBy *HouseholdMember `json:"created_by,omitempty"`
}
