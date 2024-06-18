package app

import (
	"fmt"

	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/io/routes"
)

func Run() error {
	router := routes.Bind()

	return router.Run(fmt.Sprintf(":%s", config.Envs.Port))
}
