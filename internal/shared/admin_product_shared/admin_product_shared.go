package admin_product_shared

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_product_stocks_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_products_repository"
)

func Create(
	c *gin.Context,
	tx pgx.Tx,
	newProductRequest admin_products_repository.CreateParams,
	newProductStockRequest admin_product_stocks_repository.CreateParams,
	newProductInstallmentsRequest []admin_products_repository.CreateInstallmentsParams,
) (
	admin_products_repository.Product,
	admin_product_stocks_repository.ProductStock,
	[]admin_products_repository.FinProductPaymentTypeInstallmentTime,
) {
	adminProductsRepository := repository.AdminProducts.WithTx(tx)
	adminProductStocksRepository := repository.AdminProductStocks.WithTx(tx)

	newProduct, err := adminProductsRepository.Create(c, newProductRequest)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	newProductStockRequest.ProductID = newProduct.ID

	newProductStock, err := adminProductStocksRepository.Create(c, newProductStockRequest)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	for i := range newProductInstallmentsRequest {
		newProductInstallmentsRequest[i].ProductID = newProduct.ID
	}

	installmentsBatchQuery := adminProductsRepository.CreateInstallments(c, newProductInstallmentsRequest)

	newProductInstallments := []admin_products_repository.FinProductPaymentTypeInstallmentTime{}

	installmentsBatchQuery.QueryRow(
		func(
			index int,
			installment admin_products_repository.FinProductPaymentTypeInstallmentTime,
			err error,
		) {
			if err != nil {
				panic(exception.InternalServerException(fmt.Sprintf("query of index %d failed: %v", index, err)))
			}

			newProductInstallments = append(newProductInstallments, installment)
		},
	)

	return newProduct, newProductStock, newProductInstallments
}
