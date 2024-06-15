package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/wesleyfebarretos/ticket-sale/infra/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	action := os.Args[1]
	var migrationType string

	if action == "down" {
		migrationType = "tables"
	} else {
		migrationType = os.Args[2]
	}

	stringConnect := db.GetStringConnection()

	db, err := db.OpenConnection(stringConnect)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var m *migrate.Migrate

	switch migrationType {
	case "tables":
		m, err = migrate.NewWithDatabaseInstance("file://cmd/migrations/tables", "postgres", driver)
	case "seeders":
		m, err = migrate.NewWithDatabaseInstance("file://cmd/migrations/seeders", "postgres", driver)
	default:
		log.Fatal("Invalid migration type. Use 'tables' or 'seeders'")
	}

	if err != nil {
		log.Fatal(err)
	}

	switch action {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	default:
		log.Fatal("Invalid action. Use 'Up' or 'Down'")
	}
}
