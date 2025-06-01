package usecases

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/repo"
)

type GetBalanceStorage interface {
	GetBalanceByUserID(
		ctx context.Context,
		userID objects.UserID,
	) (current, withdrawn objects.Amount, err error)
}

type GetBalance struct {
	storage GetBalanceStorage
}

func NewGetBalance(storage GetBalanceStorage) *GetBalance {
	return &GetBalance{storage: storage}
}

func (u *GetBalance) Do(
	ctx context.Context,
	rawUserID uuid.UUID,
) (current, withdrawn objects.Amount, err error) {
	const op = "getBalance"

	userID := objects.NewUserID(rawUserID)

	current, withdrawn, err = u.storage.GetBalanceByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, repo.ErrNoBalanceFound) {
			return objects.NullAmount, objects.NullAmount, nil
		}

		log.Error().
			Err(err).
			Str("layer", "storage").
			Str("op", op).
			Msg("failed to get balance from storage")

		return objects.NullAmount, objects.NullAmount, err
	}

	return current, withdrawn, nil
}
