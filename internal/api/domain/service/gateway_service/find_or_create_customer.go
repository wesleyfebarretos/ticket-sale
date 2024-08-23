package gateway_service

import (
	"context"
	"errors"
	"fmt"

	stripe_provider "github.com/wesleyfebarretos/ticket-sale/external/providers/gateways/stripe"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/gateway_provider_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/gateway_customer_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/user_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
)

func FindOrCreateCustomer(ctx context.Context, userID int32) (*gateway_customer_repository.FindOneByGatewayAndUserIdResponse, error) {
	user := user_repository.New().GetOneById(ctx, user_repository.GetOneByIdParams{
		ID:   userID,
		Role: roles_enum.USER,
	})

	if user == nil {
		return nil, errors.New("user not found")
	}

	gateway, err := GetActive(ctx)
	if err != nil {
		return nil, err
	}

	gatewayCustomerRepository := gateway_customer_repository.New()

	customer := gatewayCustomerRepository.FindOneByGatewayAndUserId(ctx, gateway_customer_repository.FindOneByGatewayAndUserIdParams{
		UserID:    userID,
		GatewayID: gateway.ID,
	})

	if customer != nil {
		return customer, nil
	}

	var gatewayCustomerId string

	switch gateway.GatewayProviderID {
	case gateway_provider_enum.STRIPE:

		customer, err := stripe_provider.CreateCustomer(&stripe_provider.CreateCustomerDTO{
			Name:  fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			Email: user.Email,
		})
		if err != nil {
			return nil, errors.New("Error on creating customer")
		}

		gatewayCustomerId = customer.ID

	default:
		return nil, errors.New("Integration not available with the provider")
	}

	newCustomer := gatewayCustomerRepository.Create(ctx, gateway_customer_repository.CreateParams{
		UserID:            userID,
		GatewayID:         gateway.ID,
		GatewayCustomerID: utils.Encrypt(gatewayCustomerId),
	})

	customer = &gateway_customer_repository.FindOneByGatewayAndUserIdResponse{
		UserID:            newCustomer.UserID,
		GatewayID:         newCustomer.GatewayID,
		GatewayCustomerID: utils.Decrypt(newCustomer.GatewayCustomerID),
		CreatedAt:         newCustomer.CreatedAt,
		UpdatedAt:         newCustomer.UpdatedAt,
	}

	return customer, nil
}
