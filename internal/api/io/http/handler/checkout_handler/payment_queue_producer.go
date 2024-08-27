package checkout_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type PaymentQueueProducerRequest struct {
	ProductUUID       *uuid.UUID
	CardUUID          *uuid.UUID
	CardID            *int32
	InstallmentTimeID int32
	PaymentTypeID     int32
	Qty               int32
	UserID            int32
}

func PaymentQueueProducer(ctx *gin.Context) {
}
