package service

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
)

func CreateUserAddress(c *gin.Context, conn *sqlc.Queries, newAddress sqlc.CreateUserAddressParams) sqlc.UsersAddress {
	newUserAddress, err := conn.CreateUserAddress(c, newAddress)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return newUserAddress
}
