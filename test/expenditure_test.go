package test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/domain/misc"
	"proletariat-budget-core/core/port"
	"proletariat-budget-core/core/usecase"
	"proletariat-budget-core/test/mocks"
	"testing"
)

func TestExpenditureUseCase_Create_ShouldReturnValidationError_WhenExpenditureCategoryNil(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()

	// Create invalid expenditure that will fail validation
	expenditure := coreentity.Expenditure{
		Category: nil, // Invalid - category is required
		Transaction: &coreentity.Transaction{
			AccountID: "account-123",
			Amount:    0, // Invalid - amount should be positive
		},
	}

	expectedError := coreentity.ErrExpenditureCategoryNil

	result, err := useCase.Create(
		ctx,
		expenditure,
		nil,
	)

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
		err,
		expectedError,
	)
}

func TestExpenditureUseCase_Create_ShouldReturnValidationError_WhenExpenditureTransactionNil(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()

	// Create invalid expenditure that will fail validation
	expenditure := coreentity.Expenditure{
		Category:    &coreentity.Category{ID: "1"},
		Transaction: nil,
	}

	expectedError := coreentity.ErrExpenditureTransactionNil

	result, err := useCase.Create(
		ctx,
		expenditure,
		nil,
	)

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
		err,
		expectedError,
	)
}

func TestExpenditureUseCase_Get_ShouldReturnExpenditureSuccessfully_WhenValidIDIsProvidedAndExpenditureExistsInRepository(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	expenditureID := "expenditure-123"
	accountID := "account-123"
	categoryID := "category-123"
	transactionID := "transaction-123"
	amount := float32(100.00)
	balanceAfter := float32(400.00)

	expectedExpenditure := &coreentity.Expenditure{
		ID: expenditureID,
		Category: &coreentity.Category{
			ID:     categoryID,
			Name:   "Test Category",
			Active: true,
		},
		Transaction: &coreentity.Transaction{
			ID:           &transactionID,
			AccountID:    accountID,
			Amount:       amount,
			BalanceAfter: &balanceAfter,
		},
		Tags: &[]*coreentity.Tag{
			{ID: "tag-1", Name: "Test Tag"},
		},
	}

	mockExpenditureRepo.EXPECT().GetByID(
		ctx,
		expenditureID,
	).Return(
		expectedExpenditure,
		nil,
	)

	result, err := useCase.Get(
		ctx,
		expenditureID,
	)

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
		expectedExpenditure,
		result,
	)
}

func TestExpenditureUseCase_Get_ShouldReturnErrExpenditureNotFound_WhenRepositoryReturnsErrRecordNotFoundForGivenID(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	expenditureID := "expenditure-123"

	mockExpenditureRepo.EXPECT().GetByID(
		ctx,
		expenditureID,
	).Return(
		nil,
		port.ErrRecordNotFound,
	)

	result, err := useCase.Get(
		ctx,
		expenditureID,
	)

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
		coreentity.ErrExpenditureNotFound,
		err,
	)
}

func TestExpenditureUseCase_Get_ShouldReturnOriginalRepositoryError_WhenRepositoryGetByIDFailsWithErrorOtherThanErrRecordNotFound(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	expenditureID := "expenditure-123"
	unexpectedError := errors.New("database connection error")

	mockExpenditureRepo.EXPECT().GetByID(
		ctx,
		expenditureID,
	).Return(
		nil,
		unexpectedError,
	)

	result, err := useCase.Get(
		ctx,
		expenditureID,
	)

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
		unexpectedError,
		err,
	)
}

