package auth_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/user_repository"
)

func Auth(c *gin.Context, email string) {
	_, err := repository.User.GetOneByEmailAndRole(c, user_repository.GetOneByEmailAndRoleParams{
		Email: email,
		Role:  enum.USER_ROLE,
	})
	if err != nil {
		fmt.Println(err)
		panic(exception.UnauthorizedException("email or password invalid"))
	}
	// This method is just an abstraction before pass responsabilty to gin jwt
}
