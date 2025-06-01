package objects

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Password []byte

type PasswordError Error

func (e *PasswordError) Error() string {
	return e.Message
}

const (
	MixPasswordLen = 12
	MaxPasswordLen = 32
)

var (
	ErrEmptyPassword    = &PasswordError{"empty password"}
	ErrPasswordTooShort = &PasswordError{"password too short"}
	ErrPasswordTooLong  = &PasswordError{"password too long"}
)

var NilPassword Password = nil

func NewPassword(value string) (Password, error) {
	passwordLen := len(value)
	if passwordLen == 0 {
		return NilPassword, ErrEmptyPassword
	}

	if passwordLen < MixPasswordLen {
		return NilPassword, ErrPasswordTooShort
	}

	if passwordLen > MaxPasswordLen {
		return NilPassword, ErrPasswordTooLong
	}

	return Password([]byte(value)), nil
}

func (o Password) String() string {
	return string(o)
}

func (o Password) Hash() ([]byte, error) {
	return bcrypt.GenerateFromPassword(o, bcrypt.DefaultCost)
}

func (o Password) IsHashMatch(hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, o)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
