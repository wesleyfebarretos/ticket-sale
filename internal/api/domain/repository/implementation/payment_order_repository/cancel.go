package payment_order_repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/payment_order_connection"
)

type CancelParams struct {
	UpdatedAt time.Time  `json:"updatedAt"`
	CancelAt  *time.Time `json:"cancelAt"`
	UpdatedBy *int32     `json:"updatedBy"`
	Uuid      uuid.UUID  `json:"uuid"`
}

func (this *CancelParams) ToEntity() payment_order_connection.CancelParams {
	return payment_order_connection.CancelParams{
		Uuid:      this.Uuid,
		CancelAt:  this.CancelAt,
		UpdatedBy: this.UpdatedBy,
		UpdatedAt: this.UpdatedAt,
	}
}

func (this *PaymentOrderRepository) Cancel(ctx context.Context, p CancelParams) {
	err := this.queries.Cancel(ctx, p.ToEntity())
	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}
}
