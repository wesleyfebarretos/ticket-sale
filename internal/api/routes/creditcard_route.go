package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/handler/creditcard_handler"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/middleware"
)

func HandleCreditcard(router *gin.RouterGroup) {
	creditcardRoute := router.Group("creditcard")

	creditcardRoute.Use(middleware.JWT.MiddlewareFunc())
	creditcardRoute.Use(middleware.Authorization(roles_enum.USER))

	creditcardRoute.POST("", creditcard_handler.Create)
	creditcardRoute.GET("user", creditcard_handler.GetAllUserCreditcards)
	creditcardRoute.PUT(":uuid", creditcard_handler.Update)
	creditcardRoute.DELETE(":uuid", creditcard_handler.SoftDelete)
}
