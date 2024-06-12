package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/wesleyfebarretos/ticket-sale/cmd/app"
	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/infra/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connector := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Envs.DBHost,
		config.Envs.DBPort,
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBName)

	_db, err := db.OpenConnection(connector)
	if err != nil {
		log.Fatal(err)
	}

	defer _db.Close()

	db.Init(_db)

	if err := app.Run(_db); err != nil {
		log.Fatal(err)
	}
}
