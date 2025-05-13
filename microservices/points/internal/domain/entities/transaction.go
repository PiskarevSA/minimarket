package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/PiskarevSA/minimarket/microservices/points/internal/domain/objects"
)

type Transaction struct {
	id              int32
	operation       objects.Operation
	orderId, userId uuid.UUID
	amount          objects.Amount
	timestamp       time.Time
}

func NewTransaction[AmountT string | pgtype.Numeric](
	id int32,
	operation string,
	orderId, userId uuid.UUID,
	amount AmountT,
	timestamp time.Time,
) (Transaction, error) {
	var (
		tx  Transaction
		err error
	)

	tx.operation, err = objects.NewOperation(operation)
	if err != nil {
		return Transaction{}, err
	}

	tx.amount, err = objects.NewAmount(amount)
	if err != nil {
		return Transaction{}, err
	}

	tx.orderId = orderId
	tx.userId = userId
	tx.timestamp = timestamp

	return tx, nil
}

func (t Transaction) Id() int32 {
	return t.id
}

func (t Transaction) Operation() objects.Operation {
	return t.operation
}

func (t Transaction) OrderId() uuid.UUID {
	return t.orderId
}

func (t Transaction) UserId() uuid.UUID {
	return t.userId
}

func (t Transaction) Amount() objects.Amount {
	return t.amount
}

func (t Transaction) Timestamp() time.Time {
	return t.timestamp
}
