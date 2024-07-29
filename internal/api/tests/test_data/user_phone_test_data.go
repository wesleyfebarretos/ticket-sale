package test_data

import (
	"context"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/phone_types_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/users_phones_repository"
)

func NewUserPhone(userID int32) users_phones_repository.UsersPhone {
	newPhone := users_phones_repository.CreateParams{
		UserID: userID,
		Ddd:    "021",
		Number: "999999999",
		Type:   phone_types_enum.PHONE,
	}

	userPhone, _ := repository.UsersPhones.Create(context.Background(), newPhone)

	return userPhone
}
