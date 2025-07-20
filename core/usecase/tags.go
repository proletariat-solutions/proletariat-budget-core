package usecase

import (
	"context"
	"errors"

	"proletariat-budget-core/core/domain/coreentity"
	"proletariat-budget-core/core/port"
)

type Tags struct {
	tagsRepo port.Tags
}

func NewTagsUseCase(tagsRepo port.Tags) *Tags {
	return &Tags{tagsRepo: tagsRepo}
}

func (u *Tags) ListTags(
	ctx context.Context,
	tagType *coreentity.TagType,
) (
	[]*coreentity.Tag,
	error,
) {
	if tagType == nil {
		tags, err := u.tagsRepo.List(ctx)
		if err != nil {
			return nil, err
		}

		return *tags, nil
	}
	tags, err := u.tagsRepo.ListByType(
		ctx,
		*tagType,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return *tags, nil
}

func (u *Tags) CreateTag(
	ctx context.Context,
	tag *coreentity.Tag,
) (
	*coreentity.Tag,
	error,
) {
	if err := tag.Validate(); err != nil {
		return nil, err
	}
	id, err := u.tagsRepo.Create(
		ctx,
		*tag,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrDuplicateKey,
		) {
			return nil, coreentity.ErrTagAlreadyExists
		}

		return nil, err
	}
	tag.ID = id

	return tag, nil
}

func (u *Tags) DeleteTag(
	ctx context.Context,
	id string,
) error {
	err := u.tagsRepo.Delete(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(
			err,
			port.ErrRecordNotFound,
		) {
			return coreentity.ErrTagNotFound
		}

		return err
	}

	return nil
}
