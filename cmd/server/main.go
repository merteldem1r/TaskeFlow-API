package main

import (
	"context"
	"crypto/tls"
	"database/sql/driver"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func main() {
	conn, err := connect()

	if err != nil {
		log.Fatalf("failed to connect to ClickHouse: %v", err)
	}


}

func connect() (driver.Conn, error) {
	var (
		ctx = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{"<CLICKHOUSE_SECURE_NATIVE_HOSTNAME>:9440"},
			Auth: clickhouse.Auth{
				
			}
		})
		
	) 
}
