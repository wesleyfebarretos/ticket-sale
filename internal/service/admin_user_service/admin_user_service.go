package admin_user_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_users_repository"
	"github.com/wesleyfebarretos/ticket-sale/utils"
)

func GetAll(c *gin.Context) []admin_users_repository.GetAllRow {
	adminUsers, err := repository.AdminUsers.GetAll(c, roles_enum.ADMIN)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return adminUsers
}

func GetOneById(c *gin.Context, id int32) admin_users_repository.GetOneByIdRow {
	adminUser, err := repository.AdminUsers.GetOneById(c, admin_users_repository.GetOneByIdParams{
		ID:   id,
		Role: roles_enum.ADMIN,
	})
	if err != nil {
		panic(exception.NotFoundException(fmt.Sprintf("admin user of id %d not found", id)))
	}

	return adminUser
}

func GetOneByEmail(c *gin.Context, email string) admin_users_repository.GetOneByEmailRow {
	adminUser, err := repository.AdminUsers.GetOneByEmail(c, admin_users_repository.GetOneByEmailParams{
		Email: email,
		Role:  roles_enum.ADMIN,
	})
	if err != nil {
		panic(exception.NotFoundException(fmt.Sprintf("admin user of email %s not found", email)))
	}

	return adminUser
}

func Create(c *gin.Context, newAdminUser admin_users_repository.CreateParams) admin_users_repository.CreateRow {
	_, err := repository.AdminUsers.GetOneByEmail(c, admin_users_repository.GetOneByEmailParams{
		Email: newAdminUser.Email,
		Role:  roles_enum.ADMIN,
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
	newAdminUser.Role = roles_enum.ADMIN

	createdAdminUser, err := repository.AdminUsers.Create(c, newAdminUser)
	if err != nil {
		panic(exception.BadRequestException(err.Error()))
	}

	return createdAdminUser
}

func Update(c *gin.Context, adminUser admin_users_repository.UpdateParams) {
	_, err := repository.AdminUsers.CheckIfEmailExists(c, admin_users_repository.CheckIfEmailExistsParams{
		Email: adminUser.Email,
		ID:    adminUser.ID,
	})

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	if err == nil {
		panic(exception.BadRequestException(fmt.Sprintf("email %s already registered", adminUser.Email)))
	}

	err = repository.AdminUsers.Update(c, adminUser)
	if err != nil {
		panic(exception.NotFoundException(err.Error()))
	}
}

func Delete(c *gin.Context, id int32) {
	err := repository.AdminUsers.Delete(c, id)
	if err != nil {
		panic(exception.NotFoundException(fmt.Sprintf("user of id %d not found", id)))
	}
}
