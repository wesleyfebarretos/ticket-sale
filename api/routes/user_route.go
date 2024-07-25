package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/api/internal/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/api/io/http/controller/user_controller"
	"github.com/wesleyfebarretos/ticket-sale/api/middleware"
)

func HandleUser(router *gin.RouterGroup) {
	userRoute := router.Group("users")

	userRoute.POST("", user_controller.Create)

	userRoute.Use(middleware.JWT.MiddlewareFunc())
	userRoute.Use(middleware.Authorization(roles_enum.USER))

	userRoute.GET("", user_controller.GetAll)
	userRoute.GET("full-profile", user_controller.GetFullProfile)
	userRoute.GET(":id", user_controller.GetOneById)
	userRoute.PUT("", user_controller.Update)
}
