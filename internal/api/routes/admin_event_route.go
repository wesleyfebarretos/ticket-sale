package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/controller/admin_event_controller"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/middleware"
)

func HandleEvent(router *gin.RouterGroup) {
	adminEventRoute := router.Group("admin/events")

	adminEventRoute.Use(middleware.JWT.MiddlewareFunc())
	adminEventRoute.Use(middleware.Authorization(roles_enum.ADMIN))

	adminEventRoute.POST("", admin_event_controller.Create)
	adminEventRoute.GET("", admin_event_controller.GetAll)
	adminEventRoute.GET(":id", admin_event_controller.GetOneById)
	adminEventRoute.PUT(":id", admin_event_controller.Update)
	adminEventRoute.DELETE(":id", admin_event_controller.SoftDelete)
}
