package test_data

import (
	"context"
	"testing"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/gateway_customer_card_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
)

func NewGatewayCustomerCard(t *testing.T, gatewayID, userID, cardID int32) gateway_customer_card_repository.CreateResponse {
	return gateway_customer_card_repository.New().Create(context.Background(), gateway_customer_card_repository.CreateParams{
		GatewayID:     gatewayID,
		UserID:        userID,
		CardID:        cardID,
		GatewayCardID: utils.Encrypt("testing"),
	})
}
