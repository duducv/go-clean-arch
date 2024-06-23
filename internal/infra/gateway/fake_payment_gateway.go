package gatewayadapters

import (
	"context"

	"github.com/duducv/go-clean-arch/internal/core/application/dto"
	"github.com/duducv/go-clean-arch/internal/core/application/gateway"
)

type FakePaymentGateway struct{}

func (f FakePaymentGateway) CreateTransaction(ctx context.Context, input dto.PaymentTransactionInput) (*dto.PaymentTransactionOutput, error) {
	return &dto.PaymentTransactionOutput{
		Status: "approved",
		TID:    "12345678",
	}, nil
}

func NewFakePaymentGateway() gateway.PaymentGateway {
	return FakePaymentGateway{}
}
