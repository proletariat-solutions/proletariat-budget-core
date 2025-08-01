package test

import (
	"github.com/stretchr/testify/assert"
	"proletariat-budget-core/core/domain/coreentity"
	"testing"
)

func TestHouseholdMember_Validate_ShouldReturnErrMemberFirstNameRequiredWhenFirstNameIsEmptyString(t *testing.T) {
	// Arrange
	householdMember := coreentity.HouseholdMember{
		FirstName: "",
		LastName:  "Doe",
		Role:      "owner",
	}

	// Act
	err := householdMember.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrMemberFirstNameRequired,
		err,
	)
}

func TestHouseholdMember_Validate_ShouldReturnErrMemberLastNameRequiredWhenLastNameIsEmptyString(t *testing.T) {
	// Arrange
	householdMember := coreentity.HouseholdMember{
		FirstName: "John",
		LastName:  "",
		Role:      "owner",
	}

	// Act
	err := householdMember.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrMemberLastNameRequired,
		err,
	)
}
func TestHouseholdMember_Validate_ShouldReturnErrMemberRoleRequiredWhenRoleIsEmptyString(t *testing.T) {
	// Arrange
	householdMember := coreentity.HouseholdMember{
		FirstName: "John",
		LastName:  "Doe",
		Role:      "",
	}

	// Act
	err := householdMember.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrMemberRoleRequired,
		err,
	)
}
func TestHouseholdMember_Validate_ShouldReturnNilWhenAllRequiredFieldsAreProvidedWithValidValues(t *testing.T) {
	// Arrange
	householdMember := coreentity.HouseholdMember{
		FirstName: "John",
		LastName:  "Doe",
		Role:      "owner",
	}

	// Act
	err := householdMember.Validate()

	// Assert
	assert.NoError(
		t,
		err,
	)
}
