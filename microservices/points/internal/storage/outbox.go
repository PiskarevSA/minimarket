package storage

import (
	"context"

	sqlc "github.com/PiskarevSA/minimarket/microservices/points/internal/gen/sqlc/postgresql"
)

func (s *Storage) createOutbox(
	ctx context.Context, query *sqlc.Queries,
	event string,
	payload []byte,
	createdBy string,
) error {
	return query.CreateOutbox(
		ctx, sqlc.CreateOutboxParams{
			Event:     event,
			Payload:   payload,
			CreatedBy: createdBy,
			UpdatedBy: createdBy,
		},
	)
}
