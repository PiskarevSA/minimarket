package postgresql

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/github.com/PiskarevSA/minimarket/pkg/pgcodes"
	"github.com/github.com/PiskarevSA/minimarket/pkg/pgx/transactor"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/entities"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/gen/sqlc/postgresql"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/repo"
)

func (r *PostgreSql) createAccount(
	ctx context.Context,
	querier *postgresql.Queries,
	account entities.Account,
) error {
	err := querier.CreateAccount(
		ctx,
		postgresql.CreateAccountParams{
			Id:           account.ID().UUID(),
			Login:        account.Login().String(),
			PasswordHash: string(account.PasswordHash()),
			CreatedAt:    account.CreatedAt(),
			UpdatedAt:    account.UpdatedAt(),
		},
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgcodes.IsUniqueViolation(pgErr.Code) {
				err = repo.ErrLoginAlreadyInUse
			}
		}

		return err
	}

	return nil
}

func (r *PostgreSql) CreateAccount(
	ctx context.Context,
	account entities.Account,
) error {
	return r.createAccount(ctx, r.querier, account)
}

func (r *PostgreSql) CreateAccountInTx(
	ctx context.Context,
	account entities.Account,
) error {
	pgxTx := transactor.TxCtxKey.Value(ctx)
	querier := r.querier.WithTx(pgxTx)

	return r.createAccount(ctx, querier, account)
}
