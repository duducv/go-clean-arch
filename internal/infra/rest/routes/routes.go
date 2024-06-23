package routes

import (
	"github.com/duducv/go-clean-arch/config"
	"github.com/go-chi/chi"
)

func ApplyRoutes(router *chi.Mux, adapters *config.RepositoryAdapters) {
	NewTicketController(router, adapters)
}

func ConfigRouter() *chi.Mux {
	return chi.NewRouter()
}
