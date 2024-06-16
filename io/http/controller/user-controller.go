package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/service"
	"github.com/wesleyfebarretos/ticket-sale/io/dto"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
)

type UserController struct {
	BaseController
	conn *pgx.Conn
}

func NewUserController(conn *pgx.Conn) *UserController {
	return &UserController{
		conn:           conn,
		BaseController: BaseController{},
	}
}

func (u *UserController) GetAll(c *gin.Context) {
	conn := sqlc.New(u.conn)

	users := service.GetUsers(c, conn)

	c.JSON(http.StatusOK, users)
}

func (u *UserController) GetOne(c *gin.Context) {
	conn := sqlc.New(u.conn)

	id := u.GetId(c)

	user := service.GetUser(c, conn, id)

	c.JSON(http.StatusOK, user)
}

func (u *UserController) Create(c *gin.Context) {
	var body dto.CreateUserRequest

	u.ReadBody(c, &body)
	conn := sqlc.New(u.conn)

	createUser := sqlc.CreateUserParams{
		Email:     body.Email,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Password:  body.Password,
	}

	newUserResponse := service.CreateUser(c, conn, createUser)

	var newUser dto.CreateUserResponse

	newUser.Id = int(newUserResponse.ID)

	createUserAddress := sqlc.CreateUserAddressParams{
		UserID:        int32(newUser.Id),
		AddressType:   newUser.Address.AddressType,
		StreetAddress: newUser.Address.StreetAddress,
		City:          newUser.Address.City,
		Complement:    newUser.Address.Complement,
		State:         newUser.Address.State,
		PostalCode:    newUser.Address.PostalCode,
		Country:       newUser.Address.Country,
		Favorite:      newUser.Address.Favorite,
	}

	newUserAddress := service.CreateUserAddress(c, conn, createUserAddress)

	newUser.Address.ID = newUserAddress.ID
	newUser.Address.UserID = newUserAddress.UserID
	newUser.Address.CreatedAt = newUserAddress.CreatedAt
	newUser.Address.UpdatedAt = newUserAddress.UpdatedAt
	c.JSON(http.StatusCreated, newUser)
}

func (u *UserController) Update(c *gin.Context) {
	id := u.GetId(c)

	conn := sqlc.New(u.conn)

	var body dto.UpdateUserRequest

	u.ReadBody(c, &body)

	updateUser := sqlc.UpdateUserParams{
		ID:        id,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
	}

	service.UpdateUser(c, conn, updateUser)

	c.JSON(http.StatusOK, true)
}

func (u *UserController) Destroy(c *gin.Context) {
	conn := sqlc.New(u.conn)

	id := u.GetId(c)

	service.DeleteUser(c, conn, id)

	c.Status(http.StatusOK)
}

func (u *UserController) GetFullProfile(c *gin.Context) {
	conn := sqlc.New(u.conn)

	id := u.GetId(c)

	user := service.GetUserFullProfile(c, conn, id)

	c.JSON(http.StatusOK, user)
}