func TestExpenditureUseCase_List_ShouldReturnExpenditureListSuccessfully_WhenRepositoryReturnsValidListWithMultipleExpenditures(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	params := coreentity.ExpenditureListParams{
		ListParams: misc.ListParams{
			Limit:  10,
			Offset: 0,
		},
	}

	expenditures := []coreentity.Expenditure{
		{
			ID: "expenditure-1",
			Category: &coreentity.Category{
				ID:   "category-1",
				Name: "Food",
			},
			Transaction: &coreentity.Transaction{
				ID:     &[]string{"transaction-1"}[0],
				Amount: 50.00,
			},
		},
		{
			ID: "expenditure-2",
			Category: &coreentity.Category{
				ID:   "category-2",
				Name: "Transport",
			},
			Transaction: &coreentity.Transaction{
				ID:     &[]string{"transaction-2"}[0],
				Amount: 25.00,
			},
		},
	}

	expectedList := &coreentity.ExpenditureList{
		Expenditures: expenditures,
		Metadata: misc.ListMetadata{
			Total:  2,
			Offset: 0,
			Limit:  10,
		},
	}

	mockExpenditureRepo.EXPECT().FindExpenditures(
		ctx,
		params,
	).Return(
		expectedList,
		nil,
	)

	result, err := useCase.List(
		ctx,
		params,
	)

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
		expectedList,
		result,
	)
	assert.Len(
		t,
		result.Expenditures,
		2,
	)
	assert.Equal(
		t,
		uint(2),
		result.Metadata.Total,
	)
}
func TestExpenditureUseCase_Create_ShouldReturnError_WhenValidateAccountFailsWithInvalidAccountID(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	accountID := "invalid-account-123"
	categoryID := "category-123"

	expenditure := coreentity.Expenditure{
		Category: &coreentity.Category{ID: categoryID},
		Transaction: &coreentity.Transaction{
			AccountID:   accountID,
			Amount:      100.00,
			Description: "test expenditure",
			Currency:    "test-currency",
		},
	}

	// Mock account validation failure
	mockAccountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		nil,
		port.ErrRecordNotFound,
	)

	result, err := useCase.Create(
		ctx,
		expenditure,
		nil,
	)

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
		coreentity.ErrAccountNotFound,
		err,
	)
}

func TestExpenditureUseCase_Create_ShouldReturnError_WhenValidateCategoryFailsWithInvalidCategoryID(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	accountID := "account-123"
	categoryID := "invalid-category-123"

	account := &coreentity.Account{
		ID:             &accountID,
		CurrentBalance: 500.00,
		Active:         true,
	}

	expenditure := coreentity.Expenditure{
		Category: &coreentity.Category{ID: categoryID},
		Transaction: &coreentity.Transaction{
			AccountID:   accountID,
			Amount:      100.00,
			Description: "test expenditure",
			Currency:    "test-currency",
		},
	}

	// Mock account validation success
	mockAccountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		account,
		nil,
	)

	// Mock category validation failure
	mockCategoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		nil,
		port.ErrRecordNotFound,
	)

	result, err := useCase.Create(
		ctx,
		expenditure,
		nil,
	)

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
		coreentity.ErrCategoryNotFound,
		err,
	)
}

