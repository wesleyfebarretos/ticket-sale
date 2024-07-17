package admin_product_shared

import (
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
) (admin_products_repository.Product, admin_product_stocks_repository.ProductStock) {
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

	return newProduct, newProductStock
}
