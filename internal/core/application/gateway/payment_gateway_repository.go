package gateway

import (
	"context"

	"github.com/duducv/go-clean-arch/internal/core/application/dto"
)

type PaymentGateway interface {
	CreateTransaction(ctx context.Context, input dto.PaymentTransactionInput) (*dto.PaymentTransactionOutput, error)
}
