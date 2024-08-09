package admin_user_controller

import (
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_user_repository"
)

func (this *GetAllResponseDto) FromEntity(p []admin_user_repository.GetAllResponse) []GetAllResponseDto {
	res := []GetAllResponseDto{}

	for _, v := range p {
		res = append(res, GetAllResponseDto{
			ID:        v.ID,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Email:     v.Email,
			Role:      v.Role,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return res
}

func (this *GetOneByIdResponseDto) FromEntity(p *admin_user_repository.GetOneByIdResponse) GetOneByIdResponseDto {
	return GetOneByIdResponseDto{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Role:      p.Role,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (this *GetOneByEmailResponseDto) FromEntity(p *admin_user_repository.GetOneByEmailResponse) GetOneByEmailResponseDto {
	return GetOneByEmailResponseDto{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Role:      p.Role,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (this *CreateRequestDto) ToEntity() admin_user_repository.CreateParams {
	return admin_user_repository.CreateParams{
		FirstName: this.FirstName,
		LastName:  this.LastName,
		Email:     this.Email,
		Password:  this.Password,
		Role:      roles_enum.ADMIN,
	}
}

func (this *CreateResponseDto) FromEntity(p admin_user_repository.CreateResponse) CreateResponseDto {
	return CreateResponseDto{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Role:      p.Role,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (this *UpdateRequestDto) ToEntity(userID int32) admin_user_repository.UpdateParams {
	return admin_user_repository.UpdateParams{
		ID:        userID,
		FirstName: this.FirstName,
		LastName:  this.LastName,
		Email:     this.Email,
		Role:      this.Role,
	}
}
