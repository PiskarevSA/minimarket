package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
)

type Order struct {
	number     objects.OrderNumber
	userID     objects.UserID
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
	userID uuid.UUID,
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

	o.userID = objects.NewUserID(userID)
	o.uploadedAt = uploadedAt

	return o, nil
}

func (o Order) Number() objects.OrderNumber {
	return o.number
}

func (o Order) UserID() objects.UserID {
	return o.userID
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

func (o *Order) SetUserID(userID objects.UserID) {
	o.userID = userID
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

func (o Order) EqualUserID(other Order) bool {
	return o.userID == other.userID
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
