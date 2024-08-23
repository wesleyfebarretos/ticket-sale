package stripe_provider

import (
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/paymentsource"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/config"
)

type CreateCardDTO struct {
	CustomerID string
	FullName   string
	Number     string
	ExpMonth   string
	CVC        string
	ExpYear    string
}

func (this *CreateCardDTO) validate() error {
	return nil
}

func CreateCard(c *CreateCardDTO) (*stripe.PaymentSource, error) {
	if config.Envs.AppEnv == "testing" {
		return &stripe.PaymentSource{
			ID: "card_1NGTaT2eZvKYlo2CZWSctn5n",
		}, nil
	}

	stripe.Key = config.Envs.Providers.Gateways.Stripe.Key

	params := &stripe.PaymentSourceParams{
		Customer: &c.CustomerID,
		Source: &stripe.PaymentSourceSourceParams{
			Card: &stripe.CardParams{
				Number:   &c.Number,
				ExpMonth: &c.ExpMonth,
				CVC:      &c.CVC,
				ExpYear:  &c.ExpYear,
				Name:     &c.FullName,
			},
		},
	}

	newCard, err := paymentsource.New(params)
	if err != nil {
		return nil, err
	}

	return newCard, nil
}
