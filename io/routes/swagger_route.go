package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleSwagger(router *gin.RouterGroup) {
	swaggerRoute := router.Group("swagger")
	swaggerRoute.GET("*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
