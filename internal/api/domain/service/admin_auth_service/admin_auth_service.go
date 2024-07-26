package admin_auth_service

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_users_repository"
)

func Auth(c *gin.Context, email string) {
	_, err := repository.AdminUsers.GetOneByEmailAndRoles(c, admin_users_repository.GetOneByEmailAndRolesParams{
		Email:  email,
		Role:   roles_enum.ADMIN,
		Role_2: roles_enum.SUPER_ADMIN,
	})
	if err != nil {
		panic(exception.UnauthorizedException("email or password invalid"))
	}
	// This method is just an abstraction before pass responsabilty to gin jwt
}
