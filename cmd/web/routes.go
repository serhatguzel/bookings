package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/serhatguzel/bookings/pkg/config"
	"github.com/serhatguzel/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer) // Middleware to recover from panics and log errors
	mux.Use(NoSurf)
	mux.Use(SessionLoad) // Middleware to load and save session data

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
