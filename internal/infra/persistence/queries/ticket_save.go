package queries

import (
	"context"

	"github.com/duducv/go-clean-arch/internal/core/domain/entity"
)

func (repository TicketQuery) Save(
	ctx context.Context,
	input entity.Ticket,
) error {
	span, ctx := repository.trace.StartSpan(ctx, "TicketQuery.Save", "infra")
	defer span.End()
	stmt, err := repository.pgConn.Prepare(ctx,
		"saveTicket",
		`INSERT INTO public.ticket (ticket_id, event_id, email, status)
		 VALUES ($1, $2, $3, $4)
		 `,
	)
	if err != nil {
		repository.trace.CaptureError(ctx, err)
		return err
	}
	_, err = repository.pgConn.Exec(
		ctx,
		stmt.SQL,
		input.TicketID,
		input.EventID,
		input.Email,
		input.Status,
	)
	if err != nil {
		repository.trace.CaptureError(ctx, err)
		return err
	}
	return nil
}
