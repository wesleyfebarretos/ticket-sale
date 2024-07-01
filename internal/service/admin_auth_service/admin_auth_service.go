package admin_auth_service

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_users_repository"
)

func Auth(c *gin.Context, email string) {
	_, err := repository.AdminUser.GetOneByEmailAndRoles(c, admin_users_repository.GetOneByEmailAndRolesParams{
		Email:  email,
		Role:   roles_enum.ADMIN,
		Role_2: roles_enum.SUPER_ADMIN,
	})
	if err != nil {
		panic(exception.UnauthorizedException("email or password invalid"))
	}
	// This method is just an abstraction before pass responsabilty to gin jwt
}
