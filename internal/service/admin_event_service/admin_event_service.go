package admin_event_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/infra/db_util"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/shared/admin_product_shared"
	"github.com/wesleyfebarretos/ticket-sale/repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_events_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_product_stocks_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_products_repository"
)

type CreateResponse struct {
	Event        admin_events_repository.Event
	Product      admin_products_repository.Product
	ProductStock admin_product_stocks_repository.ProductStock
}

func Create(
	c *gin.Context,
	newEventReq admin_events_repository.CreateParams,
	newProductReq admin_products_repository.CreateParams,
	newStockReq admin_product_stocks_repository.CreateParams,
) CreateResponse {
	return db_util.WithTransaction(c, func(tx pgx.Tx) CreateResponse {
		newProduct, newStock := admin_product_shared.Create(c, tx, newProductReq, newStockReq)

		adminEventsRepository := repository.AdminEvents.WithTx(tx)

		newEventReq.ProductID = newProduct.ID

		newEvent, err := adminEventsRepository.Create(c, newEventReq)
		if err != nil {
			panic(exception.InternalServerException(err.Error()))
		}

		return CreateResponse{
			Event:        newEvent,
			Product:      newProduct,
			ProductStock: newStock,
		}
	})
}

func Update(c *gin.Context,
	updateEventReq admin_events_repository.UpdateParams,
	updateProductReq admin_products_repository.UpdateParams,
) {
	db_util.WithTransaction(c, func(tx pgx.Tx) struct{} {
		eventRepository := repository.AdminEvents.WithTx(tx)
		productId, err := eventRepository.Update(c, updateEventReq)
		if err != nil {
			panic(exception.InternalServerException(err.Error()))
		}

		updateProductReq.ID = productId

		adminProductRepository := repository.AdminProducts.WithTx(tx)
		err = adminProductRepository.Update(c, updateProductReq)
		if err != nil {
			panic(exception.InternalServerException(err.Error()))
		}

		return struct{}{}
	})
}

func SoftDelete(c *gin.Context, id int32) {
	db_util.WithTransaction(c, func(tx pgx.Tx) struct{} {
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

func GetOneById(c *gin.Context, id int32) admin_events_repository.EventsWithRelation {
	event, err := repository.AdminEvents.GetOneById(c, id)
	if err == pgx.ErrNoRows {
		panic(exception.NotFoundException(fmt.Sprintf("event of id %d not found", id)))
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}
	return event
}
