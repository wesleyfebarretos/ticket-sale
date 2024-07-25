package admin_product_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/api/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/api/internal/shared/admin_product_shared"
	"github.com/wesleyfebarretos/ticket-sale/api/repository"
	"github.com/wesleyfebarretos/ticket-sale/api/repository/sqlc/admin_product_stocks_repository"
	"github.com/wesleyfebarretos/ticket-sale/api/repository/sqlc/admin_products_repository"
	"github.com/wesleyfebarretos/ticket-sale/api/utils"
)

type CreateResponse struct {
	Product      admin_products_repository.Product
	Stock        admin_product_stocks_repository.ProductStock
	Installments []admin_products_repository.FinProductPaymentTypeInstallmentTime
}

func Create(
	c *gin.Context,
	newProductRequest admin_products_repository.CreateParams,
	newProductStockRequest admin_product_stocks_repository.CreateParams,
	newProductInstallmentsRequest []admin_products_repository.CreateInstallmentsParams,
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

func Update(
	c *gin.Context,
	updateProductRequest admin_products_repository.UpdateParams,
	updateProductInstallmentsRequest []admin_products_repository.CreateInstallmentsParams,
) {
	utils.WithTransaction(c, func(tx pgx.Tx) struct{} {
		admin_product_shared.Update(c, tx, updateProductRequest, updateProductInstallmentsRequest)
		return struct{}{}
	})
}

func SoftDelete(c *gin.Context, params admin_products_repository.SoftDeleteParams) {
	_, err := repository.AdminProducts.GetOneById(c, params.ID)

	if err == pgx.ErrNoRows {
		panic(exception.NotFoundException(fmt.Sprintf("product of id %d not found", params.ID)))
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	err = repository.AdminProducts.SoftDelete(c, params)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}
}

func GetAll(c *gin.Context) []admin_products_repository.Product {
	products, err := repository.AdminProducts.GetAll(c)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return products
}

func GetAllWithRelations(c *gin.Context) []admin_products_repository.ProductsDetail {
	products, err := repository.AdminProducts.GetAllProductsDetails(c)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return products
}

func GetOneById(c *gin.Context, id int32) admin_products_repository.ProductsDetail {
	product, err := repository.AdminProducts.GetOneById(c, id)
	if err == pgx.ErrNoRows {
		panic(exception.NotFoundException(fmt.Sprintf("product of id %d not found", id)))
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return product
}

func GetOneByUuid(c *gin.Context, uuid uuid.UUID) admin_products_repository.ProductsDetail {
	product, err := repository.AdminProducts.GetOneByUuid(c, uuid)
	if err == pgx.ErrNoRows {
		panic(exception.NotFoundException(fmt.Sprintf("product of uuid %s not found", uuid)))
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return product
}
