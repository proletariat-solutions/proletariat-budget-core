package test

import (
	"context"
	"errors"
	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/domain/misc"
	"proletariat-budget-core/core/port"
	"proletariat-budget-core/core/usecase"
	"proletariat-budget-core/test/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAccountUseCase_Create_MemberNotFound(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	ownerID := "member-123"
	account := coreentity.Account{
		Name: "Savings Account",
		Owner: &coreentity.HouseholdMember{
			ID: ownerID,
		},
		Currency: &coreentity.Currency{
			ID:     "currency-123",
			Name:   "United States Dollar",
			Symbol: "$",
		},
	}

	householdMemberRepo.EXPECT().GetByID(
		ctx,
		ownerID,
	).Return(
		nil,
		port.ErrRecordNotFound,
	)

	// Act
	result, err := useCase.Create(
		ctx,
		account,
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

func TestAccountUseCase_Create_UnexpectedError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	ownerID := "member-123"
	account := coreentity.Account{
		Name: "Savings Account",
		Owner: &coreentity.HouseholdMember{
			ID: ownerID,
		},
		Currency: &coreentity.Currency{
			ID:     "currency-123",
			Name:   "United States Dollar",
			Symbol: "$",
		},
	}

	unexpectedError := errors.New("database connection error")

	householdMemberRepo.EXPECT().GetByID(
		ctx,
		ownerID,
	).Return(
		nil,
		unexpectedError,
	)

	// Act
	result, err := useCase.Create(
		ctx,
		account,
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
func TestAccountUseCase_Create_MemberInactive(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	ownerID := "member-123"
	account := coreentity.Account{
		Name: "Savings Account",
		Owner: &coreentity.HouseholdMember{
			ID: ownerID,
		},
		Currency: &coreentity.Currency{
			ID:     "currency-123",
			Name:   "United States Dollar",
			Symbol: "$",
		},
	}

	inactiveMember := &coreentity.HouseholdMember{
		Active: false,
	}

	householdMemberRepo.EXPECT().GetByID(
		ctx,
		ownerID,
	).Return(
		inactiveMember,
		nil,
	)

	// Act
	result, err := useCase.Create(
		ctx,
		account,
	)

	// Assert
	assert.Nil(
		t,
		result,
	)
	assert.Equal(
		t,
		coreentity.ErrMemberInactive,
		err,
	)
}

func TestAccountUseCase_Create_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	ownerID := "member-123"
	account := coreentity.Account{
		Name: "Savings Account",
		Owner: &coreentity.HouseholdMember{
			ID: ownerID,
		},
		Currency: &coreentity.Currency{
			ID:     "currency-123",
			Name:   "United States Dollar",
			Symbol: "$",
		},
	}

	activeMember := &coreentity.HouseholdMember{
		Active: true,
	}

	expectedID := "account-456"

	householdMemberRepo.EXPECT().GetByID(
		ctx,
		ownerID,
	).Return(
		activeMember,
		nil,
	)

	accountRepo.EXPECT().Create(
		ctx,
		gomock.Any(),
	).Return(
		&expectedID,
		nil,
	)

	// Act
	result, err := useCase.Create(
		ctx,
		account,
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
		*result,
	)
}

func TestAccountUseCase_Create_ForeignKeyUnknownConstraint(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	ownerID := "member-123"
	account := coreentity.Account{
		Name: "Savings Account",
		Owner: &coreentity.HouseholdMember{
			ID: ownerID,
		},
		Currency: &coreentity.Currency{
			ID:     "currency-123",
			Name:   "United States Dollar",
			Symbol: "$",
		},
	}

	activeMember := &coreentity.HouseholdMember{
		Active: true,
	}

	foreignKeyError := errors.New("foreign key constraint failed: unknown_table")

	householdMemberRepo.EXPECT().GetByID(
		ctx,
		ownerID,
	).Return(
		activeMember,
		nil,
	)

	accountRepo.EXPECT().Create(
		ctx,
		gomock.Any(),
	).Return(
		nil,
		foreignKeyError,
	)

	// Act
	result, err := useCase.Create(
		ctx,
		account,
	)

	// Assert
	assert.Nil(
		t,
		result,
	)
	assert.Equal(
		t,
		foreignKeyError,
		err,
	)
}

func TestAccountUseCase_Create_RepoCreateError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	ownerID := "member-123"
	account := coreentity.Account{
		Name: "Savings Account",
		Owner: &coreentity.HouseholdMember{
			ID: ownerID,
		},
		Currency: &coreentity.Currency{
			ID:     "currency-123",
			Name:   "United States Dollar",
			Symbol: "$",
		},
	}

	activeMember := &coreentity.HouseholdMember{
		Active: true,
	}

	unexpectedError := errors.New("database connection error")

	householdMemberRepo.EXPECT().GetByID(
		ctx,
		ownerID,
	).Return(
		activeMember,
		nil,
	)

	accountRepo.EXPECT().Create(
		ctx,
		gomock.Any(),
	).Return(
		nil,
		unexpectedError,
	)

	// Act
	result, err := useCase.Create(
		ctx,
		account,
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

func TestAccountUseCase_Create_SetsAccountOwner(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	ownerID := "member-123"
	account := coreentity.Account{
		Name: "Savings Account",
		Owner: &coreentity.HouseholdMember{
			ID: ownerID,
		},
		Currency: &coreentity.Currency{
			ID:     "currency-123",
			Name:   "United States Dollar",
			Symbol: "$",
		},
	}

	activeMember := &coreentity.HouseholdMember{
		ID:        "member-123",
		Active:    true,
		FirstName: "John",
		LastName:  "Doe",
		Role:      "owner",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	expectedID := "account-456"

	householdMemberRepo.EXPECT().GetByID(
		ctx,
		ownerID,
	).Return(
		activeMember,
		nil,
	)

	accountRepo.EXPECT().Create(
		ctx,
		gomock.Any(),
	).DoAndReturn(
		func(
			ctx context.Context,
			acc coreentity.Account,
		) (
			*string,
			error,
		) {
			// Verify that the account.Owner is set to the retrieved household member
			assert.Equal(
				t,
				activeMember,
				acc.Owner,
			)
			return &expectedID, nil
		},
	)

	// Act
	result, err := useCase.Create(
		ctx,
		account,
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
		*result,
	)
}
func TestAccountUseCase_Create_PassesCorrectContextAndAccount(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	ownerID := "member-123"
	account := coreentity.Account{
		Owner: &coreentity.HouseholdMember{
			ID: ownerID,
		},
		Name: "Test Account",
		Currency: &coreentity.Currency{
			ID:     "currency-123",
			Name:   "United States Dollar",
			Symbol: "$",
		},
	}

	activeMember := &coreentity.HouseholdMember{
		ID:     "member-123",
		Active: true,
	}

	expectedID := "account-456"

	householdMemberRepo.EXPECT().GetByID(
		ctx,
		ownerID,
	).Return(
		activeMember,
		nil,
	)

	accountRepo.EXPECT().Create(
		ctx,
		gomock.Any(),
	).DoAndReturn(
		func(
			passedCtx context.Context,
			passedAccount coreentity.Account,
		) (
			*string,
			error,
		) {
			// Verify correct context is passed
			assert.Equal(
				t,
				ctx,
				passedCtx,
			)

			// Verify correct account is passed with owner set
			assert.Equal(
				t,
				account.Name,
				passedAccount.Name,
			)
			assert.Equal(
				t,
				account.Owner.ID,
				passedAccount.Owner.ID,
			)
			assert.Equal(
				t,
				activeMember,
				passedAccount.Owner,
			)

			return &expectedID, nil
		},
	)

	// Act
	result, err := useCase.Create(
		ctx,
		account,
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
		*result,
	)
}

func TestAccountUseCase_Create_ReturnsCorrectID(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	ownerID := "member-123"
	account := coreentity.Account{
		Owner: &coreentity.HouseholdMember{
			ID: ownerID,
		},
		Currency: &coreentity.Currency{
			ID:     "currency-123",
			Name:   "United States Dollar",
			Symbol: "$",
		},
		Name: "Test Account",
	}

	activeMember := &coreentity.HouseholdMember{
		ID:     "member-123",
		Active: true,
	}

	expectedID := "account-789"

	householdMemberRepo.EXPECT().GetByID(
		ctx,
		ownerID,
	).Return(
		activeMember,
		nil,
	)

	accountRepo.EXPECT().Create(
		ctx,
		gomock.Any(),
	).Return(
		&expectedID,
		nil,
	)

	// Act
	result, err := useCase.Create(
		ctx,
		account,
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
		*result,
	)
}
func TestAccountUseCase_GetByID_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	expectedAccount := &coreentity.Account{
		ID:   &accountID,
		Name: "Test Account",
	}

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		expectedAccount,
		nil,
	)

	// Act
	result, err := useCase.GetByID(
		ctx,
		accountID,
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
		expectedAccount,
		result,
	)
}
func TestAccountUseCase_GetByID_ReturnsErrAccountNotFound(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		nil,
		port.ErrRecordNotFound,
	)

	// Act
	result, err := useCase.GetByID(
		ctx,
		accountID,
	)

	// Assert
	assert.Nil(
		t,
		result,
	)
	assert.Equal(
		t,
		coreentity.ErrAccountNotFound,
		err,
	)
}

func TestAccountUseCase_GetByID_UnexpectedError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	unexpectedError := errors.New("database connection error")

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		nil,
		unexpectedError,
	)

	// Act
	result, err := useCase.GetByID(
		ctx,
		accountID,
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
func TestAccountUseCase_GetByID_ReturnsAccountWithAllFields(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	expectedAccount := &coreentity.Account{
		ID:             &accountID,
		Name:           "Test Account",
		CurrentBalance: 1000.50,
		InitialBalance: 500.25,
		Active:         true,
		Type:           coreentity.AccountTypeBank,
		Owner: &coreentity.HouseholdMember{
			ID:        "member-456",
			FirstName: "John",
			LastName:  "Doe",
			Active:    true,
		},
	}

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		expectedAccount,
		nil,
	)

	// Act
	result, err := useCase.GetByID(
		ctx,
		accountID,
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
		expectedAccount.ID,
		result.ID,
	)
	assert.Equal(
		t,
		expectedAccount.Name,
		result.Name,
	)
	assert.Equal(
		t,
		expectedAccount.Owner.ID,
		result.Owner.ID,
	)
	assert.Equal(
		t,
		expectedAccount.CurrentBalance,
		result.CurrentBalance,
	)
	assert.Equal(
		t,
		expectedAccount.InitialBalance,
		result.InitialBalance,
	)
	assert.Equal(
		t,
		expectedAccount.Type,
		result.Type,
	)
	assert.Equal(
		t,
		expectedAccount.Active,
		result.Active,
	)
	assert.Equal(
		t,
		expectedAccount.Owner,
		result.Owner,
	)
}
func TestAccountUseCase_Update_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	account := coreentity.Account{
		ID:   &accountID,
		Name: "Updated Account",
		Currency: &coreentity.Currency{
			ID:     "currency-456",
			Name:   "United States Dollar",
			Symbol: "$",
		},
		Owner: &coreentity.HouseholdMember{
			ID:        "member-456",
			FirstName: "John",
			LastName:  "Doe",
			Active:    true,
		},
	}

	updatedAccount := &coreentity.Account{
		ID:             &accountID,
		Name:           "Updated Account",
		CurrentBalance: 1500.75,
		Active:         true,
		Currency: &coreentity.Currency{
			ID:     "currency-456",
			Name:   "United States Dollar",
			Symbol: "$",
		},
		Owner: &coreentity.HouseholdMember{
			ID:        "member-456",
			FirstName: "John",
			LastName:  "Doe",
			Active:    true,
		},
	}

	accountRepo.EXPECT().Update(
		ctx,
		account,
	).Return(
		nil,
	)

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		updatedAccount,
		nil,
	)

	// Act
	result, err := useCase.Update(
		ctx,
		account,
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
		updatedAccount,
		result,
	)
}
func TestAccountUseCase_Update_GetByIDFailsWithErrRecordNotFoundAfterSuccessfulUpdate(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	account := coreentity.Account{
		ID:   &accountID,
		Name: "Updated Account",
		Currency: &coreentity.Currency{
			ID:     "currency-456",
			Name:   "United States Dollar",
			Symbol: "$",
		},
		Owner: &coreentity.HouseholdMember{
			ID:        "member-456",
			FirstName: "John",
			LastName:  "Doe",
			Active:    true,
		},
	}

	accountRepo.EXPECT().Update(
		ctx,
		account,
	).Return(
		nil,
	)

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		nil,
		port.ErrRecordNotFound,
	)

	// Act
	result, err := useCase.Update(
		ctx,
		account,
	)

	// Assert
	assert.Nil(
		t,
		result,
	)
	assert.Equal(
		t,
		coreentity.ErrAccountNotFound,
		err,
	)
}

