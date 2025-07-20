package usecase

import (
	"context"
	"errors"

	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/port"
)

type HouseholdMember struct {
	householdMembersRepo port.HouseholdMember
}

func NewHouseholdMemberUseCase(householdMembersRepo port.HouseholdMember) *HouseholdMember {
	return &HouseholdMember{householdMembersRepo: householdMembersRepo}
}

func (u *HouseholdMember) ListHouseholdMembers(
	ctx context.Context,
	params *coreentity.HouseholdMemberListParams,
) (
	*coreentity.HouseholdMemberList,
	error,
) {
	return u.householdMembersRepo.List(
		ctx,
		params,
	)
}

func (u *HouseholdMember) CreateHouseholdMember(
	ctx context.Context,
	householdMember coreentity.HouseholdMember,
) (
	*coreentity.HouseholdMember,
	error,
) {
	id, err := u.householdMembersRepo.Create(
		ctx,
		householdMember,
	)
	if err != nil {
		return nil, err
	}
	householdMember.ID = id

	return &householdMember, nil
}

func (u *HouseholdMember) UpdateHouseholdMember(
	ctx context.Context,
	id string,
	householdMember coreentity.HouseholdMember,
) error {
	errUpdate := u.householdMembersRepo.Update(
		ctx,
		id,
		householdMember,
	)
	if errUpdate != nil {
		if errors.Is(
			errUpdate,
			port.ErrRecordNotFound,
		) {
			return coreentity.ErrMemberNotFound
		}

		return errUpdate
	}

	return nil
}

func (u *HouseholdMember) DeleteHouseholdMember(
	ctx context.Context,
	id string,
) error {
	canDelete, err := u.householdMembersRepo.CanDelete(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return coreentity.ErrMemberNotFound
		}

		return err
	}
	if !canDelete {
		return coreentity.ErrMemberHasActiveAccounts
	}

	errDelete := u.householdMembersRepo.Delete(
		ctx,
		id,
	)
	if errDelete != nil {
		return errDelete
	}

	return nil
}

func (u *HouseholdMember) GetHouseholdMemberByID(
	ctx context.Context,
	id string,
) (
	*coreentity.HouseholdMember,
	error,
) {
	member, err := u.householdMembersRepo.GetByID(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return nil, coreentity.ErrMemberNotFound
		}

		return nil, err
	}

	return member, nil
}

func (u *HouseholdMember) DeactivateHouseholdMember(
	ctx context.Context,
	id string,
) error {
	member, err := u.GetHouseholdMemberByID(
		ctx,
		id,
	)
	if err != nil {
		return err
	}
	if !member.Active {
		return coreentity.ErrMemberAlreadyInactive
	}

	return u.householdMembersRepo.Deactivate(
		ctx,
		id,
	)
}

func (u *HouseholdMember) CanDeleteHouseholdMember(
	ctx context.Context,
	id string,
) (
	bool,
	error,
) {
	_, err := u.GetHouseholdMemberByID(
		ctx,
		id,
	)
	if err != nil {
		return false, err
	}

	return u.householdMembersRepo.CanDelete(
		ctx,
		id,
	)
}

func (u *HouseholdMember) ActivateHouseholdMember(
	ctx context.Context,
	id string,
) error {
	member, err := u.GetHouseholdMemberByID(
		ctx,
		id,
	)
	if err != nil {
		return err
	}
	if member.Active {
		return coreentity.ErrMemberAlreadyActive
	}

	return u.householdMembersRepo.Activate(
		ctx,
		id,
	)
}
