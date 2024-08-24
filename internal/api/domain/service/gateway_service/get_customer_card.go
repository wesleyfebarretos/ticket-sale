package gateway_service

import (
	"context"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/gateway_customer_card_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
)

func GetCustomerCardID(ctx context.Context, cardID, gatewayID int32) *gateway_customer_card_repository.GetByCardAndGatewayIdResponse {
	card := gateway_customer_card_repository.New().GetByCardAndGatewayId(ctx, gatewayID, cardID)
	if card == nil {
		return nil
	}
	card.GatewayCardID = utils.Decrypt(card.GatewayCardID)

	return card
}
