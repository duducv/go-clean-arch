package persistenceadapters

import (
	"context"

	applicationrepository "github.com/duducv/go-clean-arch/internal/core/application/repository"
	"github.com/duducv/go-clean-arch/internal/core/domain/entity"
	"github.com/duducv/go-clean-arch/internal/infra/persistence/queries"
	"github.com/jackc/pgx/v5"
)

type TicketRepositorySQLAdapter struct {
	pgConn pgx.Tx
	trace  applicationrepository.TracingRepository
}

func NewTicketRepositorySQLAdapter(
	pgConn *pgx.Tx,
	trace applicationrepository.TracingRepository,
) entity.TicketRepository {
	return &TicketRepositorySQLAdapter{
		pgConn: *pgConn,
		trace:  trace,
	}
}

func (repository TicketRepositorySQLAdapter) Save(ctx context.Context, ticket entity.Ticket) error {
	span, ctx := repository.trace.StartSpan(ctx, "TicketRepositorySQLAdapter.Save", "infra")
	defer span.End()

	query := queries.NewTicketQuery(ctx, &repository.pgConn, repository.trace)
	err := query.Save(ctx, ticket)
	if err != nil {
		repository.trace.CaptureError(ctx, err)
		return err
	}
	return nil
}

func (repository TicketRepositorySQLAdapter) Update(ctx context.Context, ticket entity.Ticket) error {
	span, ctx := repository.trace.StartSpan(ctx, "TicketRepositorySQLAdapter.Save", "infra")
	defer span.End()

	query := queries.NewTicketQuery(ctx, &repository.pgConn, repository.trace)
	err := query.UpdateOne(ctx, ticket)
	if err != nil {
		repository.trace.CaptureError(ctx, err)
		return err
	}
	return nil
}
