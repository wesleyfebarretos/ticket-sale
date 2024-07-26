package user_address_service

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/users_addresses_repository"
)

func Create(c *gin.Context, newAddress users_addresses_repository.CreateParams) users_addresses_repository.UsersAddress {
	newUserAddress, err := repository.UsersAdresses.Create(c, newAddress)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return newUserAddress
}
