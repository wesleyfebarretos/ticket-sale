package main

import (
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/wesleyfebarretos/ticket-sale/cmd/migrations/migration"
	"github.com/wesleyfebarretos/ticket-sale/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if len(os.Args) < 2 {
		log.Fatal("Action argument is required (up or down)")
	}

	config.Init()

	action := os.Args[1]
	if action == "up" {
		migration.Up()
	} else {
		migration.Down()
	}
}
