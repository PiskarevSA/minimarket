package storage

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/PiskarevSA/minimarket/microservices/points/internal/domain/entities"
	sqlc "github.com/PiskarevSA/minimarket/microservices/points/internal/gen/sqlc/postgresql"
	"github.com/PiskarevSA/minimarket/pkg/pgx/convtype"
)

func (s *Storage) createTransaction(
	ctx context.Context, query *sqlc.Queries,
	orderId, userId uuid.UUID,
	operation string,
	amount pgtype.Numeric,
) error {
	return query.CreateTransaction(
		ctx, sqlc.CreateTransactionParams{
			OrderId:   orderId,
			UserId:    userId,
			Operation: operation,
			Amount:    amount,
		},
	)
}

func toDomainTransactions(
	txs []sqlc.Transaction,
) ([]entities.Transaction, error) {
	result := make([]entities.Transaction, len(txs))

	for i, r := range txs {
		tx, err := entities.NewTransaction(
			r.Id,
			r.Operation,
			r.OrderId,
			r.UserId,
			r.Amount,
			r.Timestamp,
		)
		if err != nil {
			return nil, err
		}

		result[i] = tx
	}

	return result, nil
}

func (s *Storage) GetTransactions(
	ctx context.Context,
	userId uuid.UUID,
	offset, limit int32,
) ([]entities.Transaction, error) {
	txs, err := s.sqlQuerier.GetTransactions(
		ctx, sqlc.GetTransactionsParams{
			UserId: userId,
			Offset: convtype.Int32ToInt4(offset),
			Limit:  convtype.Int32ToInt4(limit),
		},
	)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}

		err = &Error{"transactions not found"}

		return nil, err
	}

	return toDomainTransactions(txs)
}
