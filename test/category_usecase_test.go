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

func TestCategoryUseCase_ListCategories_Success_AllCategories(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	expectedCategories := []coreentity.Category{
		{
			ID:              "1",
			Name:            "Food",
			CategoryType:    coreentity.CategoryTypeExpenditure,
			Active:          true,
			Color:           "#FF0000",
			BackgroundColor: "#00FF00",
		},
		{
			ID:              "2",
			Name:            "Salary",
			CategoryType:    coreentity.CategoryTypeIngress,
			Color:           "#FF0000",
			BackgroundColor: "#00FF00",
			Active:          true,
		},
	}

	mockRepo.EXPECT().
		List(ctx).
		Return(
			expectedCategories,
			nil,
		).
		Times(1)

	categories, err := uc.ListCategories(
		ctx,
		nil,
	)

	assert.NoError(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedCategories,
		categories,
	)
}
func TestCategoryUseCase_ListCategories_Success_EmptyList_NilCategoryType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	var expectedCategories []coreentity.Category

	mockRepo.EXPECT().
		List(ctx).
		Return(
			expectedCategories,
			nil,
		).
		Times(1)

	categories, err := uc.ListCategories(
		ctx,
		nil,
	)

	assert.NoError(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedCategories,
		categories,
	)
	assert.Len(
		t,
		categories,
		0,
	)
}
func TestCategoryUseCase_ListCategories_Error_NilCategoryType_RepositoryFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	expectedError := errors.New("repository error")

	mockRepo.EXPECT().
		List(ctx).
		Return(
			nil,
			expectedError,
		).
		Times(1)

	categories, err := uc.ListCategories(
		ctx,
		nil,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedError,
		err,
	)
	assert.Nil(
		t,
		categories,
	)
}
func TestCategoryUseCase_ListCategories_Success_FilteredByType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	categoryType := coreentity.CategoryTypeExpenditure
	expectedCategories := []coreentity.Category{
		{
			ID:              "1",
			Name:            "Food",
			CategoryType:    coreentity.CategoryTypeExpenditure,
			Active:          true,
			Color:           "#FF0000",
			BackgroundColor: "#00FF00",
		},
		{
			ID:              "2",
			Name:            "Transport",
			CategoryType:    coreentity.CategoryTypeExpenditure,
			Active:          true,
			Color:           "#0000FF",
			BackgroundColor: "#FFFF00",
		},
	}

	mockRepo.EXPECT().
		FindByType(
			ctx,
			categoryType,
		).
		Return(
			expectedCategories,
			nil,
		).
		Times(1)

	categories, err := uc.ListCategories(
		ctx,
		&categoryType,
	)

	assert.NoError(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedCategories,
		categories,
	)
	assert.Len(
		t,
		categories,
		2,
	)
}
func TestCategoryUseCase_ListCategories_Error_WithCategoryType_RepositoryFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	categoryType := coreentity.CategoryTypeExpenditure
	expectedError := errors.New("repository error")

	mockRepo.EXPECT().
		FindByType(
			ctx,
			categoryType,
		).
		Return(
			nil,
			expectedError,
		).
		Times(1)

	categories, err := uc.ListCategories(
		ctx,
		&categoryType,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedError,
		err,
	)
	assert.Nil(
		t,
		categories,
	)
}
func TestCategoryUseCase_GetCategory_Error_CategoryNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	categoryID := "non-existent-id"

	mockRepo.EXPECT().
		GetByID(
			ctx,
			categoryID,
		).
		Return(
			nil,
			port.ErrRecordNotFound,
		).
		Times(1)

	category, err := uc.GetCategory(
		ctx,
		categoryID,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		coreentity.ErrCategoryNotFound,
		err,
	)
	assert.Nil(
		t,
		category,
	)
}

