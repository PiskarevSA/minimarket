package objects

import "errors"

type Operation string

const (
	OperationDeposit  Operation = "DEPOSIT"
	OperationWithdraw Operation = "WITHDRAW"
)

var ErrInvalidTransactionType = errors.New("invalid transaction type")

func NewOperation(value string) (Operation, error) {
	txType := Operation(value)

	if txType != OperationDeposit &&
		txType != OperationWithdraw {
		return "", ErrInvalidTransactionType
	}

	return txType, nil
}

func (o Operation) String() string {
	return string(o)
}
