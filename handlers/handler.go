package handlers

import (
	"go-learn/middleware"
	"net/http"

	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

func Handlers(r *chi.Mux) {
	r.Use(chimiddleware.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/", GetCoinBalance)
		router.Post("/transfer", PostTransferCoin)
	})
}

func PostTransferCoin(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement transfer coin handler
}
