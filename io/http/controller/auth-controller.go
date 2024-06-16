package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/service/auth"
	"github.com/wesleyfebarretos/ticket-sale/io/dto"
)

type AuthController struct {
	BaseController
}

func NewAuthController(conn *pgx.Conn) *AuthController {
	return &AuthController{
		BaseController{conn: conn},
	}
}

func (ac *AuthController) SignIn(c *gin.Context) {
	body := dto.SignInRequest{}
	ac.ReadBody(c, &body)
	conn := ac.NewConnection()

	response := auth.UserSignIn(c, conn, body)

	c.JSON(http.StatusOK, response)
}
