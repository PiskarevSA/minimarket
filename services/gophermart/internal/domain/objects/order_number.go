package objects

import (
	"strconv"

	"github.com/github.com/PiskarevSA/minimarket/pkg/luhn"
)

type OrderNumber string

type OrderNumberError Error

func (e OrderNumberError) Error() string {
	return e.Message
}

var NullOrderNumber OrderNumber = ""

var (
	EmptyOrderNumber      = &OrderNumberError{"empty order number"}
	ErrInvalidOrderNumber = &OrderNumberError{"invalid order number"}
)

func NewOrderNumber(value string) (OrderNumber, error) {
	orderNumberLen := len(value)
	if orderNumberLen == 0 {
		return NullOrderNumber, EmptyOrderNumber
	}

	orderNumber, err := strconv.Atoi(value)
	if err != nil {
		return NullOrderNumber, ErrInvalidOrderNumber
	}

	ok := luhn.IsValid(orderNumber)

	if !ok {
		return NullOrderNumber, ErrInvalidOrderNumber
	}

	return OrderNumber(value), nil
}

func (o OrderNumber) String() string {
	return string(o)
}

func (o OrderNumber) Equal(other OrderNumber) bool {
	return o == other
}
