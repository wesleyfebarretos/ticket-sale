package gateway_customer_card_repository

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/gateway_customer_card_connection"
)

type GetByUserAndGatewayIdResponse struct {
	ID            int32  `json:"id"`
	GatewayID     int32  `json:"gatewayId"`
	UserID        int32  `json:"userId"`
	CardID        int32  `json:"cardId"`
	GatewayCardID string `json:"gatewayCardId"`
}

func (this *GetByUserAndGatewayIdResponse) FromEntity(p []gateway_customer_card_connection.GetByUserAndGatewayIdRow) []GetByUserAndGatewayIdResponse {
	res := []GetByUserAndGatewayIdResponse{}

	for _, v := range p {
		res = append(res, GetByUserAndGatewayIdResponse{
			ID:            v.ID,
			GatewayID:     v.GatewayID,
			UserID:        v.UserID,
			CardID:        v.CardID,
			GatewayCardID: v.GatewayCardID,
		})
	}

	return res
}

func (this *GatewayCustomerCardRepository) GetByUserAndGatewayId(ctx context.Context, gatewayID, userID int32) []GetByUserAndGatewayIdResponse {
	customerCards, err := this.queries.GetByUserAndGatewayId(ctx, gateway_customer_card_connection.GetByUserAndGatewayIdParams{
		UserID:    userID,
		GatewayID: gatewayID,
	})
	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetByUserAndGatewayIdResponse{}

	return res.FromEntity(customerCards)
}
