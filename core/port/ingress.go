package port

import (
	"context"

	"proletariat-budget-core/core/domain/coreentity"
)

//go:generate mockgen -source=ingress.go -destination=../../test/mocks/mock_ingressrepo.go -package mocks
type Ingress interface {
	// Ingress operations
	Create(
		ctx context.Context,
		ingress coreentity.Ingress,
	) (
		string,
		error,
	)
	GetByID(
		ctx context.Context,
		id string,
	) (
		coreentity.Ingress,
		error,
	)
	List(
		ctx context.Context,
		params coreentity.IngressListParams,
	) (
		coreentity.IngressList,
		error,
	)
	UpdateTemplate(
		ctx context.Context,
		template coreentity.Ingress,
	) error
	DeleteTemplate(
		ctx context.Context,
		template coreentity.Ingress,
	) error
}
