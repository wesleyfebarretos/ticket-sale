package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
)

// TODO:
// Retirar todos os logs  fatal
// Dar marshal nas respostas para retornar json
// Criar services
type UserController struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *UserController {
	return &UserController{
		db: db,
	}
}

func (u *UserController) GetAll(c *gin.Context) {
	userRepository := sqlc.New(u.db)

	users, err := userRepository.GetUsers(c)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, users)
}

func (u *UserController) GetOne(c *gin.Context) {
	log.Println("Get one users")
	user := UserController{}
	c.JSON(http.StatusOK, user)
}

func (u *UserController) Create(c *gin.Context) {
	var b sqlc.CreateUserParams
	c.Bind(&b)
	userRepository := sqlc.New(u.db)
	newUser, err := userRepository.CreateUser(c, b)
	if err != nil {
		log.Fatal(err)
	}

	userJSON, err := json.Marshal(newUser)
	if err != nil {
		log.Fatal(err)
	}

	response := string(userJSON)

	c.JSON(http.StatusCreated, response)
}

func (u *UserController) Update(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (u *UserController) Destroy(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
