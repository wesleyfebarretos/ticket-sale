package admin_product_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/admin_product_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/handler"
)

// CreateProduct godoc
//
//	@Tags			Admin Product
//	@Summary		Create a product
//	@Description	Create a product
//	@Produce		json
//	@Param			Product	body		CreateRequestDto	true	"New product"
//	@Success		201		{object}	CreateResponseDto
//	@Failure		500		{object}	exception.HttpException
//	@Failure		400		{object}	exception.HttpException
//	@Failure		403		{object}	middleware.RolePermissionError
//	@Failure		401		{object}	middleware.AuthenticationError
//	@Router			/admin/products [post]
func Create(c *gin.Context) {
	body := CreateRequestDto{}

	handler.ReadBody(c, &body)

	adminUser := handler.GetClaims(c)

	domainObj := body.ToDomain(adminUser.Id)

	newProduct := admin_product_service.Create(c, domainObj.Product, domainObj.Stock, domainObj.Installments)

	res := CreateResponseDto{}

	c.JSON(http.StatusCreated, res.FromDomain(newProduct))
}

// UpdateProduct godoc
//
//	@Tags			Admin Product
//	@Summary		Update a product
//	@Description	Update a product
//	@Produce		json
//	@Param			id		path		int					true	"Product ID"
//	@Param			Product	body		UpdateRequestDto	true	"Update product"
//	@Success		200		{object}	bool
//	@Failure		500		{object}	exception.HttpException
//	@Failure		404		{object}	exception.HttpException
//	@Failure		400		{object}	exception.HttpException
//	@Failure		403		{object}	middleware.RolePermissionError
//	@Failure		401		{object}	middleware.AuthenticationError
//	@Router			/admin/products/{id} [put]
func Update(c *gin.Context) {
	body := UpdateRequestDto{}

	handler.ReadBody(c, &body)

	id := handler.GetId(c)

	adminUser := handler.GetClaims(c)

	domainObj := body.ToDomain(adminUser.Id, id)

	admin_product_service.Update(c, domainObj.Product, domainObj.Installments)

	c.JSON(http.StatusOK, true)
}

// SoftDeleteProduct godoc
//
//	@Tags			Admin Product
//	@Summary		Soft Delete a product
//	@Description	Soft Delete a product
//	@Produce		json
//	@Param			id	path		int	true	"Product ID"
//	@Success		200	{object}	bool
//	@Failure		500	{object}	exception.HttpException
//	@Failure		400	{object}	exception.HttpException
//	@Failure		403	{object}	middleware.RolePermissionError
//	@Failure		401	{object}	middleware.AuthenticationError
//	@Router			/admin/products/{id} [delete]
func SoftDelete(c *gin.Context) {
	id := handler.GetId(c)

	adminUser := handler.GetClaims(c)

	params := SoftDeleteRequestDto{}

	admin_product_service.SoftDelete(c, params.ToDomain(id, adminUser.Id))

	c.JSON(http.StatusOK, true)
}

// GetAllProducts godoc
//
//	@Tags			Admin Product
//	@Summary		Get all products
//	@Description	Get all products
//	@Produce		json
//	@Success		200	{object}	[]GetAllResponseDto
//	@Failure		500	{object}	exception.HttpException
//	@Failure		400	{object}	exception.HttpException
//	@Failure		403	{object}	middleware.RolePermissionError
//	@Failure		401	{object}	middleware.AuthenticationError
//	@Router			/admin/products [get]
func GetAll(c *gin.Context) {
	products := admin_product_service.GetAll(c)

	res := GetAllResponseDto{}

	c.JSON(http.StatusOK, res.FromDomain(products))
}

// GetAllProductsWithRelations godoc
//
//	@Tags			Admin Product
//	@Summary		Get all products with relations
//	@Description	Get all products with relations
//	@Produce		json
//	@Success		200	{object}	[]GetAllWithRelationsResponseDto
//	@Failure		500	{object}	exception.HttpException
//	@Failure		400	{object}	exception.HttpException
//	@Failure		403	{object}	middleware.RolePermissionError
//	@Failure		401	{object}	middleware.AuthenticationError
//	@Router			/admin/products/details [get]
func GetAllWithRelations(c *gin.Context) {
	products := admin_product_service.GetAllWithRelations(c)

	res := GetAllWithRelationsResponseDto{}

	c.JSON(http.StatusOK, res.FromDomain(products))
}

// GetOneById godoc
//
//	@Tags			Admin Product
//	@Summary		Get One By Id
//	@Description	Get One By Id
//	@Produce		json
//	@Param			id	path		int	true	"Product ID"
//	@Success		200	{object}	GetOneByIdResponseDto
//	@Failure		500	{object}	exception.HttpException
//	@Failure		404	{object}	exception.HttpException
//	@Failure		400	{object}	exception.HttpException
//	@Failure		403	{object}	middleware.RolePermissionError
//	@Failure		401	{object}	middleware.AuthenticationError
//	@Router			/admin/products/{id} [get]
func GetOneById(c *gin.Context) {
	id := handler.GetId(c)

	product := admin_product_service.GetOneById(c, id)

	res := GetOneByIdResponseDto{}

	c.JSON(http.StatusOK, res.FromDomain(product))
}

// GetOneByUuid godoc
//
//	@Tags			Admin Product
//	@Summary		Get One By UUID
//	@Description	Get One By UUID
//	@Produce		json
//	@Param			uuid	path		string	true	"Product UUID"
//	@Success		200		{object}	GetOneByUuidResponseDto
//	@Failure		500		{object}	exception.HttpException
//	@Failure		404		{object}	exception.HttpException
//	@Failure		400		{object}	exception.HttpException
//	@Failure		403		{object}	middleware.RolePermissionError
//	@Failure		401		{object}	middleware.AuthenticationError
//	@Router			/admin/products/uuid/{uuid} [get]
func GetOneByUuid(c *gin.Context) {
	uuid := handler.GetUuid(c)

	product := admin_product_service.GetOneByUuid(c, uuid)

	res := GetOneByUuidResponseDto{}

	c.JSON(http.StatusOK, res.FromDomain(product))
}