func TestAccountUseCase_Update_GetByIDFailsWithUnexpectedErrorAfterSuccessfulUpdate(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	account := coreentity.Account{
		ID:   &accountID,
		Name: "Updated Account",
		Currency: &coreentity.Currency{
			ID:     "currency-456",
			Name:   "United States Dollar",
			Symbol: "$",
		},
		Owner: &coreentity.HouseholdMember{
			ID:        "member-456",
			FirstName: "John",
			LastName:  "Doe",
			Active:    true,
		},
	}

	unexpectedError := errors.New("database connection error")

	accountRepo.EXPECT().Update(
		ctx,
		account,
	).Return(
		nil,
	)

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		nil,
		unexpectedError,
	)

	// Act
	result, err := useCase.Update(
		ctx,
		account,
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
func TestAccountUseCase_Update_ReturnsOriginalErrorWhenUpdateFailsWithNonForeignKeyViolation(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	account := coreentity.Account{
		ID:   &accountID,
		Name: "Updated Account",
		Currency: &coreentity.Currency{
			ID:     "currency-456",
			Name:   "United States Dollar",
			Symbol: "$",
		},
		Owner: &coreentity.HouseholdMember{
			ID:        "member-456",
			FirstName: "John",
			LastName:  "Doe",
			Active:    true,
		},
	}

	unexpectedError := errors.New("database connection error")

	accountRepo.EXPECT().Update(
		ctx,
		account,
	).Return(
		unexpectedError,
	)

	// Act
	result, err := useCase.Update(
		ctx,
		account,
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

func TestAccountUseCase_Deactivate_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	activeAccount := &coreentity.Account{
		ID:     &accountID,
		Name:   "Test Account",
		Active: true,
	}

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		activeAccount,
		nil,
	)

	accountRepo.EXPECT().Update(
		ctx,
		gomock.Any(),
	).DoAndReturn(
		func(
			ctx context.Context,
			acc coreentity.Account,
		) error {
			// Verify that the account is set to inactive
			assert.False(
				t,
				acc.Active,
			)
			return nil
		},
	)

	// Act
	err := useCase.Deactivate(
		ctx,
		accountID,
	)

	// Assert
	assert.NoError(
		t,
		err,
	)
}
func TestAccountUseCase_Deactivate_ReturnsErrAccountNotFound(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		nil,
		port.ErrRecordNotFound,
	)

	// Act
	err := useCase.Deactivate(
		ctx,
		accountID,
	)

	// Assert
	assert.Equal(
		t,
		coreentity.ErrAccountNotFound,
		err,
	)
}
func TestAccountUseCase_Deactivate_ReturnsUnexpectedError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	unexpectedError := errors.New("database connection error")

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		nil,
		unexpectedError,
	)

	// Act
	err := useCase.Deactivate(
		ctx,
		accountID,
	)

	// Assert
	assert.Equal(
		t,
		unexpectedError,
		err,
	)
}
func TestAccountUseCase_Deactivate_ReturnsErrorFromSetInactive(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	account := &coreentity.Account{
		ID:     &accountID,
		Name:   "Test Account",
		Active: false, // Already inactive to trigger SetInactive error
	}

	setInactiveError := errors.New("account is already inactive")

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		account,
		nil,
	)

	// Act
	err := useCase.Deactivate(
		ctx,
		accountID,
	)

	// Assert
	assert.Equal(
		t,
		setInactiveError,
		err,
	)
}
func TestAccountUseCase_Deactivate_ReturnsErrorFromAccountRepoUpdate(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	activeAccount := &coreentity.Account{
		ID:     &accountID,
		Name:   "Test Account",
		Active: true,
	}

	updateError := errors.New("failed to update account")

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		activeAccount,
		nil,
	)

	accountRepo.EXPECT().Update(
		ctx,
		gomock.Any(),
	).Return(
		updateError,
	)

	// Act
	err := useCase.Deactivate(
		ctx,
		accountID,
	)

	// Assert
	assert.Equal(
		t,
		updateError,
		err,
	)
}
func TestAccountUseCase_Deactivate_PassesDeactivatedAccountToUpdate(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	activeAccount := &coreentity.Account{
		ID:     &accountID,
		Name:   "Test Account",
		Active: true,
	}

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		activeAccount,
		nil,
	)

	accountRepo.EXPECT().Update(
		ctx,
		gomock.Any(),
	).DoAndReturn(
		func(
			ctx context.Context,
			acc coreentity.Account,
		) error {
			// Verify that the account passed to Update has Active=false
			assert.False(
				t,
				acc.Active,
			)
			assert.Equal(
				t,
				accountID,
				*acc.ID,
			)
			assert.Equal(
				t,
				"Test Account",
				acc.Name,
			)
			return nil
		},
	)

	// Act
	err := useCase.Deactivate(
		ctx,
		accountID,
	)

	// Assert
	assert.NoError(
		t,
		err,
	)
}
func TestAccountUseCase_Deactivate_SetInactiveSucceedsButUpdateFails(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	activeAccount := &coreentity.Account{
		ID:     &accountID,
		Name:   "Test Account",
		Active: true,
	}

	updateError := errors.New("database update failed")

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		activeAccount,
		nil,
	)

	accountRepo.EXPECT().Update(
		ctx,
		gomock.Any(),
	).DoAndReturn(
		func(
			ctx context.Context,
			acc coreentity.Account,
		) error {
			// Verify that the account is deactivated before update fails
			assert.False(
				t,
				acc.Active,
			)
			return updateError
		},
	)

	// Act
	err := useCase.Deactivate(
		ctx,
		accountID,
	)

	// Assert
	assert.Equal(
		t,
		updateError,
		err,
	)
}
func TestAccountUseCase_Deactivate_CallsSetInactiveOnRetrievedAccount(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	activeAccount := &coreentity.Account{
		ID:     &accountID,
		Name:   "Test Account",
		Active: true,
	}

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		activeAccount,
		nil,
	)

	accountRepo.EXPECT().Update(
		ctx,
		gomock.Any(),
	).DoAndReturn(
		func(
			ctx context.Context,
			acc coreentity.Account,
		) error {
			// Verify that SetInactive was called by checking the Active field is false
			assert.False(
				t,
				acc.Active,
			)
			// Verify it's the same account instance by checking other fields
			assert.Equal(
				t,
				accountID,
				*acc.ID,
			)
			assert.Equal(
				t,
				"Test Account",
				acc.Name,
			)
			return nil
		},
	)

	// Act
	err := useCase.Deactivate(
		ctx,
		accountID,
	)

	// Assert
	assert.NoError(
		t,
		err,
	)
}

