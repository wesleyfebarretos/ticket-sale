package app

import (
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/io/routes"
)

func Run(conn *pgx.Conn) error {
	router := routes.Bind(conn)

	return router.Run(fmt.Sprintf(":%s", config.Envs.Port))
}
