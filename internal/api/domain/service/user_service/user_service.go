package user_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/users_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
)

func GetAll(c *gin.Context) []users_repository.GetAllRow {
	users, err := repository.Users.GetAll(c, roles_enum.USER)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return users
}

func GetOneById(c *gin.Context, id int32) users_repository.GetOneByIdRow {
	user, err := repository.Users.GetOneById(c, users_repository.GetOneByIdParams{
		ID:   id,
		Role: roles_enum.USER,
	})
	if err != nil {
		panic(exception.NotFoundException(fmt.Sprintf("user of id %d not found", id)))
	}

	return user
}

func Create(c *gin.Context, newUser users_repository.CreateParams) users_repository.CreateRow {
	var createdUser users_repository.CreateRow

	_, err := repository.Users.GetOneByEmail(c, newUser.Email)
	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	if err == nil {
		panic(exception.BadRequestException(fmt.Sprintf("email %s already registered", newUser.Email)))
	}

	hashPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	newUser.Password = string(hashPassword)
	newUser.Role = roles_enum.USER

	createdUser, err = repository.Users.Create(c, newUser)
	if err != nil {
		panic(exception.BadRequestException(err.Error()))
	}

	return createdUser
}

func Update(c *gin.Context, user users_repository.UpdateParams) {
	_, err := repository.Users.CheckIfEmailExists(c, users_repository.CheckIfEmailExistsParams{
		Email: user.Email,
		ID:    user.ID,
	})

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	if err == nil {
		panic(exception.BadRequestException(fmt.Sprintf("email %s already registered", user.Email)))
	}

	user.Role = roles_enum.USER

	err = repository.Users.Update(c, user)
	if err != nil {
		panic(exception.NotFoundException(err.Error()))
	}
}

func GetFullProfile(c *gin.Context, id int32) users_repository.GetFullProfileRow {
	user, err := repository.Users.GetFullProfile(c, id)
	if err != nil {
		panic(exception.NotFoundException(fmt.Sprintf("user of id %d not found", id)))
	}

	return user
}
