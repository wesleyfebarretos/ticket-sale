package admin_gateway_provider_repository

import (
	"context"
	"time"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_gateway_provider_connection"
)

type CreateParams struct {
	Name      string `json:"name"`
	CreatedBy int32  `json:"createdBy"`
	UpdatedBy int32  `json:"updatedBy"`
}

type CreateResponse struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedBy int32     `json:"createdBy"`
	UpdatedBy int32     `json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (this *CreateParams) ToEntity() admin_gateway_provider_connection.CreateParams {
	return admin_gateway_provider_connection.CreateParams{
		Name:      this.Name,
		CreatedBy: this.CreatedBy,
		UpdatedBy: this.UpdatedBy,
	}
}

func (this *CreateResponse) FromEntity(p admin_gateway_provider_connection.FinGatewayProvider) CreateResponse {
	return CreateResponse{
		ID:        p.ID,
		Name:      p.Name,
		CreatedBy: p.CreatedBy,
		UpdatedBy: p.UpdatedBy,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (this *AdminGatewayProviderRepository) Create(ctx context.Context, p CreateParams) CreateResponse {
	provider, err := this.queries.Create(ctx, p.ToEntity())
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := CreateResponse{}

	return res.FromEntity(provider)
}

// stripe.Key = "sk_test_Ho24N7La5CVDtbmpjc377lJI"
// params := &stripe.CustomerParams{
// 	Name:  stripe.String("Jenny Rosen"),
// 	Email: stripe.String("jennyrosen@example.com"),
// }
// customer, err := customer.New(params)
// if err != nil {
// 	panic(exception.InternalServerException(err.Error()))
// }
//
// cardParams := &stripe.PaymentSourceParams{
// 	Customer: stripe.String(customer.ID),
// 	Source: &stripe.PaymentSourceSourceParams{
// 		Card: &stripe.CardParams{
// 			CVC:      stripe.String("777"),
// 			ExpMonth: stripe.String("09"),
// 			ExpYear:  stripe.String("2024"),
// 			Number:   stripe.String("4242424242424242"),
// 		},
// 	},
// }
// resultCard, err := paymentsource.New(cardParams)
// if err != nil {
// 	panic(exception.InternalServerException(err.Error()))
// }
//
// paymentIntentParams := &stripe.PaymentIntentParams{
// 	Amount:   stripe.Int64(2000),
// 	Currency: stripe.String(string(stripe.CurrencyUSD)),
// 	AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
// 		Enabled: stripe.Bool(true),
// 	},
// 	PaymentMethod: &resultCard.ID,
// 	Customer:      &customer.ID,
// }
// resultPaymentIntent, err := paymentintent.New(paymentIntentParams)
// if err != nil {
// 	panic(exception.InternalServerException(err.Error()))
// }
//
// confirmPaymentParams := &stripe.PaymentIntentConfirmParams{
// 	PaymentMethod: stripe.String(resultCard.ID),
// 	ReturnURL:     stripe.String("https://www.example.com"),
// }
//
// resultConfirmPayment, err := paymentintent.Confirm(resultPaymentIntent.ID, confirmPaymentParams)
// if err != nil {
// 	panic(exception.InternalServerException(err.Error()))
// }
// prettyPrint(customer)
// prettyPrint(resultCard)
// prettyPrint(resultPaymentIntent)
// prettyPrint(resultConfirmPayment)
//
// os.Exit(1)
