package gateway_customer_card_repository

import (
	"context"
	"time"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/gateway_customer_card_connection"
)

type CreateParams struct {
	GatewayID     int32  `json:"gatewayId"`
	UserID        int32  `json:"userId"`
	CardID        int32  `json:"cardId"`
	GatewayCardID string `json:"gatewayCardId"`
}

type CreateResponse struct {
	ID            int32     `json:"id"`
	GatewayID     int32     `json:"gatewayId"`
	UserID        int32     `json:"userId"`
	CardID        int32     `json:"cardId"`
	GatewayCardID string    `json:"gatewayCardId"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func (this *CreateParams) ToEntity() gateway_customer_card_connection.CreateParams {
	return gateway_customer_card_connection.CreateParams{
		GatewayID:     this.GatewayID,
		UserID:        this.UserID,
		CardID:        this.CardID,
		GatewayCardID: this.GatewayCardID,
	}
}

func (this *CreateResponse) FromEntity(p gateway_customer_card_connection.FinGatewayCustomerCard) CreateResponse {
	return CreateResponse{
		ID:            p.ID,
		GatewayID:     p.GatewayID,
		UserID:        p.UserID,
		CardID:        p.CardID,
		GatewayCardID: p.GatewayCardID,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

func (this *GatewayCustomerCardRepository) Create(ctx context.Context, p CreateParams) CreateResponse {
	customerCard, err := this.queries.Create(ctx, p.ToEntity())
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := CreateResponse{}

	return res.FromEntity(customerCard)
}
