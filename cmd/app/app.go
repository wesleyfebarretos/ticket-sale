package app

import (
	"fmt"

	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/routes"
)

func Run() {
	router := routes.Bind()

	router.Run(fmt.Sprintf(":%s", config.Envs.Port))
}
