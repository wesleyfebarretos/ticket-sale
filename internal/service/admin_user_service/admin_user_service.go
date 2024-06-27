package admin_user_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_user_repository"
	"github.com/wesleyfebarretos/ticket-sale/utils"
)

func GetAll(c *gin.Context) []admin_user_repository.GetAllRow {
	adminUsers, err := repository.AdminUser.GetAll(c, enum.ADMIN_ROLE)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return adminUsers
}

func GetOneById(c *gin.Context, id int32) admin_user_repository.GetOneByIdRow {
	adminUser, err := repository.AdminUser.GetOneById(c, admin_user_repository.GetOneByIdParams{
		ID:   id,
		Role: enum.ADMIN_ROLE,
	})
	if err != nil {
		panic(exception.NotFoundException(fmt.Sprintf("admin user of id %d not found", id)))
	}

	return adminUser
}

func GetOneByEmail(c *gin.Context, email string) admin_user_repository.GetOneByEmailRow {
	adminUser, err := repository.AdminUser.GetOneByEmail(c, admin_user_repository.GetOneByEmailParams{
		Email: email,
		Role:  enum.ADMIN_ROLE,
	})
	if err != nil {
		panic(exception.NotFoundException(fmt.Sprintf("admin user of email %s not found", email)))
	}

	return adminUser
}

func Create(c *gin.Context, newAdminUser admin_user_repository.CreateParams) admin_user_repository.CreateRow {
	_, err := repository.AdminUser.GetOneByEmail(c, admin_user_repository.GetOneByEmailParams{
		Email: newAdminUser.Email,
		Role:  enum.ADMIN_ROLE,
	})

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	if err == nil {
		panic(exception.BadRequestException(fmt.Sprintf("email %s already registered", newAdminUser.Email)))
	}

	hashPassword, err := utils.HashPassword(newAdminUser.Password)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	newAdminUser.Password = string(hashPassword)
	newAdminUser.Role = enum.ADMIN_ROLE

	createdAdminUser, err := repository.AdminUser.Create(c, newAdminUser)
	if err != nil {
		panic(exception.BadRequestException(err.Error()))
	}

	return createdAdminUser
}

func Update(c *gin.Context, adminUser admin_user_repository.UpdateParams) {
	_, err := repository.AdminUser.CheckIfEmailExists(c, admin_user_repository.CheckIfEmailExistsParams{
		Email: adminUser.Email,
		ID:    adminUser.ID,
	})

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	if err == nil {
		panic(exception.BadRequestException(fmt.Sprintf("email %s already registered", adminUser.Email)))
	}

	err = repository.AdminUser.Update(c, adminUser)
	if err != nil {
		panic(exception.NotFoundException(err.Error()))
	}
}
