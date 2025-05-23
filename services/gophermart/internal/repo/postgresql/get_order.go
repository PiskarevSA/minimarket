package postgresql

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/github.com/PiskarevSA/minimarket/pkg/convtype"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/entities"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/dto"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/gen/sqlc/postgresql"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/repo"
)

func (r *PostgreSql) GetOrderByNumber(
	ctx context.Context,
	number objects.OrderNumber,
) (order entities.Order, err error) {
	return order, err
}

func (r *PostgreSql) GetOrdersByUserId(
	ctx context.Context,
	userId objects.UserId,
	offset,
	limit int32,
) (orders []entities.Order, err error) {
	rows, err := r.querier.GetOrdersByUserId(
		ctx,
		postgresql.GetOrdersByUserIdParams{
			UserId: userId.Uuid(),
			Offset: convtype.Int32ToInt4(offset),
			Limit:  convtype.Int32ToInt4(limit),
		},
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = repo.ErrNoOrdersFound
		}

		return entities.NullOrders, err
	}

	orders = dto.GetOrdersByUserIdToOrders(userId, rows)

	return orders, err
}
