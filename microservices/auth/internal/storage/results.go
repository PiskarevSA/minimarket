package storage

import "github.com/google/uuid"

type UserCreds struct {
	UserId       uuid.UUID
	PasswordHash string
}
