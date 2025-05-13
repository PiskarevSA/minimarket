package usecases

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/PiskarevSA/minimarket/microservices/points/internal/domain/entities"
	"github.com/PiskarevSA/minimarket/microservices/points/internal/storage"
)

type getTransactionsStorage interface {
	GetTransactions(
		ctx context.Context,
		userId uuid.UUID,
		offset, limit int32,
	) ([]entities.Transaction, error)
}

type getTransactions struct {
	storage getTransactionsStorage
}

func NewGetTransactions(storage getTransactionsStorage) *getTransactions {
	return &getTransactions{
		storage: storage,
	}
}

func (u *getTransactions) Do(
	ctx context.Context,
	rawUserId string,
	offset, limit int32,
) ([]entities.Transaction, error) {
	const op = "transactions.get"

	args, err := newGetTransactionsArgs(
		rawUserId,
		offset, limit,
	)
	if err != nil {
		return nil, err
	}

	txs, err := u.storage.GetTransactions(
		ctx,
		args.UserId,
		args.Offset,
		args.Limit,
	)
	if err != nil {
		var storageErr *storage.Error
		if !errors.As(err, &storageErr) {
			log.Error().
				Err(err).
				Str("op", op).
				Str("layer", "storage").
				Msg("failed to get transactions")
		}

		return nil, err
	}

	return txs, nil
}
