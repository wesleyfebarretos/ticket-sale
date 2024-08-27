package test_data

import (
	"context"
	"testing"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/gateway_customer_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
)

func NewGatewayCustomer(t *testing.T, gatewayID, userID int32) gateway_customer_repository.CreateResponse {
	ctx := context.Background()
	customer := gateway_customer_repository.New().Create(ctx, gateway_customer_repository.CreateParams{
		UserID:            userID,
		GatewayID:         gatewayID,
		GatewayCustomerID: utils.Encrypt("testing"),
	})

	return customer
}