func TestCategoryUseCase_GetCategory_Error_OtherRepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	categoryID := "test-id"
	expectedError := errors.New("database connection error")

	mockRepo.EXPECT().
		GetByID(
			ctx,
			categoryID,
		).
		Return(
			nil,
			expectedError,
		).
		Times(1)

	category, err := uc.GetCategory(
		ctx,
		categoryID,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedError,
		err,
	)
	assert.Nil(
		t,
		category,
	)
}

func TestCategoryUseCase_CreateCategory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	inputCategory := coreentity.Category{
		Name:            "Food",
		CategoryType:    coreentity.CategoryTypeExpenditure,
		Active:          true,
		Color:           "#FF0000",
		BackgroundColor: "#00FF00",
	}
	expectedID := "generated-id"
	expectedCategory := &coreentity.Category{
		ID:              expectedID,
		Name:            "Food",
		CategoryType:    coreentity.CategoryTypeExpenditure,
		Active:          true,
		Color:           "#FF0000",
		BackgroundColor: "#00FF00",
	}

	mockRepo.EXPECT().
		Create(
			ctx,
			inputCategory,
		).
		Return(
			expectedID,
			nil,
		).
		Times(1)

	mockRepo.EXPECT().
		GetByID(
			ctx,
			expectedID,
		).
		Return(
			expectedCategory,
			nil,
		).
		Times(1)

	createdCategory, err := uc.CreateCategory(
		ctx,
		inputCategory,
	)

	assert.NoError(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedCategory,
		createdCategory,
	)
}
func TestCategoryUseCase_CreateCategory_Error_CreateRepositoryFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	inputCategory := coreentity.Category{
		Name:            "Food",
		CategoryType:    coreentity.CategoryTypeExpenditure,
		Active:          true,
		Color:           "#FF0000",
		BackgroundColor: "#00FF00",
	}
	expectedError := errors.New("repository create error")

	mockRepo.EXPECT().
		Create(
			ctx,
			inputCategory,
		).
		Return(
			"",
			expectedError,
		).
		Times(1)

	createdCategory, err := uc.CreateCategory(
		ctx,
		inputCategory,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedError,
		err,
	)
	assert.Nil(
		t,
		createdCategory,
	)
}

func TestCategoryUseCase_CreateCategory_Error_GetCategoryFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	inputCategory := coreentity.Category{
		Name:            "Food",
		CategoryType:    coreentity.CategoryTypeExpenditure,
		Active:          true,
		Color:           "#FF0000",
		BackgroundColor: "#00FF00",
	}
	expectedID := "generated-id"
	expectedError := errors.New("database connection error")

	mockRepo.EXPECT().
		Create(
			ctx,
			inputCategory,
		).
		Return(
			expectedID,
			nil,
		).
		Times(1)

	mockRepo.EXPECT().
		GetByID(
			ctx,
			expectedID,
		).
		Return(
			nil,
			expectedError,
		).
		Times(1)

	createdCategory, err := uc.CreateCategory(
		ctx,
		inputCategory,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedError,
		err,
	)
	assert.Nil(
		t,
		createdCategory,
	)
}

func TestCategoryUseCase_UpdateCategory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	inputCategory := coreentity.Category{
		ID:              "test-id",
		Name:            "Updated Food",
		CategoryType:    coreentity.CategoryTypeExpenditure,
		Active:          true,
		Color:           "#FF0000",
		BackgroundColor: "#00FF00",
	}
	existingCategory := &coreentity.Category{
		ID:              "test-id",
		Name:            "Food",
		CategoryType:    coreentity.CategoryTypeExpenditure,
		Active:          true,
		Color:           "#0000FF",
		BackgroundColor: "#FFFF00",
	}
	updatedCategory := &coreentity.Category{
		ID:              "test-id",
		Name:            "Updated Food",
		CategoryType:    coreentity.CategoryTypeExpenditure,
		Active:          true,
		Color:           "#FF0000",
		BackgroundColor: "#00FF00",
	}

	// First GetCategory call to check if category exists
	mockRepo.EXPECT().
		GetByID(
			ctx,
			inputCategory.ID,
		).
		Return(
			existingCategory,
			nil,
		).
		Times(1)

	// Update operation
	mockRepo.EXPECT().
		Update(
			ctx,
			inputCategory,
		).
		Return(nil).
		Times(1)

	// Second GetCategory call to return updated category
	mockRepo.EXPECT().
		GetByID(
			ctx,
			inputCategory.ID,
		).
		Return(
			updatedCategory,
			nil,
		).
		Times(1)

	result, err := uc.UpdateCategory(
		ctx,
		inputCategory,
	)

	assert.NoError(
		t,
		err,
	)
	assert.Equal(
		t,
		updatedCategory,
		result,
	)
}

