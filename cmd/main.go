package main

import (
	"context"
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

	conn, err := db.OpenConnection(stringConnect)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(context.Background())

	if err := app.Run(conn); err != nil {
		log.Fatal(err)
	}
}
