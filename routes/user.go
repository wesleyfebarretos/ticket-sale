package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUser(router *gin.Engine) {
	user := router.Group("users")

	user.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "All Users",
		})
	})

	user.GET(":id", func(c *gin.Context) {
		var user string
		if value, ok := c.Params.Get("id"); ok {
			user = value
		} else {
			user = "not found"
		}
		c.JSON(http.StatusOK, gin.H{
			"status": fmt.Sprintf("Show user %s", user),
		})
	})

	user.POST("", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"status": "create user",
		})
	})

	user.PUT(":id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "update user",
		})
	})

	user.DELETE(":id", func(c *gin.Context) {
		c.JSON(http.StatusNoContent, gin.H{
			"status": "DELETE user",
		})
	})
}
