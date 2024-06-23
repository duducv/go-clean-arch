package queries

import (
	"context"

	"github.com/duducv/go-clean-arch/internal/core/application/constants"
	"github.com/duducv/go-clean-arch/internal/core/domain/entity"
)

func (repository TicketQuery) UpdateOne(
	ctx context.Context,
	input entity.Ticket,
) error {
	span, ctx := repository.trace.StartSpan(ctx, "TicketQuery.UpdateOne", string(constants.InfraLayer))
	defer span.End()
	stmt, err := repository.pgConn.Prepare(ctx,
		"updateOneTicket",
		`UPDATE public.ticket SET status = $1
		 WHERE ticket_id = $2`,
	)
	if err != nil {
		repository.trace.CaptureError(ctx, err)
		return err
	}
	_, err = repository.pgConn.Exec(
		ctx,
		stmt.SQL,
		input.Status,
		input.TicketID,
	)
	if err != nil {
		repository.trace.CaptureError(ctx, err)
		return err
	}
	return nil
}
