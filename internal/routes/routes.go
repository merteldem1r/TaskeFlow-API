package routes

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/merteldem1r/TaskeFlow-API/internal/handlers"
)

func Setup(r *chi.Mux) {
	// Built-in Chi middlewares
	r.Use(middleware.Logger)    // Logs every request
	r.Use(middleware.Recoverer) // Recovers from panics

	// handlers
	healthHandler := handlers.NewHealthHandler()
	taskHandler := handlers.NewTaskHandler()

	r.Get("/health", healthHandler.Check)

	// API v1 routes group
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/tasks", func(r chi.Router) {
			r.Get("/", taskHandler.GetAll)
			r.Post("/", taskHandler.Create)
			r.Get("/{id}", taskHandler.GetByID)
			r.Put("/{id}", taskHandler.Update)
			r.Delete("/{id}", taskHandler.Delete)
		})
	})
}
