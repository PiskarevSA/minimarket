package usecases

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/entities"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/repo"
)

type UploadOrderNumberStorage interface {
	CreateOrder(
		ctx context.Context,
		order entities.Order,
	) error
}

type UploadOrderNumber struct {
	storage UploadOrderNumberStorage
}

func NewUploadOrderNumber(storage UploadOrderNumberStorage) *UploadOrderNumber {
	return &UploadOrderNumber{storage: storage}
}

func (u *UploadOrderNumber) Do(
	ctx context.Context,
	rawUserId uuid.UUID,
	rawOrderNumber string,
) error {
	const op = "uploadOrderNumber"

	userID, orderNumber, err := u.parseInputs(rawUserId, rawOrderNumber)
	if err != nil {
		return err
	}

	now := time.Now()
	order := u.newOrder(userID, orderNumber, now)

	return u.createOrder(ctx, op, order)
}

func (u *UploadOrderNumber) parseInputs(
	rawUserId uuid.UUID,
	rawOrderNumber string,
) (
	userId objects.UserID,
	orderNumber objects.OrderNumber,
	err error,
) {
	userId = objects.NewUserID(rawUserId)

	orderNumber, err = objects.NewOrderNumber(rawOrderNumber)
	if err != nil {
		err = &ValidationError{
			Code:    "V1107",
			Field:   "orderNumber",
			Message: err.Error(),
		}

		return objects.NullUserID, objects.NullOrderNumber, err
	}

	return userId, orderNumber, nil
}

func (u *UploadOrderNumber) newOrder(
	userId objects.UserID,
	orderNumber objects.OrderNumber,
	uploadedAt time.Time,
) entities.Order {
	var order entities.Order

	order.SetNumber(orderNumber)
	order.SetUserID(userId)

	order.SetStatus(objects.OrderStatusNew)
	order.SetUploadedAt(uploadedAt)

	return order
}

func (u *UploadOrderNumber) createOrder(
	ctx context.Context,
	op string,
	order entities.Order,
) error {
	err := u.storage.CreateOrder(ctx, order)
	if err != nil {
		if errors.Is(err, repo.ErrOrderAlreadyCreatedByOther) {
			err = &BusinessError{
				Code:    "D1426",
				Message: "order already created by other user",
			}

			return err
		}

		if errors.Is(err, repo.ErrOrderAlreadyCreatedByUser) {
			err = &BusinessError{
				Code:    "D1531",
				Message: "order already created by user",
			}

			return err
		}

		log.Error().
			Err(err).
			Str("layer", "storage").
			Str("op", op).
			Msg("failed to create order in storage")

		return err
	}

	return nil
}
