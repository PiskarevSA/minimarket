package usecases

import (
	"context"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/dto"
)

type UserStorage interface {
	CreateUser(ctx context.Context, params dto.CreateUserParams) error
}

type User struct {
	storage UserStorage
}

func NewUser(storage UserStorage) *User {
	return nil
}
