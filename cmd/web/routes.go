package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/serhatguzel/bookings/internal/config"
	"github.com/serhatguzel/bookings/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer) // Middleware to recover from panics and log errors
	mux.Use(NoSurf)
	mux.Use(SessionLoad) // Middleware to load and save session data

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)

	fileServer := http.FileServer(http.Dir("./static/"))             // File server for static files
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer)) // Serve static files from the static directory

	return mux
}
