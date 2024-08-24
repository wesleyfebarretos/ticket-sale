package payment_order_repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/payment_order_connection"
)

type CreateParams struct {
	CreditcardUuid    *uuid.UUID `json:"creditcardUuid"`
	UpdatedBy         *int32     `json:"updatedBy"`
	TotalPrice        float64    `json:"totalPrice"`
	TotalPrice_2      float64    `json:"totalPrice2"`
	AddedValue        float64    `json:"addedValue"`
	BaseValue         float64    `json:"baseValue"`
	UserID            int32      `json:"userId"`
	PaymentTypeID     int32      `json:"paymentTypeId"`
	InstallmentTimeID int32      `json:"installmentTimeId"`
	GatewayID         int32      `json:"gatewayId"`
	PaymentStatusID   int32      `json:"paymentStatusId"`
	CreatedBy         int32      `json:"createdBy"`
}

type CreateResponse struct {
	UpdatedAt             time.Time  `json:"updatedAt"`
	CreatedAt             time.Time  `json:"createdAt"`
	ExtraInfo             *string    `json:"extraInfo"`
	CreditcardUuid        *uuid.UUID `json:"creditcardUuid"`
	UpdatedBy             *int32     `json:"updatedBy"`
	ExpirationAt          *time.Time `json:"expirationAt"`
	DueAt                 *time.Time `json:"dueAt"`
	CancelAt              *time.Time `json:"cancelAt"`
	PaymentAt             *time.Time `json:"paymentAt"`
	PaymentCancelReasonID *int32     `json:"paymentCancelReasonId"`
	AddedValue            float64    `json:"addedValue"`
	ReversedValue         float64    `json:"reversedValue"`
	TotalPrice            float64    `json:"totalPrice"`
	TotalValue            float64    `json:"totalValue"`
	CanceledValue         float64    `json:"canceledValue"`
	BaseValue             float64    `json:"baseValue"`
	PaymentTypeID         int32      `json:"paymentTypeId"`
	PaymentStatusID       int32      `json:"paymentStatusId"`
	ID                    int32      `json:"id"`
	InstallmentTimeID     int32      `json:"installmentTimeId"`
	CreatedBy             int32      `json:"createdBy"`
	GatewayID             int32      `json:"gatewayId"`
	UserID                int32      `json:"userId"`
	Uuid                  uuid.UUID  `json:"uuid"`
}

func (this *CreateParams) ToEntity() payment_order_connection.CreateParams {
	return payment_order_connection.CreateParams{
		CreditcardUuid:    this.CreditcardUuid,
		UserID:            this.UserID,
		TotalPrice:        this.TotalPrice,
		PaymentTypeID:     this.PaymentTypeID,
		InstallmentTimeID: this.InstallmentTimeID,
		GatewayID:         this.GatewayID,
		PaymentStatusID:   this.PaymentStatusID,
		TotalPrice_2:      this.TotalPrice_2,
		AddedValue:        this.AddedValue,
		BaseValue:         this.BaseValue,
		CreatedBy:         this.CreatedBy,
		UpdatedBy:         this.UpdatedBy,
	}
}

func (this *CreateResponse) FromEntity(p payment_order_connection.FinPaymentOrder) CreateResponse {
	return CreateResponse{
		UpdatedAt:             p.UpdatedAt,
		CreatedAt:             p.CreatedAt,
		ExtraInfo:             p.ExtraInfo,
		CreditcardUuid:        p.CreditcardUuid,
		UpdatedBy:             p.UpdatedBy,
		ExpirationAt:          p.ExpirationAt,
		DueAt:                 p.DueAt,
		CancelAt:              p.CancelAt,
		PaymentAt:             p.PaymentAt,
		PaymentCancelReasonID: p.PaymentCancelReasonID,
		AddedValue:            p.AddedValue,
		ReversedValue:         p.ReversedValue,
		TotalPrice:            p.TotalPrice,
		TotalValue:            p.TotalValue,
		CanceledValue:         p.CanceledValue,
		BaseValue:             p.BaseValue,
		PaymentTypeID:         p.PaymentTypeID,
		PaymentStatusID:       p.PaymentStatusID,
		ID:                    p.ID,
		InstallmentTimeID:     p.InstallmentTimeID,
		CreatedBy:             p.CreatedBy,
		GatewayID:             p.GatewayID,
		UserID:                p.UserID,
		Uuid:                  p.Uuid,
	}
}

func (this *PaymentOrderRepository) Create(ctx context.Context, p CreateParams) CreateResponse {
	newPaymentOrder, err := this.queries.Create(ctx, p.ToEntity())
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := CreateResponse{}

	return res.FromEntity(newPaymentOrder)
}
