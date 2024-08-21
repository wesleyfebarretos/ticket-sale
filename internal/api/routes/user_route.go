package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/handler/user_handler"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/middleware"
)

func HandleUser(router *gin.RouterGroup) {
	userRoute := router.Group("users")

	userRoute.POST("", user_handler.Create)

	userRoute.Use(middleware.JWT.MiddlewareFunc())
	userRoute.Use(middleware.Authorization(roles_enum.USER))

	userRoute.GET("", user_handler.GetAll)
	userRoute.GET("full-profile", user_handler.GetFullProfile)
	userRoute.GET(":id", user_handler.GetOneById)
	userRoute.PUT("", user_handler.Update)
}
