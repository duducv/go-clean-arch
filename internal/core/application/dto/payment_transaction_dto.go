package dto

type PaymentTransactionInput struct {
	Email           string
	CreditCardToken string
	Price           float64
}

type PaymentTransactionOutput struct {
	Status string
	TID    string
}
