package payment_order_repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/payment_order_connection"
)

type GetAllResponse struct {
	ID                    int32      `json:"id"`
	Uuid                  uuid.UUID  `json:"uuid"`
	CreditcardUuid        *uuid.UUID `json:"creditcardUuid"`
	UserID                int32      `json:"userId"`
	TotalPrice            float64    `json:"totalPrice"`
	PaymentTypeID         int32      `json:"paymentTypeId"`
	InstallmentTimeID     int32      `json:"installmentTimeId"`
	GatewayID             int32      `json:"gatewayId"`
	PaymentStatusID       int32      `json:"paymentStatusId"`
	PaymentCancelReasonID *int32     `json:"paymentCancelReasonId"`
	ExtraInfo             *string    `json:"extraInfo"`
	PaymentAt             *time.Time `json:"paymentAt"`
	CancelAt              *time.Time `json:"cancelAt"`
	DueAt                 *time.Time `json:"dueAt"`
	ExpirationAt          *time.Time `json:"expirationAt"`
	BaseValue             float64    `json:"baseValue"`
	ReversedValue         float64    `json:"reversedValue"`
	CanceledValue         float64    `json:"canceledValue"`
	AddedValue            float64    `json:"addedValue"`
	TotalValue            float64    `json:"totalValue"`
	CreatedBy             int32      `json:"createdBy"`
	UpdatedBy             *int32     `json:"updatedBy"`
	CreatedAt             time.Time  `json:"createdAt"`
	UpdatedAt             time.Time  `json:"updatedAt"`
}

func (this *GetAllResponse) FromDomain(p []payment_order_connection.FinPaymentOrder) []GetAllResponse {
	res := []GetAllResponse{}

	for _, v := range p {
		res = append(res, GetAllResponse{
			ID:                    v.ID,
			Uuid:                  v.Uuid,
			CreditcardUuid:        v.CreditcardUuid,
			UserID:                v.UserID,
			TotalPrice:            v.TotalPrice,
			PaymentTypeID:         v.PaymentTypeID,
			InstallmentTimeID:     v.InstallmentTimeID,
			GatewayID:             v.GatewayID,
			PaymentStatusID:       v.PaymentStatusID,
			PaymentCancelReasonID: v.PaymentCancelReasonID,
			ExtraInfo:             v.ExtraInfo,
			PaymentAt:             v.PaymentAt,
			CancelAt:              v.CancelAt,
			DueAt:                 v.DueAt,
			ExpirationAt:          v.ExpirationAt,
			BaseValue:             v.BaseValue,
			ReversedValue:         v.ReversedValue,
			CanceledValue:         v.CanceledValue,
			AddedValue:            v.AddedValue,
			TotalValue:            v.TotalValue,
			CreatedBy:             v.CreatedBy,
			UpdatedBy:             v.UpdatedBy,
			CreatedAt:             v.CreatedAt,
			UpdatedAt:             v.UpdatedAt,
		})
	}

	return res
}

func (this *PaymentOrderRepository) GetAll(ctx context.Context) []GetAllResponse {
	orders, err := this.queries.GetAll(ctx)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetAllResponse{}

	return res.FromDomain(orders)
}
