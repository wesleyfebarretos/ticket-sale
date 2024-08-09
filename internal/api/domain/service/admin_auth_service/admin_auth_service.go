package admin_auth_service

import (
	"github.com/gin-gonic/gin"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_user_repository"
)

func Auth(c *gin.Context, email string) {
	user := admin_user_repository.New().GetOneByEmailAndRoles(c, admin_user_repository.GetOneByEmailAndRolesParams{
		Email:  email,
		Role:   roles_enum.ADMIN,
		Role_2: roles_enum.SUPER_ADMIN,
	})
	if user == nil {
		panic(exception.UnauthorizedException("email or password invalid"))
	}
	// This method is just an abstraction before pass responsabilty to gin jwt
}