func TestAccountUseCase_Activate_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	inactiveAccount := &coreentity.Account{
		ID:     &accountID,
		Name:   "Test Account",
		Active: false,
	}

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		inactiveAccount,
		nil,
	)

	accountRepo.EXPECT().Update(
		ctx,
		gomock.Any(),
	).DoAndReturn(
		func(
			ctx context.Context,
			acc coreentity.Account,
		) error {
			// Verify that the account is set to active
			assert.True(
				t,
				acc.Active,
			)
			return nil
		},
	)

	// Act
	err := useCase.Activate(
		ctx,
		accountID,
	)

	// Assert
	assert.NoError(
		t,
		err,
	)
}

func TestAccountUseCase_Activate_ReturnsErrAccountNotFound(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		nil,
		port.ErrRecordNotFound,
	)

	// Act
	err := useCase.Activate(
		ctx,
		accountID,
	)

	// Assert
	assert.Equal(
		t,
		coreentity.ErrAccountNotFound,
		err,
	)
}

func TestAccountUseCase_Activate_ReturnsUnexpectedError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	unexpectedError := errors.New("database connection error")

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		nil,
		unexpectedError,
	)

	// Act
	err := useCase.Activate(
		ctx,
		accountID,
	)

	// Assert
	assert.Equal(
		t,
		unexpectedError,
		err,
	)
}
func TestAccountUseCase_Activate_ReturnsErrorFromSetActive(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	account := &coreentity.Account{
		ID:     &accountID,
		Name:   "Test Account",
		Active: true, // Already active to trigger SetActive error
	}

	setActiveError := errors.New("account is already active")

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		account,
		nil,
	)

	// Act
	err := useCase.Activate(
		ctx,
		accountID,
	)

	// Assert
	assert.Equal(
		t,
		setActiveError,
		err,
	)
}

