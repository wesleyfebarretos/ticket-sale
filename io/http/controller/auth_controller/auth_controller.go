package auth_controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/service/auth_service"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller"
	"github.com/wesleyfebarretos/ticket-sale/middleware"
)

// Auth godoc
//
//	@Summary		Sign In
//	@Description	Sign In
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			login	body		SignInRequestDto	true	"Sign In"
//	@Success		200		{object}	SignInResponseDto
//	@Failure		401		{object}	middleware.AuthenticationError
//	@Router			/auth [post]
func Auth(c *gin.Context) {
	body := SignInRequestDto{}

	controller.ReadBody(c, &body)

	auth_service.Auth(c, body.Email)

	recreatedBody, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to recreate request body"})
		return
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(recreatedBody))

	middleware.JWT.LoginHandler(c)
}
