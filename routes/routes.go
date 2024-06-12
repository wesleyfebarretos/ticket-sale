package routes

import (
	"github.com/gin-gonic/gin"
)

func Bind() *gin.Engine {
	router := gin.New()
	HandleHealthCheck(router)
	HandleUser(router)
	return router
}
