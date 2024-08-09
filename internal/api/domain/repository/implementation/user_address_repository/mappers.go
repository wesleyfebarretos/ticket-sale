package user_address_repository

import (
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/user_address_connection"
)

func (this *CreateParams) ToEntity() user_address_connection.CreateParams {
	return user_address_connection.CreateParams{
		UserID:        this.UserID,
		StreetAddress: this.StreetAddress,
		City:          this.City,
		Complement:    this.Complement,
		State:         this.State,
		PostalCode:    this.PostalCode,
		Country:       this.Country,
		AddressType:   this.AddressType,
		Favorite:      this.Favorite,
	}
}

func (this *CreateResponse) FromEntity(p user_address_connection.UsersAddress) CreateResponse {
	return CreateResponse{
		ID:            p.ID,
		UserID:        p.UserID,
		StreetAddress: p.StreetAddress,
		City:          p.City,
		Complement:    p.Complement,
		State:         p.State,
		PostalCode:    p.PostalCode,
		Country:       p.Country,
		AddressType:   p.AddressType,
		Favorite:      p.Favorite,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

func (this *UpdateParams) ToEntity() user_address_connection.UpdateParams {
	return user_address_connection.UpdateParams{
		ID:            this.ID,
		StreetAddress: this.StreetAddress,
		City:          this.City,
		Complement:    this.Complement,
		State:         this.State,
		PostalCode:    this.PostalCode,
		Country:       this.Country,
		AddressType:   this.AddressType,
		Favorite:      this.Favorite,
	}
}
