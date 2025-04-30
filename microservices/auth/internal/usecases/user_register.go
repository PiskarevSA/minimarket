package usecases

import (
	"context"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/dto"
	"github.com/google/uuid"
)

func (u *User) Register(ctx context.Context, login, password string) (string, error) {
	u.storage.CreateUser(ctx,
		dto.CreateUserParams{
			Id:           uuid.New(),
			Login:        login,
			PasswordHash: password,
			PasswordSalt: password,
		})
	return "", nil
}
