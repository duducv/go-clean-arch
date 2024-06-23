package queries

import (
	"context"

	"github.com/duducv/go-clean-arch/internal/core/application/constants"
	applicationrepository "github.com/duducv/go-clean-arch/internal/core/application/repository"
	"github.com/jackc/pgx/v5"
)

type TransactionQuery struct {
	pgConn pgx.Tx
	trace  applicationrepository.TracingRepository
}

func NewTransactionQuery(
	ctx context.Context, pgConn *pgx.Tx,
	trace applicationrepository.TracingRepository,
) TransactionQuery {
	tx, ok := ctx.Value(constants.TXKey).(*pgx.Tx)
	if ok {
		pgConn = tx
	}
	return TransactionQuery{
		pgConn: *pgConn,
		trace:  trace,
	}
}
