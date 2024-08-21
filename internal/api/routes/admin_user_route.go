package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/handler/admin_user_handler"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/middleware"
)

func HandleAdminUser(router *gin.RouterGroup) {
	adminUserRoute := router.Group("admin/users")

	adminUserRoute.Use(middleware.JWT.MiddlewareFunc())
	adminUserRoute.Use(middleware.Authorization(roles_enum.ADMIN))

	adminUserRoute.POST("", admin_user_handler.Create)
	adminUserRoute.POST("/get-by-email", admin_user_handler.GetOneByEmail)
	adminUserRoute.GET("", admin_user_handler.GetAll)
	adminUserRoute.GET(":id", admin_user_handler.GetOneById)
	adminUserRoute.PUT(":id", admin_user_handler.Update)
	adminUserRoute.DELETE(":id", admin_user_handler.Delete)
}
