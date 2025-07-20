package coreentity

import (
	"errors"
	"time"
)

var (
	ErrMemberHasActiveAccounts = errors.New("member has active accounts")
	ErrMemberAlreadyActive     = errors.New("member is already active")
	ErrMemberNotFound          = errors.New("member not found")
	ErrMemberAlreadyInactive   = errors.New("member is already inactive")
	ErrMemberInactive          = errors.New("member is inactive")
)

// HouseholdMember represents a household member in the domain
type HouseholdMember struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Nickname  *string   `json:"nickname,omitempty"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// IsActive returns true if the household member is active
func (hm *HouseholdMember) IsActive() bool {
	return hm.Active
}

// FullName returns the full name of the household member
func (hm *HouseholdMember) FullName() string {
	return hm.FirstName + " " + hm.LastName
}

// DisplayName returns the display name (nickname if available, otherwise full name)
func (hm *HouseholdMember) DisplayName() string {
	if hm.Nickname != nil && *hm.Nickname != "" {
		return *hm.Nickname
	}

	return hm.FullName()
}

// Deactivate marks the household member as inactive
func (hm *HouseholdMember) Deactivate() {
	hm.Active = false
	hm.UpdatedAt = time.Now()
}

// Activate marks the household member as active
func (hm *HouseholdMember) Activate() {
	hm.Active = true
	hm.UpdatedAt = time.Now()
}

// HouseholdMemberList represents a paginated list of household members
type HouseholdMemberList struct {
	HouseholdMembers []HouseholdMember `json:"household_members"`
}

type HouseholdMemberListParams struct {
	Active *bool   `json:"active,omitempty"`
	Role   *string `json:"role,omitempty"`
}
