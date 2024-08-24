package gateway_customer_card_repository

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/gateway_customer_card_connection"
)

type GetByCardAndGatewayIdResponse struct {
	ID            int32  `json:"id"`
	GatewayID     int32  `json:"gatewayId"`
	UserID        int32  `json:"userId"`
	CardID        int32  `json:"cardId"`
	GatewayCardID string `json:"gatewayCardId"`
}

func (this *GetByCardAndGatewayIdResponse) FromEntity(p gateway_customer_card_connection.GetByCardAndGatewayIdRow) *GetByCardAndGatewayIdResponse {
	return &GetByCardAndGatewayIdResponse{
		ID:            p.ID,
		GatewayID:     p.GatewayID,
		UserID:        p.UserID,
		CardID:        p.CardID,
		GatewayCardID: p.GatewayCardID,
	}
}

func (this *GatewayCustomerCardRepository) GetByCardAndGatewayId(ctx context.Context, gatewayID, cardID int32) *GetByCardAndGatewayIdResponse {
	customerCards, err := this.queries.GetByCardAndGatewayId(ctx, gateway_customer_card_connection.GetByCardAndGatewayIdParams{
		CardID:    cardID,
		GatewayID: gatewayID,
	})
	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetByCardAndGatewayIdResponse{}

	return res.FromEntity(customerCards)
}
