package postgre

import (
	"context"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/dto"
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/storage"
)

const codeLoginConflict = 1

func (s *User) CreateUser(ctx context.Context, params dto.CreateUserParams) error {
	code, err := s.query.CreateUser(ctx, params.ToSqlc())
	if err != nil {
		return nil
	}

	if code == codeLoginConflict {
		return storage.ErrUserLoginAlreadyExists
	}

	return nil
}
