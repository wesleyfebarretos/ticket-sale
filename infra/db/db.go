package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/wesleyfebarretos/ticket-sale/config"
)

const DRIVER = "postgres"

func OpenConnection(connector string) (*sql.DB, error) {
	db, err := sql.Open(DRIVER, connector)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func GetStringConnection() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		config.Envs.DBHost,
		config.Envs.DBPort,
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBName)
}

func Init(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected")
}
