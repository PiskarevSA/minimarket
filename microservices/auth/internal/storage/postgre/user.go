package postgre

import (
	"github.com/PiskarevSA/minimarket/microservices/auth/internal/storage/postgre/sqlc/users"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	pool  *pgxpool.Conn
	query *users.Queries
}

func NewUser() *User {
	return nil
}
