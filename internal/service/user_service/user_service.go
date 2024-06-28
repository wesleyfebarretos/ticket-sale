package user_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/user_repository"
	"github.com/wesleyfebarretos/ticket-sale/utils"
)

func GetAll(c *gin.Context) []user_repository.GetAllRow {
	users, err := repository.User.GetAll(c, enum.USER_ROLE)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return users
}

func GetOneById(c *gin.Context, id int32) user_repository.GetOneByIdRow {
	user, err := repository.User.GetOneById(c, id)
	if err != nil {
		panic(exception.NotFoundException(fmt.Sprintf("user of id %d not found", id)))
	}

	return user
}

func Create(c *gin.Context, newUser user_repository.CreateParams) user_repository.CreateRow {
	var createdUser user_repository.CreateRow

	_, err := repository.User.GetOneByEmail(c, newUser.Email)
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
	newUser.Role = enum.USER_ROLE

	createdUser, err = repository.User.Create(c, newUser)
	if err != nil {
		panic(exception.BadRequestException(err.Error()))
	}

	return createdUser
}

func Update(c *gin.Context, user user_repository.UpdateParams) {
	_, err := repository.User.CheckIfEmailExists(c, user_repository.CheckIfEmailExistsParams{
		Email: user.Email,
		ID:    user.ID,
	})

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	if err == nil {
		panic(exception.BadRequestException(fmt.Sprintf("email %s already registered", user.Email)))
	}

	user.Role = enum.USER_ROLE

	err = repository.User.Update(c, user)
	if err != nil {
		panic(exception.NotFoundException(err.Error()))
	}
}

func GetFullProfile(c *gin.Context, id int32) user_repository.GetFullProfileRow {
	user, err := repository.User.GetFullProfile(c, id)
	if err != nil {
		panic(exception.NotFoundException(fmt.Sprintf("user of id %d not found", id)))
	}

	return user
}
