package stripe_provider

import (
	"errors"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/paymentintent"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/config"
)

type CreatePaymentIntentDTO struct {
	CustomerID    string
	PaymentMethod string
	Amount        int64
}

func (this *CreatePaymentIntentDTO) validate() error {
	if this.CustomerID == "" {
		return errors.New("Customer id cannot be empty on create payment intent")
	}

	if this.PaymentMethod == "" {
		return errors.New("Payment method cannot be empty on create payment intent")
	}

	if this.Amount == 0 {
		return errors.New("Amount cannot be empty on create payment intent")
	}

	return nil
}

func CreatePaymentIntent(dto *CreatePaymentIntentDTO) (*stripe.PaymentIntent, error) {
	stripe.Key = config.Envs.Providers.Gateways.Stripe.Key

	if err := dto.validate(); err != nil {
		return nil, err
	}

	if config.Envs.AppEnv == "testing" {
		return &stripe.PaymentIntent{
			ID: "pi_3MtwBwLkdIwHu7ix28a3tqPa",
		}, nil
	}

	params := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(dto.Amount),
		Customer:      &dto.CustomerID,
		PaymentMethod: &dto.PaymentMethod,
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}

	result, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	return result, nil
}