func TestAccountUseCase_Activate_ReturnsErrorFromAccountRepoUpdate(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	inactiveAccount := &coreentity.Account{
		ID:     &accountID,
		Name:   "Test Account",
		Active: false,
	}

	updateError := errors.New("failed to update account")

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		inactiveAccount,
		nil,
	)

	accountRepo.EXPECT().Update(
		ctx,
		gomock.Any(),
	).Return(
		updateError,
	)

	// Act
	err := useCase.Activate(
		ctx,
		accountID,
	)

	// Assert
	assert.Equal(
		t,
		updateError,
		err,
	)
}

func TestAccountUseCase_Activate_CallsSetActiveOnRetrievedAccount(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	inactiveAccount := &coreentity.Account{
		ID:     &accountID,
		Name:   "Test Account",
		Active: false,
	}

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		inactiveAccount,
		nil,
	)

	accountRepo.EXPECT().Update(
		ctx,
		gomock.Any(),
	).DoAndReturn(
		func(
			ctx context.Context,
			acc coreentity.Account,
		) error {
			// Verify that SetActive was called by checking the Active field is true
			assert.True(
				t,
				acc.Active,
			)
			// Verify it's the same account instance by checking other fields
			assert.Equal(
				t,
				accountID,
				*acc.ID,
			)
			assert.Equal(
				t,
				"Test Account",
				acc.Name,
			)
			return nil
		},
	)

	// Act
	err := useCase.Activate(
		ctx,
		accountID,
	)

	// Assert
	assert.NoError(
		t,
		err,
	)
}

