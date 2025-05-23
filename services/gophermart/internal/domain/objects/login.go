package objects

type Login string

type LoginError Error

func (e LoginError) Error() string {
	return e.Message
}

const (
	MinLoginLen = 5
	MaxLoginLen = 24
)

var (
	ErrEmptyLogin    = &LoginError{"empty login"}
	ErrLoginTooShort = &LoginError{"login too short"}
	ErrLoginTooLong  = &LoginError{"login too long"}
)

var NullLogin Login = ""

func NewLogin(value string) (Login, error) {
	loginLen := len(value)
	if loginLen == 0 {
		return NullLogin, ErrEmptyLogin
	}

	if loginLen < MinLoginLen {
		return NullLogin, ErrLoginTooShort
	}

	if loginLen > MaxLoginLen {
		return NullLogin, ErrLoginTooLong
	}

	return Login(value), nil
}

func (o Login) String() string {
	return string(o)
}

func (l Login) Equal(other Login) bool {
	return l == other
}
