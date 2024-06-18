package user_controller

import (
	"time"
)

type CreateUserRequest struct {
	FirstName string         `json:"firstName" binding:"required"`
	LastName  string         `json:"lastName" binding:"required"`
	Email     string         `json:"email" binding:"required,email"`
	Password  string         `json:"password" binding:"required"`
	Address   AddressRequest `json:"address"`
}

type AddressRequest struct {
	Favorite      *bool      `json:"favorite"`
	Complement    *string    `json:"complement"`
	PostalCode    *string    `json:"postalCode"`
	AddressType   *string    `json:"addressType"`
	StreetAddress string     `json:"streetAddress"`
	City          string     `json:"city"`
	State         string     `json:"state"`
	Country       string     `json:"country"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
}

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

type UpdateUserRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}
