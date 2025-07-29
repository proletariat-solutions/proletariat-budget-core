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
	"time"
)

func TestIngressUseCase_Create_ReturnsValidationErrorWhenIngressDoesNotHaveTransaction(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	ingress := coreentity.Ingress{
		Transaction: nil, // Missing transaction
		Category: &coreentity.Category{
			ID: "category-123",
		},
	}

	// Act
	result, err := useCase.Create(
		ctx,
		ingress,
		nil,
	)

	// Assert
	assert.Nil(
		t,
		result,
	)
	assert.Equal(
		t,
		coreentity.ErrTransactionRequired,
		err,
	)
}
func TestIngressUseCase_Create_ReturnsValidationErrorWhenIngressDoesNotHaveCategory(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	ingress := coreentity.Ingress{
		Transaction: &coreentity.Transaction{
			AccountID: "account-123",
		},
		Category: nil, // Missing category
	}

	// Act
	result, err := useCase.Create(
		ctx,
		ingress,
		nil,
	)

	// Assert
	assert.Nil(
		t,
		result,
	)
	assert.Equal(
		t,
		coreentity.ErrIngressCategoryRequired,
		err,
	)
}

func TestIngressUseCase_Create_ReturnsValidationErrorWhenIngressDoesNotHaveDate(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	ingress := coreentity.Ingress{
		Transaction: &coreentity.Transaction{
			AccountID: "account-123",
		},
		Category: &coreentity.Category{
			ID: "category-123",
		},
		AuditData: coreentity.AuditData{Date: nil}, // Missing date
	}

	// Act
	result, err := useCase.Create(
		ctx,
		ingress,
		nil,
	)

	// Assert
	assert.Nil(
		t,
		result,
	)
	assert.Equal(
		t,
		coreentity.ErrIngressDateRequired,
		err,
	)
}

func TestIngressUseCase_Create_ReturnsErrorWhenAccountValidationFailsWithInvalidAccountID(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	accountID := "invalid-account-123"
	date := time.Now()
	amount := float32(1000.50)
	description := "Test ingress transaction"
	transactionType := coreentity.TransactionTypeIngress

	ingress := coreentity.Ingress{
		Transaction: &coreentity.Transaction{
			AccountID:       accountID,
			Amount:          amount,
			Description:     description,
			TransactionType: transactionType,
			Currency:        "test-currency",
		},
		Category: &coreentity.Category{
			ID: "category-123",
		},
		AuditData: coreentity.AuditData{
			Date: &date,
		},
	}

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		nil,
		port.ErrRecordNotFound,
	)

	// Act
	result, err := useCase.Create(
		ctx,
		ingress,
		nil,
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
func TestIngressUseCase_Create_ReturnsErrorWhenCategoryValidationFailsWithInvalidCategoryID(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	accountID := "account-123"
	categoryID := "invalid-category-123"
	date := time.Now()
	amount := float32(1000.50)
	description := "Test ingress transaction"
	transactionType := coreentity.TransactionTypeIngress

	ingress := coreentity.Ingress{
		Transaction: &coreentity.Transaction{
			AccountID:       accountID,
			Amount:          amount,
			Description:     description,
			TransactionType: transactionType,
			Currency:        "test-currency",
		},
		Category: &coreentity.Category{
			ID: categoryID,
		},
		AuditData: coreentity.AuditData{
			Date: &date,
		},
	}

	activeAccount := &coreentity.Account{
		ID:             &accountID,
		Active:         true,
		CurrentBalance: 500.25,
	}

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		activeAccount,
		nil,
	)

	categoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		nil,
		port.ErrRecordNotFound,
	)

	// Act
	result, err := useCase.Create(
		ctx,
		ingress,
		nil,
	)

	// Assert
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

