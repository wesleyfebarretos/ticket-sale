// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package users_addresses_repository

import (
	"time"
)

type UsersAddress struct {
	ID            int32      `json:"id"`
	UserID        int32      `json:"userId"`
	StreetAddress string     `json:"streetAddress"`
	City          string     `json:"city"`
	Complement    *string    `json:"complement"`
	State         string     `json:"state"`
	PostalCode    *string    `json:"postalCode"`
	Country       string     `json:"country"`
	AddressType   *string    `json:"addressType"`
	Favorite      *bool      `json:"favorite"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
}
