package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/user_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/controller"
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

	usersResponse := GetAllResponseDto{}

	c.JSON(http.StatusOK, usersResponse.FromDomain(users))
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

	userResponse := GetOneByIdResponseDto{}

	c.JSON(http.StatusOK, userResponse.FromDomain(user))
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

	domainBody := body.ToDomain()

	newUserResponse := user_service.Create(c, domainBody)

	res := CreateResponseDto{}

	c.JSON(http.StatusCreated, res.FromDomain(newUserResponse))
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

	user_service.Update(c, body.ToDomain(user.Id))

	c.JSON(http.StatusOK, true)
}

func GetFullProfile(c *gin.Context) {
	claims := controller.GetClaims(c)

	user := user_service.GetFullProfile(c, claims.Id)

	res := GetProfileResponseDto{}

	c.JSON(http.StatusOK, res.FromDomain(user))
}
