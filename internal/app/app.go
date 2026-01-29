package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/go-chi/chi/v5"
	"github.com/merteldem1r/TaskeFlow-API/internal/config"
	"github.com/merteldem1r/TaskeFlow-API/internal/database"
	"github.com/merteldem1r/TaskeFlow-API/internal/handlers"
	"github.com/merteldem1r/TaskeFlow-API/internal/repositories"
	"github.com/merteldem1r/TaskeFlow-API/internal/routes"
	"github.com/merteldem1r/TaskeFlow-API/internal/services"
)

type App struct {
	Config *config.Config
	Router *chi.Mux
	DB     driver.Conn
}

func NewApp(cfg *config.Config) *App {
	db, err := database.Connect(cfg)

	if err != nil {
		log.Fatal("Failed to connect ClickHouse database: ", err)
	}

	// Repositories and Services

	// Task
	taskRepo := repositories.NewTaskRepository(db, cfg.DBName)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	// User
	userRepo := repositories.NewUserRepository(db, cfg.DBName)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	app := &App{
		Config: cfg,
		Router: chi.NewRouter(),
		DB:     db,
	}

	routes.Setup(app.Router, taskHandler, userHandler)

	return app
}

func (a *App) Run() error {
	addr := fmt.Sprintf(":%d", a.Config.Port)
	fmt.Printf("Server is starting on http://localhost%s\n", addr)

	return http.ListenAndServe(addr, a.Router)
}
