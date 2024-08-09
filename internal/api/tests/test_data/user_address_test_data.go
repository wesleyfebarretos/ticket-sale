package test_data

import (
	"context"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/user_address_repository"
)

func NewUserAddress(userId int32) user_address_repository.CreateResponse {
	favorite := true
	complement := "Moon"
	postalCode := "Jupiter"
	addressType := "House"

	newAddress := user_address_repository.CreateParams{
		Favorite:      &favorite,
		Complement:    &complement,
		PostalCode:    &postalCode,
		AddressType:   &addressType,
		StreetAddress: "Via Láctea",
		City:          "Dark Side",
		State:         "VL",
		Country:       "James Webb",
		UserID:        userId,
	}

	address := user_address_repository.New().Create(context.Background(), newAddress)

	return address
}
