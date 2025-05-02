package usecases

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/models"
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/storage"
)

type UserRegister struct {
	storage          UserStorage
	jwtSignKey       any
	jwtAlgo          jwt.SigningMethod
	jwtAccessExpiry  time.Duration
	jwtRefreshExpiry time.Duration
}

func NewUserRegister(
	storage UserStorage,
	jwtSignKey any,
	jwtAlgo jwt.SigningMethod,
	jwtAccessExpiry time.Duration,
	jwtRefreshExpiry time.Duration,
) *UserRegister {
	return &UserRegister{
		storage:          storage,
		jwtSignKey:       jwtSignKey,
		jwtAlgo:          jwtAlgo,
		jwtAccessExpiry:  jwtAccessExpiry,
		jwtRefreshExpiry: jwtRefreshExpiry,
	}
}

func (u *UserRegister) Do(
	ctx context.Context,
	login, password string,
) (*models.UserRegistered, error) {
	const op = "usecases.user_register.do"

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		log.Error().
			Err(err).
			Str("op", op).
			Msg("failed to generate password hash")

		return nil, err
	}

	userId := uuid.New()

	err = u.storage.CreateUser(
		ctx,
		userId,
		login,
		string(passwordHash),
	)
	if err != nil {
		if errors.Is(err, storage.ErrLoginAlreadyInUse) {
			return nil, ErrLoginAlreadyInUse
		}

		log.Error().
			Err(err).
			Str("op", op).
			Msg("failed to create user in storage")

		return nil, err
	}

	userIdStr := userId.String()

	accessToken, refreshToken, err := createTokenPair(
		userIdStr,
		u.jwtAccessExpiry,
		u.jwtRefreshExpiry,
		u.jwtSignKey,
		u.jwtAlgo,
	)
	if err != nil {
		log.Error().
			Err(err).
			Str("op", op).
			Msg("failed to create token pair")

		return nil, err
	}

	return &models.UserRegistered{
		UserId:       userIdStr,
		Login:        login,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
