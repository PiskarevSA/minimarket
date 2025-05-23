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

func (r *PostgreSql) createTransaction(
	ctx context.Context,
	querier *postgresql.Queries,
	tx entities.Transaction,
) error {
	err := querier.CreateTransaction(
		ctx,
		postgresql.CreateTransactionParams{
			UserId:      tx.UserId().Uuid(),
			OrderNumber: tx.OrderNumber().String(),
			Operation:   tx.Operation().String(),
			Sum:         tx.Sum().Numeric(),
			ProcessedAt: tx.ProcessedAt(),
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

func (r *PostgreSql) CreateTransaction(
	ctx context.Context,
	tx entities.Transaction,
) error {
	return r.createTransaction(ctx, r.querier, tx)
}

func (r *PostgreSql) CreateTransactionInTx(
	ctx context.Context,
	tx entities.Transaction,
) error {
	pgxTx := transactor.TxCtxKey.Value(ctx)
	querier := r.querier.WithTx(pgxTx)

	return r.createTransaction(ctx, querier, tx)
}