func TestExpenditureUseCase_Create_ShouldSuccessfullyCreateExpenditure_WhenAllValidationsPassAndTransactionCompletes(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)
	mockTxContext := mocks.NewMockTransactionContext(c)
	mockTxAccountRepo := mocks.NewMockAccount(c)
	mockTxTransactionRepo := mocks.NewMockTransaction(c)
	mockTxExpenditureRepo := mocks.NewMockExpenditure(c)
	mockTxTagsRepo := mocks.NewMockTags(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	accountID := "account-123"
	categoryID := "category-123"
	expenditureID := "expenditure-123"
	transactionID := "transaction-123"
	amount := float32(100.00)
	currentBalance := float32(500.00)
	balanceAfter := currentBalance - amount

	account := &coreentity.Account{
		ID:             &accountID,
		CurrentBalance: currentBalance,
		Active:         true,
	}

	category := &coreentity.Category{
		ID:     categoryID,
		Active: true,
	}

	tags := []*coreentity.Tag{
		{ID: "tag-1"},
		{ID: "tag-2"},
	}

	transaction := &coreentity.Transaction{
		AccountID:   accountID,
		Amount:      amount,
		Description: "test expenditure",
		Currency:    "test-currency",
	}

	expenditure := coreentity.Expenditure{
		Category:    &coreentity.Category{ID: categoryID},
		Transaction: transaction,
		Tags:        &tags,
	}

	expectedExpenditure := &coreentity.Expenditure{
		ID:       expenditureID,
		Category: category,
		Transaction: &coreentity.Transaction{
			ID:           &transactionID,
			AccountID:    accountID,
			Amount:       amount,
			BalanceAfter: &balanceAfter,
			Description:  "test expenditure",
			Currency:     "test-currency",
		},
		Tags: &tags,
	}

	// Mock account validation
	mockAccountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		account,
		nil,
	)

	// Mock category validation
	mockCategoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		category,
		nil,
	)

	// Mock transaction manager
	mockTxManager.EXPECT().WithDatabaseTransaction(
		ctx,
		gomock.Any(),
	).DoAndReturn(
		func(
			ctx context.Context,
			fn func(
				context.Context,
				port.TransactionContext,
			) error,
		) error {
			return fn(
				ctx,
				mockTxContext,
			)
		},
	)

	// Mock transaction context repositories
	mockTxContext.EXPECT().GetAccountRepo().Return(mockTxAccountRepo)
	mockTxContext.EXPECT().GetTransactionRepo().Return(mockTxTransactionRepo)
	mockTxContext.EXPECT().GetExpenditureRepo().Return(mockTxExpenditureRepo)
	mockTxContext.EXPECT().GetTagsRepo().Return(mockTxTagsRepo)

	// Mock transaction processing
	mockTxTransactionRepo.EXPECT().Create(
		ctx,
		gomock.Any(),
	).Return(
		transactionID,
		nil,
	)
	mockTxAccountRepo.EXPECT().Update(
		ctx,
		gomock.Any(),
	).Return(nil)

	// Mock expenditure creation
	mockTxExpenditureRepo.EXPECT().Create(
		ctx,
		gomock.Any(),
	).Return(
		expenditureID,
		nil,
	)

	// Mock tag linking
	mockTxTagsRepo.EXPECT().LinkTagsToType(
		ctx,
		expenditureID,
		&tags,
	).Return(nil)

	// Mock final expenditure retrieval
	mockExpenditureRepo.EXPECT().GetByID(
		ctx,
		expenditureID,
	).Return(
		expectedExpenditure,
		nil,
	)

	result, err := useCase.Create(
		ctx,
		expenditure,
		nil,
	)

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
		expectedExpenditure,
		result,
	)
}
func TestExpenditureUseCase_Create_ShouldReturnError_WhenTransactionAmountIsNegativeAndNotRollback(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	accountID := "account-123"
	categoryID := "category-123"
	negativeAmount := float32(-100.00)

	expenditure := coreentity.Expenditure{
		Category: &coreentity.Category{ID: categoryID},
		Transaction: &coreentity.Transaction{
			AccountID:       accountID,
			Amount:          negativeAmount,
			TransactionType: coreentity.TransactionTypeExpenditure,
			Description:     "test expenditure",
			Currency:        "test-currency",
		},
	}

	result, err := useCase.Create(
		ctx,
		expenditure,
		nil,
	)

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
		coreentity.ErrTransactionAmountMustBePositive,
		err,
	)
}

func TestExpenditureUseCase_Create_ShouldReturnError_WhenTransactionDescriptionIsEmpty(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	accountID := "account-123"
	categoryID := "category-123"

	expenditure := coreentity.Expenditure{
		Category: &coreentity.Category{ID: categoryID},
		Transaction: &coreentity.Transaction{
			AccountID:   accountID,
			Amount:      100.00,
			Description: "", // Empty description
			Currency:    "test-currency",
		},
	}

	result, err := useCase.Create(
		ctx,
		expenditure,
		nil,
	)

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
		coreentity.ErrTransactionDescriptionEmpty,
		err,
	)
}

