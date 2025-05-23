package usecases

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/entities"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/repo"
)

type GetOrdersStorage interface {
	GetOrdersByUserId(
		ctx context.Context,
		userId objects.UserId,
		offset int32,
		limit int32,
	) ([]entities.Order, error)
}

type GetOrders struct {
	storage GetOrdersStorage
}

func NewGetOrders(storage GetOrdersStorage) *GetOrders {
	return &GetOrders{storage: storage}
}

func (u *GetOrders) Do(
	ctx context.Context,
	rawUserId uuid.UUID,
	offset,
	limit int32,
) (orders []entities.Order, err error) {
	const op = "getOrders"

	userId := objects.NewUserId(rawUserId)

	orders, err = u.storage.GetOrdersByUserId(ctx, userId, offset, limit)
	if err != nil {
		if errors.Is(err, repo.ErrNoOrdersFound) {
			return entities.NullOrders, nil
		}

		log.Error().
			Err(err).
			Str("layer", "storage").
			Str("op", op).
			Msg("failed to get orders from storage")

		return entities.NullOrders, err
	}

	return orders, nil
}