func TestCategoryUseCase_UpdateCategory_Error_GetCategoryFailsDuringInitialCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	inputCategory := coreentity.Category{
		ID:              "test-id",
		Name:            "Updated Food",
		CategoryType:    coreentity.CategoryTypeExpenditure,
		Active:          true,
		Color:           "#FF0000",
		BackgroundColor: "#00FF00",
	}
	expectedError := errors.New("database connection error")

	mockRepo.EXPECT().
		GetByID(
			ctx,
			inputCategory.ID,
		).
		Return(
			nil,
			expectedError,
		).
		Times(1)

	result, err := uc.UpdateCategory(
		ctx,
		inputCategory,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedError,
		err,
	)
	assert.Nil(
		t,
		result,
	)
}

func TestCategoryUseCase_UpdateCategory_Error_UpdateRepositoryFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	inputCategory := coreentity.Category{
		ID:              "test-id",
		Name:            "Updated Food",
		CategoryType:    coreentity.CategoryTypeExpenditure,
		Active:          true,
		Color:           "#FF0000",
		BackgroundColor: "#00FF00",
	}
	existingCategory := &coreentity.Category{
		ID:              "test-id",
		Name:            "Food",
		CategoryType:    coreentity.CategoryTypeExpenditure,
		Active:          true,
		Color:           "#0000FF",
		BackgroundColor: "#FFFF00",
	}
	expectedError := errors.New("repository update error")

	// First GetByID call to check if category exists
	mockRepo.EXPECT().
		GetByID(
			ctx,
			inputCategory.ID,
		).
		Return(
			existingCategory,
			nil,
		).
		Times(1)

	// Update operation fails
	mockRepo.EXPECT().
		Update(
			ctx,
			inputCategory,
		).
		Return(expectedError).
		Times(1)

	result, err := uc.UpdateCategory(
		ctx,
		inputCategory,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedError,
		err,
	)
	assert.Nil(
		t,
		result,
	)
}

func TestCategoryUseCase_UpdateCategory_Error_GetCategoryFailsAfterUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	inputCategory := coreentity.Category{
		ID:              "test-id",
		Name:            "Updated Food",
		CategoryType:    coreentity.CategoryTypeExpenditure,
		Active:          true,
		Color:           "#FF0000",
		BackgroundColor: "#00FF00",
	}
	existingCategory := &coreentity.Category{
		ID:              "test-id",
		Name:            "Food",
		CategoryType:    coreentity.CategoryTypeExpenditure,
		Active:          true,
		Color:           "#0000FF",
		BackgroundColor: "#FFFF00",
	}
	expectedError := errors.New("database connection error")

	// First GetByID call to check if category exists
	mockRepo.EXPECT().
		GetByID(
			ctx,
			inputCategory.ID,
		).
		Return(
			existingCategory,
			nil,
		).
		Times(1)

	// Update operation succeeds
	mockRepo.EXPECT().
		Update(
			ctx,
			inputCategory,
		).
		Return(nil).
		Times(1)

	// Second GetByID call fails after successful update
	mockRepo.EXPECT().
		GetByID(
			ctx,
			inputCategory.ID,
		).
		Return(
			nil,
			expectedError,
		).
		Times(1)

	result, err := uc.UpdateCategory(
		ctx,
		inputCategory,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedError,
		err,
	)
	assert.Nil(
		t,
		result,
	)
}
func TestCategoryUseCase_DeleteCategory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	categoryID := "test-id"

	mockRepo.EXPECT().
		Delete(
			ctx,
			categoryID,
		).
		Return(nil).
		Times(1)

	err := uc.DeleteCategory(
		ctx,
		categoryID,
	)

	assert.NoError(
		t,
		err,
	)
}

