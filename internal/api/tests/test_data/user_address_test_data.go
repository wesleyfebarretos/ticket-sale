package test_data

import (
	"context"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/users_addresses_repository"
)

func NewUserAddress(userId int32) users_addresses_repository.UsersAddress {
	favorite := true
	complement := "Moon"
	postalCode := "Jupiter"
	addressType := "House"

	newAddress := users_addresses_repository.CreateParams{
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

	address, _ := repository.UsersAdresses.Create(context.Background(), newAddress)

	return address
}
