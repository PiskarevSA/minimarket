package storage

type Error struct {
	Msg string
}

func (e *Error) Error() string {
	return e.Msg
}
