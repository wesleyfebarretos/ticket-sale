package user_controller

import (
	"time"
)

// REQUESTS
type AddressRequestDto struct {
	Favorite      *bool     `json:"favorite" example:"true"`
	Complement    *string   `json:"complement" example:"Apt 101"`
	PostalCode    *string   `json:"postalCode" example:"12345"`
	AddressType   *string   `json:"addressType" example:"home"`
	StreetAddress string    `json:"streetAddress" example:"123 Main St"`
	City          string    `json:"city" example:"Springfield"`
	State         string    `json:"state" example:"IL"`
	Country       string    `json:"country" example:"USA"`
	CreatedAt     time.Time `json:"createdAt" example:"2023-01-01T00:00:00Z"`
	UpdatedAt     time.Time `json:"updatedAt" example:"2023-01-02T00:00:00Z"`
}

type UpdateRequestDto struct {
	FirstName string `json:"firstName" binding:"required" example:"John Update"`
	LastName  string `json:"lastName" binding:"required" example:"Doe update"`
	Email     string `json:"email" binding:"required,email" example:"johndoeupdate@gmail.com"`
}

type CreateRequestDto struct {
	FirstName string            `json:"firstName" binding:"required" example:"John"`
	LastName  string            `json:"lastName" binding:"required" example:"Doe"`
	Email     string            `json:"email" binding:"required,email" example:"johndoe@gmail.com"`
	Password  string            `json:"password" binding:"required" example:"123456"`
	Address   AddressRequestDto `json:"address"`
	Phone     PhoneRequestDto   `json:"phone"`
}

type PhoneRequestDto struct {
	Ddd    string `json:"ddd" binding:"required,number,max=5" example:"021"`
	Number string `json:"number" binding:"required,number,max=10" example:"999999999"`
}

// RESPONSES
type GetAllResponseDto struct {
	Id        int32     `json:"id" example:"1"`
	FirstName string    `json:"firstName" example:"John"`
	LastName  string    `json:"lastName" example:"Doe"`
	Email     string    `json:"email" example:"johndoe@gmail.com"`
	Role      string    `json:"role" example:"user"`
	CreatedAt time.Time `json:"createdAt" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2023-01-01T00:00:00Z"`
}

type GetOneByIdResponseDto struct {
	Id        int32     `json:"id" example:"1"`
	FirstName string    `json:"firstName" example:"John"`
	LastName  string    `json:"lastName" example:"Doe"`
	Email     string    `json:"email" example:"johndoe@gmail.com"`
	Role      string    `json:"role" example:"user"`
	CreatedAt time.Time `json:"createdAt" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2023-01-01T00:00:00Z"`
}

type CreateResponseDto struct {
	Id        int                `json:"id" example:"1"`
	Role      string             `json:"role" example:"user"`
	FirstName string             `json:"firstName" example:"John"`
	LastName  string             `json:"lastName" example:"Doe"`
	Email     string             `json:"email" example:"johndoe@gmail.com"`
	Address   AddressResponseDto `json:"address"`
	Phone     PhoneResponseDto   `json:"phone"`
}

type AddressResponseDto struct {
	ID            int32     `json:"id" example:"1"`
	UserID        int32     `json:"userId" example:"1"`
	Favorite      *bool     `json:"favorite" example:"true"`
	Complement    *string   `json:"complement" example:"Apt 101"`
	PostalCode    *string   `json:"postalCode" example:"12345"`
	AddressType   *string   `json:"addressType" example:"home"`
	StreetAddress string    `json:"streetAddress" example:"123 Main St"`
	City          string    `json:"city" example:"Springfield"`
	State         string    `json:"state" example:"IL"`
	Country       string    `json:"country" example:"USA"`
	CreatedAt     time.Time `json:"createdAt" example:"2023-01-01T00:00:00Z"`
	UpdatedAt     time.Time `json:"updatedAt" example:"2023-01-02T00:00:00Z"`
}

type PhoneResponseDto struct {
	ID     int32  `json:"id" example:"1"`
	UserID int32  `json:"userId" example:"2"`
	Ddd    string `json:"ddd" example:"021"`
	Number string `json:"number" example:"999999999"`
}

type GetProfileResponseDto struct {
	ID        int32                           `json:"id"`
	FirstName string                          `json:"firstName"`
	LastName  string                          `json:"lastName"`
	Email     string                          `json:"email"`
	Role      string                          `json:"role"`
	CreatedAt time.Time                       `json:"createdAt"`
	UpdatedAt time.Time                       `json:"updatedAt"`
	Addresses []UserProfileAddressResponseDto `json:"addresses"`
	Phones    []UserProfilePhoneResponseDto   `json:"phones"`
}

type UserProfilePhoneResponseDto struct {
	ID     int32  `json:"id" example:"1"`
	UserID int32  `json:"userId" example:"2"`
	DDD    string `json:"ddd" example:"021"`
	Number string `json:"number" example:"999999999"`
	Type   string `json:"type" example:"phone"`
}

type UserProfileAddressResponseDto struct {
	ID            int32   `json:"id" example:"1"`
	UserID        int32   `json:"userId" example:"1"`
	Favorite      *bool   `json:"favorite" example:"true"`
	Complement    *string `json:"complement" example:"Apt 101"`
	PostalCode    *string `json:"postalCode" example:"12345"`
	AddressType   *string `json:"addressType" example:"home"`
	StreetAddress string  `json:"streetAddress" example:"123 Main St"`
	City          string  `json:"city" example:"Springfield"`
	State         string  `json:"state" example:"IL"`
	Country       string  `json:"country" example:"USA"`
}
