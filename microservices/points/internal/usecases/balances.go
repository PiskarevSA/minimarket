package usecases

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

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

type adjustBalanceArgs struct {
	UserId    uuid.UUID
	OrderId   uuid.UUID
	Operation objects.Operation
	Amount    objects.Amount
}

func validateArgs(
	rawUserId, rawOrderId string,
	rawOperation string,
	rawAmount string,
) (adjustBalanceArgs, error) {
	var (
		args adjustBalanceArgs
		err  error
	)

	args.UserId, err = uuid.Parse(rawUserId)
	if err != nil {
		return args, err
	}

	args.OrderId, err = uuid.Parse(rawOrderId)
	if err != nil {
		return args, err
	}

	args.Operation, err = objects.NewOperation(rawOperation)
	if err != nil {
		return args, err
	}

	args.Amount, err = objects.NewAmount(rawAmount)
	if err != nil {
		return args, err
	}

	return args, nil
}

func (u *adjustBalance) Do(
	ctx context.Context,
	rawUserId, rawOrderId string,
	rawOperation string,
	rawAmount string,
	event events.Event,
) error {
	const op = "balances.adjust"

	args, err := validateArgs(
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
