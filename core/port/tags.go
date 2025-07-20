package port

import (
	"context"

	"proletariat-budget-core/core/domain/coreentity"
)

//go:generate mockgen -source=tags.go -destination=../../test/mocks/mock_tagsrepo.go -package mocks
type Tags interface {
	Create(
		ctx context.Context,
		tag coreentity.Tag,
	) (
		string,
		error,
	)
	Update(
		ctx context.Context,
		id string,
		tag coreentity.Tag,
	) error
	Delete(
		ctx context.Context,
		id string,
	) error
	GetByID(
		ctx context.Context,
		id string,
	) (
		*coreentity.Tag,
		error,
	)
	GetByIDs(
		ctx context.Context,
		ids []string,
	) (
		*[]*coreentity.Tag,
		error,
	)

	ListByType(
		ctx context.Context,
		tagType coreentity.TagType,
		ids *[]string,
	) (
		*[]*coreentity.Tag,
		error,
	)

	LinkTagsToType(
		ctx context.Context,
		foreignID string,
		tags *[]*coreentity.Tag,
	) error

	List(ctx context.Context) (
		*[]*coreentity.Tag,
		error,
	)

	GetByNameAndType(
		ctx context.Context,
		name string,
		tagType coreentity.TagType,
	) (
		*coreentity.Tag,
		error,
	)
}
