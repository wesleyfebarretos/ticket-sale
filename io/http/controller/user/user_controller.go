package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user_service "github.com/wesleyfebarretos/ticket-sale/internal/service/user"
	user_address_service "github.com/wesleyfebarretos/ticket-sale/internal/service/user_address"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
)

func GetAll(c *gin.Context) {
	users := user_service.GetAll(c)

	c.JSON(http.StatusOK, users)
}

func GetById(c *gin.Context) {
	id := controller.GetId(c)

	user := user_service.GetOneById(c, id)

	c.JSON(http.StatusOK, user)
}

func Create(c *gin.Context) {
	body := CreateUserRequest{}

	controller.ReadBody(c, &body)

	createUser := sqlc.CreateUserParams{
		Email:     body.Email,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Password:  body.Password,
	}

	newUserResponse := user_service.Create(c, createUser)

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
	newUserAddress := user_address_service.Create(c, createUserAddress)

	newUser := &CreateUserResponse{
		Id:        int(newUserResponse.ID),
		Email:     newUserResponse.Email,
		Role:      string(newUserResponse.Role),
		FirstName: newUserResponse.FirstName,
		LastName:  newUserResponse.LastName,
		Address: &AddressResponse{
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

func Update(c *gin.Context) {
	user := controller.GetClaims(c)
	body := UpdateUserRequest{}

	controller.ReadBody(c, &body)

	updateUser := sqlc.UpdateUserParams{
		ID:        user.Id,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
	}

	user_service.Update(c, updateUser)

	c.JSON(http.StatusOK, true)
}

func GetFullProfile(c *gin.Context) {
	user := user_service.GetFullProfile(c)

	c.JSON(http.StatusOK, user)
}
