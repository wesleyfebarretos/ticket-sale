package user_address_repository

import "time"

type CreateParams struct {
	UserID        int32   `json:"userId"`
	StreetAddress string  `json:"streetAddress"`
	City          string  `json:"city"`
	Complement    *string `json:"complement"`
	State         string  `json:"state"`
	PostalCode    *string `json:"postalCode"`
	Country       string  `json:"country"`
	AddressType   *string `json:"addressType"`
	Favorite      *bool   `json:"favorite"`
}

type CreateResponse struct {
	ID            int32     `json:"id"`
	UserID        int32     `json:"userId"`
	StreetAddress string    `json:"streetAddress"`
	City          string    `json:"city"`
	Complement    *string   `json:"complement"`
	State         string    `json:"state"`
	PostalCode    *string   `json:"postalCode"`
	Country       string    `json:"country"`
	AddressType   *string   `json:"addressType"`
	Favorite      *bool     `json:"favorite"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type UpdateParams struct {
	ID            int32   `json:"id"`
	StreetAddress string  `json:"streetAddress"`
	City          string  `json:"city"`
	Complement    *string `json:"complement"`
	State         string  `json:"state"`
	PostalCode    *string `json:"postalCode"`
	Country       string  `json:"country"`
	AddressType   *string `json:"addressType"`
	Favorite      *bool   `json:"favorite"`
}
