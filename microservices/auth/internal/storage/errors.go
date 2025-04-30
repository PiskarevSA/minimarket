package storage

import "errors"

const ErrMsgSomethingWentWrong = "something went wrong"

var (
	ErrUserLoginAlreadyExists = errors.New("uesr login already exists")
)
