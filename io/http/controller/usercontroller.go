package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/service"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
	"github.com/wesleyfebarretos/ticket-sale/utils"
)

// TODO:
// Retirar todos os logs  fatal
// Dar marshal nas respostas para retornar json
// Criar services
type UserController struct {
	conn *pgx.Conn
}

func NewUserController(conn *pgx.Conn) *UserController {
	return &UserController{
		conn: conn,
	}
}

func (u *UserController) GetAll(c *gin.Context) {
	conn := sqlc.New(u.conn)
	fmt.Println(sqlc.GetUserFullProfileRow{})

	users, err := service.GetUsers(c, conn)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, users)
}

func (u *UserController) GetOne(c *gin.Context) {
	conn := sqlc.New(u.conn)

	id, err := utils.GetId(c)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	user, err := service.GetUser(c, conn, id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *UserController) Create(c *gin.Context) {
	var b sqlc.CreateUserParams
	c.Bind(&b)

	c.JSON(http.StatusOK, c.ShouldBindJSON)
	// TODO: DEBUG HOW TO GET REST BODY

	return
	conn := sqlc.New(u.conn)

	user, err := service.CreateUser(c, conn, b)
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

	conn := sqlc.New(u.conn)

	var b sqlc.UpdateUserParams

	c.Bind(&b)

	b.ID = id

	err = service.UpdateUser(c, conn, b)
	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, true)
}

func (u *UserController) Destroy(c *gin.Context) {
	conn := sqlc.New(u.conn)

	id, err := utils.GetId(c)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
	}

	err = service.DestroyUser(c, conn, id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, true)
}

func (u *UserController) GetFullProfile(c *gin.Context) {
	conn := sqlc.New(u.conn)

	id, err := utils.GetId(c)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	user, err := service.GetUserFullProfile(c, conn, id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, user)
}
