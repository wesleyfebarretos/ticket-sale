package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/wesleyfebarretos/ticket-sale/cmd/app"
	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/infra/db"
	_ "github.com/wesleyfebarretos/ticket-sale/swagger"
)

// @title						Ticket Sale
// @version					1.0
// @description				This is a simple ticket selling application.
// @termsOfService				http://swagger.io/terms/
// @contact.name				Ticket Sale Support
// @contact.url				http://www.swagger.io/support
// @contact.email				support@swagger.io
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @host						localhost:8080
// @BasePath					/v1
// @securityDefinitions.apikey	ApiKeyAuth
// @in							cookie
// @name						jwt_ticket_sale
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
	defer db.Conn.Close()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