func TestAccountUseCase_Activate_SetActiveSucceedsButUpdateFails(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	inactiveAccount := &coreentity.Account{
		ID:     &accountID,
		Name:   "Test Account",
		Active: false,
	}

	updateError := errors.New("database update failed")

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		inactiveAccount,
		nil,
	)

	accountRepo.EXPECT().Update(
		ctx,
		gomock.Any(),
	).DoAndReturn(
		func(
			ctx context.Context,
			acc coreentity.Account,
		) error {
			// Verify that the account is activated before update fails
			assert.True(
				t,
				acc.Active,
			)
			return updateError
		},
	)

	// Act
	err := useCase.Activate(
		ctx,
		accountID,
	)

	// Assert
	assert.Equal(
		t,
		updateError,
		err,
	)
}

func TestAccountUseCase_Delete_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"

	accountRepo.EXPECT().HasTransactions(
		ctx,
		accountID,
	).Return(
		false,
		nil,
	)

	accountRepo.EXPECT().Delete(
		ctx,
		accountID,
	).Return(
		nil,
	)

	// Act
	err := useCase.Delete(
		ctx,
		accountID,
	)

	// Assert
	assert.NoError(
		t,
		err,
	)
}

func TestAccountUseCase_Delete_ReturnsErrAccountHasTransactions(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"

	accountRepo.EXPECT().HasTransactions(
		ctx,
		accountID,
	).Return(
		true,
		nil,
	)

	// Act
	err := useCase.Delete(
		ctx,
		accountID,
	)

	// Assert
	assert.Equal(
		t,
		coreentity.ErrAccountHasTransactions,
		err,
	)
}

