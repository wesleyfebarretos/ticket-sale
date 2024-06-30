package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/service/user_address_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/service/user_service"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller"
	"github.com/wesleyfebarretos/ticket-sale/repository/user_address_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/user_repository"
)

// GetAll godoc
//
//	@Summary		Get All
//	@Description	Get All Users
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		GetAllResponseDto
//	@Failure		500	{object}	exception.HttpException
//	@Router			/users [get]
func GetAll(c *gin.Context) {
	users := user_service.GetAll(c)

	usersResponse := []GetAllResponseDto{}

	for _, u := range users {
		usersResponse = append(usersResponse, GetAllResponseDto{
			Id:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Role:      string(u.Role),
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, usersResponse)
}

func GetOneById(c *gin.Context) {
	id := controller.GetId(c)

	user := user_service.GetOneById(c, id)

	userResponse := GetOneByIdDto{
		Id:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      string(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	c.JSON(http.StatusOK, userResponse)
}

// Create godoc
//
//	@Summary		Create
//	@Description	Create an new User
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		CreateRequestDto	true	"Add user"
//	@Success		201		{object}	CreateResponseDto
//	@Failure		500		{object}	exception.HttpException
//	@Failure		400		{object}	exception.HttpException
//	@Router			/users [post]
func Create(c *gin.Context) {
	body := CreateRequestDto{}

	controller.ReadBody(c, &body)

	createUser := user_repository.CreateParams{
		Email:     body.Email,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Password:  body.Password,
	}

	newUserResponse := user_service.Create(c, createUser)

	createUserAddress := user_address_repository.CreateParams{
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

	newUser := CreateResponseDto{
		Id:        int(newUserResponse.ID),
		Email:     newUserResponse.Email,
		FirstName: newUserResponse.FirstName,
		LastName:  newUserResponse.LastName,
		Role:      string(newUserResponse.Role),
		Address: AddressResponseDto{
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
	body := UpdateRequestDto{}

	controller.ReadBody(c, &body)

	updateUser := user_repository.UpdateParams{
		ID:        user.Id,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
	}

	user_service.Update(c, updateUser)

	c.JSON(http.StatusOK, true)
}

func GetFullProfile(c *gin.Context) {
	claims := controller.GetClaims(c)

	user := user_service.GetFullProfile(c, claims.Id)

	c.JSON(http.StatusOK, user)
}
