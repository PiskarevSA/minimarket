package entities

import (
	"time"

	"github.com/google/uuid"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
)

type Account struct {
	id           objects.UserId
	login        objects.Login
	passwordHash []byte
	createdAt    time.Time
	updatedAt    time.Time
}

var NullAccount = Account{}

func NewAccount(
	id uuid.UUID,
	login string,
	passwordHash []byte,
	createdAt time.Time,
	updatedAt time.Time,
) (Account, error) {
	var (
		a   Account
		err error
	)

	a.login, err = objects.NewLogin(login)
	if err != nil {
		return NullAccount, err
	}

	a.id = objects.NewUserId(id)
	a.passwordHash = passwordHash
	a.createdAt = createdAt
	a.updatedAt = updatedAt

	return a, nil
}

func (a Account) Id() objects.UserId {
	return a.id
}

func (a Account) Login() objects.Login {
	return a.login
}

func (a Account) PasswordHash() []byte {
	return a.passwordHash
}

func (a Account) CreatedAt() time.Time {
	return a.createdAt
}

func (a Account) UpdatedAt() time.Time {
	return a.updatedAt
}

func (a *Account) SetId(id objects.UserId) {
	a.id = id
}

func (a *Account) SetLogin(login objects.Login) {
	a.login = login
}

func (a *Account) SetPasswordHash(passwordHash []byte) {
	a.passwordHash = passwordHash
}

func (a *Account) SetCreatedAt(t time.Time) {
	a.createdAt = t
}

func (a *Account) SetUpdatedAt(t time.Time) {
	a.updatedAt = t
}
