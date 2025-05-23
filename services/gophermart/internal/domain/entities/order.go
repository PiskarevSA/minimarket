package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
)

type Order struct {
	number     objects.OrderNumber
	userId     objects.UserId
	status     objects.OrderStatus
	accrual    objects.Amount
	uploadedAt time.Time
}

var (
	NullOrder  = Order{}
	NullOrders = []Order{}
)

func NewOrder(
	number string,
	userId uuid.UUID,
	status string,
	accrual string,
	uploadedAt time.Time,
) (Order, error) {
	var (
		o   Order
		err error
	)

	o.number, err = objects.NewOrderNumber(number)
	if err != nil {
		return NullOrder, err
	}

	o.status, err = objects.NewOrderStatus(status)
	if err != nil {
		return NullOrder, err
	}

	o.accrual, err = objects.NewAmount(accrual)
	if err != nil {
		return NullOrder, err
	}

	o.userId = objects.NewUserId(userId)
	o.uploadedAt = uploadedAt

	return o, nil
}

func (o Order) Number() objects.OrderNumber {
	return o.number
}

func (o Order) UserId() objects.UserId {
	return o.userId
}

func (o Order) Status() objects.OrderStatus {
	return o.status
}

func (o Order) Accrual() objects.Amount {
	return o.accrual
}

func (o Order) HasAccrual() bool {
	dec := o.accrual.Decimal()

	return dec.GreaterThan(decimal.Zero)
}

func (o Order) UploadedAt() time.Time {
	return o.uploadedAt
}

func (o *Order) SetNumber(number objects.OrderNumber) {
	o.number = number
}

func (o *Order) SetUserId(userId objects.UserId) {
	o.userId = userId
}

func (o *Order) SetStatus(status objects.OrderStatus) {
	o.status = status
}

func (o *Order) SetAccrual(amount objects.Amount) {
	o.accrual = amount
}

func (o *Order) SetUploadedAt(uploadedAt time.Time) {
	o.uploadedAt = uploadedAt
}

func (o Order) EqualNumber(other Order) bool {
	return o.number.Equal(other.number)
}

func (o Order) EqualUserId(other Order) bool {
	return o.userId == other.userId
}

func (o Order) EqualStatus(other Order) bool {
	return o.status.Equal(other.status)
}

func (o Order) EqualAccrual(other Order) bool {
	return o.accrual.Equal(other.accrual)
}

func (o Order) EqualUploadedAt(other Order) bool {
	return o.uploadedAt.Equal(other.uploadedAt)
}
