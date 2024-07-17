package db

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/wesleyfebarretos/ticket-sale/config"
)

const DRIVER = "postgres"

var (
	Conn     *pgxpool.Pool
	initOnce sync.Once
)

func openConnection(connector string) {
	config, err := pgxpool.ParseConfig(connector)
	if err != nil {
		log.Fatalf("error on parse db config: %v", err)
	}

	config.MaxConns = 10

	insideConn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	Conn = insideConn
}

func getStringConnection() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		config.Envs.DB.Host,
		config.Envs.DB.Port,
		config.Envs.DB.User,
		config.Envs.DB.Password,
		config.Envs.DB.Name)
}

func Init() {
	initOnce.Do(func() {
		conn := getStringConnection()
		openConnection(conn)
	})
}

func TruncateAll() {
	ctx := context.Background()
	rows, err := Conn.Query(ctx, `
        SELECT table_name, table_schema
        FROM information_schema.tables
        WHERE table_schema IN('public', 'fin') AND table_type = 'BASE TABLE'
    `)
	if err != nil {
		log.Fatalf("Failed to fetch table names: %v\n", err)
	}

	// Truncate each table
	for rows.Next() {
		var tableName string
		var tableSchema string
		if err := rows.Scan(&tableName, &tableSchema); err != nil {
			log.Fatalf("Failed to scan table name: %v\n", err)
		}

		var table string

		if tableSchema != "public" {
			table = fmt.Sprintf("%s.%s", tableSchema, tableName)
		} else {
			table = tableName
		}

		_, err := Conn.Exec(ctx, fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table))
		if err != nil {
			log.Fatalf("Failed to truncate table %s: %v\n", tableName, err)
		}
		// fmt.Printf("Truncated table: %s\n", tableName)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating rows: %v\n", err)
	}

	// fmt.Println("All tables truncated successfully.")
}
