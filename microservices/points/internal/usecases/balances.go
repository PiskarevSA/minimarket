package usecases

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/PiskarevSA/minimarket/microservices/points/internal/domain/entities"
	"github.com/PiskarevSA/minimarket/microservices/points/internal/domain/objects"
	"github.com/PiskarevSA/minimarket/microservices/points/internal/events"
	"github.com/PiskarevSA/minimarket/microservices/points/internal/storage"

	json "github.com/bytedance/sonic"
)

type adjustBalanceStorage interface {
	AdjustBalance(
		ctx context.Context,
		orderId, userId uuid.UUID,
		operation objects.Operation,
		amount objects.Amount,
		event events.Event,
		payload []byte,
		createdBy string,
	) error
}

type adjustBalance struct {
	serviceName string
	storage     adjustBalanceStorage
}

func NewAdjustBalance(
	serviceName string,
	storage adjustBalanceStorage,
) *adjustBalance {
	return &adjustBalance{
		serviceName: serviceName,
		storage:     storage,
	}
}

func (u *adjustBalance) Do(
	ctx context.Context,
	rawUserId, rawOrderId string,
	rawOperation string,
	rawAmount string,
	event events.Event,
) error {
	const op = "balances.adjust"

	args, err := newAjustBalanceArgs(
		rawUserId, rawOrderId,
		rawOperation,
		rawAmount,
	)
	if err != nil {
		return err
	}

	payload, err := json.ConfigDefault.Marshal(
		events.BalanceChanged{
			OrderId: rawOrderId,
			UserId:  rawUserId,
			Amount:  rawAmount,
		},
	)
	if err != nil {
		log.Error().
			Err(err).
			Str("op", op).
			Str("layer", "usecases").
			Msg("failed to marshal event")

		return err
	}

	err = u.storage.AdjustBalance(
		ctx,
		args.OrderId,
		args.UserId,
		args.Operation,
		args.Amount,
		event,
		payload,
		u.serviceName,
	)
	if err != nil {
		var storageErr *storage.Error
		if !errors.As(err, &storageErr) {
			log.Error().
				Err(err).
				Str("op", op).
				Str("layer", "storage").
				Msg("failed to write balance changes")
		}

		return err
	}

	return nil
}

type getBalanceStorage interface {
	GetBalance(
		ctx context.Context,
		userId uuid.UUID,
	) (entities.Balance, error)
}

type getBalance struct {
	storage getBalanceStorage
}

func NewGetBalance(storage getBalanceStorage) *getBalance {
	return &getBalance{
		storage: storage,
	}
}

func (u *getBalance) Do(
	ctx context.Context,
	rawUserId string,
) (entities.Balance, error) {
	const op = "balance.get"

	userId, err := uuid.Parse(rawUserId)
	if err != nil {
		return entities.Balance{}, err
	}

	balance, err := u.storage.GetBalance(ctx, userId)
	if err != nil {
		var storageErr *storage.Error
		if !errors.As(err, &storageErr) {
			log.Error().
				Err(err).
				Str("op", op).
				Str("layer", "storage").
				Msg("failed to get balance")
		}

		return entities.Balance{}, err
	}

	return balance, nil
}
