package gateway_customer_repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/gateway_customer_connection"
)

type FindOneByGatewayAndUserIdParams struct {
	UserID    int32 `json:"userId"`
	GatewayID int32 `json:"gatewayId"`
}

type FindOneByGatewayAndUserIdResponse struct {
	UserID            int32     `json:"userId"`
	GatewayID         int32     `json:"gatewayId"`
	GatewayCustomerID string    `json:"gatewayCustomerId"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

func (this *FindOneByGatewayAndUserIdParams) ToEntity() gateway_customer_connection.FindOneByGatewayAndUserIdParams {
	return gateway_customer_connection.FindOneByGatewayAndUserIdParams{
		UserID:    this.UserID,
		GatewayID: this.GatewayID,
	}
}

func (this *FindOneByGatewayAndUserIdResponse) FromEntity(p gateway_customer_connection.FinGatewayCustomer) *FindOneByGatewayAndUserIdResponse {
	return &FindOneByGatewayAndUserIdResponse{
		UserID:            p.UserID,
		GatewayID:         p.GatewayID,
		GatewayCustomerID: p.GatewayCustomerID,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
	}
}

func (this *GatewayCustomerRepository) FindOneByGatewayAndUserId(
	c context.Context,
	param FindOneByGatewayAndUserIdParams,
) *FindOneByGatewayAndUserIdResponse {
	customer, err := this.queries.FindOneByGatewayAndUserId(c, param.ToEntity())
	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := FindOneByGatewayAndUserIdResponse{}

	return res.FromEntity(customer)
}
