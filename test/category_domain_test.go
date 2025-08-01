package test

import (
	"github.com/stretchr/testify/assert"
	"proletariat-budget-core/core/domain/coreentity"
	"testing"
)

func TestCategory_Validate_ShouldReturnErrCategoryNameEmptyWhenNameFieldIsEmptyString(t *testing.T) {
	// Arrange
	category := coreentity.Category{
		Name:            "",
		Color:           "#FF0000",
		BackgroundColor: "#FFFFFF",
		Active:          true,
		CategoryType:    coreentity.CategoryTypeExpenditure,
	}

	// Act
	err := category.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrCategoryNameEmpty,
		err,
	)
}

func TestCategory_Validate_ShouldReturnErrCategoryColorEmptyWhenColorFieldIsEmptyStringAndNameIsValid(t *testing.T) {
	// Arrange
	category := coreentity.Category{
		Name:            "Test Category",
		Color:           "",
		BackgroundColor: "#FFFFFF",
		Active:          true,
		CategoryType:    coreentity.CategoryTypeExpenditure,
	}

	// Act
	err := category.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrCategoryColorEmpty,
		err,
	)
}

func TestCategory_Validate_ShouldReturnErrCategoryBackgroundColorEmptyWhenBackgroundColorFieldIsEmptyStringAndNameAndColorAreValid(t *testing.T) {
	// Arrange
	category := coreentity.Category{
		Name:            "Test Category",
		Color:           "#FF0000",
		BackgroundColor: "",
		Active:          true,
		CategoryType:    coreentity.CategoryTypeExpenditure,
	}

	// Act
	err := category.Validate()

	// Assert
	assert.Equal(
		t,
		coreentity.ErrCategoryBackgroundColorEmpty,
		err,
	)
}

func TestCategory_Validate_ShouldReturnNilWhenAllRequiredFieldsAreNonEmpty(t *testing.T) {
	// Arrange
	category := coreentity.Category{
		Name:            "Test Category",
		Color:           "#FF0000",
		BackgroundColor: "#FFFFFF",
		Active:          true,
		CategoryType:    coreentity.CategoryTypeExpenditure,
	}

	// Act
	err := category.Validate()

	// Assert
	assert.NoError(
		t,
		err,
	)
}
