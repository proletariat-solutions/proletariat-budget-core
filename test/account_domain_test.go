package test

import (
	"github.com/stretchr/testify/assert"
	"proletariat-budget-core/core/domain/coreentity"
	"testing"
)

func TestAccount_Validate_ReturnsErrAccountNameRequiredWhenNameIsEmpty(t *testing.T) {
	// Arrange
	account := coreentity.Account{
		Name: "",
		Currency: &coreentity.Currency{
			ID: "USD",
		},
		Owner: &coreentity.HouseholdMember{
			ID: "member-123",
		},
	}

	// Act
	err := account.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrAccountNameRequired,
		err,
	)
}
func TestAccount_Validate_ReturnsErrAccountOwnerRequiredWhenOwnerIsNil(t *testing.T) {
	// Arrange
	account := coreentity.Account{
		Name: "Test Account",
		Currency: &coreentity.Currency{
			ID: "USD",
		},
		Owner: nil,
	}

	// Act
	err := account.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrAccountOwnerRequired,
		err,
	)
}

func TestAccount_Validate_ReturnsErrCurrencyIDRequiredWhenCurrencyValidationFailsWithErrCurrencyRequired(t *testing.T) {
	// Arrange
	account := coreentity.Account{
		Name: "Test Account",
		Currency: &coreentity.Currency{
			ID: "", // Empty currency ID to trigger validation error
		},
		Owner: &coreentity.HouseholdMember{
			ID: "member-123",
		},
	}

	// Act
	err := account.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrAccountCurrencyRequired,
		err,
	)
}
