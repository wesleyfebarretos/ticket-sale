package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/controller/admin_auth_controller"
)

func HandleAdminAuth(router *gin.RouterGroup) {
	authRoute := router.Group("admin/auth")

	authRoute.POST("", admin_auth_controller.Auth)
}
