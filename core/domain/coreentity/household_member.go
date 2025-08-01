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
	ErrMemberFirstNameRequired = errors.New("member first name is required")
	ErrMemberLastNameRequired  = errors.New("member last name is required")
	ErrMemberRoleRequired      = errors.New("member role is required")
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

// Validate checks if the household member is valid
func (hm *HouseholdMember) Validate() error {
	if hm.FirstName == "" {
		return ErrMemberFirstNameRequired
	}

	if hm.LastName == "" {
		return ErrMemberLastNameRequired
	}

	if hm.Role == "" {
		return ErrMemberRoleRequired
	}

	return nil
}

// HouseholdMemberList represents a paginated list of household members
type HouseholdMemberList struct {
	HouseholdMembers []HouseholdMember `json:"household_members"`
}

type HouseholdMemberListParams struct {
	Active *bool   `json:"active,omitempty"`
	Role   *string `json:"role,omitempty"`
}
