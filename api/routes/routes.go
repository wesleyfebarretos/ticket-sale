package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/api/config"
	"github.com/wesleyfebarretos/ticket-sale/api/middleware"
)

func Bind() *gin.Engine {
	router := gin.New()
	if config.Envs.AppEnv != "testing" {
		router.Use(middleware.Logger)
	}
	router.Use(gin.CustomRecovery(middleware.ExceptionHandler))

	// Init JWT instance
	middleware.InitJWT()

	v1 := router.Group("/v1")

	HandleHealthCheck(v1)
	HandleSwagger(v1)
	HandleAuth(v1)
	HandleAdminAuth(v1)
	HandleUser(v1)
	HandleAdminUser(v1)
	HandleAdminProduct(v1)
	HandleEvent(v1)

	return router
}