func TestCategoryUseCase_DeleteCategory_Error_CategoryNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	categoryID := "non-existent-id"

	mockRepo.EXPECT().
		Delete(
			ctx,
			categoryID,
		).
		Return(port.ErrRecordNotFound).
		Times(1)

	err := uc.DeleteCategory(
		ctx,
		categoryID,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		coreentity.ErrCategoryNotFound,
		err,
	)
}

func TestCategoryUseCase_DeleteCategory_Error_OtherRepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	categoryID := "test-id"
	expectedError := errors.New("database connection error")

	mockRepo.EXPECT().
		Delete(
			ctx,
			categoryID,
		).
		Return(expectedError).
		Times(1)

	err := uc.DeleteCategory(
		ctx,
		categoryID,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedError,
		err,
	)
}

func TestCategoryUseCase_Activate_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	categoryID := "test-id"
	category := &coreentity.Category{
		ID:              categoryID,
		Name:            "Food",
		CategoryType:    coreentity.CategoryTypeExpenditure,
		Active:          false,
		Color:           "#FF0000",
		BackgroundColor: "#00FF00",
	}

	expectedCategory := *category
	expectedCategory.Active = true

	mockRepo.EXPECT().
		GetByID(
			ctx,
			categoryID,
		).
		Return(
			category,
			nil,
		).
		Times(1)

	mockRepo.EXPECT().
		Update(
			ctx,
			expectedCategory,
		).
		Return(nil).
		Times(1)

	err := uc.Activate(
		ctx,
		categoryID,
	)

	assert.NoError(
		t,
		err,
	)
}

func TestCategoryUseCase_Activate_Error_GetByIDFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	categoryID := "test-id"
	expectedError := errors.New("database connection error")

	mockRepo.EXPECT().
		GetByID(
			ctx,
			categoryID,
		).
		Return(
			nil,
			expectedError,
		).
		Times(1)

	err := uc.Activate(
		ctx,
		categoryID,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedError,
		err,
	)
}

func TestCategoryUseCase_Activate_Error_CategoryNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	categoryID := "non-existent-id"

	mockRepo.EXPECT().
		GetByID(
			ctx,
			categoryID,
		).
		Return(
			nil,
			port.ErrRecordNotFound,
		).
		Times(1)

	err := uc.Activate(
		ctx,
		categoryID,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		port.ErrRecordNotFound,
		err,
	)
}

func TestCategoryUseCase_Activate_Error_UpdateRepositoryFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	categoryID := "test-id"
	category := &coreentity.Category{
		ID:              categoryID,
		Name:            "Food",
		CategoryType:    coreentity.CategoryTypeExpenditure,
		Active:          false,
		Color:           "#FF0000",
		BackgroundColor: "#00FF00",
	}
	expectedError := errors.New("repository update error")

	expectedCategory := *category
	expectedCategory.Active = true

	mockRepo.EXPECT().
		GetByID(
			ctx,
			categoryID,
		).
		Return(
			category,
			nil,
		).
		Times(1)

	mockRepo.EXPECT().
		Update(
			ctx,
			expectedCategory,
		).
		Return(expectedError).
		Times(1)

	err := uc.Activate(
		ctx,
		categoryID,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedError,
		err,
	)
}

