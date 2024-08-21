package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/handler/auth_handler"
)

func HandleAuth(router *gin.RouterGroup) {
	authRoute := router.Group("auth")

	authRoute.POST("", auth_handler.Auth)
}
