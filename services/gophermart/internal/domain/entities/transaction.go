package entities

import (
	"time"

	"github.com/google/uuid"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
)

type Transaction struct {
	id           int32
	userID       objects.UserID
	orderNumber  objects.OrderNumber
	sum          objects.Amount
	opetaion     objects.Operation
	proccessedAt time.Time
}

var (
	NullTransaction = Transaction{}
	Nullransactions = []Transaction{}
)

func NewTransaction(
	id int32,
	userID uuid.UUID,
	orderNum string,
	sum string,
	operation string,
	processedAt time.Time,
) (Transaction, error) {
	var (
		tx  Transaction
		err error
	)

	tx.orderNumber, err = objects.NewOrderNumber(orderNum)
	if err != nil {
		return NullTransaction, err
	}

	tx.sum, err = objects.NewAmount(sum)
	if err != nil {
		return NullTransaction, err
	}

	tx.opetaion, err = objects.NewOperation(operation)
	if err != nil {
		return NullTransaction, err
	}

	tx.id = id
	tx.userID = objects.NewUserID(userID)
	tx.proccessedAt = processedAt

	return tx, nil
}

func (t Transaction) ID() int32 {
	return t.id
}

func (t Transaction) UserID() objects.UserID {
	return t.userID
}

func (t Transaction) OrderNumber() objects.OrderNumber {
	return t.orderNumber
}

func (t Transaction) Sum() objects.Amount {
	return t.sum
}

func (t Transaction) Operation() objects.Operation {
	return t.opetaion
}

func (t Transaction) ProcessedAt() time.Time {
	return t.proccessedAt
}

func (t *Transaction) SetID(id int32) {
	t.id = id
}

func (t *Transaction) SetUserID(userID objects.UserID) {
	t.userID = userID
}

func (t *Transaction) SetOrderNumber(orderNumber objects.OrderNumber) {
	t.orderNumber = orderNumber
}

func (t *Transaction) SetSum(sum objects.Amount) {
	t.sum = sum
}

func (t *Transaction) SetOperation(op objects.Operation) {
	t.opetaion = op
}

func (t *Transaction) SetProcessedAt(processedAt time.Time) {
	t.proccessedAt = processedAt
}
