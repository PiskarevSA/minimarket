package usecases

import (
	"github.com/google/uuid"

	"github.com/PiskarevSA/minimarket/microservices/points/internal/domain/objects"
)

type adjustBalanceArgs struct {
	UserId    uuid.UUID
	OrderId   uuid.UUID
	Operation objects.Operation
	Amount    objects.Amount
}

func newAjustBalanceArgs(
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
		return adjustBalanceArgs{}, err
	}

	args.OrderId, err = uuid.Parse(rawOrderId)
	if err != nil {
		return adjustBalanceArgs{}, err
	}

	args.Operation, err = objects.NewOperation(rawOperation)
	if err != nil {
		return adjustBalanceArgs{}, err
	}

	args.Amount, err = objects.NewAmount(rawAmount)
	if err != nil {
		return adjustBalanceArgs{}, err
	}

	return args, nil
}

type getTransactionsArgs struct {
	UserId uuid.UUID
	Offset int32
	Limit  int32
}

func newGetTransactionsArgs(
	rawUserId string,
	offset, limit int32,
) (getTransactionsArgs, error) {
	var (
		args getTransactionsArgs
		err  error
	)

	args.UserId, err = uuid.Parse(rawUserId)
	if err != nil {
		return getTransactionsArgs{}, err
	}

	args.Offset = offset
	args.Limit = limit

	return args, nil
}
