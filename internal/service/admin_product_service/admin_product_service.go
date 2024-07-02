package admin_product_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_products_repository"
)

func Create(c *gin.Context, newProductRequest admin_products_repository.CreateParams) admin_products_repository.Product {
	uuid, err := uuid.NewV7()
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	newProductRequest.Uuid = uuid

	newProduct, err := repository.AdminProducts.Create(c, newProductRequest)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return newProduct
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

func GetAllWithRelations(c *gin.Context) []admin_products_repository.GetAllWithRelationsRow {
	products, err := repository.AdminProducts.GetAllWithRelations(c)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return products
}

func GetOneById(c *gin.Context, id int32) admin_products_repository.GetOneByIdRow {
	product, err := repository.AdminProducts.GetOneById(c, id)
	if err == pgx.ErrNoRows {
		panic(exception.NotFoundException(fmt.Sprintf("product of id %d not found", id)))
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return product
}

func GetOneByUuid(c *gin.Context, uuid uuid.UUID) admin_products_repository.GetOneByUuidRow {
	product, err := repository.AdminProducts.GetOneByUuid(c, uuid)
	if err == pgx.ErrNoRows {
		panic(exception.NotFoundException(fmt.Sprintf("product of uuid %d not found", uuid)))
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return product
}
