package admin_product_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/infra/db"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_product_stocks_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_products_repository"
)

func Create(
	c *gin.Context,
	newProductRequest admin_products_repository.CreateParams,
	newProductStockRequest admin_product_stocks_repository.CreateParams,
) (admin_products_repository.Product, admin_product_stocks_repository.ProductStock) {
	uuid, err := uuid.NewV7()
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	newProductRequest.Uuid = uuid

	tx, err := db.Conn.Begin(c)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}
	defer tx.Rollback(c)

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

	if err := tx.Commit(c); err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return newProduct, newProductStock
}

func Update(c *gin.Context, updateProductRequest admin_products_repository.UpdateParams) {
	_, err := repository.AdminProducts.GetOneById(c, updateProductRequest.ID)

	if err == pgx.ErrNoRows {
		panic(exception.NotFoundException(fmt.Sprintf("product of id %d not found", updateProductRequest.ID)))
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	err = repository.AdminProducts.Update(c, updateProductRequest)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}
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

func GetAllWithRelations(c *gin.Context) []admin_products_repository.ProductsWithRelation {
	products, err := repository.AdminProducts.GetAllWithRelations(c)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return products
}

func GetOneById(c *gin.Context, id int32) admin_products_repository.ProductsWithRelation {
	product, err := repository.AdminProducts.GetOneById(c, id)
	if err == pgx.ErrNoRows {
		panic(exception.NotFoundException(fmt.Sprintf("product of id %d not found", id)))
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return product
}

func GetOneByUuid(c *gin.Context, uuid uuid.UUID) admin_products_repository.ProductsWithRelation {
	product, err := repository.AdminProducts.GetOneByUuid(c, uuid)
	if err == pgx.ErrNoRows {
		panic(exception.NotFoundException(fmt.Sprintf("product of uuid %s not found", uuid)))
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return product
}
