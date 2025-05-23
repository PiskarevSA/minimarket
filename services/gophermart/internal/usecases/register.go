package usecases

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/entities"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/repo"
)

type RegisterStorage interface {
	CreateAccount(
		ctx context.Context,
		account entities.Account,
	) error
}

type Register struct {
	storage RegisterStorage
	idp     IdentityProvider
}

func NewRegister(storage RegisterStorage, idp IdentityProvider) *Register {
	return &Register{
		storage: storage,
		idp:     idp,
	}
}

func (u *Register) Do(
	ctx context.Context,
	rawLogin,
	rawPassword string,
) (token string, err error) {
	const op = "register"

	login, password, err := u.parseInputs(rawLogin, rawPassword)
	if err != nil {
		return "", err
	}

	now := time.Now()
	passwordHash := password.Hash()

	account := u.newAccount(login, passwordHash, now)
	rawUserId := account.Id().Uuid()

	token, err = u.idp.IssueToken(rawUserId, now)
	if err != nil {
		log.Error().
			Err(err).
			Str("op", op).
			Str("layer", "usecase").
			Msg("failed to create access token")

		return "", err
	}

	err = u.createAccount(ctx, op, account)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *Register) parseInputs(
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

func (u *Register) newAccount(
	login objects.Login,
	passwordHash []byte,
	createdAt time.Time,
) entities.Account {
	var account entities.Account

	rawUserId := uuid.New()
	userId := objects.NewUserId(rawUserId)
	account.SetId(userId)

	account.SetLogin(login)
	account.SetPasswordHash(passwordHash)

	account.SetCreatedAt(createdAt)
	account.SetUpdatedAt(createdAt)

	return account
}

func (u *Register) createAccount(
	ctx context.Context,
	op string,
	account entities.Account,
) error {
	err := u.storage.CreateAccount(ctx, account)
	if err != nil {
		if errors.Is(err, repo.ErrLoginAlreadyInUse) {
			err = &BusinessError{
				Code:    "D1002",
				Message: "login already in use",
			}

			return err
		}

		log.Error().
			Err(err).
			Str("op", op).
			Str("layer", "storage").
			Msg("failed to create account in storage")

		return err
	}

	return nil
}