func TestExpenditureUseCase_Create_ShouldReturnError_WhenTransactionCurrencyIsEmpty(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	accountID := "account-123"
	categoryID := "category-123"

	expenditure := coreentity.Expenditure{
		Category: &coreentity.Category{ID: categoryID},
		Transaction: &coreentity.Transaction{
			AccountID:   accountID,
			Amount:      100.00,
			Description: "test expenditure",
			Currency:    "", // Empty currency
		},
	}

	result, err := useCase.Create(
		ctx,
		expenditure,
		nil,
	)

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
		coreentity.ErrTransactionCurrencyEmpty,
		err,
	)
}
func TestExpenditureUseCase_Create_ShouldReturnError_WhenProcessTransactionFailsDuringDatabaseTransaction(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)
	mockTxContext := mocks.NewMockTransactionContext(c)
	mockTxAccountRepo := mocks.NewMockAccount(c)
	mockTxTransactionRepo := mocks.NewMockTransaction(c)
	mockTxExpenditureRepo := mocks.NewMockExpenditure(c)
	mockTxTagsRepo := mocks.NewMockTags(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	accountID := "account-123"
	categoryID := "category-123"
	amount := float32(600.00)
	currentBalance := float32(500.00)

	account := &coreentity.Account{
		ID:             &accountID,
		CurrentBalance: currentBalance,
		Active:         true,
	}

	category := &coreentity.Category{
		ID:     categoryID,
		Active: true,
	}

	expenditure := coreentity.Expenditure{
		Category: &coreentity.Category{ID: categoryID},
		Transaction: &coreentity.Transaction{
			AccountID:   accountID,
			Amount:      amount,
			Description: "test expenditure",
			Currency:    "test-currency",
		},
	}

	// Mock account validation
	mockAccountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		account,
		nil,
	)

	// Mock category validation
	mockCategoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		category,
		nil,
	)

	// Mock transaction manager
	mockTxManager.EXPECT().WithDatabaseTransaction(
		ctx,
		gomock.Any(),
	).DoAndReturn(
		func(
			ctx context.Context,
			fn func(
				context.Context,
				port.TransactionContext,
			) error,
		) error {
			return fn(
				ctx,
				mockTxContext,
			)
		},
	)

	// Mock transaction context repositories
	mockTxContext.EXPECT().GetAccountRepo().Return(mockTxAccountRepo)
	mockTxContext.EXPECT().GetTransactionRepo().Return(mockTxTransactionRepo)
	mockTxContext.EXPECT().GetExpenditureRepo().Return(mockTxExpenditureRepo)
	mockTxContext.EXPECT().GetTagsRepo().Return(mockTxTagsRepo)

	result, err := useCase.Create(
		ctx,
		expenditure,
		nil,
	)

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
		coreentity.ErrInsufficientBalance,
		err,
	)
}

