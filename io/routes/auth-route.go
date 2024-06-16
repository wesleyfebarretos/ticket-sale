package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller"
)

func HandleAuth(router *gin.Engine, authController *controller.AuthController) {
	authRoute := router.Group("auth")

	authRoute.POST("", authController.SignIn)
}
