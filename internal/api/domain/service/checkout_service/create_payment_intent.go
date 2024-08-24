package checkout_service

import (
	"context"
	"errors"
	"fmt"

	stripe_provider "github.com/wesleyfebarretos/ticket-sale/external/providers/gateways/stripe"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/gateway_provider_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/gateway_customer_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
)

type GatewayDTO struct {
	ID         int32
	ProviderID int32
}

type CreatePaymentIntentDTO struct {
	PaymentMethod string
	Amount        int64
	UserID        int32
	Gateway       GatewayDTO
}

func (this *CreatePaymentIntentDTO) validate() error {
	if this.UserID == 0 {
		return errors.New("customer id cannot be empty on create payment intent")
	}

	if this.Gateway.ID == 0 {
		return errors.New("gateway id cannot be empty on create payment intent")
	}

	if this.Gateway.ProviderID == 0 {
		return errors.New("gateway provider id cannot be empty on create payment intent")
	}

	if this.PaymentMethod == "" {
		return errors.New("payment method cannot be empty on create payment intent")
	}

	if this.Amount == 0 {
		return errors.New("amount cannot be empty on create payment intent")
	}

	return nil
}

func CreatePaymentIntent(ctx context.Context, dto *CreatePaymentIntentDTO) (bool, error) {
	if err := dto.validate(); err != nil {
		return false, err
	}

	customer := gateway_customer_repository.New().FindOneByGatewayAndUserId(ctx, gateway_customer_repository.FindOneByGatewayAndUserIdParams{
		UserID:    dto.UserID,
		GatewayID: dto.Gateway.ID,
	})

	if customer == nil {
		return false, fmt.Errorf("no customer found to user %d", dto.UserID)
	}

	switch dto.Gateway.ProviderID {
	case gateway_provider_enum.STRIPE:
		_, err := stripe_provider.CreatePaymentIntent(&stripe_provider.CreatePaymentIntentDTO{
			CustomerID:    utils.Decrypt(customer.GatewayCustomerID),
			PaymentMethod: dto.PaymentMethod,
			Amount:        dto.Amount,
		})
		if err != nil {
			return false, err
		}
	default:
		return false, errors.New("no active gateways")
	}

	return true, nil
}
