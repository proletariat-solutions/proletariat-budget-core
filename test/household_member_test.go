package test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/port"
	"proletariat-budget-core/core/usecase"
	"proletariat-budget-core/test/mocks"
	"testing"
)

func TestHouseholdMemberUseCase_ListHouseholdMembers_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	isActive := true
	role := "owner"

	params := &coreentity.HouseholdMemberListParams{
		Active: &isActive,
		Role:   &role,
	}

	expectedHouseholdMemberList := &coreentity.HouseholdMemberList{
		HouseholdMembers: []coreentity.HouseholdMember{
			{
				ID:        "member-1",
				FirstName: "John",
				LastName:  "Doe",
				Active:    true,
				Role:      "owner",
			},
			{
				ID:        "member-2",
				FirstName: "Jane",
				LastName:  "Smith",
				Active:    true,
				Role:      "member",
			},
		},
	}

	householdMembersRepo.EXPECT().List(
		ctx,
		params,
	).Return(
		expectedHouseholdMemberList,
		nil,
	)

	// Act
	result, err := useCase.ListHouseholdMembers(
		ctx,
		params,
	)

	// Assert
	assert.NoError(
		t,
		err,
	)
	assert.NotNil(
		t,
		result,
	)
	assert.Equal(
		t,
		expectedHouseholdMemberList,
		result,
	)
}

func TestHouseholdMemberUseCase_ListHouseholdMembers_ReturnsErrorWhenRepositoryListOperationFails(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	isActive := true
	role := "owner"

	params := &coreentity.HouseholdMemberListParams{
		Active: &isActive,
		Role:   &role,
	}

	expectedError := errors.New("database connection error")

	householdMembersRepo.EXPECT().List(
		ctx,
		params,
	).Return(
		nil,
		expectedError,
	)

	// Act
	result, err := useCase.ListHouseholdMembers(
		ctx,
		params,
	)

	// Assert
	assert.Error(
		t,
		err,
	)
	assert.Nil(
		t,
		result,
	)
	assert.Equal(
		t,
		expectedError,
		err,
	)
}

func TestHouseholdMemberUseCase_CreateHouseholdMember_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	householdMember := coreentity.HouseholdMember{
		FirstName: "John",
		LastName:  "Doe",
		Active:    true,
		Role:      "owner",
	}

	expectedID := "member-123"

	householdMembersRepo.EXPECT().Create(
		ctx,
		householdMember,
	).Return(
		expectedID,
		nil,
	)

	// Act
	result, err := useCase.CreateHouseholdMember(
		ctx,
		householdMember,
	)

	// Assert
	assert.NoError(
		t,
		err,
	)
	assert.NotNil(
		t,
		result,
	)
	assert.Equal(
		t,
		expectedID,
		result.ID,
	)
	assert.Equal(
		t,
		householdMember.FirstName,
		result.FirstName,
	)
	assert.Equal(
		t,
		householdMember.LastName,
		result.LastName,
	)
	assert.Equal(
		t,
		householdMember.Active,
		result.Active,
	)
	assert.Equal(
		t,
		householdMember.Role,
		result.Role,
	)
}

func TestHouseholdMemberUseCase_CreateHouseholdMember_ReturnsErrorWhenRepositoryCreateOperationFails(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	householdMember := coreentity.HouseholdMember{
		FirstName: "John",
		LastName:  "Doe",
		Active:    true,
		Role:      "owner",
	}

	expectedError := errors.New("database connection error")

	householdMembersRepo.EXPECT().Create(
		ctx,
		householdMember,
	).Return(
		"",
		expectedError,
	)

	// Act
	result, err := useCase.CreateHouseholdMember(
		ctx,
		householdMember,
	)

	// Assert
	assert.Error(
		t,
		err,
	)
	assert.Nil(
		t,
		result,
	)
	assert.Equal(
		t,
		expectedError,
		err,
	)
}

