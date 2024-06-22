package routes

import (
	"github.com/gin-gonic/gin"
	user_controller "github.com/wesleyfebarretos/ticket-sale/io/http/controller/user"
	"github.com/wesleyfebarretos/ticket-sale/middleware"
)

func HandleUser(router *gin.Engine) {
	userRoute := router.Group("users")

	userRoute.POST("", user_controller.Create)

	userRoute.Use(middleware.JWT.MiddlewareFunc())

	userRoute.GET("", user_controller.GetAll)
	userRoute.GET("full-profile", user_controller.GetFullProfile)
	userRoute.GET(":id", user_controller.GetById)
	userRoute.PUT("", user_controller.Update)
}
