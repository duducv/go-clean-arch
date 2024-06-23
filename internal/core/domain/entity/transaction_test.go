package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	transaction := NewTransaction("abc", "cde", "fgh", 100.0, "approved")
	assert.Equal(t, transaction.TicketID, "abc")
	assert.Equal(t, transaction.EventID, "cde")
	assert.Equal(t, transaction.TID, "fgh")
	assert.Equal(t, transaction.Price, 100.0)
	assert.Equal(t, transaction.Status, "approved")
}
