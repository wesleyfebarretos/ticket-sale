package user_address_service

import (
	"github.com/gin-gonic/gin"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/user_address_repository"
)

func Create(c *gin.Context, newAddress user_address_repository.CreateParams) user_address_repository.CreateResponse {
	newUserAddress := user_address_repository.New().Create(c, newAddress)

	return newUserAddress
}
