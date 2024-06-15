package userservice

import (
	"database/sql"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	enum "github.com/wesleyfebarretos/ticket-sale/domain/enums"
	"github.com/wesleyfebarretos/ticket-sale/domain/service/auth"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
	"github.com/wesleyfebarretos/ticket-sale/utils"
)

func GetAll(c *gin.Context, dbCon *sqlc.Queries) ([]sqlc.GetUsersRow, error) {
	users, err := dbCon.GetUsers(c)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return nil, err
	}

	return users, nil
}

func GetOne(c *gin.Context, dbCon *sqlc.Queries, id int32) (sqlc.GetUserRow, error) {
	user, err := dbCon.GetUser(c, id)
	if err != nil {
		utils.WriteError(c, http.StatusNotFound, fmt.Errorf("user of id %d not found", id))
		return user, err
	}

	return user, nil
}

func Create(c *gin.Context, dbCon *sqlc.Queries, user sqlc.CreateUserParams) (sqlc.CreateUserRow, error) {
	var u sqlc.CreateUserRow
	_, err := dbCon.GetUserByEmail(c, user.Email)
	if err != nil && err != sql.ErrNoRows {
		utils.WriteError(c, http.StatusNotFound, fmt.Errorf("email %s already registered", user.Email))
		return u, err
	}

	roles := []string{enum.ADMIN_ROLE, enum.USER_ROLE, enum.WEBSERVICE_ROLE}

	// TODO:
	// Validar a Role do usuário antes de cadastrar
	// Fazer separação de responsabilidade, por role
	// Esse cadastro vai ser role comum, porém terá rota de ADMIN
	// Que poderá administrar
	if user.Role == "" || slices.Index(roles, string(user.Role)) == -1 {
		err := fmt.Errorf("expected role as %v and receive %v", strings.Join(roles, ","), user.Role)
		utils.WriteError(c, http.StatusBadRequest, err)
		return u, err
	}

	hashPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return u, err
	}

	user.Password = string(hashPassword)

	u, err = dbCon.CreateUser(c, user)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return u, err
	}

	return u, nil
}

func Update(c *gin.Context, dbCon *sqlc.Queries, user sqlc.UpdateUserParams) error {
	_, err := GetOne(c, dbCon, user.ID)
	if err != nil {
		return err
	}

	err = dbCon.UpdateUser(c, user)
	if err != nil {
		utils.WriteError(c, http.StatusNotFound, err)
		return err
	}

	return nil
}

func Destroy(c *gin.Context, dbCon *sqlc.Queries, id int32) error {
	_, err := GetOne(c, dbCon, id)
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