func TestIngressUseCase_Create_ReturnsErrorWhenTransactionCreationFailsInDatabase(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	txContext := mocks.NewMockTransactionContext(controller)
	transactionRepo := mocks.NewMockTransaction(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	accountID := "account-123"
	categoryID := "category-123"
	amount := float32(1000.50)
	date := time.Now()
	description := "Test ingress transaction"
	transactionType := coreentity.TransactionTypeIngress

	ingress := coreentity.Ingress{
		Transaction: &coreentity.Transaction{
			AccountID:       accountID,
			Amount:          amount,
			Description:     description,
			TransactionType: transactionType,
			Currency:        "test-currency",
		},
		Category: &coreentity.Category{
			ID: categoryID,
		},
		AuditData: coreentity.AuditData{
			Date: &date,
		},
	}

	activeAccount := &coreentity.Account{
		ID:             &accountID,
		Active:         true,
		CurrentBalance: 500.25,
	}

	activeCategory := &coreentity.Category{
		ID:     categoryID,
		Active: true,
		Name:   "Test Category",
	}

	transactionError := errors.New("database transaction creation failed")

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		activeAccount,
		nil,
	)

	categoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		activeCategory,
		nil,
	)

	txManager.EXPECT().WithDatabaseTransaction(
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
			txContext.EXPECT().GetTransactionRepo().Return(transactionRepo)

			transactionRepo.EXPECT().Create(
				ctx,
				gomock.Any(),
			).Return(
				"",
				transactionError,
			)

			return fn(
				ctx,
				txContext,
			)
		},
	)

	// Act
	result, err := useCase.Create(
		ctx,
		ingress,
		nil,
	)

	// Assert
	assert.Nil(
		t,
		result,
	)
	assert.Equal(
		t,
		transactionError,
		err,
	)
}

func TestIngressUseCase_Create_ReturnsErrorWhenAccountUpdateFailsInDatabase(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	txContext := mocks.NewMockTransactionContext(controller)
	transactionRepo := mocks.NewMockTransaction(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	accountID := "account-123"
	categoryID := "category-123"
	amount := float32(1000.50)
	date := time.Now()
	description := "Test ingress transaction"
	transactionType := coreentity.TransactionTypeIngress

	ingress := coreentity.Ingress{
		Transaction: &coreentity.Transaction{
			AccountID:       accountID,
			Amount:          amount,
			Description:     description,
			TransactionType: transactionType,
			Currency:        "test-currency",
		},
		Category: &coreentity.Category{
			ID: categoryID,
		},
		AuditData: coreentity.AuditData{
			Date: &date,
		},
	}

	activeAccount := &coreentity.Account{
		ID:             &accountID,
		Active:         true,
		CurrentBalance: 500.25,
	}

	activeCategory := &coreentity.Category{
		ID:     categoryID,
		Active: true,
		Name:   "Test Category",
	}

	expectedTransactionID := "transaction-456"
	accountUpdateError := errors.New("failed to update account")

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		activeAccount,
		nil,
	)

	categoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		activeCategory,
		nil,
	)

	txManager.EXPECT().WithDatabaseTransaction(
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
			txContext.EXPECT().GetTransactionRepo().Return(transactionRepo)
			txContext.EXPECT().GetAccountRepo().Return(accountRepo)

			transactionRepo.EXPECT().Create(
				ctx,
				gomock.Any(),
			).Return(
				expectedTransactionID,
				nil,
			)

			accountRepo.EXPECT().Update(
				ctx,
				gomock.Any(),
			).Return(
				accountUpdateError,
			)

			return fn(
				ctx,
				txContext,
			)
		},
	)

	// Act
	result, err := useCase.Create(
		ctx,
		ingress,
		nil,
	)

	// Assert
	assert.Nil(
		t,
		result,
	)
	assert.Equal(
		t,
		accountUpdateError,
		err,
	)
}

