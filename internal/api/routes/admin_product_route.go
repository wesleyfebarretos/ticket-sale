package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/handler/admin_product_handler"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/middleware"
)

func HandleAdminProduct(router *gin.RouterGroup) {
	adminProductRoute := router.Group("admin/products")

	adminProductRoute.Use(middleware.JWT.MiddlewareFunc())
	adminProductRoute.Use(middleware.Authorization(roles_enum.ADMIN))

	adminProductRoute.POST("", admin_product_handler.Create)
	adminProductRoute.GET("", admin_product_handler.GetAll)
	adminProductRoute.GET("details", admin_product_handler.GetAllWithRelations)
	adminProductRoute.GET("uuid/:uuid", admin_product_handler.GetOneByUuid)
	adminProductRoute.GET(":id", admin_product_handler.GetOneById)
	adminProductRoute.PUT(":id", admin_product_handler.Update)
	adminProductRoute.DELETE(":id", admin_product_handler.SoftDelete)
}
