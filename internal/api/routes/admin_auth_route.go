package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/handler/admin_auth_handler"
)

func HandleAdminAuth(router *gin.RouterGroup) {
	authRoute := router.Group("admin/auth")

	authRoute.POST("", admin_auth_handler.Auth)
}
