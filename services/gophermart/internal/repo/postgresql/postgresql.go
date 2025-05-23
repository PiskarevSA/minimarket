package postgresql

import "github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/gen/sqlc/postgresql"

type PostgreSql struct {
	querier *postgresql.Queries
}

func New(dbtx postgresql.DBTX) *PostgreSql {
	return &PostgreSql{
		querier: postgresql.New(dbtx),
	}
}
