package postgresql

import "github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/gen/sqlc/postgresql"

type PostgreSQL struct {
	querier *postgresql.Queries
}

func New(dbtx postgresql.DBTX) *PostgreSQL {
	return &PostgreSQL{
		querier: postgresql.New(dbtx),
	}
}
