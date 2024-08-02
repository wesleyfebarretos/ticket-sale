package admin_product_shared

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_product_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_product_stocks_repository"
)

func Create(
	c *gin.Context,
	tx pgx.Tx,
	newProductRequest admin_product_repository.CreateParams,
	newProductStockRequest admin_product_stocks_repository.CreateParams,
	newProductInstallmentsRequest []admin_product_repository.CreateInstallmentsParams,
) (
	admin_product_repository.CreateResponse,
	admin_product_stocks_repository.ProductStock,
	[]admin_product_repository.CreateInstallmentsResponse,
) {
	adminProductRepository := admin_product_repository.New().WithTx(tx)
	adminProductStocksRepository := repository.AdminProductStocks.WithTx(tx)

	newProduct := adminProductRepository.Create(c, newProductRequest)

	newProductStockRequest.ProductID = newProduct.ID

	newProductStock, err := adminProductStocksRepository.Create(c, newProductStockRequest)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	for i := range newProductInstallmentsRequest {
		newProductInstallmentsRequest[i].ProductID = newProduct.ID
	}

	newProductInstallments := adminProductRepository.CreateInstallments(c, newProductInstallmentsRequest)

	return newProduct, newProductStock, newProductInstallments
}

func Update(
	c *gin.Context,
	tx pgx.Tx,
	updateProductRequest admin_product_repository.UpdateParams,
	updateProductInstallmentsRequest []admin_product_repository.CreateInstallmentsParams,
) {
	adminProductRepository := admin_product_repository.New().WithTx(tx)

	product := adminProductRepository.GetOneById(c, updateProductRequest.ID)

	if product == nil {
		panic(exception.NotFoundException(fmt.Sprintf("product of id %d not found", updateProductRequest.ID)))
	}

	adminProductRepository.Update(c, updateProductRequest)

	persistedInstallments := adminProductRepository.GetAllInstallmentTimes(
		c,
		updateProductRequest.ID,
	)

	if someInstallmentChanged(updateProductInstallmentsRequest, persistedInstallments) {
		adminProductRepository.DeleteAllInstallmentTimes(c, updateProductRequest.ID)
		adminProductRepository.CreateInstallments(c, updateProductInstallmentsRequest)
	}
}

func someInstallmentChanged(
	newInstallments []admin_product_repository.CreateInstallmentsParams,
	oldInstallments []admin_product_repository.GetAllInstallmentTimeResponse,
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
