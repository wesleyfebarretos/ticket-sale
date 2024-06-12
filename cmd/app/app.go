package app

import (
	"database/sql"
	"fmt"

	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/routes"
)

func Run(db *sql.DB) error {
	router := routes.Bind()

	return router.Run(fmt.Sprintf(":%s", config.Envs.Port))
}
