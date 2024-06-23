package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/duducv/go-clean-arch/internal/core/application/dto"
	"github.com/duducv/go-clean-arch/internal/core/domain/entity"
	"github.com/duducv/go-clean-arch/internal/infra/tracing"
	"github.com/duducv/go-clean-arch/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewTicketPurchaseUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketRepoMock := mock.NewMockTicketRepository(ctrl)
	eventRepoMock := mock.NewMockEventRepository(ctrl)
	transactionRepoMock := mock.NewMockTransactionRepository(ctrl)
	paymentGatewayMock := mock.NewMockPaymentGateway(ctrl)
	tracingMock := mock.NewMockTracingRepository(ctrl)
	uowRepoMock := mock.NewMockUnitOfWorkRepository(ctrl)
	usecase := NewTicketPurchaseUseCase(
		ticketRepoMock,
		eventRepoMock,
		transactionRepoMock,
		paymentGatewayMock,
		uowRepoMock,
		tracingMock,
	)
	assert.NotEmpty(t, usecase)
}

func TestTicketPurchaseUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketRepoMock := mock.NewMockTicketRepository(ctrl)
	eventRepoMock := mock.NewMockEventRepository(ctrl)
	transactionRepoMock := mock.NewMockTransactionRepository(ctrl)
	paymentGatewayMock := mock.NewMockPaymentGateway(ctrl)
	tracingMock := tracing.NewNotracingAdapter()
	uowRepoMock := mock.NewMockUnitOfWorkRepository(ctrl)
	usecase := NewTicketPurchaseUseCase(
		ticketRepoMock,
		eventRepoMock,
		transactionRepoMock,
		paymentGatewayMock,
		uowRepoMock,
		tracingMock,
	)
	ctx := context.TODO()
	uowRepoMock.EXPECT().Begin(ctx).Return(ctx, nil)
	eventRepoMock.EXPECT().Get(ctx, gomock.Any()).Return(&entity.Event{}, nil)
	ticketRepoMock.EXPECT().Save(ctx, gomock.Any())
	paymentGatewayMock.EXPECT().CreateTransaction(ctx, gomock.Any()).Return(&dto.PaymentTransactionOutput{}, nil)
	transactionRepoMock.EXPECT().Save(ctx, gomock.Any()).Return(nil)
	ticketRepoMock.EXPECT().Update(ctx, gomock.Any()).Return(nil)
	uowRepoMock.EXPECT().Commit(ctx).Return(nil)
	_, err := usecase.Execute(ctx, dto.TicketPurchaseInput{
		EventId:         gomock.Any().String(),
		Email:           gomock.Any().String(),
		CreditCardToken: gomock.Any().String(),
	})

	assert.Nil(t, err)
}

func TestTicketPurchaseUseCase_ExecuteTicketNotFoundError(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketRepoMock := mock.NewMockTicketRepository(ctrl)
	eventRepoMock := mock.NewMockEventRepository(ctrl)
	transactionRepoMock := mock.NewMockTransactionRepository(ctrl)
	paymentGatewayMock := mock.NewMockPaymentGateway(ctrl)
	tracingMock := tracing.NewNotracingAdapter()
	uowRepoMock := mock.NewMockUnitOfWorkRepository(ctrl)
	usecase := NewTicketPurchaseUseCase(
		ticketRepoMock,
		eventRepoMock,
		transactionRepoMock,
		paymentGatewayMock,
		uowRepoMock,
		tracingMock,
	)
	ctx := context.TODO()
	uowRepoMock.EXPECT().Begin(ctx).Return(ctx, nil)
	eventRepoMock.EXPECT().Get(ctx, gomock.Any()).Return(nil, errors.New("Not found"))
	uowRepoMock.EXPECT().Rollback(ctx).Return(nil)
	_, err := usecase.Execute(ctx, dto.TicketPurchaseInput{
		EventId:         gomock.Any().String(),
		Email:           gomock.Any().String(),
		CreditCardToken: gomock.Any().String(),
	})

	assert.NotNil(t, err)
}

func TestTicketPurchaseUseCase_ExecuteTicketSaveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketRepoMock := mock.NewMockTicketRepository(ctrl)
	eventRepoMock := mock.NewMockEventRepository(ctrl)
	transactionRepoMock := mock.NewMockTransactionRepository(ctrl)
	paymentGatewayMock := mock.NewMockPaymentGateway(ctrl)
	tracingMock := tracing.NewNotracingAdapter()
	uowRepoMock := mock.NewMockUnitOfWorkRepository(ctrl)
	usecase := NewTicketPurchaseUseCase(
		ticketRepoMock,
		eventRepoMock,
		transactionRepoMock,
		paymentGatewayMock,
		uowRepoMock,
		tracingMock,
	)
	ctx := context.TODO()
	uowRepoMock.EXPECT().Begin(ctx).Return(ctx, nil)
	eventRepoMock.EXPECT().Get(ctx, gomock.Any()).Return(&entity.Event{}, nil)
	ticketRepoMock.EXPECT().Save(ctx, gomock.Any()).Return(errors.New(""))
	uowRepoMock.EXPECT().Rollback(ctx).Return(nil)
	_, err := usecase.Execute(ctx, dto.TicketPurchaseInput{
		EventId:         gomock.Any().String(),
		Email:           gomock.Any().String(),
		CreditCardToken: gomock.Any().String(),
	})

	assert.NotNil(t, err)
}