func TestHouseholdMemberUseCase_UpdateHouseholdMember_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"
	householdMember := coreentity.HouseholdMember{
		FirstName: "John",
		LastName:  "Smith",
		Active:    true,
		Role:      "member",
	}

	householdMembersRepo.EXPECT().Update(
		ctx,
		memberID,
		householdMember,
	).Return(
		nil,
	)

	// Act
	err := useCase.UpdateHouseholdMember(
		ctx,
		memberID,
		householdMember,
	)

	// Assert
	assert.NoError(
		t,
		err,
	)
}

func TestHouseholdMemberUseCase_UpdateHouseholdMember_ReturnsErrMemberNotFoundWhenRepositoryReturnsErrRecordNotFound(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"
	householdMember := coreentity.HouseholdMember{
		FirstName: "John",
		LastName:  "Smith",
		Active:    true,
		Role:      "member",
	}

	householdMembersRepo.EXPECT().Update(
		ctx,
		memberID,
		householdMember,
	).Return(
		port.ErrRecordNotFound,
	)

	// Act
	err := useCase.UpdateHouseholdMember(
		ctx,
		memberID,
		householdMember,
	)

	// Assert
	assert.Equal(
		t,
		coreentity.ErrMemberNotFound,
		err,
	)
}

func TestHouseholdMemberUseCase_UpdateHouseholdMember_ReturnsOriginalErrorWhenRepositoryUpdateFailsWithDatabaseConnectionError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"
	householdMember := coreentity.HouseholdMember{
		FirstName: "John",
		LastName:  "Smith",
		Active:    true,
		Role:      "member",
	}

	unexpectedError := errors.New("database connection error")

	householdMembersRepo.EXPECT().Update(
		ctx,
		memberID,
		householdMember,
	).Return(
		unexpectedError,
	)

	// Act
	err := useCase.UpdateHouseholdMember(
		ctx,
		memberID,
		householdMember,
	)

	// Assert
	assert.Equal(
		t,
		unexpectedError,
		err,
	)
}

func TestHouseholdMemberUseCase_DeleteHouseholdMember_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"

	householdMembersRepo.EXPECT().CanDelete(
		ctx,
		memberID,
	).Return(
		true,
		nil,
	)

	householdMembersRepo.EXPECT().Delete(
		ctx,
		memberID,
	).Return(
		nil,
	)

	// Act
	err := useCase.DeleteHouseholdMember(
		ctx,
		memberID,
	)

	// Assert
	assert.NoError(
		t,
		err,
	)
}

func TestHouseholdMemberUseCase_DeleteHouseholdMember_ReturnsErrMemberHasActiveAccountsWhenCanDeleteReturnsFalse(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"

	householdMembersRepo.EXPECT().CanDelete(
		ctx,
		memberID,
	).Return(
		false,
		nil,
	)

	// Act
	err := useCase.DeleteHouseholdMember(
		ctx,
		memberID,
	)

	// Assert
	assert.Equal(
		t,
		coreentity.ErrMemberHasActiveAccounts,
		err,
	)
}

func TestHouseholdMemberUseCase_DeleteHouseholdMember_ReturnsOriginalErrorWhenCanDeleteOperationFailsWithDatabaseError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"
	unexpectedError := errors.New("database connection error")

	householdMembersRepo.EXPECT().CanDelete(
		ctx,
		memberID,
	).Return(
		true,
		unexpectedError,
	)

	// Act
	err := useCase.DeleteHouseholdMember(
		ctx,
		memberID,
	)

	// Assert
	assert.Equal(
		t,
		unexpectedError,
		err,
	)
}

func TestHouseholdMemberUseCase_DeleteHouseholdMember_ReturnsOriginalErrorWhenDeleteOperationFailsWithDatabaseConnectionError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"
	unexpectedError := errors.New("database connection error")

	householdMembersRepo.EXPECT().CanDelete(
		ctx,
		memberID,
	).Return(
		true,
		nil,
	)

	householdMembersRepo.EXPECT().Delete(
		ctx,
		memberID,
	).Return(
		unexpectedError,
	)

	// Act
	err := useCase.DeleteHouseholdMember(
		ctx,
		memberID,
	)

	// Assert
	assert.Equal(
		t,
		unexpectedError,
		err,
	)
}

