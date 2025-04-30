package dto

import (
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/storage/postgre/sqlc/users"
	"github.com/google/uuid"
)

type CreateUserParams struct {
	Id           uuid.UUID
	Login        string
	PasswordHash string
	PasswordSalt string
}

func (p CreateUserParams) ToSqlc() users.CreateUserParams {
	return users.CreateUserParams{
		Id:           p.Id,
		Login:        p.Login,
		PasswordHash: p.PasswordHash,
		PasswordSalt: p.PasswordSalt,
	}
}