func TestCategoryUseCase_Activate_Success_AlreadyActiveCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	categoryID := "test-id"
	expectedError := coreentity.ErrCategoryAlreadyActive
	category := &coreentity.Category{
		ID:              categoryID,
		Name:            "Food",
		CategoryType:    coreentity.CategoryTypeExpenditure,
		Active:          true,
		Color:           "#FF0000",
		BackgroundColor: "#00FF00",
	}

	mockRepo.EXPECT().
		GetByID(
			ctx,
			categoryID,
		).
		Return(
			category,
			nil,
		).
		Times(1)

	err := uc.Activate(
		ctx,
		categoryID,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedError,
		err,
	)
}
func TestCategoryUseCase_Deactivate_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	categoryID := "test-id"
	category := &coreentity.Category{
		ID:              categoryID,
		Name:            "Food",
		CategoryType:    coreentity.CategoryTypeExpenditure,
		Active:          true,
		Color:           "#FF0000",
		BackgroundColor: "#00FF00",
	}

	expectedCategory := *category
	expectedCategory.Active = false

	mockRepo.EXPECT().
		GetByID(
			ctx,
			categoryID,
		).
		Return(
			category,
			nil,
		).
		Times(1)

	mockRepo.EXPECT().
		Update(
			ctx,
			expectedCategory,
		).
		Return(nil).
		Times(1)

	err := uc.Deactivate(
		ctx,
		categoryID,
	)

	assert.NoError(
		t,
		err,
	)
}

func TestCategoryUseCase_Deactivate_Error_GetByIDFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	categoryID := "test-id"
	expectedError := errors.New("database connection error")

	mockRepo.EXPECT().
		GetByID(
			ctx,
			categoryID,
		).
		Return(
			nil,
			expectedError,
		).
		Times(1)

	err := uc.Deactivate(
		ctx,
		categoryID,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedError,
		err,
	)
}

func TestCategoryUseCase_Deactivate_Error_CategoryNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	categoryID := "non-existent-id"

	mockRepo.EXPECT().
		GetByID(
			ctx,
			categoryID,
		).
		Return(
			nil,
			port.ErrRecordNotFound,
		).
		Times(1)

	err := uc.Deactivate(
		ctx,
		categoryID,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		port.ErrRecordNotFound,
		err,
	)
}

func TestCategoryUseCase_Deactivate_Error_UpdateRepositoryFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	categoryID := "test-id"
	category := &coreentity.Category{
		ID:              categoryID,
		Name:            "Food",
		CategoryType:    coreentity.CategoryTypeExpenditure,
		Active:          true,
		Color:           "#FF0000",
		BackgroundColor: "#00FF00",
	}
	expectedError := errors.New("repository update error")

	expectedCategory := *category
	expectedCategory.Active = false

	mockRepo.EXPECT().
		GetByID(
			ctx,
			categoryID,
		).
		Return(
			category,
			nil,
		).
		Times(1)

	mockRepo.EXPECT().
		Update(
			ctx,
			expectedCategory,
		).
		Return(expectedError).
		Times(1)

	err := uc.Deactivate(
		ctx,
		categoryID,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedError,
		err,
	)
}

func TestCategoryUseCase_Deactivate_Error_AlreadyInactiveCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCategory(ctrl)
	uc := usecase.NewCategoryUseCase(mockRepo)

	ctx := context.Background()
	categoryID := "test-id"
	expectedError := coreentity.ErrCategoryAlreadyInactive
	category := &coreentity.Category{
		ID:              categoryID,
		Name:            "Food",
		CategoryType:    coreentity.CategoryTypeExpenditure,
		Active:          false,
		Color:           "#FF0000",
		BackgroundColor: "#00FF00",
	}

	mockRepo.EXPECT().
		GetByID(
			ctx,
			categoryID,
		).
		Return(
			category,
			nil,
		).
		Times(1)

	err := uc.Deactivate(
		ctx,
		categoryID,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		expectedError,
		err,
	)
}
