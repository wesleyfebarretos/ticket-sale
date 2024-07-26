package admin_product_controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/admin_product_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/controller"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_product_stocks_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_products_repository"
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

	controller.ReadBody(c, &body)

	adminUser := controller.GetClaims(c)

	newProductRequest := admin_products_repository.CreateParams{
		Name:           body.Name,
		Description:    body.Description,
		Price:          body.Price,
		DiscountPrice:  body.DiscountPrice,
		Active:         body.Active,
		Image:          body.Image,
		ImageMobile:    body.ImageMobile,
		ImageThumbnail: body.ImageThumbnail,
		CategoryID:     body.CategoryID,
		CreatedBy:      adminUser.Id,
	}

	newProductStockRequest := admin_product_stocks_repository.CreateParams{
		Qty:       body.Stock.Qty,
		MinQty:    body.Stock.MinQty,
		CreatedBy: adminUser.Id,
	}

	newProductInstallments := []admin_products_repository.CreateInstallmentsParams{}

	for _, installment := range body.Installments {
		newInstallment := admin_products_repository.CreateInstallmentsParams{
			PaymentTypeID:     installment.PaymentTypeID,
			InstallmentTimeID: installment.ID,
			Fee:               0,
			Tariff:            0,
			CreatedBy:         adminUser.Id,
		}

		if installment.Fee != nil {
			newInstallment.Fee = *installment.Fee
		}

		if installment.Tariff != nil {
			newInstallment.Tariff = *installment.Tariff
		}

		newProductInstallments = append(newProductInstallments, newInstallment)
	}

	res := admin_product_service.Create(c, newProductRequest, newProductStockRequest, newProductInstallments)

	newProduct := res.Product
	newProductStock := res.Stock

	newInstallments := []CreateInstallmentsResponseDto{}

	for _, i := range res.Installments {
		newInstallments = append(newInstallments, CreateInstallmentsResponseDto{
			ID:            i.ID,
			PaymentTypeID: i.PaymentTypeID,
			Fee:           i.Fee,
			Tariff:        i.Tariff,
			InstallmentID: i.InstallmentTimeID,
		})
	}

	newProductResponse := CreateResponseDto{
		ID:             newProduct.ID,
		Name:           newProduct.Name,
		Description:    newProduct.Description,
		Price:          newProduct.Price,
		DiscountPrice:  newProduct.DiscountPrice,
		Active:         newProduct.Active,
		Image:          newProduct.Image,
		ImageMobile:    newProduct.ImageMobile,
		ImageThumbnail: newProduct.ImageThumbnail,
		CategoryID:     newProduct.CategoryID,
		CreatedBy:      newProduct.CreatedBy,
		Uuid:           newProduct.Uuid,
		IsDeleted:      newProduct.IsDeleted,
		UpdatedBy:      newProduct.UpdatedBy,
		CreatedAt:      newProduct.CreatedAt,
		UpdatedAt:      newProduct.UpdatedAt,
		Stock: CreateStockResponseDto{
			ID:        newProductStock.ID,
			ProductID: newProductStock.ProductID,
			Qty:       newProductStock.Qty,
			MinQty:    newProductStock.MinQty,
		},
		Installments: newInstallments,
	}

	c.JSON(http.StatusCreated, newProductResponse)
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

	controller.ReadBody(c, &body)

	id := controller.GetId(c)

	adminUser := controller.GetClaims(c)

	updateProduct := admin_products_repository.UpdateParams{
		Name:           body.Name,
		Description:    body.Description,
		Price:          body.Price,
		DiscountPrice:  body.DiscountPrice,
		Active:         body.Active,
		Image:          body.Image,
		ImageMobile:    body.ImageMobile,
		ImageThumbnail: body.ImageThumbnail,
		CategoryID:     body.CategoryID,
		UpdatedBy:      &adminUser.Id,
		ID:             id,
	}

	updateInstallments := []admin_products_repository.CreateInstallmentsParams{}

	for _, installment := range body.Installments {
		updatedInstallment := admin_products_repository.CreateInstallmentsParams{
			PaymentTypeID:     installment.PaymentTypeID,
			InstallmentTimeID: installment.ID,
			Fee:               0,
			Tariff:            0,
			CreatedBy:         adminUser.Id,
			ProductID:         id,
		}
		if installment.Fee != nil {
			updatedInstallment.Fee = *installment.Fee
		}

		if installment.Tariff != nil {
			updatedInstallment.Tariff = *installment.Tariff
		}

		updateInstallments = append(updateInstallments, updatedInstallment)
	}

	admin_product_service.Update(c, updateProduct, updateInstallments)

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
	id := controller.GetId(c)

	adminUser := controller.GetClaims(c)

	admin_product_service.SoftDelete(c, admin_products_repository.SoftDeleteParams{
		ID:        id,
		UpdatedBy: &adminUser.Id,
	})

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

	productsResponse := []GetAllResponseDto{}

	for _, product := range products {
		productsResponse = append(productsResponse, GetAllResponseDto{
			ID:             product.ID,
			Name:           product.Name,
			Description:    product.Description,
			Uuid:           product.Uuid,
			Price:          product.Price,
			DiscountPrice:  product.DiscountPrice,
			Active:         product.Active,
			IsDeleted:      product.IsDeleted,
			Image:          product.Image,
			ImageMobile:    product.ImageMobile,
			ImageThumbnail: product.ImageThumbnail,
			CategoryID:     product.CategoryID,
			CreatedBy:      product.CreatedBy,
			UpdatedBy:      product.UpdatedBy,
			CreatedAt:      product.CreatedAt,
			UpdatedAt:      product.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, productsResponse)
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

	productsResponse := []GetAllWithRelationsResponseDto{}

	for _, product := range products {
		stock := &StockResponseDto{}
		category := &CategoryResponseDto{}
		installments := InstallmentsResponseDto{}

		bStock, err := json.Marshal(product.Stock)
		if err != nil {
			panic(exception.InternalServerException(err.Error()))
		}

		bCategory, err := json.Marshal(product.Category)
		if err != nil {
			panic(exception.InternalServerException(err.Error()))
		}

		bInstallments, err := json.Marshal(product.Installments)
		if err != nil {
			panic(exception.InternalServerException(err.Error()))
		}

		if err := json.Unmarshal(bStock, &stock); err != nil {
			panic(exception.InternalServerException(err.Error()))
		}

		if err := json.Unmarshal(bCategory, &category); err != nil {
			panic(exception.InternalServerException(err.Error()))
		}

		if err := json.Unmarshal(bInstallments, &installments); err != nil {
			panic(exception.InternalServerException(err.Error()))
		}

		productsResponse = append(productsResponse, GetAllWithRelationsResponseDto{
			ID:             product.ID,
			Name:           product.Name,
			Description:    product.Description,
			Uuid:           product.Uuid,
			Price:          product.Price,
			DiscountPrice:  product.DiscountPrice,
			Active:         product.Active,
			IsDeleted:      product.IsDeleted,
			Image:          product.Image,
			ImageMobile:    product.ImageMobile,
			ImageThumbnail: product.ImageThumbnail,
			CategoryID:     product.CategoryID,
			CreatedBy:      product.CreatedBy,
			UpdatedBy:      product.UpdatedBy,
			CreatedAt:      product.CreatedAt,
			UpdatedAt:      product.UpdatedAt,
			Stock:          stock,
			Category:       category,
			Installments:   installments,
		})
	}
	c.JSON(http.StatusOK, productsResponse)
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
	id := controller.GetId(c)

	product := admin_product_service.GetOneById(c, id)

	stock := &StockResponseDto{}
	category := &CategoryResponseDto{}

	bStock, err := json.Marshal(product.Stock)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	bCategory, err := json.Marshal(product.Category)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	if err := json.Unmarshal(bStock, &stock); err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	if err := json.Unmarshal(bCategory, &category); err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	productResponse := GetOneByIdResponseDto{
		ID:             product.ID,
		Name:           product.Name,
		Description:    product.Description,
		Uuid:           product.Uuid,
		Price:          product.Price,
		DiscountPrice:  product.DiscountPrice,
		Active:         product.Active,
		IsDeleted:      product.IsDeleted,
		Image:          product.Image,
		ImageMobile:    product.ImageMobile,
		ImageThumbnail: product.ImageThumbnail,
		CategoryID:     product.CategoryID,
		CreatedBy:      product.CreatedBy,
		UpdatedBy:      product.UpdatedBy,
		CreatedAt:      product.CreatedAt,
		UpdatedAt:      product.UpdatedAt,
		Stock:          stock,
		Category:       category,
	}

	c.JSON(http.StatusOK, productResponse)
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
	uuid := controller.GetUuid(c)

	product := admin_product_service.GetOneByUuid(c, uuid)

	stock := &StockResponseDto{}
	category := &CategoryResponseDto{}

	bStock, err := json.Marshal(product.Stock)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	bCategory, err := json.Marshal(product.Category)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	if err := json.Unmarshal(bStock, &stock); err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	if err := json.Unmarshal(bCategory, &category); err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	productResponse := GetOneByUuidResponseDto{
		ID:             product.ID,
		Name:           product.Name,
		Description:    product.Description,
		Uuid:           product.Uuid,
		Price:          product.Price,
		DiscountPrice:  product.DiscountPrice,
		Active:         product.Active,
		IsDeleted:      product.IsDeleted,
		Image:          product.Image,
		ImageMobile:    product.ImageMobile,
		ImageThumbnail: product.ImageThumbnail,
		CategoryID:     product.CategoryID,
		CreatedBy:      product.CreatedBy,
		UpdatedBy:      product.UpdatedBy,
		CreatedAt:      product.CreatedAt,
		UpdatedAt:      product.UpdatedAt,
		Stock:          stock,
		Category:       category,
	}

	c.JSON(http.StatusOK, productResponse)
}
