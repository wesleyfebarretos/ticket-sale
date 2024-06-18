package user_address_service

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/infra/db"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
)

func Create(c *gin.Context, newAddress sqlc.CreateUserAddressParams) sqlc.UsersAddress {
	newUserAddress, err := db.Query.CreateUserAddress(c, newAddress)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return newUserAddress
}
