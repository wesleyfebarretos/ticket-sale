package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/middleware"
)

func Bind() *gin.Engine {
	router := gin.New()
	router.Use(gin.CustomRecovery(middleware.ExceptionMiddleware))
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/metrics"},
	}))

	HandleHealthCheck(router)
	HandleAuth(router)
	HandleUser(router)

	return router
}
