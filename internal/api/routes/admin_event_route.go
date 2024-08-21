package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/handler/admin_event_handler"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/middleware"
)

func HandleEvent(router *gin.RouterGroup) {
	adminEventRoute := router.Group("admin/events")

	adminEventRoute.Use(middleware.JWT.MiddlewareFunc())
	adminEventRoute.Use(middleware.Authorization(roles_enum.ADMIN))

	adminEventRoute.POST("", admin_event_handler.Create)
	adminEventRoute.GET("", admin_event_handler.GetAll)
	adminEventRoute.GET(":id", admin_event_handler.GetOneById)
	adminEventRoute.PUT(":id", admin_event_handler.Update)
	adminEventRoute.DELETE(":id", admin_event_handler.SoftDelete)
}
