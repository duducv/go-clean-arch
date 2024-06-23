package dto

type TicketPurchaseInput struct {
	EventId         string `json:"eventId"`
	Email           string `json:"email"`
	CreditCardToken string `json:"creditCardToken"`
}

type TicketPurchaseOutput struct {
	TicketId string  `json:"ticketId"`
	Status   string  `json:"status"`
	TID      string  `json:"tid"`
	Price    float64 `json:"price"`
}
