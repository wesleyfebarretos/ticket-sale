package admin_auth_service

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_user_repository"
)

func Auth(c *gin.Context, email string) {
	_, err := repository.AdminUser.GetOneByEmailAndRoles(c, admin_user_repository.GetOneByEmailAndRolesParams{
		Email:  email,
		Role:   enum.ADMIN_ROLE,
		Role_2: enum.SUPER_ADMIN,
	})
	if err != nil {
		panic(exception.UnauthorizedException("email or password invalid"))
	}
	// This method is just an abstraction before pass responsabilty to gin jwt
}
