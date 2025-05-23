package objects

import "github.com/github.com/PiskarevSA/minimarket/pkg/damm"

type OrderNumber string

type OrderNumberError Error

func (e OrderNumberError) Error() string {
	return e.Message
}

var NullOrderNumber OrderNumber = ""

var (
	EmptyOrderNumber         = &OrderNumberError{"empty order number"}
	ErrInvalidOrderNumber    = &OrderNumberError{"invalid order number"}
	ErrInvalidOrderNumberLen = &OrderNumberError{"invalid order number len"}
)

const OrderNubmerLen = 12

func NewOrderNumber(value string) (OrderNumber, error) {
	orderNumberLen := len(value)
	if orderNumberLen == 0 {
		return NullOrderNumber, EmptyOrderNumber
	}

	if orderNumberLen != OrderNubmerLen {
		return NullOrderNumber, ErrInvalidOrderNumberLen
	}

	ok, err := damm.Verify(value)
	if err != nil {
		return NullOrderNumber, ErrInvalidOrderNumber
	}

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
