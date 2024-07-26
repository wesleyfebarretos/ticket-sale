package auth_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/users_repository"
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
