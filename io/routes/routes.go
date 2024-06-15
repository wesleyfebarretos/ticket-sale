package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller"
	"github.com/wesleyfebarretos/ticket-sale/middleware"
)

func Bind(conn *pgx.Conn) *gin.Engine {
	router := gin.New()
	router.Use(gin.CustomRecovery(middleware.ExceptionMiddleware))
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/metrics"},
	}))

	HandleHealthCheck(router)

	userController := controller.NewUserController(conn)
	HandleUser(router, userController)
	return router
}