func TestTicketPurchaseUseCase_ExecuteCreateTransactioError(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketRepoMock := mock.NewMockTicketRepository(ctrl)
	eventRepoMock := mock.NewMockEventRepository(ctrl)
	transactionRepoMock := mock.NewMockTransactionRepository(ctrl)
	paymentGatewayMock := mock.NewMockPaymentGateway(ctrl)
	tracingMock := tracing.NewNotracingAdapter()
	uowRepoMock := mock.NewMockUnitOfWorkRepository(ctrl)
	usecase := NewTicketPurchaseUseCase(
		ticketRepoMock,
		eventRepoMock,
		transactionRepoMock,
		paymentGatewayMock,
		uowRepoMock,
		tracingMock,
	)
	ctx := context.TODO()
	uowRepoMock.EXPECT().Begin(ctx).Return(ctx, nil)
	eventRepoMock.EXPECT().Get(ctx, gomock.Any()).Return(&entity.Event{}, nil)
	ticketRepoMock.EXPECT().Save(ctx, gomock.Any())
	paymentGatewayMock.EXPECT().CreateTransaction(ctx, gomock.Any()).Return(nil, errors.New(""))
	uowRepoMock.EXPECT().Rollback(ctx).Return(nil)
	_, err := usecase.Execute(ctx, dto.TicketPurchaseInput{
		EventId:         gomock.Any().String(),
		Email:           gomock.Any().String(),
		CreditCardToken: gomock.Any().String(),
	})

	assert.NotNil(t, err)
}

func TestTicketPurchaseUseCase_ExecuteTransactionSaveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketRepoMock := mock.NewMockTicketRepository(ctrl)
	eventRepoMock := mock.NewMockEventRepository(ctrl)
	transactionRepoMock := mock.NewMockTransactionRepository(ctrl)
	paymentGatewayMock := mock.NewMockPaymentGateway(ctrl)
	tracingMock := tracing.NewNotracingAdapter()
	uowRepoMock := mock.NewMockUnitOfWorkRepository(ctrl)
	usecase := NewTicketPurchaseUseCase(
		ticketRepoMock,
		eventRepoMock,
		transactionRepoMock,
		paymentGatewayMock,
		uowRepoMock,
		tracingMock,
	)
	ctx := context.TODO()
	uowRepoMock.EXPECT().Begin(ctx).Return(ctx, nil)
	eventRepoMock.EXPECT().Get(ctx, gomock.Any()).Return(&entity.Event{}, nil)
	ticketRepoMock.EXPECT().Save(ctx, gomock.Any())
	paymentGatewayMock.EXPECT().CreateTransaction(ctx, gomock.Any()).Return(&dto.PaymentTransactionOutput{}, nil)
	transactionRepoMock.EXPECT().Save(ctx, gomock.Any()).Return(errors.New(""))
	uowRepoMock.EXPECT().Rollback(ctx).Return(nil)
	_, err := usecase.Execute(ctx, dto.TicketPurchaseInput{
		EventId:         gomock.Any().String(),
		Email:           gomock.Any().String(),
		CreditCardToken: gomock.Any().String(),
	})

	assert.NotNil(t, err)
}

func TestTicketPurchaseUseCase_ExecuteStatusApproved(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketRepoMock := mock.NewMockTicketRepository(ctrl)
	eventRepoMock := mock.NewMockEventRepository(ctrl)
	transactionRepoMock := mock.NewMockTransactionRepository(ctrl)
	paymentGatewayMock := mock.NewMockPaymentGateway(ctrl)
	tracingMock := tracing.NewNotracingAdapter()
	uowRepoMock := mock.NewMockUnitOfWorkRepository(ctrl)
	usecase := NewTicketPurchaseUseCase(
		ticketRepoMock,
		eventRepoMock,
		transactionRepoMock,
		paymentGatewayMock,
		uowRepoMock,
		tracingMock,
	)
	ctx := context.TODO()
	uowRepoMock.EXPECT().Begin(ctx).Return(ctx, nil)
	eventRepoMock.EXPECT().Get(ctx, gomock.Any()).Return(&entity.Event{}, nil)
	ticketRepoMock.EXPECT().Save(ctx, gomock.Any())
	paymentGatewayMock.EXPECT().CreateTransaction(ctx, gomock.Any()).Return(&dto.PaymentTransactionOutput{
		Status: "approved",
	}, nil)
	transactionRepoMock.EXPECT().Save(ctx, gomock.Any()).Return(nil)
	ticketRepoMock.EXPECT().Update(ctx, gomock.Any()).Return(nil)
	uowRepoMock.EXPECT().Commit(ctx).Return(nil)
	_, err := usecase.Execute(ctx, dto.TicketPurchaseInput{
		EventId:         gomock.Any().String(),
		Email:           gomock.Any().String(),
		CreditCardToken: gomock.Any().String(),
	})

	assert.Nil(t, err)
}

