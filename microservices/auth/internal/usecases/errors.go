package usecases

type UsecaseError struct{ Msg string }

func (e *UsecaseError) Error() string {
	return e.Msg
}

var ErrInvalidLoginOrPassword = &UsecaseError{
	"invalid login or password",
}

var ErrLoginAlreadyInUse = &UsecaseError{
	"login already in use",
}