func TestHouseholdMemberUseCase_DeleteHouseholdMember_ReturnsErrMemberNotFoundWhenDeleteOperationReturnsErrRecordNotFound(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"

	householdMembersRepo.EXPECT().CanDelete(
		ctx,
		memberID,
	).Return(
		false,
		port.ErrRecordNotFound,
	)

	// Act
	err := useCase.DeleteHouseholdMember(
		ctx,
		memberID,
	)

	// Assert
	assert.Equal(
		t,
		coreentity.ErrMemberNotFound,
		err,
	)
}

func TestHouseholdMemberUseCase_GetHouseholdMemberByID_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"
	expectedMember := &coreentity.HouseholdMember{
		ID:        memberID,
		FirstName: "John",
		LastName:  "Doe",
		Active:    true,
		Role:      "owner",
	}

	householdMembersRepo.EXPECT().GetByID(
		ctx,
		memberID,
	).Return(
		expectedMember,
		nil,
	)

	// Act
	result, err := useCase.GetHouseholdMemberByID(
		ctx,
		memberID,
	)

	// Assert
	assert.NoError(
		t,
		err,
	)
	assert.NotNil(
		t,
		result,
	)
	assert.Equal(
		t,
		expectedMember,
		result,
	)
}

func TestHouseholdMemberUseCase_GetHouseholdMemberByID_ReturnsErrMemberNotFoundWhenRepositoryReturnsErrRecordNotFound(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"

	householdMembersRepo.EXPECT().GetByID(
		ctx,
		memberID,
	).Return(
		nil,
		port.ErrRecordNotFound,
	)

	// Act
	result, err := useCase.GetHouseholdMemberByID(
		ctx,
		memberID,
	)

	// Assert
	assert.Nil(
		t,
		result,
	)
	assert.Equal(
		t,
		coreentity.ErrMemberNotFound,
		err,
	)
}

func TestHouseholdMemberUseCase_GetHouseholdMemberByID_ReturnsOriginalErrorWhenRepositoryGetByIDFailsWithDatabaseConnectionError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"
	unexpectedError := errors.New("database connection error")

	householdMembersRepo.EXPECT().GetByID(
		ctx,
		memberID,
	).Return(
		nil,
		unexpectedError,
	)

	// Act
	result, err := useCase.GetHouseholdMemberByID(
		ctx,
		memberID,
	)

	// Assert
	assert.Nil(
		t,
		result,
	)
	assert.Equal(
		t,
		unexpectedError,
		err,
	)
}

func TestHouseholdMemberUseCase_DeactivateHouseholdMember_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"
	activeMember := &coreentity.HouseholdMember{
		ID:        memberID,
		FirstName: "John",
		LastName:  "Doe",
		Active:    true,
		Role:      "owner",
	}

	householdMembersRepo.EXPECT().GetByID(
		ctx,
		memberID,
	).Return(
		activeMember,
		nil,
	)

	householdMembersRepo.EXPECT().Deactivate(
		ctx,
		memberID,
	).Return(
		nil,
	)

	// Act
	err := useCase.DeactivateHouseholdMember(
		ctx,
		memberID,
	)

	// Assert
	assert.NoError(
		t,
		err,
	)
}

func TestHouseholdMemberUseCase_DeactivateHouseholdMember_ReturnsErrorFromGetHouseholdMemberByIDWhenRepositoryGetByIDOperationFailsWithDatabaseConnectionError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"
	unexpectedError := errors.New("database connection error")

	householdMembersRepo.EXPECT().GetByID(
		ctx,
		memberID,
	).Return(
		nil,
		unexpectedError,
	)

	// Act
	err := useCase.DeactivateHouseholdMember(
		ctx,
		memberID,
	)

	// Assert
	assert.Equal(
		t,
		unexpectedError,
		err,
	)
}

