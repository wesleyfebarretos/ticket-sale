package auth_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/users_repository"
)

func Auth(c *gin.Context, email string) {
	_, err := repository.Users.GetOneByEmailAndRole(c, users_repository.GetOneByEmailAndRoleParams{
		Email: email,
		Role:  roles_enum.USER,
	})
	if err != nil {
		fmt.Println(err)
		panic(exception.UnauthorizedException("email or password invalid"))
	}
	// This method is just an abstraction before pass responsabilty to gin jwt
}
