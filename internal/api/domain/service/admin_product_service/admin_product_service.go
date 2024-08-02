package admin_product_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_product_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_product_stock_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/shared/admin_product_shared"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
)

type CreateResponse struct {
	Product      admin_product_repository.CreateResponse
	Stock        admin_product_stock_repository.CreateResponse
	Installments []admin_product_repository.CreateInstallmentsResponse
}

type CreateParams struct {
	Product      admin_product_repository.CreateParams
	Stock        admin_product_stock_repository.CreateParams
	Installments []admin_product_repository.CreateInstallmentsParams
}

func Create(
	c *gin.Context,
	newProductRequest admin_product_repository.CreateParams,
	newProductStockRequest admin_product_stock_repository.CreateParams,
	newProductInstallmentsRequest []admin_product_repository.CreateInstallmentsParams,
) CreateResponse {
	return utils.WithTransaction(c, func(tx pgx.Tx) CreateResponse {
		newProduct, newProductStock, newProductInstallments := admin_product_shared.Create(
			c,
			tx,
			newProductRequest,
			newProductStockRequest,
			newProductInstallmentsRequest,
		)

		return CreateResponse{
			Product:      newProduct,
			Stock:        newProductStock,
			Installments: newProductInstallments,
		}
	})
}

type UpdateParams struct {
	Product      admin_product_repository.UpdateParams
	Installments []admin_product_repository.CreateInstallmentsParams
}

func Update(
	c *gin.Context,
	updateProductRequest admin_product_repository.UpdateParams,
	updateProductInstallmentsRequest []admin_product_repository.CreateInstallmentsParams,
) {
	utils.WithTransaction(c, func(tx pgx.Tx) struct{} {
		admin_product_shared.Update(c, tx, updateProductRequest, updateProductInstallmentsRequest)
		return struct{}{}
	})
}

func SoftDelete(c *gin.Context, params admin_product_repository.SoftDeleteParams) {
	repository := admin_product_repository.New()

	product := repository.GetOneById(c, params.ID)

	if product == nil {
		panic(exception.NotFoundException(fmt.Sprintf("product of id %d not found", params.ID)))
	}

	repository.SoftDelete(c, params)
}

func GetAll(c *gin.Context) []admin_product_repository.GetAllResponse {
	return admin_product_repository.New().GetAll(c)
}

func GetAllWithRelations(c *gin.Context) []admin_product_repository.GetAllWithDetailsResponse {
	return admin_product_repository.New().GetAllWithDetails(c)
}

func GetOneById(c *gin.Context, id int32) admin_product_repository.GetOneByIdResponse {
	repository := admin_product_repository.New()

	product := repository.GetOneById(c, id)

	if product == nil {
		panic(exception.NotFoundException(fmt.Sprintf("product of id %d not found", id)))
	}

	return *product
}

func GetOneByUuid(c *gin.Context, uuid uuid.UUID) admin_product_repository.GetOneByUuidResponse {
	repository := admin_product_repository.New()

	product := repository.GetOneByUuid(c, uuid)
	if product == nil {
		panic(exception.NotFoundException(fmt.Sprintf("product of uuid %s not found", uuid)))
	}

	return *product
}
