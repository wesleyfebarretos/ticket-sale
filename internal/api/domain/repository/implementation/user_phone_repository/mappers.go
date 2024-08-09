package user_phone_repository

import (
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/user_phone_connection"
)

func (this *CreateParams) ToEntity() user_phone_connection.CreateParams {
	return user_phone_connection.CreateParams{
		UserID: this.UserID,
		Ddd:    this.Ddd,
		Number: this.Number,
		Type:   user_phone_connection.PhoneTypes(this.Type),
	}
}

func (this *CreateResponse) FromEntity(p user_phone_connection.UsersPhone) CreateResponse {
	return CreateResponse{
		ID:     p.ID,
		UserID: p.UserID,
		Ddd:    p.Ddd,
		Number: p.Number,
		Type:   string(p.Type),
	}
}