func TestHouseholdMemberUseCase_DeactivateHouseholdMember_ReturnsErrMemberNotFoundWhenGetHouseholdMemberByIDReturnsErrMemberNotFoundForNonExistentMemberID(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "non-existent-member-123"

	householdMembersRepo.EXPECT().GetByID(
		ctx,
		memberID,
	).Return(
		nil,
		port.ErrRecordNotFound,
	)

	// Act
	err := useCase.DeactivateHouseholdMember(
		ctx,
		memberID,
	)

	// Assert
	assert.Equal(
		t,
		coreentity.ErrMemberNotFound,
		err,
	)
}

func TestHouseholdMemberUseCase_DeactivateHouseholdMember_ReturnsErrMemberAlreadyInactiveWhenAttemptingToDeactivateInactiveMember(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"
	inactiveMember := &coreentity.HouseholdMember{
		ID:        memberID,
		FirstName: "John",
		LastName:  "Doe",
		Active:    false,
		Role:      "owner",
	}

	householdMembersRepo.EXPECT().GetByID(
		ctx,
		memberID,
	).Return(
		inactiveMember,
		nil,
	)

	// Act
	err := useCase.DeactivateHouseholdMember(
		ctx,
		memberID,
	)

	// Assert
	assert.Equal(
		t,
		coreentity.ErrMemberAlreadyInactive,
		err,
	)
}

func TestHouseholdMemberUseCase_CanDeleteHouseholdMember_ShouldReturnTrueWhenGetHouseholdMemberByIDSucceedsAndCanDeleteReturnsTrue(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"
	expectedMember := &coreentity.HouseholdMember{
		ID:        memberID,
		FirstName: "John",
		LastName:  "Doe",
		Active:    true,
		Role:      "owner",
	}

	householdMembersRepo.EXPECT().GetByID(
		ctx,
		memberID,
	).Return(
		expectedMember,
		nil,
	)

	householdMembersRepo.EXPECT().CanDelete(
		ctx,
		memberID,
	).Return(
		true,
		nil,
	)

	// Act
	result, err := useCase.CanDeleteHouseholdMember(
		ctx,
		memberID,
	)

	// Assert
	assert.NoError(
		t,
		err,
	)
	assert.True(
		t,
		result,
	)
}

func TestHouseholdMemberUseCase_CanDeleteHouseholdMember_ShouldReturnErrorWhenGetHouseholdMemberByIDFailsWithErrRecordNotFound(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"

	householdMembersRepo.EXPECT().GetByID(
		ctx,
		memberID,
	).Return(
		nil,
		port.ErrRecordNotFound,
	)

	// Act
	result, err := useCase.CanDeleteHouseholdMember(
		ctx,
		memberID,
	)

	// Assert
	assert.False(
		t,
		result,
	)
	assert.Equal(
		t,
		coreentity.ErrMemberNotFound,
		err,
	)
}

func TestHouseholdMemberUseCase_CanDeleteHouseholdMember_ShouldReturnFalseWhenGetHouseholdMemberByIDSucceedsButCanDeleteReturnsFalse(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"
	expectedMember := &coreentity.HouseholdMember{
		ID:        memberID,
		FirstName: "John",
		LastName:  "Doe",
		Active:    true,
		Role:      "owner",
	}

	householdMembersRepo.EXPECT().GetByID(
		ctx,
		memberID,
	).Return(
		expectedMember,
		nil,
	)

	householdMembersRepo.EXPECT().CanDelete(
		ctx,
		memberID,
	).Return(
		false,
		nil,
	)

	// Act
	result, err := useCase.CanDeleteHouseholdMember(
		ctx,
		memberID,
	)

	// Assert
	assert.NoError(
		t,
		err,
	)
	assert.False(
		t,
		result,
	)
}

func TestHouseholdMemberUseCase_CanDeleteHouseholdMember_ReturnsErrorWhenGetHouseholdMemberByIDFailsWithDatabaseConnectionError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"
	unexpectedError := errors.New("database connection error")

	householdMembersRepo.EXPECT().GetByID(
		ctx,
		memberID,
	).Return(
		nil,
		unexpectedError,
	)

	// Act
	result, err := useCase.CanDeleteHouseholdMember(
		ctx,
		memberID,
	)

	// Assert
	assert.False(
		t,
		result,
	)
	assert.Equal(
		t,
		unexpectedError,
		err,
	)
}

