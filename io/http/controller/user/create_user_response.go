package user_controller

import (
	"time"
)

type CreateUserResponse struct {
	Id        int              `json:"id"`
	Role      string           `json:"role"`
	FirstName string           `json:"firstName"`
	LastName  string           `json:"lastName"`
	Email     string           `json:"email"`
	Address   *AddressResponse `json:"address"`
}

type AddressResponse struct {
	ID            int32      `json:"id"`
	UserID        int32      `json:"userId"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
	Favorite      *bool      `json:"favorite"`
	Complement    *string    `json:"complement"`
	PostalCode    *string    `json:"postalCode"`
	AddressType   *string    `json:"addressType"`
	StreetAddress string     `json:"streetAddress"`
	City          string     `json:"city"`
	State         string     `json:"state"`
	Country       string     `json:"country"`
}
