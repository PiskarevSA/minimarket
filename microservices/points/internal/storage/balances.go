package storage

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/PiskarevSA/minimarket/microservices/points/internal/domain/entities"
	"github.com/PiskarevSA/minimarket/microservices/points/internal/domain/objects"
	"github.com/PiskarevSA/minimarket/microservices/points/internal/events"
	sqlc "github.com/PiskarevSA/minimarket/microservices/points/internal/gen/sqlc/postgresql"
	"github.com/PiskarevSA/minimarket/pkg/pgcodes"
)

func (s *Storage) createOrUpdateBalance(
	ctx context.Context, query *sqlc.Queries,
	userId uuid.UUID,
	txType string,
	amount pgtype.Numeric,
) error {
	err := query.CreateOrUpdateBalance(
		ctx, sqlc.CreateOrUpdateBalanceParams{
			UserId:    userId,
			Operation: txType,
			Amount:    amount,
		},
	)
	if err != nil {
		var pgxErr *pgconn.PgError
		if !errors.As(err, &pgxErr) &&
			!pgcodes.IsCheckViolation(pgxErr.Code) {
			return err
		}

		return &Error{"insufficient points"}
	}

	return nil
}

func (s *Storage) AdjustBalance(
	ctx context.Context,
	orderId, userId uuid.UUID,
	operation objects.Operation,
	amount objects.Amount,
	event events.Event,
	payload []byte,
	createdBy string,
) error {
	fn := func(query *sqlc.Queries) error {
		operation := operation.String()
		amount := amount.Numeric()

		err := s.createTransaction(
			ctx, query,
			orderId, userId,
			operation,
			amount,
		)
		if err != nil {
			return err
		}

		err = s.createOrUpdateBalance(
			ctx, query,
			userId,
			operation,
			amount,
		)
		if err != nil {
			return err
		}

		return s.createOutbox(
			ctx, query,
			event.String(),
			payload,
			createdBy,
		)
	}

	pgxTxOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}

	return s.sqlTransactor.Transact(ctx, pgxTxOpts, fn)
}

func (s *Storage) GetBalance(
	ctx context.Context,
	userId uuid.UUID,
) (entities.Balance, error) {
	balance, err := s.sqlQuerier.GetBalance(ctx, userId)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return entities.Balance{}, err
		}

		err = &Error{"balance not found"}

		return entities.Balance{}, err
	}

	return entities.NewBalance(
		userId,
		balance.Available,
		balance.Withdrawn,
		balance.UpdatedAt,
	)
}
