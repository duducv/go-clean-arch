package queries

import (
	"context"

	"github.com/duducv/go-clean-arch/internal/core/application/constants"
	applicationrepository "github.com/duducv/go-clean-arch/internal/core/application/repository"
	"github.com/jackc/pgx/v5"
)

type EventQuery struct {
	pgConn pgx.Tx
	trace  applicationrepository.TracingRepository
}

func NewEventQuery(
	ctx context.Context, pgConn *pgx.Tx,
	trace applicationrepository.TracingRepository,
) EventQuery {
	tx, ok := ctx.Value(constants.TXKey).(*pgx.Tx)
	if ok {
		pgConn = tx
	}
	return EventQuery{
		pgConn: *pgConn,
		trace:  trace,
	}
}
