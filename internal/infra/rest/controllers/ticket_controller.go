package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/duducv/go-clean-arch/internal/core/application/constants"
	"github.com/duducv/go-clean-arch/internal/core/application/dto"
	applicationrepository "github.com/duducv/go-clean-arch/internal/core/application/repository"
	"github.com/duducv/go-clean-arch/internal/core/application/usecases"
)

type TicketController struct {
	ticketPurchaseUseCase usecases.TicketPurchaseUseCase
	trace                 applicationrepository.TracingRepository
}

func NewTicketController(
	ticketPurchaseUseCase usecases.TicketPurchaseUseCase,
	trace applicationrepository.TracingRepository,
) *TicketController {
	return &TicketController{
		ticketPurchaseUseCase: ticketPurchaseUseCase,
		trace:                 trace,
	}
}

func (controller *TicketController) PurchaseTicket() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		span, ctx := controller.trace.StartSpan(
			r.Context(),
			"TicketController.PurchaseTicket",
			string(constants.InfraLayer),
		)
		defer span.End()
		w.Header().Set("Content-Type", "application/json")
		input := dto.TicketPurchaseInput{}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			controller.trace.CaptureError(ctx, err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("wrong body"))
			return
		}
		output, errOutput := controller.ticketPurchaseUseCase.Execute(ctx, input)
		if errOutput != nil {
			if errOutputAsBytes, err := json.Marshal(errOutput); err != nil {
				controller.trace.CaptureError(ctx, err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			} else {
				w.WriteHeader(errOutput.StatusCode)
				w.Write(errOutputAsBytes)
				return
			}
		}
		if outputAsBytes, err := json.Marshal(output); err != nil {
			controller.trace.CaptureError(ctx, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			w.WriteHeader(http.StatusCreated)
			w.Write(outputAsBytes)
			return
		}
	}
}