func TestExpenditureUseCase_Create_ShouldReturnError_WhenExpenditureRepositoryCreateOperationFailsDuringDatabaseTransaction(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)
	mockTxContext := mocks.NewMockTransactionContext(c)
	mockTxAccountRepo := mocks.NewMockAccount(c)
	mockTxTransactionRepo := mocks.NewMockTransaction(c)
	mockTxExpenditureRepo := mocks.NewMockExpenditure(c)
	mockTxTagsRepo := mocks.NewMockTags(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	accountID := "account-123"
	categoryID := "category-123"
	transactionID := "transaction-123"
	amount := float32(100.00)
	currentBalance := float32(500.00)

	account := &coreentity.Account{
		ID:             &accountID,
		CurrentBalance: currentBalance,
		Active:         true,
	}

	category := &coreentity.Category{
		ID:     categoryID,
		Active: true,
	}

	transaction := &coreentity.Transaction{
		AccountID:   accountID,
		Amount:      amount,
		Description: "test expenditure",
		Currency:    "test-currency",
	}

	expenditure := coreentity.Expenditure{
		Category:    &coreentity.Category{ID: categoryID},
		Transaction: transaction,
	}

	expectedError := errors.New("expenditure creation failed")

	// Mock account validation
	mockAccountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		account,
		nil,
	)

	// Mock category validation
	mockCategoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		category,
		nil,
	)

	// Mock transaction manager
	mockTxManager.EXPECT().WithDatabaseTransaction(
		ctx,
		gomock.Any(),
	).DoAndReturn(
		func(
			ctx context.Context,
			fn func(
				context.Context,
				port.TransactionContext,
			) error,
		) error {
			return fn(
				ctx,
				mockTxContext,
			)
		},
	)

	// Mock transaction context repositories
	mockTxContext.EXPECT().GetAccountRepo().Return(mockTxAccountRepo)
	mockTxContext.EXPECT().GetTransactionRepo().Return(mockTxTransactionRepo)
	mockTxContext.EXPECT().GetExpenditureRepo().Return(mockTxExpenditureRepo)
	mockTxContext.EXPECT().GetTagsRepo().Return(mockTxTagsRepo)

	// Mock transaction processing
	mockTxTransactionRepo.EXPECT().Create(
		ctx,
		gomock.Any(),
	).Return(
		transactionID,
		nil,
	)
	mockTxAccountRepo.EXPECT().Update(
		ctx,
		gomock.Any(),
	).Return(nil)

	// Mock expenditure creation failure
	mockTxExpenditureRepo.EXPECT().Create(
		ctx,
		gomock.Any(),
	).Return(
		"",
		expectedError,
	)

	result, err := useCase.Create(
		ctx,
		expenditure,
		nil,
	)

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

func TestExpenditureUseCase_Create_ShouldReturnError_WhenLinkTagsFailsDuringDatabaseTransaction(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)
	mockTxContext := mocks.NewMockTransactionContext(c)
	mockTxAccountRepo := mocks.NewMockAccount(c)
	mockTxTransactionRepo := mocks.NewMockTransaction(c)
	mockTxExpenditureRepo := mocks.NewMockExpenditure(c)
	mockTxTagsRepo := mocks.NewMockTags(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	accountID := "account-123"
	categoryID := "category-123"
	expenditureID := "expenditure-123"
	transactionID := "transaction-123"
	amount := float32(100.00)
	currentBalance := float32(500.00)

	account := &coreentity.Account{
		ID:             &accountID,
		CurrentBalance: currentBalance,
		Active:         true,
	}

	category := &coreentity.Category{
		ID:     categoryID,
		Active: true,
	}

	tags := []*coreentity.Tag{
		{ID: "tag-1"},
		{ID: "tag-2"},
	}

	transaction := &coreentity.Transaction{
		AccountID:   accountID,
		Amount:      amount,
		Description: "test expenditure",
		Currency:    "test-currency",
	}

	expenditure := coreentity.Expenditure{
		Category:    &coreentity.Category{ID: categoryID},
		Transaction: transaction,
		Tags:        &tags,
	}

	expectedError := errors.New("tag linking failed")

	// Mock account validation
	mockAccountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		account,
		nil,
	)

	// Mock category validation
	mockCategoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		category,
		nil,
	)

	// Mock transaction manager
	mockTxManager.EXPECT().WithDatabaseTransaction(
		ctx,
		gomock.Any(),
	).DoAndReturn(
		func(
			ctx context.Context,
			fn func(
				context.Context,
				port.TransactionContext,
			) error,
		) error {
			return fn(
				ctx,
				mockTxContext,
			)
		},
	)

	// Mock transaction context repositories
	mockTxContext.EXPECT().GetAccountRepo().Return(mockTxAccountRepo)
	mockTxContext.EXPECT().GetTransactionRepo().Return(mockTxTransactionRepo)
	mockTxContext.EXPECT().GetExpenditureRepo().Return(mockTxExpenditureRepo)
	mockTxContext.EXPECT().GetTagsRepo().Return(mockTxTagsRepo)

	// Mock transaction processing
	mockTxTransactionRepo.EXPECT().Create(
		ctx,
		gomock.Any(),
	).Return(
		transactionID,
		nil,
	)
	mockTxAccountRepo.EXPECT().Update(
		ctx,
		gomock.Any(),
	).Return(nil)

	// Mock expenditure creation
	mockTxExpenditureRepo.EXPECT().Create(
		ctx,
		gomock.Any(),
	).Return(
		expenditureID,
		nil,
	)

	// Mock tag linking failure
	mockTxTagsRepo.EXPECT().LinkTagsToType(
		ctx,
		expenditureID,
		&tags,
	).Return(expectedError)

	result, err := useCase.Create(
		ctx,
		expenditure,
		nil,
	)

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

