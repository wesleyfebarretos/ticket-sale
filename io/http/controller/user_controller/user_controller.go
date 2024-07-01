package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum/phone_types_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/service/user_address_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/service/user_phone_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/service/user_service"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller"
	"github.com/wesleyfebarretos/ticket-sale/repository/users_addresses_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/users_phones_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/users_repository"
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

// GetOneById godoc
//
//	@Summary		Get One By Id
//	@Description	Get one user by id
//	@Tags			Users
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	GetOneByIdResponseDto
//	@Failure		500	{object}	exception.HttpException
//	@Failure		404	{object}	exception.HttpException
//	@Failure		401	{object}	middleware.AuthenticationError
//	@Failure		403	{object}	middleware.RolePermissionError
//	@Router			/users/{id} [get]
func GetOneById(c *gin.Context) {
	id := controller.GetId(c)

	user := user_service.GetOneById(c, id)

	userResponse := GetOneByIdResponseDto{
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

	createUser := users_repository.CreateParams{
		Email:     body.Email,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Password:  body.Password,
	}

	newUserResponse := user_service.Create(c, createUser)

	createUserAddress := users_addresses_repository.CreateParams{
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

	newUserPhone := user_phone_service.Create(c, users_phones_repository.CreateParams{
		UserID: newUserResponse.ID,
		Ddd:    body.Phone.Ddd,
		Number: body.Phone.Number,
		Type:   phone_types_enum.PHONE,
	})

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
		Phone: PhoneResponseDto{
			ID:     newUserPhone.ID,
			UserID: newUserPhone.UserID,
			Ddd:    newUserPhone.Ddd,
			Number: newUserPhone.Number,
		},
	}

	c.JSON(http.StatusCreated, newUser)
}

// UpdateUser godoc
//
//	@Summary		Update An User
//	@Description	Update an user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"User ID"
//	@Param			user	body		UpdateRequestDto	true	"Add user"
//	@Success		200		{object}	bool
//	@Failure		500		{object}	exception.HttpException
//	@Failure		404		{object}	exception.HttpException
//	@Failure		401		{object}	middleware.AuthenticationError
//	@Failure		403		{object}	middleware.RolePermissionError
//	@Router			/users/{id} [put]
func Update(c *gin.Context) {
	user := controller.GetClaims(c)
	body := UpdateRequestDto{}

	controller.ReadBody(c, &body)

	updateUser := users_repository.UpdateParams{
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
