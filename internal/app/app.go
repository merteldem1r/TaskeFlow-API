package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/merteldem1r/TaskeFlow-API/internal/config"
)

type App struct {
	Config *config.Config
	Router *chi.Mux
}

func NewApp(cfg *config.Config) *App {
	app := &App{
		Config: cfg,
		Router: chi.NewRouter(),
	}

	app.setupRoutes()

	return app
}

func (app *App) setupRoutes() {
	app.Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
}
