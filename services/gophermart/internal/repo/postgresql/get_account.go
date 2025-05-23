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

func (r *PostgreSql) GetAccountByUserId(
	ctx context.Context,
	userId objects.UserId,
) (account entities.Account, err error) {
	row, err := r.querier.GetAccountByUserId(ctx, userId.Uuid())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = repo.ErrLoginNotExists
		}

		return entities.NullAccount, err
	}

	account = dto.GetAccountByUserIdToAccount(userId, row)

	return account, err
}

func (r *PostgreSql) GetAccountByLogin(
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
