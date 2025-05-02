package usecases

import (
	"context"

	"github.com/google/uuid"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/storage"
)

type UserStorage interface {
	CreateUser(
		ctx context.Context,
		userId uuid.UUID,
		login, passwordHash string,
	) error
	GetUserCreds(
		ctx context.Context,
		login string,
	) (*storage.UserCreds, error)
}
