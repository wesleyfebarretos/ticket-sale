package checkout_service

import (
	"context"
	"errors"

	"github.com/stripe/stripe-go/v79"
	stripe_provider "github.com/wesleyfebarretos/ticket-sale/external/providers/gateways/stripe"
)

type StripeCreatePaymentIntentDTO struct {
	PaymentMethod string
	CustomerID    string
	Amount        int64
}

func (this *StripeCreatePaymentIntentDTO) validate() error {
	if this.PaymentMethod == "" {
		return errors.New("payment method cannot be empty on create payment intent")
	}

	if this.Amount == 0 {
		return errors.New("amount cannot be empty on create payment intent")
	}

	return nil
}

func StripeCreatePaymentIntent(ctx context.Context, dto *StripeCreatePaymentIntentDTO) (*stripe.PaymentIntent, error) {
	if err := dto.validate(); err != nil {
		return nil, err
	}

	newPaymentIntent, err := stripe_provider.CreatePaymentIntent(&stripe_provider.CreatePaymentIntentDTO{
		CustomerID:    dto.CustomerID,
		PaymentMethod: dto.PaymentMethod,
		Amount:        dto.Amount,
	})
	if err != nil {
		return nil, err
	}
	return newPaymentIntent, nil
}
