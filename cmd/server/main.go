package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/merteldem1r/TaskeFlow-API/internal/app"
	"github.com/merteldem1r/TaskeFlow-API/internal/config"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()
	application := app.NewApp(cfg)

	err := application.Run()

	if err != nil {
		log.Fatal("Server failed to start", err)
	}
}
