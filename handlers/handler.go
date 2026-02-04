package handlers

import (
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"go-learn/middleware"

)

func Handlers(r *chi.Mux){
	r.Use(chimiddleware.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/", GetCoinBalance)
		router.Post("/transfer", PostTransferCoin)
	})
}

