package main

import (
	"context"
	"flag"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/wesleyfebarretos/ticket-sale/cmd/app"
	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/infra/db"
)

func main() {
	testMode := flag.Bool("test", false, "")

	flag.Parse()

	var err error
	if *testMode {
		err = godotenv.Load(".env.test")
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.Init()

	db.Init()
	defer db.Conn.Close(context.Background())

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
