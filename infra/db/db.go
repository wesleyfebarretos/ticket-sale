package db

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
)

const DRIVER = "postgres"

var (
	Query    *sqlc.Queries
	Conn     *pgx.Conn
	initOnce sync.Once
)

func OpenConnection(connector string) {
	initOnce.Do(func() {
		insideConn, err := pgx.Connect(context.Background(), connector)
		if err != nil {
			log.Fatal(err)
		}
		Conn = insideConn
		Query = sqlc.New(Conn)
	})
}

func GetStringConnection() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		config.Envs.DBHost,
		config.Envs.DBPort,
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBName)
}
