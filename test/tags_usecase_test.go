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

func TestTagsUseCase_ListTags_Success_AllTags_WhenTagTypeIsNil(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTags(ctrl)
	uc := usecase.NewTagsUseCase(mockRepo)

	ctx := context.Background()
	expectedTags := &[]*coreentity.Tag{
		{
			ID:      "tag-1",
			Name:    "Food",
			TagType: coreentity.TagTypeExpenditure,
		},
		{
			ID:      "tag-2",
			Name:    "Salary",
			TagType: coreentity.TagTypeIngress,
		},
	}

	mockRepo.EXPECT().
		List(ctx).
		Return(
			expectedTags,
			nil,
		).
		Times(1)

	tags, err := uc.ListTags(
		ctx,
		nil,
	)

	assert.NoError(
		t,
		err,
	)
	assert.Equal(
		t,
		*expectedTags,
		tags,
	)
	assert.Len(
		t,
		tags,
		2,
	)
}

func TestTagsUseCase_ListTags_Error_WhenTagTypeIsNilAndRepositoryListMethodFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTags(ctrl)
	uc := usecase.NewTagsUseCase(mockRepo)

	ctx := context.Background()
	expectedError := errors.New("repository list error")

	mockRepo.EXPECT().
		List(ctx).
		Return(
			nil,
			expectedError,
		).
		Times(1)

	tags, err := uc.ListTags(
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
		tags,
	)
}

func TestTagsUseCase_ListTags_Success_FilteredByType_WhenTagTypeIsProvidedAndRepositoryCallSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTags(ctrl)
	uc := usecase.NewTagsUseCase(mockRepo)

	ctx := context.Background()
	tagType := coreentity.TagTypeExpenditure
	expectedTags := &[]*coreentity.Tag{
		{
			ID:      "tag-1",
			Name:    "Food",
			TagType: coreentity.TagTypeExpenditure,
		},
		{
			ID:      "tag-2",
			Name:    "Transport",
			TagType: coreentity.TagTypeExpenditure,
		},
	}

	mockRepo.EXPECT().
		ListByType(
			ctx,
			tagType,
			nil,
		).
		Return(
			expectedTags,
			nil,
		).
		Times(1)

	tags, err := uc.ListTags(
		ctx,
		&tagType,
	)

	assert.NoError(
		t,
		err,
	)
	assert.Equal(
		t,
		*expectedTags,
		tags,
	)
	assert.Len(
		t,
		tags,
		2,
	)
}

func TestTagsUseCase_ListTags_Error_WhenTagTypeIsProvidedAndRepositoryListByTypeMethodFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTags(ctrl)
	uc := usecase.NewTagsUseCase(mockRepo)

	ctx := context.Background()
	tagType := coreentity.TagTypeExpenditure
	expectedError := errors.New("repository listByType error")

	mockRepo.EXPECT().
		ListByType(
			ctx,
			tagType,
			nil,
		).
		Return(
			nil,
			expectedError,
		).
		Times(1)

	tags, err := uc.ListTags(
		ctx,
		&tagType,
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
		tags,
	)
}

func TestTagsUseCase_CreateTag_Success_WhenValidationPassesAndRepositoryCreateSucceeds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTags(ctrl)
	uc := usecase.NewTagsUseCase(mockRepo)

	ctx := context.Background()
	inputTag := &coreentity.Tag{
		Name:    "Food",
		TagType: coreentity.TagTypeExpenditure,
	}
	expectedID := "generated-tag-id"

	mockRepo.EXPECT().
		Create(
			ctx,
			*inputTag,
		).
		Return(
			expectedID,
			nil,
		).
		Times(1)

	result, err := uc.CreateTag(
		ctx,
		inputTag,
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
		expectedID,
		result.ID,
	)
	assert.Equal(
		t,
		inputTag.Name,
		result.Name,
	)
	assert.Equal(
		t,
		inputTag.TagType,
		result.TagType,
	)
}

func TestTagsUseCase_CreateTag_Error_WhenTagValidationFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTags(ctrl)
	uc := usecase.NewTagsUseCase(mockRepo)

	ctx := context.Background()
	inputTag := &coreentity.Tag{
		Name:    "", // Invalid empty name to trigger validation error
		TagType: coreentity.TagTypeExpenditure,
	}
	expectedError := coreentity.ErrTagNameEmpty

	// No repository call should be made since validation fails first

	result, err := uc.CreateTag(
		ctx,
		inputTag,
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

func TestTagsUseCase_CreateTag_Error_WhenRepositoryReturnsErrDuplicateKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTags(ctrl)
	uc := usecase.NewTagsUseCase(mockRepo)

	ctx := context.Background()
	inputTag := &coreentity.Tag{
		Name:    "Food",
		TagType: coreentity.TagTypeExpenditure,
	}

	mockRepo.EXPECT().
		Create(
			ctx,
			*inputTag,
		).
		Return(
			"",
			port.ErrDuplicateKey,
		).
		Times(1)

	result, err := uc.CreateTag(
		ctx,
		inputTag,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		coreentity.ErrTagAlreadyExists,
		err,
	)
	assert.Nil(
		t,
		result,
	)
}

func TestTagsUseCase_CreateTag_Error_WhenRepositoryCreateFailsWithNonDuplicateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTags(ctrl)
	uc := usecase.NewTagsUseCase(mockRepo)

	ctx := context.Background()
	inputTag := &coreentity.Tag{
		Name:    "Food",
		TagType: coreentity.TagTypeExpenditure,
	}
	expectedError := errors.New("database connection error")

	mockRepo.EXPECT().
		Create(
			ctx,
			*inputTag,
		).
		Return(
			"",
			expectedError,
		).
		Times(1)

	result, err := uc.CreateTag(
		ctx,
		inputTag,
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

func TestTagsUseCase_DeleteTag_Success_WhenRepositoryDeleteMethodReturnsNoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTags(ctrl)
	uc := usecase.NewTagsUseCase(mockRepo)

	ctx := context.Background()
	tagID := "tag-123"

	mockRepo.EXPECT().
		Delete(
			ctx,
			tagID,
		).
		Return(nil).
		Times(1)

	err := uc.DeleteTag(
		ctx,
		tagID,
	)

	assert.NoError(
		t,
		err,
	)
}

func TestTagsUseCase_DeleteTag_Error_WhenRepositoryDeleteMethodReturnsErrRecordNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTags(ctrl)
	uc := usecase.NewTagsUseCase(mockRepo)

	ctx := context.Background()
	tagID := "tag-123"

	mockRepo.EXPECT().
		Delete(
			ctx,
			tagID,
		).
		Return(port.ErrRecordNotFound).
		Times(1)

	err := uc.DeleteTag(
		ctx,
		tagID,
	)

	assert.Error(
		t,
		err,
	)
	assert.Equal(
		t,
		coreentity.ErrTagNotFound,
		err,
	)
}

func TestTagsUseCase_DeleteTag_Error_WhenRepositoryDeleteMethodReturnsUnexpectedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTags(ctrl)
	uc := usecase.NewTagsUseCase(mockRepo)

	ctx := context.Background()
	tagID := "tag-123"
	expectedError := errors.New("database connection error")

	mockRepo.EXPECT().
		Delete(
			ctx,
			tagID,
		).
		Return(expectedError).
		Times(1)

	err := uc.DeleteTag(
		ctx,
		tagID,
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
