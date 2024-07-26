package admin_user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/admin_user_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/controller"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_users_repository"
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

	adminUsersResponse := []GetAllResponseDto{}

	for _, u := range adminUsers {
		adminUsersResponse = append(adminUsersResponse, GetAllResponseDto{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Role:      string(u.Role),
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	c.JSON(http.StatusOK, adminUsersResponse)
}

// GetAdminUserById godoc
//
//	@Tags			Admin Users
//	@Summary		Get One By Id
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

	adminUserResponse := GetOneByIdResponseDto{
		ID:        adminUser.ID,
		FirstName: adminUser.FirstName,
		LastName:  adminUser.LastName,
		Email:     adminUser.Email,
		Role:      string(adminUser.Role),
		CreatedAt: adminUser.CreatedAt,
		UpdatedAt: adminUser.UpdatedAt,
	}

	c.JSON(http.StatusOK, adminUserResponse)
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

	adminUserResponse := GetOneByEmailResponseDto{
		ID:        adminUser.ID,
		FirstName: adminUser.FirstName,
		LastName:  adminUser.LastName,
		Email:     adminUser.Email,
		Role:      string(adminUser.Role),
		CreatedAt: adminUser.CreatedAt,
		UpdatedAt: adminUser.UpdatedAt,
	}

	c.JSON(http.StatusOK, adminUserResponse)
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

	newAdminUser := admin_user_service.Create(c, admin_users_repository.CreateParams{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
	})

	newAdminUserResponse := CreateResponseDto{
		ID:        newAdminUser.ID,
		FirstName: newAdminUser.FirstName,
		LastName:  newAdminUser.LastName,
		Email:     newAdminUser.Email,
		Role:      string(newAdminUser.Role),
		CreatedAt: newAdminUser.CreatedAt,
		UpdatedAt: newAdminUser.UpdatedAt,
	}

	c.JSON(http.StatusCreated, newAdminUserResponse)
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

	admin_user_service.Update(c, admin_users_repository.UpdateParams{
		ID:        id,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Role:      admin_users_repository.Roles(body.Role),
	})

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
