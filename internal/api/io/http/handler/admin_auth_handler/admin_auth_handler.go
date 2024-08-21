package admin_auth_handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/admin_auth_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/handler"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/middleware"
)

// AdminAuth godoc
//
//	@Summary		Sign In
//	@Description	Sign In
//	@Tags			Admin Users
//	@Accept			json
//	@Produce		json
//	@Param			login	body		SignInRequestDto	true	"Sign In"
//	@Success		200		{object}	SignInResponseDto
//	@Failure		401		{object}	middleware.AuthenticationError
//	@Router			/admin/auth [post]
func Auth(c *gin.Context) {
	body := SignInRequestDto{}

	handler.ReadBody(c, &body)

	admin_auth_service.Auth(c, body.Email)

	recreatedBody, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to recreate request body"})
		return
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(recreatedBody))

	middleware.JWT.LoginHandler(c)
}