func TestExpenditureUseCase_Create_ShouldReturnError_WhenTxManagerWithDatabaseTransactionFailsToExecuteTransaction(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	accountID := "account-123"
	categoryID := "category-123"
	amount := float32(100.00)
	currentBalance := float32(500.00)

	account := &coreentity.Account{
		ID:             &accountID,
		CurrentBalance: currentBalance,
		Active:         true,
	}

	category := &coreentity.Category{
		ID:     categoryID,
		Active: true,
	}

	expenditure := coreentity.Expenditure{
		Category: &coreentity.Category{ID: categoryID},
		Transaction: &coreentity.Transaction{
			AccountID:   accountID,
			Amount:      amount,
			Description: "test expenditure",
			Currency:    "test-currency",
		},
	}

	expectedError := errors.New("database transaction failed to execute")

	// Mock account validation
	mockAccountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		account,
		nil,
	)

	// Mock category validation
	mockCategoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		category,
		nil,
	)

	// Mock transaction manager failure
	mockTxManager.EXPECT().WithDatabaseTransaction(
		ctx,
		gomock.Any(),
	).Return(expectedError)

	result, err := useCase.Create(
		ctx,
		expenditure,
		nil,
	)

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

func TestExpenditureUseCase_Create_ShouldReturnError_WhenExpenditureRepositoryGetByIDFailsAfterSuccessfulCreation(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)
	mockTxContext := mocks.NewMockTransactionContext(c)
	mockTxAccountRepo := mocks.NewMockAccount(c)
	mockTxTransactionRepo := mocks.NewMockTransaction(c)
	mockTxExpenditureRepo := mocks.NewMockExpenditure(c)
	mockTxTagsRepo := mocks.NewMockTags(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	accountID := "account-123"
	categoryID := "category-123"
	expenditureID := "expenditure-123"
	transactionID := "transaction-123"
	amount := float32(100.00)
	currentBalance := float32(500.00)

	account := &coreentity.Account{
		ID:             &accountID,
		CurrentBalance: currentBalance,
		Active:         true,
	}

	category := &coreentity.Category{
		ID:     categoryID,
		Active: true,
	}

	transaction := &coreentity.Transaction{
		AccountID:   accountID,
		Amount:      amount,
		Description: "test expenditure",
		Currency:    "test-currency",
	}

	expenditure := coreentity.Expenditure{
		Category:    &coreentity.Category{ID: categoryID},
		Transaction: transaction,
	}

	expectedError := errors.New("expenditure retrieval failed after creation")

	// Mock account validation
	mockAccountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		account,
		nil,
	)

	// Mock category validation
	mockCategoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		category,
		nil,
	)

	// Mock transaction manager
	mockTxManager.EXPECT().WithDatabaseTransaction(
		ctx,
		gomock.Any(),
	).DoAndReturn(
		func(
			ctx context.Context,
			fn func(
				context.Context,
				port.TransactionContext,
			) error,
		) error {
			return fn(
				ctx,
				mockTxContext,
			)
		},
	)

	// Mock transaction context repositories
	mockTxContext.EXPECT().GetAccountRepo().Return(mockTxAccountRepo)
	mockTxContext.EXPECT().GetTransactionRepo().Return(mockTxTransactionRepo)
	mockTxContext.EXPECT().GetExpenditureRepo().Return(mockTxExpenditureRepo)
	mockTxContext.EXPECT().GetTagsRepo().Return(mockTxTagsRepo)

	// Mock transaction processing
	mockTxTransactionRepo.EXPECT().Create(
		ctx,
		gomock.Any(),
	).Return(
		transactionID,
		nil,
	)
	mockTxAccountRepo.EXPECT().Update(
		ctx,
		gomock.Any(),
	).Return(nil)

	// Mock expenditure creation success
	mockTxExpenditureRepo.EXPECT().Create(
		ctx,
		gomock.Any(),
	).Return(
		expenditureID,
		nil,
	)

	// Mock final expenditure retrieval failure
	mockExpenditureRepo.EXPECT().GetByID(
		ctx,
		expenditureID,
	).Return(
		nil,
		expectedError,
	)

	result, err := useCase.Create(
		ctx,
		expenditure,
		nil,
	)

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

