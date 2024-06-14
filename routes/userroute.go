package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller"
)

func HandleUser(router *gin.Engine, userController *controller.UserController) {
	userRoute := router.Group("users")

	userRoute.GET("", userController.GetAll)
	userRoute.GET(":id", userController.GetOne)
	userRoute.POST("", userController.Create)
	userRoute.PUT(":id", userController.Update)
	userRoute.DELETE(":id", userController.Destroy)

	// USE BY REFERENCE
	// user.GET(":id", func(c *gin.Context) {
	// 	var user string
	// 	if value, ok := c.Params.Get("id"); ok {
	// 		user = value
	// 	} else {https://docs.sqlc.dev/
	// 		user = "not found"
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": fmt.Sprintf("Show user %s", user),
	// 	})
	// })
}
