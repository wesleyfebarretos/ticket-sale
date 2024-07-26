package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/controller/creditcard_controller"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/middleware"
)

func HandleCreditcard(router *gin.RouterGroup) {
	creditcardRoute := router.Group("admin/events")

	creditcardRoute.Use(middleware.JWT.MiddlewareFunc())
	creditcardRoute.Use(middleware.Authorization(roles_enum.USER))

	creditcardRoute.POST("", creditcard_controller.Create)
	creditcardRoute.GET("", creditcard_controller.GetAllUserCreditcards)
	creditcardRoute.PUT(":uuid", creditcard_controller.Update)
	creditcardRoute.DELETE(":uuid", creditcard_controller.SoftDelete)
}
