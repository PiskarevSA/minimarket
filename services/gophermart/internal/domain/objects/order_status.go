package objects

import "errors"

type OrderStatus string

const (
	OrderStatusNew        OrderStatus = "NEW"
	OrderStatusProcessing OrderStatus = "PROCESSING"
	OrderStatusInvalid    OrderStatus = "INVALID"
	OrderStatusProcessed  OrderStatus = "PROCESSED"
)

var NullOrderStatus OrderStatus = ""

var ErrInvalidOrderStatus = errors.New("invalid order status")

func NewOrderStatus(value string) (OrderStatus, error) {
	status := OrderStatus(value)
	if status != OrderStatusNew && status != OrderStatusProcessing &&
		status != OrderStatusInvalid && status != OrderStatusProcessed {
		return NullOrderStatus, ErrInvalidOrderStatus
	}

	return status, nil
}

func (o OrderStatus) String() string {
	return string(o)
}

func (o OrderStatus) Equal(other OrderStatus) bool {
	return o == other
}
