package db

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
)

const DRIVER = "postgres"

var (
	Query    *sqlc.Queries
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
	Query = sqlc.New(Conn)
}

func getStringConnection() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		config.Envs.DBHost,
		config.Envs.DBPort,
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBName)
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
        SELECT table_name
        FROM information_schema.tables
        WHERE table_schema = 'public' AND table_type = 'BASE TABLE'
    `)
	if err != nil {
		log.Fatalf("Failed to fetch table names: %v\n", err)
	}

	// Truncate each table
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("Failed to scan table name: %v\n", err)
		}
		_, err := Conn.Exec(ctx, fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", tableName))
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
