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
}

func NewUserController(conn *pgx.Conn) *UserController {
	return &UserController{
		BaseController{conn: conn},
	}
}

func (u *UserController) GetAll(c *gin.Context) {
	conn := u.NewConnection()

	users := service.GetUsers(c, conn)

	c.JSON(http.StatusOK, users)
}

func (u *UserController) GetOne(c *gin.Context) {
	conn := u.NewConnection()

	id := u.GetId(c)

	user := service.GetUser(c, conn, id)

	c.JSON(http.StatusOK, user)
}

func (u *UserController) Create(c *gin.Context) {
	body := dto.CreateUserRequest{}

	u.ReadBody(c, &body)
	conn := u.NewConnection()

	createUser := sqlc.CreateUserParams{
		Email:     body.Email,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Password:  body.Password,
	}

	newUserResponse := service.CreateUser(c, conn, createUser)

	createUserAddress := sqlc.CreateUserAddressParams{
		UserID:        int32(newUserResponse.ID),
		AddressType:   body.Address.AddressType,
		StreetAddress: body.Address.StreetAddress,
		City:          body.Address.City,
		Complement:    body.Address.Complement,
		State:         body.Address.State,
		PostalCode:    body.Address.PostalCode,
		Country:       body.Address.Country,
		Favorite:      body.Address.Favorite,
	}
	newUserAddress := service.CreateUserAddress(c, conn, createUserAddress)

	newUser := &dto.CreateUserResponse{
		Id:        int(newUserResponse.ID),
		Email:     newUserResponse.Email,
		Role:      string(newUserResponse.Role),
		FirstName: newUserResponse.FirstName,
		LastName:  newUserResponse.LastName,
		Address: &dto.AddressResponse{
			ID:            newUserAddress.ID,
			UserID:        newUserAddress.UserID,
			City:          newUserAddress.City,
			State:         newUserAddress.State,
			Country:       newUserAddress.Country,
			Complement:    newUserAddress.Complement,
			Favorite:      newUserAddress.Favorite,
			PostalCode:    newUserAddress.PostalCode,
			AddressType:   newUserAddress.AddressType,
			StreetAddress: newUserAddress.StreetAddress,
		},
	}

	c.JSON(http.StatusCreated, newUser)
}

func (u *UserController) Update(c *gin.Context) {
	id := u.GetId(c)

	conn := u.NewConnection()

	body := dto.UpdateUserRequest{}

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
	conn := u.NewConnection()

	id := u.GetId(c)

	service.DeleteUser(c, conn, id)

	c.Status(http.StatusOK)
}

func (u *UserController) GetFullProfile(c *gin.Context) {
	conn := u.NewConnection()

	id := u.GetId(c)

	user := service.GetUserFullProfile(c, conn, id)

	c.JSON(http.StatusOK, user)
}
