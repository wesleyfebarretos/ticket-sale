package auth_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	// body := SignInRequest{}
	// controller.ReadBody(c, &body)
	//
	// response := UserSignIn(c, body)
	response := true

	c.JSON(http.StatusOK, response)
}
