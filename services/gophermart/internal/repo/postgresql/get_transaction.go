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

func (r *PostgreSQL) getTransactionsByUserID(
	ctx context.Context,
	userID objects.UserID,
	operation objects.Operation,
	offset,
	limit int32,
) (txs []entities.Transaction, err error) {
	rows, err := r.querier.GetTransactionsByUserID(
		ctx,
		postgresql.GetTransactionsByUserIDParams{
			UserId:    userID.UUID(),
			Operation: operation.String(),
			Offset:    convtype.Int32ToInt4(offset),
			Limit:     convtype.Int32ToInt4(limit),
		},
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = repo.ErrNoTransactionsFound
		}

		return entities.Nullransactions, err
	}

	txs = dto.GetTxsByUserIDToTxs(userID, rows)

	return txs, nil
}

func (r *PostgreSQL) GetDepositsByUserID(
	ctx context.Context,
	userID objects.UserID,
	offset,
	limit int32,
) (txs []entities.Transaction, err error) {
	return r.getTransactionsByUserID(
		ctx,
		userID,
		objects.OperationDeposit,
		offset,
		limit,
	)
}

func (r *PostgreSQL) GetWithdrawalsByUserID(
	ctx context.Context,
	userID objects.UserID,
	offset,
	limit int32,
) (txs []entities.Transaction, err error) {
	return r.getTransactionsByUserID(
		ctx,
		userID,
		objects.OperationWithdraw,
		offset,
		limit,
	)
}
