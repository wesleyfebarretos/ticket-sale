package admin_product_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/infra/db_util"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/shared/admin_product_shared"
	"github.com/wesleyfebarretos/ticket-sale/repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_product_stocks_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_products_repository"
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
	return db_util.WithTransaction(c, func(tx pgx.Tx) CreateResponse {
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
	db_util.WithTransaction(c, func(tx pgx.Tx) struct{} {
		adminProductRepository := repository.AdminProducts.WithTx(tx)

		_, err := adminProductRepository.GetOneById(c, updateProductRequest.ID)

		if err == pgx.ErrNoRows {
			panic(exception.NotFoundException(fmt.Sprintf("product of id %d not found", updateProductRequest.ID)))
		}

		if err != nil {
			panic(exception.InternalServerException(err.Error()))
		}

		err = adminProductRepository.Update(c, updateProductRequest)
		if err != nil {
			panic(exception.InternalServerException(err.Error()))
		}

		persistedInstallments, err := adminProductRepository.GetAllProductInstallmentTimes(
			c,
			updateProductRequest.ID,
		)

		if err != nil && err != pgx.ErrNoRows {
			panic(exception.InternalServerException(err.Error()))
		}

		if someInstallmentChanged(updateProductInstallmentsRequest, persistedInstallments) {
			adminProductRepository.DeleteAllProductInstallmentTimes(c, updateProductRequest.ID)

			installmentsBatchQuery := adminProductRepository.CreateInstallments(c, updateProductInstallmentsRequest)

			installmentsBatchQuery.QueryRow(
				func(
					index int,
					_ admin_products_repository.FinProductPaymentTypeInstallmentTime,
					err error,
				) {
					if err != nil {
						panic(exception.InternalServerException(fmt.Sprintf("query of index %d failed: %v", index, err)))
					}
				},
			)
		}

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

func someInstallmentChanged(
	newInstallments []admin_products_repository.CreateInstallmentsParams,
	oldInstallments []admin_products_repository.GetAllProductInstallmentTimesRow,
) bool {
	if len(newInstallments) != len(oldInstallments) {
		return true
	}

	for _, newInstallment := range newInstallments {
		changed := true
		for _, oldInstallment := range oldInstallments {
			if newInstallment.PaymentTypeID == oldInstallment.PaymentTypeID &&
				newInstallment.Fee == oldInstallment.Fee &&
				newInstallment.Tariff == oldInstallment.Tariff {
				changed = false
				break
			}
		}
		if changed {
			return true
		}
	}

	return false
}
