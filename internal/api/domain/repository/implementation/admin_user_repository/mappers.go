package admin_user_repository

import (
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_user_connection"
)

func (this *CreateParams) ToEntity() admin_user_connection.CreateParams {
	return admin_user_connection.CreateParams{
		FirstName: this.FirstName,
		LastName:  this.LastName,
		Email:     this.Email,
		Password:  this.Password,
		Role:      admin_user_connection.Roles(this.Role),
	}
}

func (this *UpdateParams) ToEntity() admin_user_connection.UpdateParams {
	return admin_user_connection.UpdateParams{
		ID:        this.ID,
		FirstName: this.FirstName,
		LastName:  this.LastName,
		Email:     this.Email,
		Role:      admin_user_connection.Roles(this.Role),
	}
}

func (this *CheckIfEmailExistsParams) ToEntity() admin_user_connection.CheckIfEmailExistsParams {
	return admin_user_connection.CheckIfEmailExistsParams{
		Email: this.Email,
		ID:    this.ID,
	}
}

func (this *CheckIfEmailExistsResponse) FromEntity(p admin_user_connection.CheckIfEmailExistsRow) *CheckIfEmailExistsResponse {
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

func (this *GetOneByIdParams) ToEntity() admin_user_connection.GetOneByIdParams {
	return admin_user_connection.GetOneByIdParams{
		ID:   this.ID,
		Role: admin_user_connection.Roles(this.Role),
	}
}

func (this *CreateResponse) FromEntity(p admin_user_connection.CreateRow) CreateResponse {
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

func (this *GetAllResponse) FromEntity(p []admin_user_connection.GetAllRow) []GetAllResponse {
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

func (this *GetOneByIdResponse) FromEntity(p admin_user_connection.GetOneByIdRow) *GetOneByIdResponse {
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

func (this *GetOneByEmailResponse) FromEntity(p admin_user_connection.GetOneByEmailRow) *GetOneByEmailResponse {
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

func (this *GetOneByEmailParams) ToEntity() admin_user_connection.GetOneByEmailParams {
	return admin_user_connection.GetOneByEmailParams{
		Email: this.Email,
		Role:  admin_user_connection.Roles(this.Role),
	}
}

func (this *GetOneByEmailAndRolesParams) ToEntity() admin_user_connection.GetOneByEmailAndRolesParams {
	return admin_user_connection.GetOneByEmailAndRolesParams{
		Email:  this.Email,
		Role:   admin_user_connection.Roles(this.Role),
		Role_2: admin_user_connection.Roles(this.Role_2),
	}
}

func (this *GetOneByEmailAndRolesResponse) FromEntity(p admin_user_connection.GetOneByEmailAndRolesRow) *GetOneByEmailAndRolesResponse {
	return &GetOneByEmailAndRolesResponse{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Role:      string(p.Role),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
