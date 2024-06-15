package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/config"
)

const DRIVER = "postgres"

func OpenConnection(connector string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), connector)
	if err != nil {
		log.Fatal(err)
	}

	return conn, nil
}

func GetStringConnection() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		config.Envs.DBHost,
		config.Envs.DBPort,
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBName)
}

// func Init(db *sql.DB) {
// 	if err := db.Ping(); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	log.Println("DB: Successfully connected")
// }
