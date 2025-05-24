package dto

import (
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/entities"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/gen/sqlc/postgresql"
)

func GetAccountByUserIDToAccount(
	userID objects.UserID,
	row postgresql.GetAccountByUserIdRow,
) entities.Account {
	var account entities.Account

	account.SetID(userID)

	login := objects.Login(row.Login)
	account.SetLogin(login)

	passwordHash := []byte(row.PasswordHash)
	account.SetPasswordHash(passwordHash)

	account.SetCreatedAt(row.CreatedAt)
	account.SetUpdatedAt(row.UpdatedAt)

	return account
}

func GetAccountByLoginToAccount(
	login objects.Login,
	row postgresql.GetAccountByLoginRow,
) entities.Account {
	var account entities.Account

	userID := objects.UserID(row.Id)
	account.SetID(userID)

	account.SetLogin(login)

	passwordHash := []byte(row.PasswordHash)
	account.SetPasswordHash(passwordHash)

	account.SetCreatedAt(row.CreatedAt)
	account.SetUpdatedAt(row.UpdatedAt)

	return account
}