func TestIngressUseCase_Create_ReturnsErrorWhenIngressCreationFailsInDatabase(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	txContext := mocks.NewMockTransactionContext(controller)
	transactionRepo := mocks.NewMockTransaction(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	accountID := "account-123"
	categoryID := "category-123"
	amount := float32(1000.50)
	date := time.Now()
	description := "Test ingress transaction"
	transactionType := coreentity.TransactionTypeIngress

	ingress := coreentity.Ingress{
		Transaction: &coreentity.Transaction{
			AccountID:       accountID,
			Amount:          amount,
			Description:     description,
			TransactionType: transactionType,
			Currency:        "test-currency",
		},
		Category: &coreentity.Category{
			ID: categoryID,
		},
		AuditData: coreentity.AuditData{
			Date: &date,
		},
	}

	activeAccount := &coreentity.Account{
		ID:             &accountID,
		Active:         true,
		CurrentBalance: 500.25,
	}

	activeCategory := &coreentity.Category{
		ID:     categoryID,
		Active: true,
		Name:   "Test Category",
	}

	expectedTransactionID := "transaction-456"
	ingressCreationError := errors.New("failed to create ingress")

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		activeAccount,
		nil,
	)

	categoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		activeCategory,
		nil,
	)

	txManager.EXPECT().WithDatabaseTransaction(
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
			txContext.EXPECT().GetTransactionRepo().Return(transactionRepo)
			txContext.EXPECT().GetAccountRepo().Return(accountRepo)
			txContext.EXPECT().GetIngressRepo().Return(ingressRepo)

			transactionRepo.EXPECT().Create(
				ctx,
				gomock.Any(),
			).Return(
				expectedTransactionID,
				nil,
			)

			accountRepo.EXPECT().Update(
				ctx,
				gomock.Any(),
			).Return(
				nil,
			)

			ingressRepo.EXPECT().Create(
				ctx,
				gomock.Any(),
			).Return(
				"",
				ingressCreationError,
			)

			return fn(
				ctx,
				txContext,
			)
		},
	)

	// Act
	result, err := useCase.Create(
		ctx,
		ingress,
		nil,
	)

	// Assert
	assert.Nil(
		t,
		result,
	)
	assert.Equal(
		t,
		ingressCreationError,
		err,
	)
}

func TestIngressUseCase_Create_ReturnsErrorWhenAccountIsInactive(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	accountID := "account-123"
	date := time.Now()
	amount := float32(1000.50)
	description := "Test ingress transaction"
	transactionType := coreentity.TransactionTypeIngress

	ingress := coreentity.Ingress{
		Transaction: &coreentity.Transaction{
			AccountID:       accountID,
			Amount:          amount,
			Description:     description,
			TransactionType: transactionType,
			Currency:        "test-currency",
		},
		Category: &coreentity.Category{
			ID: "category-123",
		},
		AuditData: coreentity.AuditData{
			Date: &date,
		},
	}

	inactiveAccount := &coreentity.Account{
		ID:             &accountID,
		Active:         false,
		CurrentBalance: 500.25,
	}

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		inactiveAccount,
		nil,
	)

	// Act
	result, err := useCase.Create(
		ctx,
		ingress,
		nil,
	)

	// Assert
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

func TestIngressUseCase_Create_ReturnsErrorWhenCategoryIsInactive(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	accountID := "account-123"
	categoryID := "category-123"
	date := time.Now()
	amount := float32(1000.50)
	description := "Test ingress transaction"
	transactionType := coreentity.TransactionTypeIngress

	ingress := coreentity.Ingress{
		Transaction: &coreentity.Transaction{
			AccountID:       accountID,
			Amount:          amount,
			Description:     description,
			TransactionType: transactionType,
			Currency:        "test-currency",
		},
		Category: &coreentity.Category{
			ID: categoryID,
		},
		AuditData: coreentity.AuditData{
			Date: &date,
		},
	}

	activeAccount := &coreentity.Account{
		ID:             &accountID,
		Active:         true,
		CurrentBalance: 500.25,
	}

	inactiveCategory := &coreentity.Category{
		ID:     categoryID,
		Active: false,
		Name:   "Test Category",
	}

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		activeAccount,
		nil,
	)

	categoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		inactiveCategory,
		nil,
	)

	// Act
	result, err := useCase.Create(
		ctx,
		ingress,
		nil,
	)

	// Assert
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

