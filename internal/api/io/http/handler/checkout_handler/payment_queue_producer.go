package checkout_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/checkout_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/handler"
)

type PaymentQueueProducerRequest struct {
	ProductUUID       *uuid.UUID `json:"productUuid" binding:"required" example:"edb13b5a-a82d-4af4-9309-47dbf9cc81db"`
	CardUUID          *uuid.UUID `json:"cardUuid" example:"edb13b5a-a82d-4af4-9309-47dbf9cc81db"`
	InstallmentTimeID int32      `json:"installmentTimeId" binding:"required" example:"1"`
	PaymentTypeID     int32      `json:"paymentTypeId" binding:"required" example:"1"`
	Qty               int32      `json:"qty" binding:"required" example:"7"`
}

func (this *PaymentQueueProducerRequest) ToDomain(userID int32) checkout_service.PaymentQueueProducerDTO {
	return checkout_service.PaymentQueueProducerDTO{
		ProductUUID:       this.ProductUUID,
		CardUUID:          this.CardUUID,
		InstallmentTimeID: this.InstallmentTimeID,
		PaymentTypeID:     this.PaymentTypeID,
		Qty:               this.Qty,
		UserID:            userID,
	}
}

func PaymentQueueProducer(c *gin.Context) {
	user := handler.GetClaims(c)

	body := &PaymentQueueProducerRequest{}

	handler.ReadBody(c, body)

	checkout_service.PaymentQueueProducer(c, body.ToDomain(user.Id))

	c.JSON(http.StatusOK, true)
}
