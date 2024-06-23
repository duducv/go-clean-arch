package routes

import (
	"github.com/duducv/go-clean-arch/config"
	"github.com/duducv/go-clean-arch/internal/core/application/usecases"
	"github.com/duducv/go-clean-arch/internal/infra/rest/controllers"
	"github.com/go-chi/chi"
)

func NewTicketController(router chi.Router, adapters *config.RepositoryAdapters) {
	ticketPurchaseUseCase := usecases.NewTicketPurchaseUseCase(
		adapters.TicketAdapter,
		adapters.EventAdapter,
		adapters.TransactionAdapter,
		adapters.PaymentGateway,
		adapters.UnitOfWorkAdapter,
		adapters.TraceAdapter,
	)
	purchaseController := controllers.NewTicketController(
		ticketPurchaseUseCase,
		adapters.TraceAdapter,
	)
	router.Post("/purchase_ticket", purchaseController.PurchaseTicket())
}