func TestIngressUseCase_Create_SuccessfullyCreatesIngressWithoutRecurrencePattern(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	txContext := mocks.NewMockTransactionContext(controller)
	transactionRepo := mocks.NewMockTransaction(controller)
	tagsRepo := mocks.NewMockTags(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	accountID := "account-123"
	categoryID := "category-123"
	amount := float32(1000.50)
	date := time.Now()
	description := "Test ingress transaction"
	transactionType := coreentity.TransactionTypeIngress

	ingress := coreentity.Ingress{
		Transaction: &coreentity.Transaction{
			AccountID:       accountID,
			Amount:          amount,
			Description:     description,
			TransactionType: transactionType,
			Currency:        "USD",
		},
		Category: &coreentity.Category{
			ID: categoryID,
		},
		AuditData: coreentity.AuditData{
			Date: &date,
		},
	}

	activeAccount := &coreentity.Account{
		ID:             &accountID,
		Active:         true,
		CurrentBalance: 500.25,
	}

	activeCategory := &coreentity.Category{
		ID:     categoryID,
		Active: true,
		Name:   "Test Category",
	}

	expectedTransactionID := "transaction-456"
	expectedIngressID := "ingress-789"

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		activeAccount,
		nil,
	)

	categoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		activeCategory,
		nil,
	)

	txManager.EXPECT().WithDatabaseTransaction(
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
			txContext.EXPECT().GetTransactionRepo().Return(transactionRepo)
			txContext.EXPECT().GetAccountRepo().Return(accountRepo)
			txContext.EXPECT().GetIngressRepo().Return(ingressRepo)
			txContext.EXPECT().GetTagsRepo().Return(tagsRepo)

			transactionRepo.EXPECT().Create(
				ctx,
				gomock.Any(),
			).DoAndReturn(
				func(
					ctx context.Context,
					transaction coreentity.Transaction,
				) (
					string,
					error,
				) {
					// Verify transaction fields are set correctly
					assert.Equal(
						t,
						accountID,
						transaction.AccountID,
					)
					assert.Equal(
						t,
						amount,
						transaction.Amount,
					)
					assert.Equal(
						t,
						description,
						transaction.Description,
					)
					assert.Equal(
						t,
						transactionType,
						transaction.TransactionType,
					)
					assert.Equal(
						t,
						float32(1500.75),
						*transaction.BalanceAfter,
					)
					assert.Equal(
						t,
						coreentity.TransactionStatusCompleted,
						*transaction.Status,
					)
					return expectedTransactionID, nil
				},
			)

			accountRepo.EXPECT().Update(
				ctx,
				gomock.Any(),
			).DoAndReturn(
				func(
					ctx context.Context,
					account coreentity.Account,
				) error {
					// Verify account balance was credited
					assert.Equal(
						t,
						float32(1500.75),
						account.CurrentBalance,
					)
					return nil
				},
			)

			ingressRepo.EXPECT().Create(
				ctx,
				gomock.Any(),
			).DoAndReturn(
				func(
					ctx context.Context,
					ingress coreentity.Ingress,
				) (
					string,
					error,
				) {
					// Verify ingress fields are set correctly
					assert.Equal(
						t,
						expectedTransactionID,
						*ingress.Transaction.ID,
					)
					assert.Equal(
						t,
						activeCategory,
						ingress.Category,
					)
					assert.Nil(
						t,
						ingress.RecurrenceTransactionInfo,
					)
					return expectedIngressID, nil
				},
			)

			return fn(
				ctx,
				txContext,
			)
		},
	)

	// Act
	result, err := useCase.Create(
		ctx,
		ingress,
		nil,
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
		expectedIngressID,
		result.ID,
	)
	assert.Equal(
		t,
		expectedTransactionID,
		*result.Transaction.ID,
	)
	assert.Equal(
		t,
		activeCategory,
		result.Category,
	)
	assert.Equal(
		t,
		amount,
		result.Transaction.Amount,
	)
	assert.Equal(
		t,
		description,
		result.Transaction.Description,
	)
	assert.Equal(
		t,
		transactionType,
		result.Transaction.TransactionType,
	)
	assert.Equal(
		t,
		float32(1500.75),
		*result.Transaction.BalanceAfter,
	)
	assert.Equal(
		t,
		coreentity.TransactionStatusCompleted,
		*result.Transaction.Status,
	)
	assert.Nil(
		t,
		result.RecurrenceTransactionInfo,
	)
}

