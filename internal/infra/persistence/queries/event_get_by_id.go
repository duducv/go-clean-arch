package queries

import (
	"context"
	"database/sql"
	"errors"

	"github.com/duducv/go-clean-arch/internal/core/application/constants"
	"github.com/duducv/go-clean-arch/internal/core/domain/entity"
)

func (repository EventQuery) GetID(ctx context.Context, eventID string) (*entity.Event, error) {
	span, ctx := repository.trace.StartSpan(ctx, "EventQuery.GetID", string(constants.InfraLayer))
	defer span.End()
	stmt, err := repository.pgConn.Prepare(
		ctx,
		"getOneEvent",
		`
		SELECT
		  event.event_id,
		  event.description, 
		  event.price,
		  event.capacity
		  FROM public.event
		WHERE event.event_id = $1
		LIMIT 1
		  `,
	)
	if err != nil {
		repository.trace.CaptureError(ctx, err)
		return nil, err
	}
	eventReturn := entity.Event{}
	err = repository.pgConn.QueryRow(ctx, stmt.SQL, eventID).Scan(
		&eventReturn.EventID,
		&eventReturn.Description,
		&eventReturn.Price,
		&eventReturn.Capacity,
	)
	if err != nil {
		repository.trace.CaptureError(ctx, err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &eventReturn, nil
}
