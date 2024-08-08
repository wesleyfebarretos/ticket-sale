package user_controller

import (
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/user_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/users_addresses_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/users_phones_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/user_service"
)

func (_ *GetAllResponseDto) FromDomain(p []user_repository.GetAllResponse) []GetAllResponseDto {
	res := []GetAllResponseDto{}

	for _, u := range p {
		res = append(res, GetAllResponseDto{
			Id:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Role:      string(u.Role),
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}

	return res
}

func (_ *GetOneByIdResponseDto) FromDomain(p *user_repository.GetOneByIdResponse) GetOneByIdResponseDto {
	return GetOneByIdResponseDto{
		Id:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Role:      string(p.Role),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (this *CreateRequestDto) ToDomain() user_service.CreateParams {
	return user_service.CreateParams{
		User: user_repository.CreateParams{
			FirstName: this.FirstName,
			LastName:  this.LastName,
			Email:     this.Email,
			Password:  this.Password,
		},
		Address: users_addresses_repository.CreateParams{
			StreetAddress: this.Address.StreetAddress,
			City:          this.Address.City,
			Complement:    this.Address.Complement,
			State:         this.Address.State,
			PostalCode:    this.Address.PostalCode,
			Country:       this.Address.Country,
			AddressType:   this.Address.AddressType,
			Favorite:      this.Address.Favorite,
		},
		Phone: users_phones_repository.CreateParams{
			Ddd:    this.Phone.Ddd,
			Number: this.Phone.Number,
		},
	}
}

func (_ *CreateResponseDto) FromDomain(p user_service.CreateResponse) CreateResponseDto {
	return CreateResponseDto{
		Id:        int(p.User.ID),
		Role:      p.User.Role,
		FirstName: p.User.FirstName,
		LastName:  p.User.LastName,
		Email:     p.User.Email,
		Address: AddressResponseDto{
			ID:            p.Address.ID,
			UserID:        p.Address.UserID,
			Favorite:      p.Address.Favorite,
			Complement:    p.Address.Complement,
			PostalCode:    p.Address.PostalCode,
			AddressType:   p.Address.AddressType,
			StreetAddress: p.Address.StreetAddress,
			City:          p.Address.City,
			State:         p.Address.State,
			Country:       p.Address.Country,
			CreatedAt:     p.Address.CreatedAt,
			UpdatedAt:     p.Address.UpdatedAt,
		},
		Phone: PhoneResponseDto{
			ID:     p.Phone.ID,
			UserID: p.Phone.UserID,
			Ddd:    p.Phone.Ddd,
			Number: p.Phone.Number,
		},
	}
}

func (this *UpdateRequestDto) ToDomain(userID int32) user_repository.UpdateParams {
	return user_repository.UpdateParams{
		ID:        userID,
		FirstName: this.FirstName,
		LastName:  this.LastName,
		Email:     this.Email,
	}
}

func (_ *GetProfileResponseDto) FromDomain(p *user_repository.GetProfileResponse) *GetProfileResponseDto {
	addresses := []UserProfileAddressResponseDto{}
	phones := []UserProfilePhoneResponseDto{}
	for _, v := range p.Addresses {
		addresses = append(addresses, UserProfileAddressResponseDto{
			ID:            v.ID,
			UserID:        v.UserID,
			Favorite:      &v.Favorite,
			Complement:    &v.Complement,
			PostalCode:    &v.PostalCode,
			AddressType:   &v.AddressType,
			StreetAddress: v.StreetAddress,
			City:          v.City,
			State:         v.State,
			Country:       v.Country,
		})
	}

	for _, v := range p.Phones {
		phones = append(phones, UserProfilePhoneResponseDto{
			ID:     v.ID,
			UserID: v.UserID,
			DDD:    v.DDD,
			Number: v.Number,
			Type:   v.Type,
		})
	}

	return &GetProfileResponseDto{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Role:      p.Role,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Addresses: addresses,
		Phones:    phones,
	}
}
