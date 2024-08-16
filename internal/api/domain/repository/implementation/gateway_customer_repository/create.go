package gateway_customer_repository

import (
	"context"
	"time"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/gateway_customer_connection"
)

type CreateParams struct {
	UserID            int32  `json:"userId"`
	GatewayID         int32  `json:"gatewayId"`
	GatewayCustomerID string `json:"gatewayCustomerId"`
}

type CreateResponse struct {
	UserID            int32     `json:"userId"`
	GatewayID         int32     `json:"gatewayId"`
	GatewayCustomerID string    `json:"gatewayCustomerId"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

func (this *CreateParams) ToEntity() gateway_customer_connection.CreateParams {
	return gateway_customer_connection.CreateParams{
		UserID:            this.UserID,
		GatewayID:         this.GatewayID,
		GatewayCustomerID: this.GatewayCustomerID,
	}
}

func (this *CreateResponse) FromEntity(p gateway_customer_connection.FinGatewayCustomer) CreateResponse {
	return CreateResponse{
		UserID:            p.UserID,
		GatewayID:         p.GatewayID,
		GatewayCustomerID: p.GatewayCustomerID,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
	}
}

func (this *GatewayCustomerRepository) Create(c context.Context, param CreateParams) CreateResponse {
	customer, err := this.queries.Create(c, param.ToEntity())
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := CreateResponse{}

	return res.FromEntity(customer)
}
