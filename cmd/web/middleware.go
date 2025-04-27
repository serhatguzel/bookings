package main

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request method and URL
		log.Printf("Request Method: %s, Request URL: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // Call the next handler in the chain
	})
}

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction, // true ise sadece HTTPS üzerinden gönderir
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler

}

// SessionLoad loads and saves the session on each request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
