package user_controller

import (
	"time"
)

type GetAllResponseDto struct {
	Id        int32      `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type GetOneByIdDto struct {
	Id        int32      `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type CreateRequestDto struct {
	FirstName string            `json:"firstName" binding:"required"`
	LastName  string            `json:"lastName" binding:"required"`
	Email     string            `json:"email" binding:"required,email"`
	Password  string            `json:"password" binding:"required"`
	Address   AddressRequestDto `json:"address"`
}

type AddressRequestDto struct {
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

type CreateResponseDto struct {
	Id        int                `json:"id"`
	Role      string             `json:"role"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Email     string             `json:"email"`
	Address   AddressResponseDto `json:"address"`
}

type AddressResponseDto struct {
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

type UpdateRequestDto struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}
