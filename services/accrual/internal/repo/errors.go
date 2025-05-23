package repo

type Error struct {
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

var (
	ErrRewardStrategyAlreadyExists = &Error{"reward strategy already exists"}
)
