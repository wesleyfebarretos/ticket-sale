package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/handler/checkout_handler"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/middleware"
)

func HandleCheckout(router *gin.RouterGroup) {
	checkoutRoute := router.Group("checkout")

	checkoutRoute.Use(middleware.JWT.MiddlewareFunc())
	checkoutRoute.Use(middleware.Authorization(roles_enum.USER))

	checkoutRoute.POST("payment", checkout_handler.PaymentQueueProducer)
}
