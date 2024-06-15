package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller"
)

func HandleUser(router *gin.Engine, userController *controller.UserController) {
	userRoute := router.Group("users")

	userRoute.GET("", userController.GetAll)
	userRoute.GET("full-profile/:id", userController.GetFullProfile)
	userRoute.GET(":id", userController.GetOne)
	userRoute.POST("", userController.Create)
	userRoute.PUT(":id", userController.Update)
	userRoute.DELETE(":id", userController.Destroy)
}