func TestIngressUseCase_Create_SuccessfullyCreatesIngressWithRecurrencePatternAndSetsRecurrenceTransactionInfo(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	txContext := mocks.NewMockTransactionContext(controller)
	transactionRepo := mocks.NewMockTransaction(controller)
	tagsRepo := mocks.NewMockTags(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	accountID := "account-123"
	categoryID := "category-123"
	amount := float32(1000.50)
	date := time.Now()
	description := "Test recurring ingress transaction"
	transactionType := coreentity.TransactionTypeIngress

	ingress := coreentity.Ingress{
		Transaction: &coreentity.Transaction{
			AccountID:       accountID,
			Amount:          amount,
			Description:     description,
			TransactionType: transactionType,
			Currency:        "USD",
		},
		Category: &coreentity.Category{
			ID: categoryID,
		},
		AuditData: coreentity.AuditData{
			Date: &date,
		},
	}

	recurrencePatternID := "recurrence-pattern-456"
	recurrencePattern := &coreentity.RecurrencePattern{
		ID: recurrencePatternID,
	}

	activeAccount := &coreentity.Account{
		ID:             &accountID,
		Active:         true,
		CurrentBalance: 500.25,
	}

	activeCategory := &coreentity.Category{
		ID:     categoryID,
		Active: true,
		Name:   "Test Category",
	}

	expectedTransactionID := "transaction-789"
	expectedIngressID := "ingress-012"

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		activeAccount,
		nil,
	)

	categoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		activeCategory,
		nil,
	)

	txManager.EXPECT().WithDatabaseTransaction(
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
			txContext.EXPECT().GetTransactionRepo().Return(transactionRepo)
			txContext.EXPECT().GetAccountRepo().Return(accountRepo)
			txContext.EXPECT().GetIngressRepo().Return(ingressRepo)
			txContext.EXPECT().GetTagsRepo().Return(tagsRepo)

			transactionRepo.EXPECT().Create(
				ctx,
				gomock.Any(),
			).Return(
				expectedTransactionID,
				nil,
			)

			accountRepo.EXPECT().Update(
				ctx,
				gomock.Any(),
			).Return(
				nil,
			)

			ingressRepo.EXPECT().Create(
				ctx,
				gomock.Any(),
			).DoAndReturn(
				func(
					ctx context.Context,
					ingress coreentity.Ingress,
				) (
					string,
					error,
				) {
					// Verify RecurrenceTransactionInfo is set correctly
					assert.NotNil(
						t,
						ingress.RecurrenceTransactionInfo,
					)
					assert.Equal(
						t,
						&recurrencePatternID,
						ingress.RecurrenceTransactionInfo.FromRecurrencePatternID,
					)
					assert.False(
						t,
						ingress.RecurrenceTransactionInfo.IsTemplate,
					)
					return expectedIngressID, nil
				},
			)

			return fn(
				ctx,
				txContext,
			)
		},
	)

	// Act
	result, err := useCase.Create(
		ctx,
		ingress,
		recurrencePattern,
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
		expectedIngressID,
		result.ID,
	)
	assert.NotNil(
		t,
		result.RecurrenceTransactionInfo,
	)
	assert.Equal(
		t,
		&recurrencePatternID,
		result.RecurrenceTransactionInfo.FromRecurrencePatternID,
	)
	assert.False(
		t,
		result.RecurrenceTransactionInfo.IsTemplate,
	)
}

