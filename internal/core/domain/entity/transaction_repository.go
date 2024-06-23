package entity

import "context"

type TransactionRepository interface {
	Save(ctx context.Context, transaction Transaction) error
}
