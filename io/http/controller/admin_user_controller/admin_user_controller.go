package admin_user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/service/admin_user_service"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_user_repository"
)

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

func Create(c *gin.Context) {
	body := CreateRequestDto{}

	controller.ReadBody(c, &body)

	newAdminUser := admin_user_service.Create(c, admin_user_repository.CreateParams{
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

func Update(c *gin.Context) {
	id := controller.GetId(c)

	body := UpdateRequestDto{}

	controller.ReadBody(c, &body)

	admin_user_service.Update(c, admin_user_repository.UpdateParams{
		ID:        id,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Role:      admin_user_repository.Roles(body.Role),
	})

	c.JSON(http.StatusOK, true)
}

func Delete(c *gin.Context) {
	id := controller.GetId(c)

	admin_user_service.Delete(c, id)

	c.JSON(http.StatusOK, true)
}
