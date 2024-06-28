package app

import (
	"fmt"

	"github.com/wesleyfebarretos/ticket-sale/config"
	"github.com/wesleyfebarretos/ticket-sale/io/routes"
	"github.com/wesleyfebarretos/ticket-sale/repository"
)

func Run() error {
	router := routes.Bind()
	repository.Bind()

	return router.Run(fmt.Sprintf(":%s", config.Envs.Port))
}
