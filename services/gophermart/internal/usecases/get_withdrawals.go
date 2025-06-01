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

type GetWithdrawalsStorage interface {
	GetWithdrawalsByUserID(
		ctx context.Context,
		userID objects.UserID,
		offset,
		limit int32,
	) ([]entities.Transaction, error)
}

type GetWithdrawals struct {
	storage GetWithdrawalsStorage
}

func NewGetWithdrawals(storage GetWithdrawalsStorage) *GetWithdrawals {
	return &GetWithdrawals{storage: storage}
}

func (u *GetWithdrawals) Do(
	ctx context.Context,
	rawUserID uuid.UUID,
	offset, limit int32,
) (txs []entities.Transaction, err error) {
	const op = "getWithdrawals"

	userID := objects.NewUserID(rawUserID)

	txs, err = u.storage.GetWithdrawalsByUserID(ctx, userID, offset, limit)
	if err != nil {
		if errors.Is(err, repo.ErrNoTransactionsFound) {
			return entities.Nullransactions, nil
		}

		log.Error().
			Err(err).
			Str("layer", "storage").
			Str("op", op).
			Msg("failed to get transactions from storage")

		return entities.Nullransactions, err
	}

	return txs, nil
}
