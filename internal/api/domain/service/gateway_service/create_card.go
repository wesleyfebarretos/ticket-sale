package gateway_service

import (
	"errors"
	"fmt"

	stripe_provider "github.com/wesleyfebarretos/ticket-sale/external/providers/gateways/stripe"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/gateway_provider_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/gateway_customer_card_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/gateway_customer_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/user_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
	"golang.org/x/net/context"
)

type CreateCardDTO struct {
	Number   string
	ExpMonth string
	CVC      string
	ExpYear  string
	CardID   int32
	userID   int32
}

func CreateCard(ctx context.Context, dto CreateCardDTO) (*gateway_customer_card_repository.GetByUserAndGatewayIdResponse, error) {
	gateway, err := GetActive(ctx)
	if err != nil {
		return nil, err
	}

	gatewayCustomerRepository := gateway_customer_repository.New()

	customer := gatewayCustomerRepository.FindOneByGatewayAndUserId(ctx, gateway_customer_repository.FindOneByGatewayAndUserIdParams{
		UserID:    dto.userID,
		GatewayID: gateway.ID,
	})

	if customer == nil {
		return nil, errors.New("gateway customer not found")
	}

	user := user_repository.New().GetOneById(ctx, user_repository.GetOneByIdParams{
		ID:   dto.userID,
		Role: roles_enum.USER,
	})

	gatewayCardID := ""

	switch gateway.GatewayProviderID {
	case gateway_provider_enum.STRIPE:
		newCard, err := stripe_provider.CreateCard(&stripe_provider.CreateCardDTO{
			CustomerID: utils.Decrypt(customer.GatewayCustomerID),
			FullName:   fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			Number:     dto.Number,
			ExpMonth:   dto.ExpMonth,
			CVC:        dto.CVC,
			ExpYear:    dto.ExpYear,
		})
		if err != nil {
			return nil, err
		}

		gatewayCardID = newCard.ID
	default:
		return nil, errors.New("Integration not available with the provider")
	}

	newCard := gateway_customer_card_repository.New().Create(ctx, gateway_customer_card_repository.CreateParams{
		GatewayID:     gateway.ID,
		UserID:        dto.userID,
		CardID:        dto.CardID,
		GatewayCardID: utils.Encrypt(gatewayCardID),
	})

	res := &gateway_customer_card_repository.GetByUserAndGatewayIdResponse{
		ID:            newCard.ID,
		GatewayID:     newCard.GatewayID,
		UserID:        newCard.UserID,
		CardID:        newCard.CardID,
		GatewayCardID: utils.Decrypt(newCard.GatewayCardID),
	}

	return res, nil
}
