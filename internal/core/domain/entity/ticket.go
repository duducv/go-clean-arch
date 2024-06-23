package entity

import "github.com/lucsky/cuid"

type TicketStatus string

const (
	TicketStatusReserved TicketStatus = "reserved"
	TicketStatusApproved TicketStatus = "approved"
	TicketStatusCanceled TicketStatus = "canceled"
)

type Ticket struct {
	TicketID string
	EventID  string
	Email    string
	Status   TicketStatus
}

func NewTicket(
	eventID string,
	email string,
) Ticket {
	return Ticket{
		TicketID: cuid.New(),
		EventID:  eventID,
		Email:    email,
		Status:   TicketStatusReserved,
	}
}

func (t *Ticket) Approve() {
	t.Status = TicketStatusApproved
}

func (t *Ticket) Cancel() {
	t.Status = TicketStatusCanceled
}
