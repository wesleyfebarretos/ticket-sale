package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/middleware"
)

func HandleAuth(router *gin.Engine) {
	authRoute := router.Group("auth")

	authRoute.POST("", middleware.Jwt.LoginHandler)
}
