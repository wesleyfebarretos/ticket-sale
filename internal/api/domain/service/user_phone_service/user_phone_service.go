package user_phone_service

import (
	"github.com/gin-gonic/gin"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/users_phones_repository"
)

func Create(c *gin.Context, newUserPhone users_phones_repository.CreateParams) users_phones_repository.UsersPhone {
	userPhone, err := repository.UsersPhones.Create(c, users_phones_repository.CreateParams{
		UserID: newUserPhone.UserID,
		Ddd:    newUserPhone.Ddd,
		Number: newUserPhone.Number,
		Type:   newUserPhone.Type,
	})
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return userPhone
}
