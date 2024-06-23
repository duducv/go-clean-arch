package queries

import (
	"context"

	"github.com/duducv/go-clean-arch/internal/core/domain/entity"
)

// create table public.transaction (
//     transaction_id text,
//     ticket_id text,
//     event_id text,
//     tid text,
//     price numeric,
//     status text
// );

func (repository TransactionQuery) Save(
	ctx context.Context,
	input entity.Transaction,
) error {
	span, ctx := repository.trace.StartSpan(ctx, "TransactionQuery.Save", "infra")
	defer span.End()
	stmt, err := repository.pgConn.Prepare(ctx,
		"saveTransaction",
		`INSERT INTO public.transaction (transaction_id, ticket_id, event_id, tid, price, status)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 `,
	)
	if err != nil {
		repository.trace.CaptureError(ctx, err)
		return err
	}
	_, err = repository.pgConn.Exec(
		ctx,
		stmt.SQL,
		input.TransactionID,
		input.TicketID,
		input.EventID,
		input.TID,
		input.Price,
		input.Status,
	)
	if err != nil {
		repository.trace.CaptureError(ctx, err)
		return err
	}
	return nil
}
