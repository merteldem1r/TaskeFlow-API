package main

import (
	"fmt"

	"github.com/merteldem1r/TaskeFlow-API/internal/config"
)

func main() {
	cfg := config.Load()

	fmt.Println(cfg.AppName, "is starting...")
	fmt.Println("Server will run on port: ", cfg.Port)
	fmt.Println("Connecting to database at", cfg.DBHost, "on port", cfg.DBPort)
}
