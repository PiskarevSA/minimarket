package postgresql

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/PiskarevSA/minimarket/microservices/auth/internal/storage/postgresql/sqlc/users"
)

type User struct {
	pool  *pgxpool.Pool
	query *users.Queries
}

func NewUser(pool *pgxpool.Pool) *User {
	return &User{
		pool:  pool,
		query: users.New(pool),
	}
}
