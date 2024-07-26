package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/controller/admin_user_controller"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/middleware"
)

func HandleAdminUser(router *gin.RouterGroup) {
	adminUserRoute := router.Group("admin/users")

	adminUserRoute.Use(middleware.JWT.MiddlewareFunc())
	adminUserRoute.Use(middleware.Authorization(roles_enum.ADMIN))

	adminUserRoute.POST("", admin_user_controller.Create)
	adminUserRoute.POST("/get-by-email", admin_user_controller.GetOneByEmail)
	adminUserRoute.GET("", admin_user_controller.GetAll)
	adminUserRoute.GET(":id", admin_user_controller.GetOneById)
	adminUserRoute.PUT(":id", admin_user_controller.Update)
	adminUserRoute.DELETE(":id", admin_user_controller.Delete)
}
