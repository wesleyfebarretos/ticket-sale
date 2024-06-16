package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/service/auth"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
)

func GetUsers(c *gin.Context, query *sqlc.Queries) []sqlc.GetUsersRow {
	users, err := query.GetUsers(c)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return users
}

func GetUser(c *gin.Context, query *sqlc.Queries, id int32) sqlc.GetUserRow {
	user, err := query.GetUser(c, id)
	if err != nil {
		panic(exception.NotFoundException("user of id %d not found"))
	}

	return user
}

func CreateUser(c *gin.Context, query *sqlc.Queries, newUser sqlc.CreateUserParams) sqlc.CreateUserRow {
	var createdUser sqlc.CreateUserRow

	_, err := query.GetUserByEmail(c, newUser.Email)
	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	if err == nil {
		panic(exception.BadRequestException(fmt.Sprintf("email %s already registered", newUser.Email)))
	}

	hashPassword, err := auth.HashPassword(newUser.Password)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	newUser.Password = string(hashPassword)
	newUser.Role = enum.USER_ROLE

	createdUser, err = query.CreateUser(c, newUser)
	if err != nil {
		panic(exception.BadRequestException(err.Error()))
	}

	return createdUser
}

func UpdateUser(c *gin.Context, query *sqlc.Queries, user sqlc.UpdateUserParams) {
	_, err := query.GetDifferentUserByEmail(c, sqlc.GetDifferentUserByEmailParams{
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

	err = query.UpdateUser(c, user)
	if err != nil {
		panic(exception.NotFoundException(err.Error()))
	}
}

func DeleteUser(c *gin.Context, query *sqlc.Queries, id int32) {
	err := query.DeleteUser(c, id)
	if err != nil {
		panic(exception.NotFoundException(fmt.Sprintf("user of id %d not found", id)))
	}
}

func GetUserFullProfile(c *gin.Context, query *sqlc.Queries, id int32) sqlc.GetUserFullProfileRow {
	user, err := query.GetUserFullProfile(c, id)
	if err != nil {
		panic(exception.NotFoundException(fmt.Sprintf("user of id %d not found", id)))
	}

	return user
}
