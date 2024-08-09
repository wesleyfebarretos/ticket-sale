package admin_user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/admin_user_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/controller"
)

// GetAdminUsers godoc
//
//	@Tags			Admin Users
//	@Summary		Get All
//	@Description	Get All Admin Users
//	@Produce		json
//	@Success		200	{object}	[]GetAllResponseDto
//	@Failure		500	{object}	exception.HttpException
//	@Failure		401	{object}	middleware.AuthenticationError
//	@Failure		403	{object}	middleware.RolePermissionError
//	@Router			/admin/users [get]
func GetAll(c *gin.Context) {
	adminUsers := admin_user_service.GetAll(c)

	res := GetAllResponseDto{}

	c.JSON(http.StatusOK, res.FromEntity(adminUsers))
}

// GetAdminUserById godoc
//
//	@Tags			Admin Users
//	@Summary		Get One By Id
//
//	@Description	Get one admin user by id
//	@Produce		json
//	@Param			id	path		int	true	"Admin User ID"
//	@Success		200	{object}	GetOneByIdResponseDto
//	@Failure		500	{object}	exception.HttpException
//	@Failure		404	{object}	exception.HttpException
//	@Failure		403	{object}	middleware.RolePermissionError
//	@Failure		401	{object}	middleware.AuthenticationError
//	@Router			/admin/users/{id} [get]
func GetOneById(c *gin.Context) {
	id := controller.GetId(c)

	adminUser := admin_user_service.GetOneById(c, id)

	res := GetOneByIdResponseDto{}

	c.JSON(http.StatusOK, res.FromEntity(adminUser))
}

// GetAdminUserByEmail godoc
//
//	@Tags			Admin Users
//	@Summary		Get One By Email
//	@Description	Get one admin user by email
//	@Produce		json
//	@Param			email	body		GetOneByEmailRequestDto	true	"Admin User Email"
//	@Success		200		{object}	GetOneByEmailResponseDto
//	@Failure		500		{object}	exception.HttpException
//	@Failure		404		{object}	exception.HttpException
//	@Failure		403		{object}	middleware.RolePermissionError
//	@Failure		401		{object}	middleware.AuthenticationError
//	@Router			/admin/users/get-by-email [post]
func GetOneByEmail(c *gin.Context) {
	body := GetOneByEmailRequestDto{}

	controller.ReadBody(c, &body)

	adminUser := admin_user_service.GetOneByEmail(c, body.Email)

	res := GetOneByEmailResponseDto{}

	c.JSON(http.StatusOK, res.FromEntity(adminUser))
}

// CreateAdminUser godoc
//
//	@Tags			Admin Users
//	@Summary		Create An Admin User
//	@Description	Create an admin user
//	@Produce		json
//	@Param			AdminUser	body		CreateRequestDto	true	"New Admin User"
//	@Success		200			{object}	CreateResponseDto
//	@Failure		500			{object}	exception.HttpException
//	@Failure		400			{object}	exception.HttpException
//	@Failure		403			{object}	middleware.RolePermissionError
//	@Failure		401			{object}	middleware.AuthenticationError
//	@Router			/admin/users [post]
func Create(c *gin.Context) {
	body := CreateRequestDto{}

	controller.ReadBody(c, &body)

	newAdminUser := admin_user_service.Create(c, body.ToEntity())

	res := CreateResponseDto{}

	c.JSON(http.StatusCreated, res.FromEntity(newAdminUser))
}

// UpdateAdminUser godoc
//
//	@Tags			Admin Users
//	@Summary		Update An Admin User
//	@Description	Update an admin user
//	@Produce		json
//	@Param			AdminUser	body		UpdateRequestDto	true	"Update Admin User"
//	@Param			id			path		int					true	"Admin User ID"
//	@Success		200			{object}	bool
//	@Failure		500			{object}	exception.HttpException
//	@Failure		400			{object}	exception.HttpException
//	@Failure		404			{object}	exception.HttpException
//	@Failure		403			{object}	middleware.RolePermissionError
//	@Failure		401			{object}	middleware.AuthenticationError
//	@Router			/admin/users/{id} [put]
func Update(c *gin.Context) {
	id := controller.GetId(c)

	body := UpdateRequestDto{}

	controller.ReadBody(c, &body)

	admin_user_service.Update(c, body.ToEntity(id))

	c.JSON(http.StatusOK, true)
}

// DeleteAdminUser godoc
//
//	@Tags			Admin Users
//	@Summary		Delete An Admin User
//	@Description	Delete an admin user
//	@Produce		json
//	@Param			id	path		int	true	"Admin User ID"
//	@Success		200	{object}	bool
//	@Failure		500	{object}	exception.HttpException
//	@Failure		404	{object}	exception.HttpException
//	@Failure		403	{object}	middleware.RolePermissionError
//	@Failure		401	{object}	middleware.AuthenticationError
//	@Router			/admin/users/{id} [delete]
func Delete(c *gin.Context) {
	id := controller.GetId(c)

	admin_user_service.Delete(c, id)

	c.JSON(http.StatusOK, true)
}
