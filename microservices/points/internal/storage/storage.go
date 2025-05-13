package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"

	sqlc "github.com/PiskarevSA/minimarket/microservices/points/internal/gen/sqlc/postgresql"
	"github.com/PiskarevSA/minimarket/pkg/pgx/transactor"
)

type Storage struct {
	sqlQuerier    *sqlc.Queries
	sqlTransactor *transactor.Transactor[*sqlc.Queries]
}

func New(driver *pgxpool.Pool) *Storage {
	sqlQuerier := sqlc.New(driver)

	return &Storage{
		sqlQuerier:    sqlQuerier,
		sqlTransactor: transactor.New(driver, sqlQuerier.WithTx),
	}
}