func TestIngressUseCase_Create_HandlesTagLinkingFailureGracefullyWithRecurrencePattern(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	txContext := mocks.NewMockTransactionContext(controller)
	transactionRepo := mocks.NewMockTransaction(controller)
	tagsRepo := mocks.NewMockTags(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	accountID := "account-123"
	categoryID := "category-123"
	amount := float32(1000.50)
	date := time.Now()
	description := "Test recurring ingress transaction"
	transactionType := coreentity.TransactionTypeIngress

	tags := []*coreentity.Tag{
		{ID: "tag-1", Name: "Test Tag 1"},
		{ID: "tag-2", Name: "Test Tag 2"},
	}

	ingress := coreentity.Ingress{
		Transaction: &coreentity.Transaction{
			AccountID:       accountID,
			Amount:          amount,
			Description:     description,
			TransactionType: transactionType,
			Currency:        "USD",
		},
		Category: &coreentity.Category{
			ID: categoryID,
		},
		AuditData: coreentity.AuditData{
			Date: &date,
		},
		Tags: &tags,
	}

	recurrencePatternID := "recurrence-pattern-456"
	recurrencePattern := &coreentity.RecurrencePattern{
		ID: recurrencePatternID,
	}

	activeAccount := &coreentity.Account{
		ID:             &accountID,
		Active:         true,
		CurrentBalance: 500.25,
	}

	activeCategory := &coreentity.Category{
		ID:     categoryID,
		Active: true,
		Name:   "Test Category",
	}

	expectedTransactionID := "transaction-789"
	expectedIngressID := "ingress-012"
	tagLinkingError := errors.New("failed to link tags")

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		activeAccount,
		nil,
	)

	categoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		activeCategory,
		nil,
	)

	txManager.EXPECT().WithDatabaseTransaction(
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
			txContext.EXPECT().GetTransactionRepo().Return(transactionRepo)
			txContext.EXPECT().GetAccountRepo().Return(accountRepo)
			txContext.EXPECT().GetIngressRepo().Return(ingressRepo)
			txContext.EXPECT().GetTagsRepo().Return(tagsRepo)

			transactionRepo.EXPECT().Create(
				ctx,
				gomock.Any(),
			).Return(
				expectedTransactionID,
				nil,
			)

			accountRepo.EXPECT().Update(
				ctx,
				gomock.Any(),
			).Return(
				nil,
			)

			ingressRepo.EXPECT().Create(
				ctx,
				gomock.Any(),
			).Return(
				expectedIngressID,
				nil,
			)

			tagsRepo.EXPECT().LinkTagsToType(
				ctx,
				expectedIngressID,
				&tags,
			).Return(
				tagLinkingError,
			)

			return fn(
				ctx,
				txContext,
			)
		},
	)

	// Act
	result, err := useCase.Create(
		ctx,
		ingress,
		recurrencePattern,
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
		expectedIngressID,
		result.ID,
	)
	assert.Equal(
		t,
		expectedTransactionID,
		*result.Transaction.ID,
	)
	assert.NotNil(
		t,
		result.RecurrenceTransactionInfo,
	)
	assert.Equal(
		t,
		&recurrencePatternID,
		result.RecurrenceTransactionInfo.FromRecurrencePatternID,
	)
}

