package stripe_provider

import (
	"errors"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/customer"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/config"
)

type CreateCustomerDTO struct {
	Name  string
	Email string
}

func (this *CreateCustomerDTO) validate() error {
	if this.Name == "" || this.Email == "" {
		return errors.New("missing fields in stripe create customer struct")
	}
	return nil
}

func CreateCustomer(c *CreateCustomerDTO) (*stripe.Customer, error) {
	err := c.validate()
	if err != nil {
		return nil, err
	}

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
