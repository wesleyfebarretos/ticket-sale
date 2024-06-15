package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/domain/service"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
	"github.com/wesleyfebarretos/ticket-sale/utils"
)

// TODO:
// Retirar todos os logs  fatal
// Dar marshal nas respostas para retornar json
// Criar services
type UserController struct {
	db *sql.DB
}

func NewUserController(db *sql.DB) *UserController {
	return &UserController{
		db: db,
	}
}

func (u *UserController) GetAll(c *gin.Context) {
	dbCon := sqlc.New(u.db)

	users, err := service.GetUsers(c, dbCon)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, users)
}

func (u *UserController) GetOne(c *gin.Context) {
	dbCon := sqlc.New(u.db)

	id, err := utils.GetId(c)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	user, err := service.GetUser(c, dbCon, id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *UserController) Create(c *gin.Context) {
	var b sqlc.CreateUserParams
	c.Bind(&b)
	dbCon := sqlc.New(u.db)

	user, err := service.CreateUser(c, dbCon, b)
	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (u *UserController) Update(c *gin.Context) {
	id, err := utils.GetId(c)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
	}

	dbCon := sqlc.New(u.db)

	var b sqlc.UpdateUserParams

	c.Bind(&b)

	b.ID = id

	err = service.UpdateUser(c, dbCon, b)
	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, true)
}

func (u *UserController) Destroy(c *gin.Context) {
	dbCon := sqlc.New(u.db)

	id, err := utils.GetId(c)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
	}

	err = service.DestroyUser(c, dbCon, id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, true)
}
