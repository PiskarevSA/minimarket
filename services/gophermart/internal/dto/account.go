package dto

import (
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/entities"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/gen/sqlc/postgresql"
)

func GetAccountByUserIdToAccount(
	userId objects.UserId,
	row postgresql.GetAccountByUserIdRow,
) entities.Account {
	var account entities.Account

	account.SetId(userId)

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

	userId := objects.UserId(row.Id)
	account.SetId(userId)

	account.SetLogin(login)

	passwordHash := []byte(row.PasswordHash)
	account.SetPasswordHash(passwordHash)

	account.SetCreatedAt(row.CreatedAt)
	account.SetUpdatedAt(row.UpdatedAt)

	return account
}
