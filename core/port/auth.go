package port

import (
	"context"

	"proletariat-budget-core/core/domain/auth"
)

//go:generate mockgen -source=auth.go -destination=../../test/mocks/mock_authrepo.go -package mocks
type Auth interface {
	CreateUser(
		ctx context.Context,
		user auth.User,
	) (
		string,
		error,
	)
	GetUserByEmail(
		ctx context.Context,
		email string,
	) (
		*auth.User,
		error,
	)
	GetUserByID(
		ctx context.Context,
		id string,
	) (
		*auth.User,
		error,
	)
	UpdateUser(
		ctx context.Context,
		id string,
		user auth.User,
	) error

	// Token management
	CreateToken(
		ctx context.Context,
		userID string,
	) (
		*auth.AuthToken,
		error,
	)
	ValidateToken(
		ctx context.Context,
		token string,
	) (
		*auth.User,
		error,
	)
	RevokeToken(
		ctx context.Context,
		token string,
	) error
}
