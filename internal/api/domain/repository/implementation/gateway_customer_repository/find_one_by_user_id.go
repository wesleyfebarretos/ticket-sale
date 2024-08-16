package gateway_customer_repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/gateway_customer_connection"
)

type FindOneByUserIdResponse struct {
	UserID            int32     `json:"userId"`
	GatewayID         int32     `json:"gatewayId"`
	GatewayCustomerID string    `json:"gatewayCustomerId"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

func (this *FindOneByUserIdResponse) FromEntity(p gateway_customer_connection.FinGatewayCustomer) *FindOneByUserIdResponse {
	return &FindOneByUserIdResponse{
		UserID:            p.UserID,
		GatewayID:         p.GatewayID,
		GatewayCustomerID: p.GatewayCustomerID,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
	}
}

func (this *GatewayCustomerRepository) FindOneByUserId(c context.Context, id int32) *FindOneByUserIdResponse {
	customer, err := this.queries.FindOneByUserId(c, id)
	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := FindOneByUserIdResponse{}

	return res.FromEntity(customer)
}
