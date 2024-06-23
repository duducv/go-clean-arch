package config

import (
	"github.com/duducv/go-clean-arch/internal/core/application/gateway"
	applicationrepository "github.com/duducv/go-clean-arch/internal/core/application/repository"
	"github.com/duducv/go-clean-arch/internal/core/domain/entity"
)

type RepositoryAdapters struct {
	TicketAdapter      entity.TicketRepository
	EventAdapter       entity.EventRepository
	TransactionAdapter entity.TransactionRepository
	PaymentGateway     gateway.PaymentGateway
	UnitOfWorkAdapter  applicationrepository.UnitOfWorkRepository
	TraceAdapter       applicationrepository.TracingRepository
}

func NewRepositoryAdapters(
	ticketRepository entity.TicketRepository,
	eventAdapter entity.EventRepository,
	transactionAdapter entity.TransactionRepository,
	paymentGateway gateway.PaymentGateway,
	unitOfWorkAdapter applicationrepository.UnitOfWorkRepository,
	traceAdapter applicationrepository.TracingRepository,
) *RepositoryAdapters {
	return &RepositoryAdapters{
		TicketAdapter:      ticketRepository,
		EventAdapter:       eventAdapter,
		PaymentGateway:     paymentGateway,
		TransactionAdapter: transactionAdapter,
		UnitOfWorkAdapter:  unitOfWorkAdapter,
		TraceAdapter:       traceAdapter,
	}
}