func TestIngressUseCase_Create_ReturnsErrorWhenTagLinkingFailsAndNoRecurrencePatternProvided(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	txContext := mocks.NewMockTransactionContext(controller)
	transactionRepo := mocks.NewMockTransaction(controller)
	tagsRepo := mocks.NewMockTags(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	accountID := "account-123"
	categoryID := "category-123"
	amount := float32(1000.50)
	date := time.Now()
	description := "Test ingress transaction"
	transactionType := coreentity.TransactionTypeIngress

	tags := []*coreentity.Tag{
		{ID: "tag-1", Name: "Test Tag 1"},
		{ID: "tag-2", Name: "Test Tag 2"},
	}

	ingress := coreentity.Ingress{
		Transaction: &coreentity.Transaction{
			AccountID:       accountID,
			Amount:          amount,
			Description:     description,
			TransactionType: transactionType,
			Currency:        "USD",
		},
		Category: &coreentity.Category{
			ID: categoryID,
		},
		AuditData: coreentity.AuditData{
			Date: &date,
		},
		Tags: &tags,
	}

	activeAccount := &coreentity.Account{
		ID:             &accountID,
		Active:         true,
		CurrentBalance: 500.25,
	}

	activeCategory := &coreentity.Category{
		ID:     categoryID,
		Active: true,
		Name:   "Test Category",
	}

	expectedTransactionID := "transaction-789"
	expectedIngressID := "ingress-012"
	tagLinkingError := errors.New("failed to link tags")

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		activeAccount,
		nil,
	)

	categoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		activeCategory,
		nil,
	)

	txManager.EXPECT().WithDatabaseTransaction(
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
			txContext.EXPECT().GetTransactionRepo().Return(transactionRepo)
			txContext.EXPECT().GetAccountRepo().Return(accountRepo)
			txContext.EXPECT().GetIngressRepo().Return(ingressRepo)
			txContext.EXPECT().GetTagsRepo().Return(tagsRepo)

			transactionRepo.EXPECT().Create(
				ctx,
				gomock.Any(),
			).Return(
				expectedTransactionID,
				nil,
			)

			accountRepo.EXPECT().Update(
				ctx,
				gomock.Any(),
			).Return(
				nil,
			)

			ingressRepo.EXPECT().Create(
				ctx,
				gomock.Any(),
			).Return(
				expectedIngressID,
				nil,
			)

			tagsRepo.EXPECT().LinkTagsToType(
				ctx,
				expectedIngressID,
				&tags,
			).Return(
				tagLinkingError,
			)

			return fn(
				ctx,
				txContext,
			)
		},
	)

	// Act
	result, err := useCase.Create(
		ctx,
		ingress,
		nil,
	)

	// Assert
	assert.Nil(
		t,
		result,
	)
	assert.Equal(
		t,
		tagLinkingError,
		err,
	)
}

