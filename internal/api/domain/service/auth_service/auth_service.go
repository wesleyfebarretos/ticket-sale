package auth_service

import (
	"github.com/gin-gonic/gin"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/user_repository"
)

func Auth(c *gin.Context, email string) {
	user := user_repository.New().GetOneByEmailAndRole(c, user_repository.GetOneByEmailAndRoleParams{
		Email: email,
		Role:  roles_enum.USER,
	})
	if user == nil {
		panic(exception.UnauthorizedException("email or password invalid"))
	}
}
