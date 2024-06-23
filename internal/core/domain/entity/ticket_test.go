package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTicket_NewTicket(t *testing.T) {
	ticket := NewTicket("abc", "test@test.com")

	assert.Equal(t, ticket.EventID, "abc")
	assert.Equal(t, ticket.Email, "test@test.com")
	assert.NotEmpty(t, ticket.TicketID)
	assert.Equal(t, string(ticket.Status), "reserved")
}

func TestTicket_Approve(t *testing.T) {
	ticket := NewTicket("abc", "test@test.com")
	ticket.Approve()

	assert.Equal(t, string(ticket.Status), "approved")
}

func TestTicket_Cancel(t *testing.T) {
	ticket := NewTicket("abc", "test@test.com")
	ticket.Cancel()

	assert.Equal(t, string(ticket.Status), "canceled")
}
