package stripe_provider

import (
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/customer"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/config"
)

func CreateCustomer() (*stripe.Customer, error) {
	stripe.Key = config.Envs.Providers.Gateways.Stripe.Key

	params := &stripe.CustomerParams{
		Name:  stripe.String("Jenny Rosen"),
		Email: stripe.String("jennyrosen@example.com"),
	}
	newCustomer, err := customer.New(params)
	if err != nil {
		return nil, err
	}

	return newCustomer, nil
}