func TestHouseholdMemberUseCase_CanDeleteHouseholdMember_ReturnsErrorWhenGetHouseholdMemberByIDSucceedsButCanDeleteFailsWithDatabaseConnectionError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"
	expectedMember := &coreentity.HouseholdMember{
		ID:        memberID,
		FirstName: "John",
		LastName:  "Doe",
		Active:    true,
		Role:      "owner",
	}

	unexpectedError := errors.New("database connection error")

	householdMembersRepo.EXPECT().GetByID(
		ctx,
		memberID,
	).Return(
		expectedMember,
		nil,
	)

	householdMembersRepo.EXPECT().CanDelete(
		ctx,
		memberID,
	).Return(
		false,
		unexpectedError,
	)

	// Act
	result, err := useCase.CanDeleteHouseholdMember(
		ctx,
		memberID,
	)

	// Assert
	assert.False(
		t,
		result,
	)
	assert.Equal(
		t,
		unexpectedError,
		err,
	)
}

func TestHouseholdMemberUseCase_ActivateHouseholdMember_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"
	inactiveMember := &coreentity.HouseholdMember{
		ID:        memberID,
		FirstName: "John",
		LastName:  "Doe",
		Active:    false,
		Role:      "owner",
	}

	householdMembersRepo.EXPECT().GetByID(
		ctx,
		memberID,
	).Return(
		inactiveMember,
		nil,
	)

	householdMembersRepo.EXPECT().Activate(
		ctx,
		memberID,
	).Return(
		nil,
	)

	// Act
	err := useCase.ActivateHouseholdMember(
		ctx,
		memberID,
	)

	// Assert
	assert.NoError(
		t,
		err,
	)
}

func TestHouseholdMemberUseCase_ActivateHouseholdMember_ReturnsErrorFromGetHouseholdMemberByIDWhenRepositoryGetByIDOperationFailsWithDatabaseConnectionError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"
	unexpectedError := errors.New("database connection error")

	householdMembersRepo.EXPECT().GetByID(
		ctx,
		memberID,
	).Return(
		nil,
		unexpectedError,
	)

	// Act
	err := useCase.ActivateHouseholdMember(
		ctx,
		memberID,
	)

	// Assert
	assert.Equal(
		t,
		unexpectedError,
		err,
	)
}

func TestHouseholdMemberUseCase_ActivateHouseholdMember_ReturnsErrMemberAlreadyActiveWhenAttemptingToActivateAlreadyActiveMember(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"
	activeMember := &coreentity.HouseholdMember{
		ID:        memberID,
		FirstName: "John",
		LastName:  "Doe",
		Active:    true,
		Role:      "owner",
	}

	householdMembersRepo.EXPECT().GetByID(
		ctx,
		memberID,
	).Return(
		activeMember,
		nil,
	)

	// Act
	err := useCase.ActivateHouseholdMember(
		ctx,
		memberID,
	)

	// Assert
	assert.Equal(
		t,
		coreentity.ErrMemberAlreadyActive,
		err,
	)
}

func TestHouseholdMemberUseCase_ActivateHouseholdMember_ReturnsErrorWhenRepositoryActivateOperationFailsWithDatabaseConnectionErrorAfterSuccessfulMemberRetrieval(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	householdMembersRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewHouseholdMemberUseCase(householdMembersRepo)

	memberID := "member-123"
	inactiveMember := &coreentity.HouseholdMember{
		ID:        memberID,
		FirstName: "John",
		LastName:  "Doe",
		Active:    false,
		Role:      "owner",
	}

	unexpectedError := errors.New("database connection error")

	householdMembersRepo.EXPECT().GetByID(
		ctx,
		memberID,
	).Return(
		inactiveMember,
		nil,
	)

	householdMembersRepo.EXPECT().Activate(
		ctx,
		memberID,
	).Return(
		unexpectedError,
	)

	// Act
	err := useCase.ActivateHouseholdMember(
		ctx,
		memberID,
	)

	// Assert
	assert.Equal(
		t,
		unexpectedError,
		err,
	)
}
