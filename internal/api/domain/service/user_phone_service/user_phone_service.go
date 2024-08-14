package user_phone_service

import (
	"github.com/gin-gonic/gin"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/user_phone_repository"
)

func Create(c *gin.Context, newUserPhone user_phone_repository.CreateParams) user_phone_repository.CreateResponse {
	userPhone := user_phone_repository.New().Create(c, newUserPhone)

	return userPhone
}
