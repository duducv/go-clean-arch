package usecases

import (
	"context"

	"github.com/duducv/go-clean-arch/internal/core/application/constants"
	"github.com/duducv/go-clean-arch/internal/core/application/dto"
	"github.com/duducv/go-clean-arch/internal/core/application/gateway"
	applicationrepository "github.com/duducv/go-clean-arch/internal/core/application/repository"
	"github.com/duducv/go-clean-arch/internal/core/domain/entity"
)

type TicketPurchaseUseCase struct {
	ticketRepository      entity.TicketRepository
	eventRepository       entity.EventRepository
	transactionRepository entity.TransactionRepository
	paymentGateway        gateway.PaymentGateway
	unitOfWorkRepository  applicationrepository.UnitOfWorkRepository
	trace                 applicationrepository.TracingRepository
}

func NewTicketPurchaseUseCase(
	ticketRepository entity.TicketRepository,
	eventRepository entity.EventRepository,
	transactionRepository entity.TransactionRepository,
	paymentGateway gateway.PaymentGateway,
	unitOfWorkRepository applicationrepository.UnitOfWorkRepository,
	trace applicationrepository.TracingRepository,
) TicketPurchaseUseCase {
	return TicketPurchaseUseCase{
		ticketRepository:      ticketRepository,
		eventRepository:       eventRepository,
		paymentGateway:        paymentGateway,
		transactionRepository: transactionRepository,
		unitOfWorkRepository:  unitOfWorkRepository,
		trace:                 trace,
	}
}

func (usecase TicketPurchaseUseCase) Execute(
	ctx context.Context,
	input dto.TicketPurchaseInput,
) (*dto.TicketPurchaseOutput, *dto.ErrorOutput) {
	span, ctx := usecase.trace.StartSpan(ctx, "TicketPurchaseUseCase.Execute", string(constants.ApplicationLayer))
	defer span.End()
	ctx, err := usecase.unitOfWorkRepository.Begin(ctx)
	if err != nil {
		_ = usecase.unitOfWorkRepository.Rollback(ctx)
		return nil, dto.NewErrorOutput(err.Error(), constants.InfraLayer, 500, constants.ErrInternal)
	}

	event, err := usecase.eventRepository.Get(ctx, input.EventId)
	if err != nil {
		_ = usecase.unitOfWorkRepository.Rollback(ctx)
		return nil, dto.NewErrorOutput(err.Error(), constants.InfraLayer, 500, constants.ErrInternal)
	}

	ticket := entity.NewTicket(input.EventId, input.Email)
	err = usecase.ticketRepository.Save(ctx, ticket)
	if err != nil {
		_ = usecase.unitOfWorkRepository.Rollback(ctx)
		return nil, dto.NewErrorOutput(err.Error(), constants.InfraLayer, 500, constants.ErrInternal)
	}

	transactionOutput, err := usecase.paymentGateway.CreateTransaction(ctx, dto.PaymentTransactionInput{
		Email: input.Email, CreditCardToken: input.CreditCardToken, Price: event.Price,
	})
	if err != nil {
		_ = usecase.unitOfWorkRepository.Rollback(ctx)
		return nil, dto.NewErrorOutput(err.Error(), constants.InfraLayer, 500, constants.ErrInternal)
	}
	transaction := entity.NewTransaction(
		ticket.TicketID,
		event.EventID,
		transactionOutput.TID,
		event.Price,
		transactionOutput.Status,
	)
	err = usecase.transactionRepository.Save(ctx, transaction)
	if err != nil {
		_ = usecase.unitOfWorkRepository.Rollback(ctx)
		return nil, dto.NewErrorOutput(err.Error(), constants.InfraLayer, 500, constants.ErrInternal)
	}

	if transactionOutput.Status == "approved" {
		ticket.Approve()
	} else {
		ticket.Cancel()
	}

	err = usecase.ticketRepository.Update(ctx, ticket)
	if err != nil {
		_ = usecase.unitOfWorkRepository.Rollback(ctx)
		return nil, dto.NewErrorOutput(err.Error(), constants.InfraLayer, 500, constants.ErrInternal)
	}

	err = usecase.unitOfWorkRepository.Commit(ctx)
	if err != nil {
		usecase.trace.CaptureError(ctx, err)
		_ = usecase.unitOfWorkRepository.Rollback(ctx)
		return nil, dto.NewErrorOutput(err.Error(), constants.InfraLayer, 500, constants.ErrInternal)
	}

	return &dto.TicketPurchaseOutput{
		TicketId: ticket.TicketID,
		Status:   string(ticket.Status),
		TID:      transactionOutput.TID,
		Price:    event.Price,
	}, nil
}
