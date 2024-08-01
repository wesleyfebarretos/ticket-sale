package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/controller/admin_gateway_controller"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/middleware"
)

func HandleAdminGateway(router *gin.RouterGroup) {
	adminGateway := router.Group("admin/gateway")

	adminGateway.Use(middleware.JWT.MiddlewareFunc())
	adminGateway.Use(middleware.Authorization(roles_enum.ADMIN))

	adminGateway.POST("", admin_gateway_controller.Create)
	adminGateway.GET("", admin_gateway_controller.GetAll)
	adminGateway.GET(":id", admin_gateway_controller.GetOneById)
	adminGateway.PUT(":id", admin_gateway_controller.Update)
	adminGateway.DELETE(":id", admin_gateway_controller.SoftDelete)
}
