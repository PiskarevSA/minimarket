package usecases

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/config"
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/models"
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/storage"
)

type UserLogIn struct {
	storage          UserStorage
	jwtSignKey       any
	jwtAccessExpiry  time.Duration
	jwtRefreshExpiry time.Duration
}

func NewUserLogIn(
	storage UserStorage,
	jwtSignKey any,
	jwtAccessExpiry time.Duration,
	jwtRefreshExpiry time.Duration,
) *UserLogIn {
	return &UserLogIn{
		storage:          storage,
		jwtSignKey:       jwtSignKey,
		jwtAccessExpiry:  jwtAccessExpiry,
		jwtRefreshExpiry: jwtRefreshExpiry,
	}
}

func (u *UserLogIn) Do(
	ctx context.Context,
	login, password string,
) (*models.UserLoggedIn, error) {
	const op = "usecases.user_login.do"

	creds, err := u.storage.GetUserCreds(ctx, login)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			return nil, ErrInvalidLoginOrPassword
		}

		log.Error().
			Err(err).
			Str("op", op).
			Msg("failed to get user credentials")

		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(creds.PasswordHash),
		[]byte(password),
	)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, ErrInvalidLoginOrPassword
		}

		log.Error().
			Err(err).
			Str("op", op).
			Msg("failed to compare password hash")

		return nil, err
	}

	userId := uuid.New()
	userIdStr := userId.String()

	accessToken, refreshToken, err := createTokenPair(
		userIdStr,
		u.jwtAccessExpiry,
		u.jwtRefreshExpiry,
		u.jwtSignKey,
		config.JwtAlgo(),
	)
	if err != nil {
		log.Error().
			Err(err).
			Str("op", op).
			Msg("failed to create token pair")

		return nil, err
	}

	return &models.UserLoggedIn{
		UserId:       userIdStr,
		Login:        login,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
