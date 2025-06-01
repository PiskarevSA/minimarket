package repo

type Error struct {
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

var (
	ErrLoginAlreadyInUse          = &Error{"login already in use"}
	ErrLoginNotExists             = &Error{"login not exists"}
	ErrNoBalanceFound             = &Error{"no balance found"}
	ErrNoOrdersFound              = &Error{"no orders found"}
	ErrNoTransactionsFound        = &Error{"no transactions found"}
	ErrNotEnoughtBalance          = &Error{"not enought balance"}
	ErrOrderAlreadyCreatedByOther = &Error{"order already created by other"}
	ErrOrderAlreadyCreatedByUser  = &Error{"order already created by user"}
)
