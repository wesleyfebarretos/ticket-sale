package user_repository

import (
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/user_connection"
)

func (this *CreateParams) ToEntity() user_connection.CreateParams {
	return user_connection.CreateParams{
		FirstName: this.FirstName,
		LastName:  this.LastName,
		Email:     this.Email,
		Password:  this.Password,
		Role:      user_connection.Roles(this.Role),
	}
}

func (this *UpdateParams) ToEntity() user_connection.UpdateParams {
	return user_connection.UpdateParams{
		ID:        this.ID,
		FirstName: this.FirstName,
		LastName:  this.LastName,
		Email:     this.Email,
		Role:      user_connection.Roles(this.Role),
	}
}

func (this *CheckIfEmailExistsParams) ToEntity() user_connection.CheckIfEmailExistsParams {
	return user_connection.CheckIfEmailExistsParams{
		Email: this.Email,
		ID:    this.ID,
	}
}

func (this *GetOneByIdParams) ToEntity() user_connection.GetOneByIdParams {
	return user_connection.GetOneByIdParams{
		ID:   this.ID,
		Role: user_connection.Roles(this.Role),
	}
}

func (this *GetOneByEmailAndRoleParams) ToEntity() user_connection.GetOneByEmailAndRoleParams {
	return user_connection.GetOneByEmailAndRoleParams{
		Email: this.Email,
		Role:  user_connection.Roles(this.Role),
	}
}

func (this *CreateResponse) FromEntity(p user_connection.CreateRow) CreateResponse {
	return CreateResponse{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Role:      string(p.Role),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (this *GetOneByEmailAndRoleResponse) FromEntity(p user_connection.GetOneByEmailAndRoleRow) *GetOneByEmailAndRoleResponse {
	return &GetOneByEmailAndRoleResponse{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Role:      string(p.Role),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (this *GetAllResponse) FromEntity(p []user_connection.GetAllRow) []GetAllResponse {
	res := []GetAllResponse{}

	for _, v := range p {
		res = append(res, GetAllResponse{
			ID:        v.ID,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Email:     v.Email,
			Role:      string(v.Role),
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res
}

func (this *GetProfileResponse) FromEntity(p user_connection.UserProfile) *GetProfileResponse {
	addresses := []UserProfileAddressResponse{}
	phones := []UserProfilePhoneResponse{}

	for _, v := range p.Addresses {
		addresses = append(addresses, UserProfileAddressResponse{
			ID:            v.ID,
			UserID:        v.UserID,
			StreetAddress: v.StreetAddress,
			City:          v.City,
			Complement:    v.Complement,
			State:         v.State,
			PostalCode:    v.PostalCode,
			Country:       v.Country,
			AddressType:   v.AddressType,
			Favorite:      v.Favorite,
		})
	}

	for _, v := range p.Phones {
		phones = append(phones, UserProfilePhoneResponse{
			ID:     v.ID,
			UserID: v.UserID,
			DDD:    v.DDD,
			Number: v.Number,
			Type:   v.Type,
		})
	}

	return &GetProfileResponse{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Role:      string(p.Role),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Addresses: addresses,
		Phones:    phones,
	}
}

func (this *CheckIfEmailExistsResponse) FromEntity(p user_connection.CheckIfEmailExistsRow) *CheckIfEmailExistsResponse {
	return &CheckIfEmailExistsResponse{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Role:      string(p.Role),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (this *GetOneByIdResponse) FromEntity(p user_connection.GetOneByIdRow) *GetOneByIdResponse {
	return &GetOneByIdResponse{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Role:      string(p.Role),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (this *GetOneByEmailResponse) FromEntity(p user_connection.GetOneByEmailRow) *GetOneByEmailResponse {
	return &GetOneByEmailResponse{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Role:      string(p.Role),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (this *GetOneWithPasswordByEmailResponse) FromEntity(p user_connection.GetOneWithPasswordByEmailRow) *GetOneWithPasswordByEmailResponse {
	return &GetOneWithPasswordByEmailResponse{
		ID:        p.ID,
		Password:  p.Password,
		Role:      string(p.Role),
		Email:     p.Email,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
