package usecase

import (
	"context"

	"golang.org/x/crypto/bcrypt"
	"proletariat-budget-core/core/domain/auth"
	"proletariat-budget-core/core/port"
)

type Auth struct {
	authRepo port.Auth
}

func NewAuthUseCase(authRepo port.Auth) *Auth {
	return &Auth{
		authRepo: authRepo,
	}
}

func (uc *Auth) Login(
	ctx context.Context,
	email, password string,
) (
	*auth.AuthToken,
	*auth.User,
	error,
) {
	user, err := uc.authRepo.GetUserByEmail(
		ctx,
		email,
	)
	if err != nil {
		return nil, nil, auth.ErrInvalidCredentials
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)
	if err != nil {
		return nil, nil, auth.ErrInvalidCredentials
	}

	// Generate token
	token, err := uc.authRepo.CreateToken(
		ctx,
		user.ID,
	)
	if err != nil {
		return nil, nil, err
	}

	return token, user, nil
}

func (uc *Auth) ValidateToken(
	ctx context.Context,
	tokenStr string,
) (
	*auth.User,
	error,
) {
	return uc.authRepo.ValidateToken(
		ctx,
		tokenStr,
	)
}

func (uc *Auth) RefreshToken(
	ctx context.Context,
	userID string,
) (
	*auth.AuthToken,
	*auth.User,
	error,
) {
	user, err := uc.authRepo.GetUserByID(
		ctx,
		userID,
	)
	if err != nil {
		return nil, nil, err
	}

	token, err := uc.authRepo.CreateToken(
		ctx,
		userID,
	)
	if err != nil {
		return nil, nil, err
	}

	return token, user, nil
}
