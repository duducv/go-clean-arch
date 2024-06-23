package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/duducv/go-clean-arch/config"
	gatewayadapters "github.com/duducv/go-clean-arch/internal/infra/gateway"
	persistenceadapters "github.com/duducv/go-clean-arch/internal/infra/persistence"
	"github.com/duducv/go-clean-arch/internal/infra/rest/routes"
	"github.com/duducv/go-clean-arch/internal/infra/tracing"
	"github.com/go-chi/chi"
)

func main() {
	dbUrl, port := config.LoadEnv("../../.env")
	pgConn := config.OpenPGXConnection(dbUrl)
	router := routes.ConfigRouter()
	tracing := tracing.NewNotracingAdapter()
	adapters := config.NewRepositoryAdapters(
		persistenceadapters.NewTicketRepositorySQLAdapter(pgConn, tracing),
		persistenceadapters.NewEventRepositorySQLAdapter(pgConn, tracing),
		persistenceadapters.NewTransactionRepositorySQLAdapter(pgConn, tracing),
		gatewayadapters.NewFakePaymentGateway(),
		persistenceadapters.NewUnitOfWorkSQLAdapter(*pgConn),
		tracing,
	)
	routes.ApplyRoutes(router, adapters)
	listen(port, router)
}

func listen(port string, router *chi.Mux) {
	fmt.Println("\033[32m✔\033[0m ", "Iniciando servidor na porta", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
