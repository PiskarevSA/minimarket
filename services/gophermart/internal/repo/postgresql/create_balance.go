package postgresql

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/github.com/PiskarevSA/minimarket/pkg/pgcodes"
	"github.com/github.com/PiskarevSA/minimarket/pkg/pgx/transactor"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/gen/sqlc/postgresql"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/repo"
)

func (r PostgreSql) createOrUpdateBalance(
	ctx context.Context,
	querier *postgresql.Queries,
	userId objects.UserId,
	operation objects.Operation,
	sum objects.Amount,
) error {
	err := querier.CreateOrUpdateBalance(
		ctx,
		postgresql.CreateOrUpdateBalanceParams{
			UserId:    userId.Uuid(),
			Operation: operation.String(),
			Sum:       sum.Numeric(),
		},
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgcodes.IsCheckViolation(pgErr.Code) {
				err = repo.ErrNotEnoughtBalance
			}
		}

		return err
	}

	return nil
}

func (r *PostgreSql) CreateOrUpdateBalance(
	ctx context.Context,
	userId objects.UserId,
	operation objects.Operation,
	sum objects.Amount,
) error {
	return r.createOrUpdateBalance(ctx, r.querier, userId, operation, sum)
}

func (r *PostgreSql) CreateOrUpdateBalanceInTx(
	ctx context.Context,
	userId objects.UserId,
	operation objects.Operation,
	sum objects.Amount,
) error {
	pgxTx := transactor.TxCtxKey.Value(ctx)
	querier := r.querier.WithTx(pgxTx)

	return r.createOrUpdateBalance(ctx, querier, userId, operation, sum)
}
