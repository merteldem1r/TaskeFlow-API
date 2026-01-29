package database

import (
	"context"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/merteldem1r/TaskeFlow-API/internal/config"
)

func Connect(cfg *config.Config) (driver.Conn, error) {
	ctx := context.Background()

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", cfg.DBHost, cfg.DBPort)},
		Auth: clickhouse.Auth{
			Username: "default",
			Password: "",
		},
		Debug: true,
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout:     5 * time.Second,
		ConnMaxLifetime: 1 * time.Hour,
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to open connections %w", err)
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Failed to ping database: %w", err)
	}

	fmt.Println("Successfully connected to ClickHouse")
	return conn, nil
}