func TestExpenditureUseCase_Create_ShouldReturnError_WhenCategoryIsInactive(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	accountID := "account-123"
	categoryID := "category-123"

	account := &coreentity.Account{
		ID:             &accountID,
		CurrentBalance: 500.00,
		Active:         true,
	}

	category := &coreentity.Category{
		ID:     categoryID,
		Active: false, // Inactive category
	}

	expenditure := coreentity.Expenditure{
		Category: &coreentity.Category{ID: categoryID},
		Transaction: &coreentity.Transaction{
			AccountID:   accountID,
			Amount:      100.00,
			Description: "test expenditure",
			Currency:    "test-currency",
		},
	}

	// Mock account validation success
	mockAccountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		account,
		nil,
	)

	// Mock category validation - returns inactive category
	mockCategoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		category,
		nil,
	)

	result, err := useCase.Create(
		ctx,
		expenditure,
		nil,
	)

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
		coreentity.ErrCategoryInactive,
		err,
	)
}

func TestExpenditureUseCase_Create_ShouldReturnError_WhenAccountIsInactive(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	accountID := "account-123"
	categoryID := "category-123"

	account := &coreentity.Account{
		ID:             &accountID,
		CurrentBalance: 500.00,
		Active:         false, // Inactive account
	}

	expenditure := coreentity.Expenditure{
		Category: &coreentity.Category{ID: categoryID},
		Transaction: &coreentity.Transaction{
			AccountID:   accountID,
			Amount:      100.00,
			Description: "test expenditure",
			Currency:    "test-currency",
		},
	}

	// Mock account validation - returns inactive account
	mockAccountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		account,
		nil,
	)

	result, err := useCase.Create(
		ctx,
		expenditure,
		nil,
	)

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
		coreentity.ErrAccountInactive,
		err,
	)
}
func TestExpenditureUseCase_Create_ShouldReturnError_WhenAccountRepoGetByIDFailsWithDatabaseError(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	accountID := "account-123"
	categoryID := "category-123"

	expenditure := coreentity.Expenditure{
		Category: &coreentity.Category{ID: categoryID},
		Transaction: &coreentity.Transaction{
			AccountID:   accountID,
			Amount:      100.00,
			Description: "test expenditure",
			Currency:    "test-currency",
		},
	}

	expectedError := errors.New("database connection error")

	// Mock account validation failure with database error
	mockAccountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		nil,
		expectedError,
	)

	result, err := useCase.Create(
		ctx,
		expenditure,
		nil,
	)

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
func TestExpenditureUseCase_Create_ShouldReturnError_WhenCategoryRepoGetByIDFailsWithDatabaseError(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	accountID := "account-123"
	categoryID := "category-123"

	account := &coreentity.Account{
		ID:             &accountID,
		CurrentBalance: 500.00,
		Active:         true,
	}

	expenditure := coreentity.Expenditure{
		Category: &coreentity.Category{ID: categoryID},
		Transaction: &coreentity.Transaction{
			AccountID:   accountID,
			Amount:      100.00,
			Description: "test expenditure",
			Currency:    "test-currency",
		},
	}

	expectedError := errors.New("database connection error")

	// Mock account validation success
	mockAccountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		account,
		nil,
	)

	// Mock category validation failure with database error
	mockCategoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		nil,
		expectedError,
	)

	result, err := useCase.Create(
		ctx,
		expenditure,
		nil,
	)

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

