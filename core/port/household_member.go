package port

import (
	"context"

	"proletariat-budget-core/core/domain/coreentity"
)

//go:generate mockgen -source=household_member.go -destination=../../test/mocks/mock_household_memberrepo.go -package mocks
type HouseholdMember interface {
	Create(
		ctx context.Context,
		householdMember coreentity.HouseholdMember,
	) (
		string,
		error,
	)
	Update(
		ctx context.Context,
		id string,
		householdMember coreentity.HouseholdMember,
	) error
	Delete(
		ctx context.Context,
		id string,
	) error
	Deactivate(
		ctx context.Context,
		id string,
	) error
	Activate(
		ctx context.Context,
		id string,
	) error
	CanDelete(
		ctx context.Context,
		id string,
	) (
		bool,
		error,
	)
	GetByID(
		ctx context.Context,
		id string,
	) (
		*coreentity.HouseholdMember,
		error,
	)
	List(
		ctx context.Context,
		params *coreentity.HouseholdMemberListParams,
	) (
		*coreentity.HouseholdMemberList,
		error,
	)
}
