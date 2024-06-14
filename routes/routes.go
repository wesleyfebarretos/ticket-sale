package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller"
)

func Bind(db *sql.DB) *gin.Engine {
	router := gin.New()
	HandleHealthCheck(router)

	userController := controller.NewUser(db)
	HandleUser(router, userController)
	return router
}
