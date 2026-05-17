package handlers

import (
	"github.com/go-chi/chi"
	chmiddle "github.com/go-chi/chi/middleware"
	"myapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chmiddle.StripSlashes)
	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}
