package admin_user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/service/admin_user_service"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller"
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

	controller.ReadBody(c, body)

	adminUser := admin_user_service.GetOneByEmail(c, body.email)

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