func TestExpenditureUseCase_Create_ShouldReturnError_WhenTransactionRepoCreateFailsDuringDatabaseTransaction(t *testing.T) {
	c := gomock.NewController(t)
	mockExpenditureRepo := mocks.NewMockExpenditure(c)
	mockAccountRepo := mocks.NewMockAccount(c)
	mockTagsRepo := mocks.NewMockTags(c)
	mockCategoryRepo := mocks.NewMockCategory(c)
	mockTransactionRepo := mocks.NewMockTransaction(c)
	mockTxManager := mocks.NewMockTransactionManager(c)
	mockTxContext := mocks.NewMockTransactionContext(c)
	mockTxAccountRepo := mocks.NewMockAccount(c)
	mockTxTransactionRepo := mocks.NewMockTransaction(c)
	mockTxExpenditureRepo := mocks.NewMockExpenditure(c)
	mockTxTagsRepo := mocks.NewMockTags(c)

	useCase := usecase.NewExpenditureUseCase(
		mockExpenditureRepo,
		mockAccountRepo,
		mockTagsRepo,
		mockCategoryRepo,
		mockTransactionRepo,
		mockTxManager,
	)

	ctx := context.Background()
	accountID := "account-123"
	categoryID := "category-123"
	amount := float32(100.00)
	currentBalance := float32(500.00)

	account := &coreentity.Account{
		ID:             &accountID,
		CurrentBalance: currentBalance,
		Active:         true,
	}

	category := &coreentity.Category{
		ID:     categoryID,
		Active: true,
	}

	transaction := &coreentity.Transaction{
		AccountID:   accountID,
		Amount:      amount,
		Description: "test expenditure",
		Currency:    "test-currency",
	}

	expenditure := coreentity.Expenditure{
		Category:    &coreentity.Category{ID: categoryID},
		Transaction: transaction,
	}

	expectedError := errors.New("transaction creation failed")

	// Mock account validation
	mockAccountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		account,
		nil,
	)

	// Mock category validation
	mockCategoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		category,
		nil,
	)

	// Mock transaction manager
	mockTxManager.EXPECT().WithDatabaseTransaction(
		ctx,
		gomock.Any(),
	).DoAndReturn(
		func(
			ctx context.Context,
			fn func(
				context.Context,
				port.TransactionContext,
			) error,
		) error {
			return fn(
				ctx,
				mockTxContext,
			)
		},
	)

	// Mock transaction context repositories
	mockTxContext.EXPECT().GetAccountRepo().Return(mockTxAccountRepo)
	mockTxContext.EXPECT().GetTransactionRepo().Return(mockTxTransactionRepo)
	mockTxContext.EXPECT().GetExpenditureRepo().Return(mockTxExpenditureRepo)
	mockTxContext.EXPECT().GetTagsRepo().Return(mockTxTagsRepo)

	// Mock transaction creation failure
	mockTxTransactionRepo.EXPECT().Create(
		ctx,
		gomock.Any(),
	).Return(
		"",
		expectedError,
	)

	result, err := useCase.Create(
		ctx,
		expenditure,
		nil,
	)

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
