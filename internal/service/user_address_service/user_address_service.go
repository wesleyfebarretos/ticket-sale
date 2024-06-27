package user_address_service

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/user_address_repository"
)

func Create(c *gin.Context, newAddress user_address_repository.CreateParams) user_address_repository.UsersAddress {
	newUserAddress, err := repository.UserAdress.Create(c, newAddress)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return newUserAddress
}
