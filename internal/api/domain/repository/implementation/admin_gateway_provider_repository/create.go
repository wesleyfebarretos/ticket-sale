package admin_gateway_provider_repository

import (
	"context"
	"time"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_gateway_provider_connection"
)

type CreateParams struct {
	Name      string `json:"name"`
	CreatedBy int32  `json:"createdBy"`
	UpdatedBy int32  `json:"updatedBy"`
}

type CreateResponse struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedBy int32     `json:"createdBy"`
	UpdatedBy int32     `json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (this *CreateParams) ToEntity() admin_gateway_provider_connection.CreateParams {
	return admin_gateway_provider_connection.CreateParams{
		Name:      this.Name,
		CreatedBy: this.CreatedBy,
		UpdatedBy: this.UpdatedBy,
	}
}

func (this *CreateResponse) FromEntity(p admin_gateway_provider_connection.FinGatewayProvider) CreateResponse {
	return CreateResponse{
		ID:        p.ID,
		Name:      p.Name,
		CreatedBy: p.CreatedBy,
		UpdatedBy: p.UpdatedBy,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (this *AdminGatewayProviderRepository) Create(ctx context.Context, p CreateParams) CreateResponse {
	provider, err := this.queries.Create(ctx, p.ToEntity())
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := CreateResponse{}

	return res.FromEntity(provider)
}
