package main

import (
	"net/http"

	"github.com/joeyjurjens/peekaboo/pkg/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Kinda a hack, but it works, lol, stupid they don't have a normal way to do this.
	mux.HandleFunc("GET /", handlers.NotFoundHandler)

	mux.HandleFunc("GET /{$}", handlers.HomeHandler)
	mux.HandleFunc("GET /dashboard/{$}", handlers.DashboardIndexHandler)
	mux.HandleFunc("GET /projects/{$}", handlers.ProjectIndexHandler)
	mux.HandleFunc("GET /projects/{id}/{$}", handlers.ProjectShowHandler)
	mux.HandleFunc("GET /projects/create/{$}", handlers.ProjectCreateFormHandler)
	mux.HandleFunc("POST /projects/create/{$}", handlers.ProjectCreateHandler)
	mux.HandleFunc("GET /projects/{id}/edit/{$}", handlers.ProjectEditFormHandler)
	mux.HandleFunc("POST /projects/{id}/edit/{$}", handlers.ProjectEditHandler)
	mux.HandleFunc("GET /projects/{id}/delete/{$}", handlers.ProjectDeleteFormHandler)
	mux.HandleFunc("POST /projects/{id}/delete/{$}", handlers.ProjectDeleteHandler)

	http.ListenAndServe(":8000", mux)
}
