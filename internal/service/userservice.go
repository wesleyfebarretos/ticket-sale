package service

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/service/auth"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
	"github.com/wesleyfebarretos/ticket-sale/utils"
)

func GetUsers(c *gin.Context, conn *sqlc.Queries) ([]sqlc.GetUsersRow, error) {
	users, err := conn.GetUsers(c)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return []sqlc.GetUsersRow{}, err
	}

	return users, nil
}

func GetUser(c *gin.Context, conn *sqlc.Queries, id int32) (sqlc.GetUserRow, error) {
	user, err := conn.GetUser(c, id)
	if err != nil {
		utils.WriteError(c, http.StatusNotFound, fmt.Errorf("user of id %d not found", id))
		return sqlc.GetUserRow{}, err
	}

	return user, nil
}

func CreateUser(c *gin.Context, conn *sqlc.Queries, newUser sqlc.CreateUserParams) (sqlc.CreateUserRow, error) {
	var createdUser sqlc.CreateUserRow

	_, err := conn.GetUserByEmail(c, newUser.Email)
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

	createdUser, err = conn.CreateUser(c, newUser)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return createdUser, err
	}

	return createdUser, nil
}

func UpdateUser(c *gin.Context, conn *sqlc.Queries, user sqlc.UpdateUserParams) error {
	_, err := GetUser(c, conn, user.ID)
	if err != nil {
		return err
	}

	_, err = conn.GetDifferentUserByEmail(c, sqlc.GetDifferentUserByEmailParams{
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

	err = conn.UpdateUser(c, user)
	if err != nil {
		utils.WriteError(c, http.StatusNotFound, err)
		return err
	}

	return nil
}

func DestroyUser(c *gin.Context, conn *sqlc.Queries, id int32) error {
	_, err := GetUser(c, conn, id)
	if err != nil {
		return err
	}

	err = conn.DestroyUser(c, id)
	if err != nil {
		utils.WriteError(c, http.StatusNotFound, fmt.Errorf("user of id %d not found", id))
		return err
	}

	return nil
}

func GetUserFullProfile(c *gin.Context, conn *sqlc.Queries, id int32) (sqlc.GetUserFullProfileRow, error) {
	user, err := conn.GetUserFullProfile(c, id)
	if err != nil {
		utils.WriteError(c, http.StatusNotFound, fmt.Errorf("user of id %d not found", id))
		return sqlc.GetUserFullProfileRow{}, err
	}

	return user, nil
}
