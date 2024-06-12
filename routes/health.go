package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHealthCheck(router *gin.Engine) {
	router.GET("health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
}
