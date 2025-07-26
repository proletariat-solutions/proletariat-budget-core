package port

import "context"

//go:generate mockgen -source=app_config.go -destination=../../test/mocks/mock_app_config_repo.go -package mocks
type AppConfig interface {
	Create(
		ctx context.Context,
		appConfig AppConfig,
	) (
		*AppConfig,
		error,
	)
	Get(ctx context.Context) (
		*AppConfig,
		error,
	)
	Update(
		ctx context.Context,
		appConfig AppConfig,
	) error
}
