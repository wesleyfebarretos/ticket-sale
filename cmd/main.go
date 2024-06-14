package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/wesleyfebarretos/ticket-sale/cmd/app"
	"github.com/wesleyfebarretos/ticket-sale/infra/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	stringConnect := db.GetStringConnection()

	_db, err := db.OpenConnection(stringConnect)
	if err != nil {
		log.Fatal(err)
	}

	defer _db.Close()

	db.Init(_db)

	if err := app.Run(_db); err != nil {
		log.Fatal(err)
	}
}