func TestAccountUseCase_Delete_ReturnsErrorFromHasTransactions(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	unexpectedError := errors.New("database connection error")

	accountRepo.EXPECT().HasTransactions(
		ctx,
		accountID,
	).Return(
		false,
		unexpectedError,
	)

	// Act
	err := useCase.Delete(
		ctx,
		accountID,
	)

	// Assert
	assert.Equal(
		t,
		unexpectedError,
		err,
	)
}

func TestAccountUseCase_Delete_ReturnsErrAccountNotFound(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"

	accountRepo.EXPECT().HasTransactions(
		ctx,
		accountID,
	).Return(
		false,
		nil,
	)

	accountRepo.EXPECT().Delete(
		ctx,
		accountID,
	).Return(
		port.ErrRecordNotFound,
	)

	// Act
	err := useCase.Delete(
		ctx,
		accountID,
	)

	// Assert
	assert.Equal(
		t,
		coreentity.ErrAccountNotFound,
		err,
	)
}

func TestAccountUseCase_Delete_ReturnsOriginalErrorFromDelete(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	unexpectedError := errors.New("database connection error")

	accountRepo.EXPECT().HasTransactions(
		ctx,
		accountID,
	).Return(
		false,
		nil,
	)

	accountRepo.EXPECT().Delete(
		ctx,
		accountID,
	).Return(
		unexpectedError,
	)

	// Act
	err := useCase.Delete(
		ctx,
		accountID,
	)

	// Assert
	assert.Equal(
		t,
		unexpectedError,
		err,
	)
}
func TestAccountUseCase_List_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)
	params := coreentity.AccountListParams{
		ListParams: misc.ListParams{
			Limit:  10,
			Offset: 0,
		},
	}

	expectedAccounts := &coreentity.AccountList{
		Accounts: []coreentity.Account{
			{
				ID:             stringPtr("account-1"),
				Name:           "Test Account 1",
				CurrentBalance: 1000.50,
				Active:         true,
			},
			{
				ID:             stringPtr("account-2"),
				Name:           "Test Account 2",
				CurrentBalance: 2000.75,
				Active:         true,
			},
		},
		Metadata: misc.ListMetadata{
			Total:  2,
			Limit:  10,
			Offset: 0,
		},
	}

	accountRepo.EXPECT().List(
		ctx,
		params,
	).Return(
		expectedAccounts,
		nil,
	)

	// Act
	result, err := useCase.List(
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
		uint(10),
		result.Metadata.Limit,
	)
	assert.Equal(
		t,
		uint(0),
		result.Metadata.Offset,
	)
	assert.Equal(
		t,
		uint(2),
		result.Metadata.Total,
	)
	assert.Equal(
		t,
		expectedAccounts,
		result,
	)
}

