package usecases

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/entities"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/repo"
)

type LoginStorage interface {
	GetAccountByLogin(
		ctx context.Context,
		login objects.Login,
	) (account entities.Account, err error)
}

type Login struct {
	storage LoginStorage
	idp     IdentityProvider
}

func NewLogin(storage LoginStorage, idp IdentityProvider) *Login {
	return &Login{
		storage: storage,
		idp:     idp,
	}
}

func (u *Login) Do(
	ctx context.Context,
	rawLogin,
	rawPassword string,
) (token string, err error) {
	const op = "login"

	login, password, err := u.parseInputs(rawLogin, rawPassword)
	if err != nil {
		return "", err
	}

	account, err := u.getAccountByLogin(ctx, op, login)
	if err != nil {
		return "", err
	}

	hash := account.PasswordHash()

	err = u.matchPasswordAndHash(op, password, hash)
	if err != nil {
		return "", err
	}

	rawUserId := account.Id().Uuid()
	now := time.Now()

	token, err = u.idp.IssueToken(rawUserId, now)
	if err != nil {
		log.Error().
			Err(err).
			Str("op", op).
			Str("layer", "usecase").
			Msg("failed to issue jwt")

		return "", err
	}

	return token, nil
}

func (u *Login) parseInputs(
	rawLogin string,
	rawPassword string,
) (
	login objects.Login,
	password objects.Password,
	err error,
) {
	login, err = objects.NewLogin(rawLogin)
	if err != nil {
		err = &ValidationError{
			Code:    "V1042",
			Field:   "login",
			Message: err.Error(),
		}

		return objects.NullLogin, objects.NilPassword, err
	}

	password, err = objects.NewPassword(rawPassword)
	if err != nil {
		err = &ValidationError{
			Code:    "V1078",
			Field:   "password",
			Message: err.Error(),
		}

		return objects.NullLogin, objects.NilPassword, err
	}

	return login, password, nil
}

func (u *Login) getAccountByLogin(
	ctx context.Context,
	op string,
	login objects.Login,
) (account entities.Account, err error) {
	account, err = u.storage.GetAccountByLogin(ctx, login)
	if err != nil {
		if errors.Is(err, repo.ErrLoginNotExists) {
			err = &BusinessError{
				Code:    "D1126",
				Message: "invalid login or password",
			}

			return entities.NullAccount, err
		}

		log.Error().
			Err(err).
			Str("op", op).
			Str("layer", "storage").
			Msg("failed to get account auth from storage")

		return entities.NullAccount, err
	}

	return account, nil
}

func (u *Login) matchPasswordAndHash(
	op string,
	password objects.Password,
	hash []byte,
) error {
	match, err := password.IsHashMatch(hash)
	if err != nil {
		log.Error().
			Err(err).
			Str("op", op).
			Str("layer", "usecase").
			Msg("failed to match password and its hash")

		return err
	}

	if !match {
		return &BusinessError{
			Code:    "D1126",
			Message: "invalid login or password",
		}
	}

	return nil
}
