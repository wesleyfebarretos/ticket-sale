package service

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/domain/enum"
	"github.com/wesleyfebarretos/ticket-sale/domain/service/auth"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
	"github.com/wesleyfebarretos/ticket-sale/utils"
)

func GetUsers(c *gin.Context, dbCon *sqlc.Queries) ([]sqlc.GetUsersRow, error) {
	users, err := dbCon.GetUsers(c)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return nil, err
	}

	return users, nil
}

func GetUser(c *gin.Context, dbCon *sqlc.Queries, id int32) (sqlc.GetUserRow, error) {
	user, err := dbCon.GetUser(c, id)
	if err != nil {
		utils.WriteError(c, http.StatusNotFound, fmt.Errorf("user of id %d not found", id))
		return user, err
	}

	return user, nil
}

func CreateUser(c *gin.Context, dbCon *sqlc.Queries, newUser sqlc.CreateUserParams) (sqlc.CreateUserRow, error) {
	var createdUser sqlc.CreateUserRow

	_, err := dbCon.GetUserByEmail(c, newUser.Email)
	if err != nil && err != sql.ErrNoRows {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return createdUser, err
	}

	if err == nil {
		err := fmt.Errorf("email %s already registered", newUser.Email)
		utils.WriteError(c, http.StatusBadRequest, err)
		return createdUser, err
	}

	hashPassword, err := auth.HashPassword(newUser.Password)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return createdUser, err
	}

	newUser.Password = string(hashPassword)
	newUser.Role = enum.USER_ROLE

	createdUser, err = dbCon.CreateUser(c, newUser)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return createdUser, err
	}

	return createdUser, nil
}

func UpdateUser(c *gin.Context, dbCon *sqlc.Queries, user sqlc.UpdateUserParams) error {
	_, err := GetUser(c, dbCon, user.ID)
	if err != nil {
		return err
	}

	_, err = dbCon.GetDifferentUserByEmail(c, sqlc.GetDifferentUserByEmailParams{
		Email: user.Email,
		ID:    user.ID,
	})

	if err != nil && err != sql.ErrNoRows {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return err
	}

	if err == nil {
		err := fmt.Errorf("email %s already registered", user.Email)
		utils.WriteError(c, http.StatusBadRequest, err)
		return err
	}

	user.Role = enum.USER_ROLE

	err = dbCon.UpdateUser(c, user)
	if err != nil {
		utils.WriteError(c, http.StatusNotFound, err)
		return err
	}

	return nil
}

func DestroyUser(c *gin.Context, dbCon *sqlc.Queries, id int32) error {
	_, err := GetUser(c, dbCon, id)
	if err != nil {
		return err
	}

	err = dbCon.DestroyUser(c, id)
	if err != nil {
		utils.WriteError(c, http.StatusNotFound, fmt.Errorf("user of id %d not found", id))
		return err
	}

	return nil
}
