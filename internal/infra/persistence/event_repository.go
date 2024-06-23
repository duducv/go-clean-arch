package persistenceadapters

import (
	"context"

	"github.com/duducv/go-clean-arch/internal/core/application/constants"
	applicationrepository "github.com/duducv/go-clean-arch/internal/core/application/repository"
	"github.com/duducv/go-clean-arch/internal/core/domain/entity"
	"github.com/duducv/go-clean-arch/internal/infra/persistence/queries"
	"github.com/jackc/pgx/v5"
)

type EventRepositorySQLAdapter struct {
	pgConn pgx.Tx
	trace  applicationrepository.TracingRepository
}

func NewEventRepositorySQLAdapter(
	pgConn *pgx.Tx,
	trace applicationrepository.TracingRepository,
) entity.EventRepository {
	return &EventRepositorySQLAdapter{
		pgConn: *pgConn,
		trace:  trace,
	}
}

func (repository *EventRepositorySQLAdapter) Get(
	ctx context.Context,
	eventID string,
) (*entity.Event, error) {
	span, ctx := repository.trace.StartSpan(
		ctx, "EventRepositorySQLAdapter.Get",
		string(constants.InfraLayer),
	)
	defer span.End()

	event, err := queries.
		NewEventQuery(ctx, &repository.pgConn, repository.trace).
		GetID(ctx, eventID)
	if err != nil {
		repository.trace.CaptureError(ctx, err)
		return nil, err
	}
	return event, nil
}
