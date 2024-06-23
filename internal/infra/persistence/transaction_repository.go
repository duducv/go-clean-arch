package persistenceadapters

import (
	"context"

	"github.com/duducv/go-clean-arch/internal/core/application/constants"
	applicationrepository "github.com/duducv/go-clean-arch/internal/core/application/repository"
	"github.com/duducv/go-clean-arch/internal/core/domain/entity"
	"github.com/duducv/go-clean-arch/internal/infra/persistence/queries"
	"github.com/jackc/pgx/v5"
)

type TransactionRepositorySQLAdapter struct {
	pgConn *pgx.Tx
	trace  applicationrepository.TracingRepository
}

func (repository *TransactionRepositorySQLAdapter) Save(ctx context.Context, transaction entity.Transaction) error {
	span, ctx := repository.trace.StartSpan(ctx, "TransactionRepositorySQLAdapter.Save", string(constants.InfraLayer))
	defer span.End()

	query := queries.NewTransactionQuery(ctx, repository.pgConn, repository.trace)
	err := query.Save(ctx, transaction)
	if err != nil {
		repository.trace.CaptureError(ctx, err)
		return err
	}
	return nil
}

func NewTransactionRepositorySQLAdapter(
	pgConn *pgx.Tx,
	trace applicationrepository.TracingRepository,
) entity.TransactionRepository {
	return &TransactionRepositorySQLAdapter{
		pgConn: pgConn,
		trace:  trace,
	}
}
