package checkout_service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/gateway_provider_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/payment_status_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_product_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/creditcard_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/payment_order_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/gateway_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
)

type OrderQueueProducerDTO struct {
	ProductUUID       *uuid.UUID
	CardUUID          *uuid.UUID
	InstallmentTimeID int32
	PaymentTypeID     int32
	Qty               int32
	UserID            int32
}

func OrderQueueProducer(ctx context.Context, dto OrderQueueProducerDTO) {
	gateway, err := gateway_service.GetActive(ctx)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	customer, err := gateway_service.FindOrCreateCustomer(ctx, dto.UserID)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	customer.GatewayCustomerID = utils.Decrypt(customer.GatewayCustomerID)

	paymentMethod := ""

	card := creditcard_repository.New().GetByUuid(ctx, *dto.CardUUID)

	if card == nil {
		panic(exception.NotFoundException("card not found"))
	}

	customerCard := gateway_service.GetCustomerCardID(ctx, card.ID, gateway.ID)
	paymentMethod = customerCard.GatewayCardID

	product := admin_product_repository.New().GetOneByUuid(ctx, *dto.ProductUUID)

	if product == nil {
		panic(exception.NotFoundException("product not found"))
	}

	//  TODO: implement installment payment
	amount := product.Price * float64(dto.Qty)

	paymentOrderRepository := payment_order_repository.New()

	order := paymentOrderRepository.Create(ctx, payment_order_repository.CreateParams{
		CreditcardUuid:    dto.CardUUID,
		UpdatedBy:         &dto.UserID,
		TotalPrice:        float64(amount),
		BaseValue:         float64(amount),
		UserID:            dto.UserID,
		PaymentTypeID:     dto.PaymentTypeID,
		InstallmentTimeID: dto.InstallmentTimeID,
		GatewayID:         gateway.ID,
		PaymentStatusID:   payment_status_enum.AWAITING_PAYMENT,
		CreatedBy:         dto.UserID,
	})

	//  TODO: send data of created payment intent to kafka queue to proceed checkout
	switch gateway.GatewayProviderID {
	case gateway_provider_enum.STRIPE:
		_, err := StripeCreatePaymentIntent(ctx, &StripeCreatePaymentIntentDTO{
			PaymentMethod: paymentMethod,
			CustomerID:    customer.GatewayCustomerID,
			Amount:        int64(amount),
		})
		if err != nil {
			now := time.Now().UTC()
			paymentOrderRepository.Cancel(ctx, payment_order_repository.CancelParams{
				UpdatedAt: now,
				CancelAt:  &now,
				UpdatedBy: &dto.UserID,
				Uuid:      order.Uuid,
			})
			panic(exception.InternalServerException(err.Error()))
		}
	default:
		panic(exception.InternalServerException("no active gateway"))
	}
}
