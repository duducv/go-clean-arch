package entity

import "github.com/lucsky/cuid"

type Transaction struct {
	TransactionID string
	TicketID      string
	EventID       string
	TID           string
	Price         float64
	Status        string
}

func NewTransaction(
	ticketID string,
	eventID string,
	TID string,
	price float64,
	status string,
) Transaction {
	return Transaction{
		TransactionID: cuid.New(),
		TicketID:      ticketID,
		EventID:       eventID,
		TID:           TID,
		Price:         price,
		Status:        status,
	}
}
