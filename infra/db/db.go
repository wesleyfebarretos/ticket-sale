package db

import (
	"database/sql"
	"log"
)

const DRIVER = "postgres"

func OpenConnection(connector string) (*sql.DB, error) {
	db, err := sql.Open(DRIVER, connector)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func Init(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected")
}
