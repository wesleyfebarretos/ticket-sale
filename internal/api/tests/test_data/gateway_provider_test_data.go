package test_data

import (
	"context"
	"testing"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_gateway_provider_repository"
)

func NewGatewayProvider(t *testing.T, userID int32) admin_gateway_provider_repository.CreateResponse {
	provider := admin_gateway_provider_repository.New().Create(context.Background(), admin_gateway_provider_repository.CreateParams{
		Name:      "testing",
		CreatedBy: userID,
		UpdatedBy: userID,
	})

	return provider
}
