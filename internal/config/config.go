package config

import (
	"os"
	"strconv"
)

type Config struct {
	AppName   string
	Port      int
	DBHost    string
	DBPort    int
	DBName    string
	JWTSecret string
}

func Load() *Config {
	return &Config{
		AppName:   os.Getenv("APP_NAME"),
		Port:      mustAtoi(os.Getenv("PORT")),
		DBHost:    os.Getenv("DB_HOST"),
		DBPort:    mustAtoi(os.Getenv("DB_PORT")),
		DBName:    os.Getenv("DB_NAME"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("Invalid integer value in environment variable: " + s)
	}
	return i
}
