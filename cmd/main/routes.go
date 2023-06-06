package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/handlers"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)

	return mux
}
