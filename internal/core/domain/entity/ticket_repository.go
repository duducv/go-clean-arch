package entity

import "context"

type TicketRepository interface {
	Save(ctx context.Context, ticket Ticket) error
	Update(ctx context.Context, ticket Ticket) error
}
