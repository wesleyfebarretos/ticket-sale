package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller/auth_controller"
)

func HandleAuth(router *gin.RouterGroup) {
	authRoute := router.Group("auth")

	authRoute.POST("", auth_controller.Auth)
}
