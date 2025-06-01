package postgresql

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/entities"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/dto"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/repo"
)

func (r *PostgreSQL) GetAccountByUserID(
	ctx context.Context,
	userID objects.UserID,
) (account entities.Account, err error) {
	row, err := r.querier.GetAccountByUserID(ctx, userID.UUID())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = repo.ErrLoginNotExists
		}

		return entities.NullAccount, err
	}

	account = dto.GetAccountByUserIDToAccount(userID, row)

	return account, err
}

func (r *PostgreSQL) GetAccountByLogin(
	ctx context.Context,
	login objects.Login,
) (account entities.Account, err error) {
	row, err := r.querier.GetAccountByLogin(ctx, login.String())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = repo.ErrLoginNotExists
		}

		return entities.NullAccount, err
	}

	account = dto.GetAccountByLoginToAccount(login, row)

	return account, err
}
