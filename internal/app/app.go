package app

import (
	"fmt"
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

func (a *App) setupRoutes() {
	a.Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
}

func (a *App) Run() error {
	addr := fmt.Sprintf(":%d", a.Config.Port)
	fmt.Printf("Server is starting on http://localhost%s\n", addr)

	return http.ListenAndServe(addr, a.Router)
}
