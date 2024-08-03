package admin_event_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_event_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_product_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_product_stock_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/shared/admin_product_shared"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
)

type CreateResponse struct {
	Event               admin_event_repository.CreateResponse
	Product             admin_product_repository.CreateResponse
	ProductStock        admin_product_stock_repository.CreateResponse
	ProductInstallments []admin_product_repository.CreateInstallmentsResponse
}

type CreateParams struct {
	NewEvent               admin_event_repository.CreateParams
	NewProduct             admin_product_repository.CreateParams
	NewStock               admin_product_stock_repository.CreateParams
	NewProductInstallments []admin_product_repository.CreateInstallmentsParams
}

type UpdateParams struct {
	UpdateEvent               admin_event_repository.UpdateParams
	UpdateProduct             admin_product_repository.UpdateParams
	UpdateProductInstallments []admin_product_repository.CreateInstallmentsParams
}

func Create(
	c *gin.Context,
	newEventReq admin_event_repository.CreateParams,
	newProductReq admin_product_repository.CreateParams,
	newStockReq admin_product_stock_repository.CreateParams,
	newProductInstallments []admin_product_repository.CreateInstallmentsParams,
) CreateResponse {
	return utils.WithTransaction(c, func(tx pgx.Tx) CreateResponse {
		newProduct, newStock, newInstallments := admin_product_shared.Create(c, tx, newProductReq, newStockReq, newProductInstallments)

		adminEventsRepository := admin_event_repository.New().WithTx(tx)

		newEventReq.ProductID = newProduct.ID

		newEvent := adminEventsRepository.Create(c, newEventReq)

		return CreateResponse{
			Event:               newEvent,
			Product:             newProduct,
			ProductStock:        newStock,
			ProductInstallments: newInstallments,
		}
	})
}

func Update(c *gin.Context,
	updateEventReq admin_event_repository.UpdateParams,
	updateProductReq admin_product_repository.UpdateParams,
	updateProductInstallmentsReq []admin_product_repository.CreateInstallmentsParams,
) {
	utils.WithTransaction(c, func(tx pgx.Tx) struct{} {
		adminEventsRepository := admin_event_repository.New().WithTx(tx)
		productId := adminEventsRepository.Update(c, updateEventReq)

		updateProductReq.ID = productId

		for i := range updateProductInstallmentsReq {
			updateProductInstallmentsReq[i].ProductID = productId
		}

		admin_product_shared.Update(c, tx, updateProductReq, updateProductInstallmentsReq)

		return struct{}{}
	})
}

func SoftDelete(c *gin.Context, id int32) {
	adminEventsRepository := admin_event_repository.New()
	event := adminEventsRepository.GetOneById(c, id)
	if event == nil {
		panic(exception.NotFoundException(fmt.Sprintf("event of id %d not found", id)))
	}

	adminEventsRepository.SoftDelete(c, id)
}

func GetAll(c *gin.Context) []admin_event_repository.GetAllResponse {
	return admin_event_repository.New().GetAll(c)

}

func GetOneById(c *gin.Context, id int32) *admin_event_repository.GetOneByIdResponse {
	event := admin_event_repository.New().GetOneById(c, id)
	if event == nil {
		panic(exception.NotFoundException(fmt.Sprintf("event of id %d not found", id)))
	}

	return event
}
