package postgresql

import (
	"context"

	"github.com/github.com/PiskarevSA/minimarket/pkg/pgx/transactor"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/entities"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/gen/sqlc/postgresql"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/repo"
)

func (r *PostgreSql) createOrder(
	ctx context.Context,
	querier *postgresql.Queries,
	order entities.Order,
) error {
	row, err := querier.CreateOrder(
		ctx,
		postgresql.CreateOrderParams{
			Number:     order.Number().String(),
			UserId:     order.UserID().UUID(),
			Status:     order.Status().String(),
			Accrual:    order.Accrual().Numeric(),
			UploadedAt: order.UploadedAt(),
		},
	)
	if err != nil {
		return err
	}

	if !row.Inserted {
		orderCreator := objects.NewUserID(row.UserId)
		if order.UserID().Equal(orderCreator) {
			return repo.ErrOrderAlreadyCreatedByUser
		}

		return repo.ErrOrderAlreadyCreatedByOther
	}

	return nil
}

func (r *PostgreSql) CreateOrder(
	ctx context.Context,
	order entities.Order,
) error {
	return r.createOrder(ctx, r.querier, order)
}

func (r *PostgreSql) CreateOrderInTx(
	ctx context.Context,
	order entities.Order,
) error {
	pgxTx := transactor.TxCtxKey.Value(ctx)
	querier := r.querier.WithTx(pgxTx)

	return r.createOrder(ctx, querier, order)
}
