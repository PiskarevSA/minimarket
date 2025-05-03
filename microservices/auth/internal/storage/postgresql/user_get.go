package postgresql

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/storage"
)

func (s *User) GetUserCreds(
	ctx context.Context,
	login string,
) (*storage.UserCreds, error) {
	credsRow, err := s.query.GetUserCreds(ctx, login)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrUserNotFound
		}

		return nil, err
	}

	return &storage.UserCreds{
		UserId:       credsRow.UserId,
		PasswordHash: credsRow.PasswordHash,
	}, nil
}
