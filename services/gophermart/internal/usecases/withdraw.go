package usecases

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/entities"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/repo"
)

type WithdrawRepo interface {
	BalanceStorage
	TransactionStorage
}

type Withdraw struct {
	storage    WithdrawRepo
	transactor Transactor
}

func NewWithdraw(storage WithdrawRepo, transactor Transactor) *Withdraw {
	return &Withdraw{
		storage:    storage,
		transactor: transactor,
	}
}

func (u *Withdraw) Do(
	ctx context.Context,
	rawUserId uuid.UUID,
	rawOrderNumber,
	rawSum string,
) error {
	const op = "withdraw"

	userID := objects.NewUserID(rawUserId)

	orderNumber, sum, err := u.parseInputs(rawOrderNumber, rawSum)
	if err != nil {
		return err
	}

	now := time.Now()
	tx := u.newTransaction(orderNumber, userID, sum, now)

	return u.createWithdraw(ctx, op, tx)
}

func (u *Withdraw) parseInputs(
	rawOrderNumber,
	rawSum string,
) (
	orderNumber objects.OrderNumber,
	amount objects.Amount,
	err error,
) {
	orderNumber, err = objects.NewOrderNumber(rawOrderNumber)
	if err != nil {
		err = &ValidationError{
			Code:    "V1107",
			Field:   "orderNumber",
			Message: err.Error(),
		}

		return objects.NullOrderNumber, objects.NullAmount, err
	}

	amount, err = objects.NewAmount(rawSum)
	if err != nil {
		err = &ValidationError{
			Code:    "V1142",
			Field:   "amount",
			Message: err.Error(),
		}

		return objects.NullOrderNumber, objects.NullAmount, err
	}

	return orderNumber, amount, nil
}

func (u *Withdraw) newTransaction(
	orderNumber objects.OrderNumber,
	userId objects.UserID,
	sum objects.Amount,
	pocessedAt time.Time,
) entities.Transaction {
	var tx entities.Transaction

	tx.SetOrderNumber(orderNumber)
	tx.SetUserID(userId)

	tx.SetSum(sum)
	tx.SetOperation(objects.OperationWithdraw)
	tx.SetProcessedAt(pocessedAt)

	return tx
}

func (u *Withdraw) createWithdraw(
	ctx context.Context,
	op string,
	tx entities.Transaction,
) error {
	pgxTxOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}

	transactFn := func(ctx context.Context) error {
		err := u.storage.CreateTransactionInTx(ctx, tx)
		if err != nil {
			log.Error().
				Err(err).
				Str("layer", "storage").
				Str("op", op).
				Msg("failed to write transaction to storage")

			return err
		}

		err = u.storage.CreateOrUpdateBalanceInTx(
			ctx,
			tx.UserID(),
			tx.Operation(),
			tx.Sum(),
		)
		if err != nil {
			if errors.Is(err, repo.ErrNotEnoughtBalance) {
				err = &BusinessError{
					Code:    "D1215",
					Message: "insufficient balance",
				}

				return err
			}

			log.Error().
				Err(err).
				Str("layer", "storage").
				Str("op", op).
				Msg("failed to write balance change to storage")

			return err
		}

		return nil
	}

	return u.transactor.Transact(ctx, pgxTxOpts, transactFn)
}
