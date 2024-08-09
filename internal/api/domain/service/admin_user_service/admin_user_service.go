package admin_user_service

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_user_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
)

func GetAll(c *gin.Context) []admin_user_repository.GetAllResponse {
	return admin_user_repository.New().GetAll(c)
}

func GetOneById(c *gin.Context, id int32) *admin_user_repository.GetOneByIdResponse {
	adminUser := admin_user_repository.New().GetOneById(c, admin_user_repository.GetOneByIdParams{
		ID:   id,
		Role: roles_enum.ADMIN,
	})
	if adminUser == nil {
		panic(exception.NotFoundException(fmt.Sprintf("admin user of id %d not found", id)))
	}

	return adminUser
}

func GetOneByEmail(c *gin.Context, email string) *admin_user_repository.GetOneByEmailResponse {
	adminUser := admin_user_repository.New().GetOneByEmail(c, admin_user_repository.GetOneByEmailParams{
		Email: email,
		Role:  roles_enum.ADMIN,
	})

	if adminUser == nil {
		panic(exception.NotFoundException(fmt.Sprintf("admin user of email %s not found", email)))
	}

	return adminUser
}

func Create(c *gin.Context, newAdminUser admin_user_repository.CreateParams) admin_user_repository.CreateResponse {
	repository := admin_user_repository.New()
	adminUser := repository.GetOneByEmail(c, admin_user_repository.GetOneByEmailParams{
		Email: newAdminUser.Email,
		Role:  roles_enum.ADMIN,
	})

	if adminUser != nil {
		panic(exception.BadRequestException(fmt.Sprintf("email %s already registered", newAdminUser.Email)))
	}

	hashPassword, err := utils.HashPassword(newAdminUser.Password)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	newAdminUser.Password = string(hashPassword)
	newAdminUser.Role = roles_enum.ADMIN

	createdAdminUser := repository.Create(c, newAdminUser)

	return createdAdminUser
}

func Update(c *gin.Context, adminUser admin_user_repository.UpdateParams) bool {
	repository := admin_user_repository.New()
	user := repository.CheckIfEmailExists(c, admin_user_repository.CheckIfEmailExistsParams{
		Email: adminUser.Email,
		ID:    adminUser.ID,
	})

	if user != nil {
		panic(exception.BadRequestException(fmt.Sprintf("email %s already registered", adminUser.Email)))
	}

	repository.Update(c, adminUser)

	return true
}

func Delete(c *gin.Context, id int32) bool {
	admin_user_repository.New().Delete(c, id)
	return true
}