func TestIngressUseCase_Create_ReturnsErrorWhenAccountRepositoryReturnsUnexpectedErrorDuringValidation(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	accountID := "account-123"
	date := time.Now()
	amount := float32(1000.50)
	description := "Test ingress transaction"
	transactionType := coreentity.TransactionTypeIngress

	ingress := coreentity.Ingress{
		Transaction: &coreentity.Transaction{
			AccountID:       accountID,
			Amount:          amount,
			Description:     description,
			TransactionType: transactionType,
			Currency:        "test-currency",
		},
		Category: &coreentity.Category{
			ID: "category-123",
		},
		AuditData: coreentity.AuditData{
			Date: &date,
		},
	}

	unexpectedError := errors.New("database connection error")

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		nil,
		unexpectedError,
	)

	// Act
	result, err := useCase.Create(
		ctx,
		ingress,
		nil,
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

func TestIngressUseCase_Create_ReturnsErrorWhenCategoryRepositoryReturnsUnexpectedErrorDuringValidation(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	accountID := "account-123"
	categoryID := "category-123"
	date := time.Now()
	amount := float32(1000.50)
	description := "Test ingress transaction"
	transactionType := coreentity.TransactionTypeIngress

	ingress := coreentity.Ingress{
		Transaction: &coreentity.Transaction{
			AccountID:       accountID,
			Amount:          amount,
			Description:     description,
			TransactionType: transactionType,
			Currency:        "test-currency",
		},
		Category: &coreentity.Category{
			ID: categoryID,
		},
		AuditData: coreentity.AuditData{
			Date: &date,
		},
	}

	activeAccount := &coreentity.Account{
		ID:             &accountID,
		Active:         true,
		CurrentBalance: 500.25,
	}

	unexpectedError := errors.New("database connection error")

	accountRepo.EXPECT().GetByID(
		ctx,
		accountID,
	).Return(
		activeAccount,
		nil,
	)

	categoryRepo.EXPECT().GetByID(
		ctx,
		categoryID,
	).Return(
		nil,
		unexpectedError,
	)

	// Act
	result, err := useCase.Create(
		ctx,
		ingress,
		nil,
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

func TestIngressUseCase_GetByID_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	ingressID := "ingress-123"
	expectedIngress := coreentity.Ingress{
		ID: ingressID,
		Transaction: &coreentity.Transaction{
			ID:          &ingressID,
			AccountID:   "account-456",
			Amount:      1000.50,
			Description: "Test ingress",
		},
		Category: &coreentity.Category{
			ID:   "category-789",
			Name: "Test Category",
		},
	}

	ingressRepo.EXPECT().GetByID(
		ctx,
		ingressID,
	).Return(
		expectedIngress,
		nil,
	)

	// Act
	result, err := useCase.GetByID(
		ctx,
		ingressID,
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
		&expectedIngress,
		result,
	)
}

func TestIngressUseCase_GetByID_ReturnsErrIngressNotFound(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	ingressID := "ingress-123"

	ingressRepo.EXPECT().GetByID(
		ctx,
		ingressID,
	).Return(
		coreentity.Ingress{},
		port.ErrRecordNotFound,
	)

	// Act
	result, err := useCase.GetByID(
		ctx,
		ingressID,
	)

	// Assert
	assert.Nil(
		t,
		result,
	)
	assert.Equal(
		t,
		coreentity.ErrIngressNotFound,
		err,
	)
}

func TestIngressUseCase_GetByID_UnexpectedError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	ingressID := "ingress-123"
	unexpectedError := errors.New("database connection error")

	ingressRepo.EXPECT().GetByID(
		ctx,
		ingressID,
	).Return(
		coreentity.Ingress{},
		unexpectedError,
	)

	// Act
	result, err := useCase.GetByID(
		ctx,
		ingressID,
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

func TestIngressUseCase_List_SuccessfullyReturnsIngressListWhenRepositoryReturnsValidData(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	params := coreentity.IngressListParams{
		ListParams: misc.ListParams{
			Offset: 0,
			Limit:  10,
		},
	}

	expectedIngressList := coreentity.IngressList{
		Ingresses: []coreentity.Ingress{
			{
				ID: "ingress-1",
				Transaction: &coreentity.Transaction{
					ID:          func() *string { s := "transaction-1"; return &s }(),
					AccountID:   "account-1",
					Amount:      1000.50,
					Description: "Test ingress 1",
				},
				Category: &coreentity.Category{
					ID:   "category-1",
					Name: "Test Category 1",
				},
			},
			{
				ID: "ingress-2",
				Transaction: &coreentity.Transaction{
					ID:          func() *string { s := "transaction-2"; return &s }(),
					AccountID:   "account-2",
					Amount:      500.25,
					Description: "Test ingress 2",
				},
				Category: &coreentity.Category{
					ID:   "category-2",
					Name: "Test Category 2",
				},
			},
		},
		Metadata: misc.ListMetadata{
			Total:  2,
			Limit:  10,
			Offset: 0,
		},
	}

	ingressRepo.EXPECT().List(
		ctx,
		params,
	).Return(
		expectedIngressList,
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
		&expectedIngressList,
		result,
	)
}

func TestIngressUseCase_List_ReturnsErrorWhenRepositoryListMethodFailsWithDatabaseConnectionError(t *testing.T) {
	// Arrange
	ctx := context.Background()
	controller := gomock.NewController(t)
	accountRepo := mocks.NewMockAccount(controller)
	categoryRepo := mocks.NewMockCategory(controller)
	ingressRepo := mocks.NewMockIngress(controller)
	txManager := mocks.NewMockTransactionManager(controller)
	useCase := usecase.NewIngressUseCase(
		accountRepo,
		categoryRepo,
		ingressRepo,
		txManager,
	)

	params := coreentity.IngressListParams{
		ListParams: misc.ListParams{
			Offset: 0,
			Limit:  10,
		},
	}

	databaseError := errors.New("database connection error")

	ingressRepo.EXPECT().List(
		ctx,
		params,
	).Return(
		coreentity.IngressList{},
		databaseError,
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
		databaseError,
		err,
	)
}
