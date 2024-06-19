package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/middleware"
)

func Bind() *gin.Engine {
	router := gin.New()
	router.Use(gin.CustomRecovery(middleware.ExceptionHandler))
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/metrics"},
	}))

	// Init JWT instance
	middleware.InitJWT()

	HandleHealthCheck(router)
	HandleAuth(router)
	HandleUser(router)

	return router
}
