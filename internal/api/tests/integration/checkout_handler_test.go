package integration_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/payment_status_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/payment_type_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/payment_order_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/handler/checkout_handler"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/tests/test_data"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/tests/test_utils"
)

func TestCheckoutHandler(t *testing.T) {
	t.Run("it should create a payment request", TRun(func(t *testing.T) {
		user := test_data.NewUser(roles_enum.USER)

		TSetCookieWithUser(t, user)

		product := test_data.NewProduct(t, user.ID)
		card := test_data.NewCreditCard(t, user.ID)
		test_data.NewGatewayCustomer(t, 1, user.ID)
		test_data.NewGatewayCustomerCard(t, 1, user.ID, card.ID)

		request := checkout_handler.PaymentQueueProducerRequest{
			ProductUUID:       &product.Uuid,
			CardUUID:          &card.Uuid,
			InstallmentTimeID: 1,
			PaymentTypeID:     payment_type_enum.CREDITCARD,
			Qty:               5,
		}

		res := TMakeRequest(t, http.MethodPost, "checkout/payment", request)

		expect := false

		test_utils.Decode(t, res.Body, &expect)

		assert.True(t, expect)
		assert.Equal(t, http.StatusOK, res.StatusCode)

		orders := payment_order_repository.New().GetAll(context.Background())

		assert.Len(t, orders, 1)
		assert.Equal(t, int32(payment_status_enum.AWAITING_PAYMENT), orders[0].PaymentStatusID)
	}))
}
