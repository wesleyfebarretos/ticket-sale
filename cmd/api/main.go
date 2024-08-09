package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/config"
	_ "github.com/wesleyfebarretos/ticket-sale/internal/api/docs"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/routes"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

// Swagger entrypoint godoc
//
//	@title			Ticket Sale
//
//	@version		1.0
//	@description	This is a simple ticket sales application.
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	Ticket Sale Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:8080
//	@BasePath		/v1
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config.Init()

	db.Init()
	defer db.Conn.Close()

	router := routes.Bind()

	if err := router.Run(fmt.Sprintf(":%s", config.Envs.Port)); err != nil {
		log.Fatalf("Error on starting API: %v", err)
	}
}
