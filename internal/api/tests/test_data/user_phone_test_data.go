package test_data

import (
	"context"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/phone_types_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/user_phone_repository"
)

func NewUserPhone(userID int32) user_phone_repository.CreateResponse {
	newPhone := user_phone_repository.CreateParams{
		UserID: userID,
		Ddd:    "021",
		Number: "999999999",
		Type:   phone_types_enum.PHONE,
	}

	userPhone := user_phone_repository.New().Create(context.Background(), newPhone)

	return userPhone
}
