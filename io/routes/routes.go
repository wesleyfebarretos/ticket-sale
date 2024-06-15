package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller"
)

func Bind(conn *pgx.Conn) *gin.Engine {
	router := gin.New()
	HandleHealthCheck(router)

	userController := controller.NewUserController(conn)
	HandleUser(router, userController)
	return router
}
