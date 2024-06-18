package user_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/infra/db"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
	"github.com/wesleyfebarretos/ticket-sale/utils"
)

func GetAll(c *gin.Context) []sqlc.GetUsersRow {
	users, err := db.Query.GetUsers(c)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return users
}

func GetOneById(c *gin.Context, id int32) sqlc.GetUserRow {
	user, err := db.Query.GetUser(c, id)
	if err != nil {
		panic(exception.NotFoundException("user of id %d not found"))
	}

	return user
}

func Create(c *gin.Context, newUser sqlc.CreateUserParams) sqlc.CreateUserRow {
	var createdUser sqlc.CreateUserRow

	_, err := db.Query.GetUserByEmail(c, newUser.Email)
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

	createdUser, err = db.Query.CreateUser(c, newUser)
	if err != nil {
		panic(exception.BadRequestException(err.Error()))
	}

	return createdUser
}

func Update(c *gin.Context, user sqlc.UpdateUserParams) {
	_, err := db.Query.GetDifferentUserByEmail(c, sqlc.GetDifferentUserByEmailParams{
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

	err = db.Query.UpdateUser(c, user)
	if err != nil {
		panic(exception.NotFoundException(err.Error()))
	}
}

func Delete(c *gin.Context, id int32) {
	err := db.Query.DeleteUser(c, id)
	if err != nil {
		panic(exception.NotFoundException(fmt.Sprintf("user of id %d not found", id)))
	}
}

func GetFullProfile(c *gin.Context) sqlc.GetUserFullProfileRow {
	claims := controller.GetClaims(c)
	user, err := db.Query.GetUserFullProfile(c, claims.Id)
	if err != nil {
		panic(exception.NotFoundException(fmt.Sprintf("user of id %d not found", claims.Id)))
	}

	return user
}