func TestAccountUseCase_List_ReturnsErrorFromRepo(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	params := coreentity.AccountListParams{
		ListParams: misc.ListParams{
			Limit:  10,
			Offset: 0,
		},
	}

	unexpectedError := errors.New("database connection error")

	accountRepo.EXPECT().List(
		ctx,
		params,
	).Return(
		nil,
		unexpectedError,
	)

	// Act
	result, err := useCase.List(
		ctx,
		params,
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

func TestAccountUseCase_HasTransactions_ReturnsTrue(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"

	accountRepo.EXPECT().HasTransactions(
		ctx,
		accountID,
	).Return(
		true,
		nil,
	)

	// Act
	result, err := useCase.HasTransactions(
		ctx,
		accountID,
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
func TestAccountUseCase_HasTransactions_ReturnsFalse(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"

	accountRepo.EXPECT().HasTransactions(
		ctx,
		accountID,
	).Return(
		false,
		nil,
	)

	// Act
	result, err := useCase.HasTransactions(
		ctx,
		accountID,
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

func TestAccountUseCase_HasTransactions_ReturnsErrorFromRepo(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	accountID := "account-123"
	unexpectedError := errors.New("database connection error")

	accountRepo.EXPECT().HasTransactions(
		ctx,
		accountID,
	).Return(
		false,
		unexpectedError,
	)

	// Act
	result, err := useCase.HasTransactions(
		ctx,
		accountID,
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
func TestAccountUseCase_Create_ReturnsValidationError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	account := coreentity.Account{
		// Invalid account - missing required fields to trigger validation error
		Name: "", // Empty name should trigger validation error
	}

	validationError := errors.New("account name is required")

	// Act
	result, err := useCase.Create(
		ctx,
		account,
	)

	// Assert
	assert.Nil(
		t,
		result,
	)
	assert.Equal(
		t,
		validationError,
		err,
	)
}

func TestAccountUseCase_Create_EmptyOwnerID(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	householdMemberRepo := mocks.NewMockHouseholdMember(controller)
	useCase := usecase.NewAccountUseCase(
		accountRepo,
		householdMemberRepo,
	)

	account := coreentity.Account{
		Owner: &coreentity.HouseholdMember{
			ID: "", // Empty owner ID
		},
		Currency: &coreentity.Currency{
			ID:     "currency-123",
			Name:   "United States Dollar",
			Symbol: "$",
		},
		Name: "Test Account",
	}

	// Act
	result, err := useCase.Create(
		ctx,
		account,
	)

	// Assert
	assert.Nil(
		t,
		result,
	)
	assert.Equal(
		t,
		coreentity.ErrAccountOwnerRequired,
		err,
	)
}
func stringPtr(s string) *string {
	return &s
}
