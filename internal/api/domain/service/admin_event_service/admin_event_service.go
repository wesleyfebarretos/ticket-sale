package admin_event_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_product_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_product_stock_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_events_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/shared/admin_product_shared"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
)

type CreateResponse struct {
	Event               admin_events_repository.Event
	Product             admin_product_repository.CreateResponse
	ProductStock        admin_product_stock_repository.CreateResponse
	ProductInstallments []admin_product_repository.CreateInstallmentsResponse
}

func Create(
	c *gin.Context,
	newEventReq admin_events_repository.CreateParams,
	newProductReq admin_product_repository.CreateParams,
	newStockReq admin_product_stock_repository.CreateParams,
	newProductInstallments []admin_product_repository.CreateInstallmentsParams,
) CreateResponse {
	return utils.WithTransaction(c, func(tx pgx.Tx) CreateResponse {
		newProduct, newStock, newInstallments := admin_product_shared.Create(c, tx, newProductReq, newStockReq, newProductInstallments)

		adminEventsRepository := repository.AdminEvents.WithTx(tx)

		newEventReq.ProductID = newProduct.ID

		newEvent, err := adminEventsRepository.Create(c, newEventReq)
		if err != nil {
			panic(exception.InternalServerException(err.Error()))
		}

		return CreateResponse{
			Event:               newEvent,
			Product:             newProduct,
			ProductStock:        newStock,
			ProductInstallments: newInstallments,
		}
	})
}

func Update(c *gin.Context,
	updateEventReq admin_events_repository.UpdateParams,
	updateProductReq admin_product_repository.UpdateParams,
	updateProductInstallmentsReq []admin_product_repository.CreateInstallmentsParams,
) {
	utils.WithTransaction(c, func(tx pgx.Tx) struct{} {
		eventRepository := repository.AdminEvents.WithTx(tx)
		productId, err := eventRepository.Update(c, updateEventReq)
		if err != nil {
			panic(exception.InternalServerException(err.Error()))
		}

		updateProductReq.ID = productId

		for i := range updateProductInstallmentsReq {
			updateProductInstallmentsReq[i].ProductID = productId
		}

		admin_product_shared.Update(c, tx, updateProductReq, updateProductInstallmentsReq)

		return struct{}{}
	})
}

func SoftDelete(c *gin.Context, id int32) {
	utils.WithTransaction(c, func(tx pgx.Tx) struct{} {
		adminEventsRepository := repository.AdminEvents.WithTx(tx)
		_, err := adminEventsRepository.GetOneById(c, id)
		if err == pgx.ErrNoRows {
			panic(exception.NotFoundException(fmt.Sprintf("event of id %d not found", id)))
		}

		if err != nil {
			panic(exception.InternalServerException(err.Error()))
		}

		adminEventsRepository.SoftDelete(c, id)

		return struct{}{}
	})
}

func GetAll(c *gin.Context) []admin_events_repository.EventsGetAll {
	events, err := repository.AdminEvents.GetAll(c)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return events
}

func GetOneById(c *gin.Context, id int32) admin_events_repository.EventsDetail {
	event, err := repository.AdminEvents.GetOneById(c, id)
	if err == pgx.ErrNoRows {
		panic(exception.NotFoundException(fmt.Sprintf("event of id %d not found", id)))
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}
	return event
}
