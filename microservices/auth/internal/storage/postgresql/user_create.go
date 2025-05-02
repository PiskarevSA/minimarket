package postgresql

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/storage/postgresql/sqlc/users"
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/usecases"
	"github.com/PiskarevSA/minimarket/pkg/pgcodes"
)

func handleCreateUserError(err error) error {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		if pgErr.Code == pgcodes.UniqueViolation {
			return usecases.ErrLoginAlreadyInUse
		}
	}

	return err
}

func (s *User) CreateUser(
	ctx context.Context,
	userId uuid.UUID,
	login, passwordHash string,
) error {
	err := s.query.CreateUser(ctx,
		users.CreateUserParams{
			Id:           userId,
			Login:        login,
			PasswordHash: passwordHash,
		},
	)
	if err != nil {
		return handleCreateUserError(err)
	}

	return nil
}
