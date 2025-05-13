package transactor

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type driver interface {
	BeginTx(ctx context.Context, otps pgx.TxOptions) (pgx.Tx, error)
}

type Transactor[Querier any] struct {
	driver      driver
	querySource func(tx pgx.Tx) Querier
}

func New[Querier any](
	driver driver,
	querySource func(pgxTx pgx.Tx) Querier,
) *Transactor[Querier] {
	return &Transactor[Querier]{
		driver:      driver,
		querySource: querySource,
	}
}

func (t *Transactor[Querier]) Transact(
	ctx context.Context,
	opts pgx.TxOptions,
	fn func(q Querier) error,
) error {
	tx, err := t.driver.BeginTx(ctx, opts)
	if err != nil {
		return err
	}

	defer func() {
		p := recover()
		if p != nil {
			// For simplicity we assume that the rollback will be executed in
			// 100% cases. For production use it's needed to implement a
			// mechanism to handle rollback errors. Descriptions of strategies
			// that help to deal with such errors can be in this paper:
			// https://nthu-datalab.github.io/db/slides/10_Transaction_Recovery.pdf
			tx.Rollback(ctx)
			panic(p)
		}
	}()

	qtx := t.querySource(tx)

	err = fn(qtx)
	if err != nil {
		tx.Rollback(ctx)

		return err
	}

	return tx.Commit(ctx)
}
