package payment_order_repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/payment_order_connection"
)

type GetOneByUuidResponse struct {
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

func (this *GetOneByUuidResponse) FromDomain(p payment_order_connection.FinPaymentOrder) *GetOneByUuidResponse {
	return &GetOneByUuidResponse{
		ID:                    p.ID,
		Uuid:                  p.Uuid,
		CreditcardUuid:        p.CreditcardUuid,
		UserID:                p.UserID,
		TotalPrice:            p.TotalPrice,
		PaymentTypeID:         p.PaymentTypeID,
		InstallmentTimeID:     p.InstallmentTimeID,
		GatewayID:             p.GatewayID,
		PaymentStatusID:       p.PaymentStatusID,
		PaymentCancelReasonID: p.PaymentCancelReasonID,
		ExtraInfo:             p.ExtraInfo,
		PaymentAt:             p.PaymentAt,
		CancelAt:              p.CancelAt,
		DueAt:                 p.DueAt,
		ExpirationAt:          p.ExpirationAt,
		BaseValue:             p.BaseValue,
		ReversedValue:         p.ReversedValue,
		CanceledValue:         p.CanceledValue,
		AddedValue:            p.AddedValue,
		TotalValue:            p.TotalValue,
		CreatedBy:             p.CreatedBy,
		UpdatedBy:             p.UpdatedBy,
		CreatedAt:             p.CreatedAt,
		UpdatedAt:             p.UpdatedAt,
	}
}

func (this *PaymentOrderRepository) GetOneByUuid(ctx context.Context, uuid uuid.UUID) *GetOneByUuidResponse {
	order, err := this.queries.GetOneByUuid(ctx, uuid)
	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetOneByUuidResponse{}

	return res.FromDomain(order)
}