func TestTicketPurchaseUseCase_ExecuteTicketUpdateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketRepoMock := mock.NewMockTicketRepository(ctrl)
	eventRepoMock := mock.NewMockEventRepository(ctrl)
	transactionRepoMock := mock.NewMockTransactionRepository(ctrl)
	paymentGatewayMock := mock.NewMockPaymentGateway(ctrl)
	tracingMock := tracing.NewNotracingAdapter()
	uowRepoMock := mock.NewMockUnitOfWorkRepository(ctrl)
	usecase := NewTicketPurchaseUseCase(
		ticketRepoMock,
		eventRepoMock,
		transactionRepoMock,
		paymentGatewayMock,
		uowRepoMock,
		tracingMock,
	)
	ctx := context.TODO()
	uowRepoMock.EXPECT().Begin(ctx).Return(ctx, nil)
	eventRepoMock.EXPECT().Get(ctx, gomock.Any()).Return(&entity.Event{}, nil)
	ticketRepoMock.EXPECT().Save(ctx, gomock.Any())
	paymentGatewayMock.EXPECT().CreateTransaction(ctx, gomock.Any()).Return(&dto.PaymentTransactionOutput{}, nil)
	transactionRepoMock.EXPECT().Save(ctx, gomock.Any()).Return(nil)
	ticketRepoMock.EXPECT().Update(ctx, gomock.Any()).Return(errors.New(""))
	uowRepoMock.EXPECT().Rollback(ctx).Return(nil)
	_, err := usecase.Execute(ctx, dto.TicketPurchaseInput{
		EventId:         gomock.Any().String(),
		Email:           gomock.Any().String(),
		CreditCardToken: gomock.Any().String(),
	})

	assert.NotNil(t, err)
}

func TestTicketPurchaseUseCase_ExecuteUowBeginError(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketRepoMock := mock.NewMockTicketRepository(ctrl)
	eventRepoMock := mock.NewMockEventRepository(ctrl)
	transactionRepoMock := mock.NewMockTransactionRepository(ctrl)
	paymentGatewayMock := mock.NewMockPaymentGateway(ctrl)
	tracingMock := tracing.NewNotracingAdapter()
	uowRepoMock := mock.NewMockUnitOfWorkRepository(ctrl)
	usecase := NewTicketPurchaseUseCase(
		ticketRepoMock,
		eventRepoMock,
		transactionRepoMock,
		paymentGatewayMock,
		uowRepoMock,
		tracingMock,
	)
	ctx := context.TODO()
	uowRepoMock.EXPECT().Begin(ctx).Return(ctx, errors.New(""))
	uowRepoMock.EXPECT().Rollback(ctx).Return(nil)
	_, err := usecase.Execute(ctx, dto.TicketPurchaseInput{
		EventId:         gomock.Any().String(),
		Email:           gomock.Any().String(),
		CreditCardToken: gomock.Any().String(),
	})

	assert.NotNil(t, err)
}

func TestTicketPurchaseUseCase_ExecuteUowCommitError(t *testing.T) {
	ctrl := gomock.NewController(t)
	ticketRepoMock := mock.NewMockTicketRepository(ctrl)
	eventRepoMock := mock.NewMockEventRepository(ctrl)
	transactionRepoMock := mock.NewMockTransactionRepository(ctrl)
	paymentGatewayMock := mock.NewMockPaymentGateway(ctrl)
	tracingMock := tracing.NewNotracingAdapter()
	uowRepoMock := mock.NewMockUnitOfWorkRepository(ctrl)
	usecase := NewTicketPurchaseUseCase(
		ticketRepoMock,
		eventRepoMock,
		transactionRepoMock,
		paymentGatewayMock,
		uowRepoMock,
		tracingMock,
	)
	ctx := context.TODO()
	uowRepoMock.EXPECT().Begin(ctx).Return(ctx, nil)
	eventRepoMock.EXPECT().Get(ctx, gomock.Any()).Return(&entity.Event{}, nil)
	ticketRepoMock.EXPECT().Save(ctx, gomock.Any())
	paymentGatewayMock.EXPECT().CreateTransaction(ctx, gomock.Any()).Return(&dto.PaymentTransactionOutput{
		Status: "approved",
	}, nil)
	transactionRepoMock.EXPECT().Save(ctx, gomock.Any()).Return(nil)
	ticketRepoMock.EXPECT().Update(ctx, gomock.Any()).Return(nil)
	uowRepoMock.EXPECT().Commit(ctx).Return(errors.New(""))
	uowRepoMock.EXPECT().Rollback(ctx).Return(nil)
	_, err := usecase.Execute(ctx, dto.TicketPurchaseInput{
		EventId:         gomock.Any().String(),
		Email:           gomock.Any().String(),
		CreditCardToken: gomock.Any().String(),
	})

	assert.NotNil(t, err)
}
