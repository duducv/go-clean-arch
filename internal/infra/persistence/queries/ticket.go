package queries

import (
	"context"

	"github.com/duducv/go-clean-arch/internal/core/application/constants"
	applicationrepository "github.com/duducv/go-clean-arch/internal/core/application/repository"
	"github.com/jackc/pgx/v5"
)

type TicketQuery struct {
	pgConn pgx.Tx
	trace  applicationrepository.TracingRepository
}

func NewTicketQuery(
	ctx context.Context, pgConn *pgx.Tx,
	trace applicationrepository.TracingRepository,
) TicketQuery {
	tx, ok := ctx.Value(constants.TXKey).(*pgx.Tx)
	if ok {
		pgConn = tx
	}
	return TicketQuery{
		pgConn: *pgConn,
		trace:  trace,
	}
}
