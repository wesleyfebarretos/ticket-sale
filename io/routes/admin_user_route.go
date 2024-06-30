package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller/admin_user_controller"
	"github.com/wesleyfebarretos/ticket-sale/middleware"
)

func HandleAdminUser(router *gin.RouterGroup) {
	userRoute := router.Group("admin/users")

	userRoute.Use(middleware.JWT.MiddlewareFunc())
	userRoute.Use(middleware.Authorization(enum.ADMIN_ROLE))

	userRoute.POST("", admin_user_controller.Create)
	userRoute.POST("/get-by-email", admin_user_controller.GetOneByEmail)
	userRoute.GET("", admin_user_controller.GetAll)
	userRoute.GET(":id", admin_user_controller.GetOneById)
	userRoute.PUT(":id", admin_user_controller.Update)
	userRoute.DELETE(":id", admin_user_controller.Delete)
}
