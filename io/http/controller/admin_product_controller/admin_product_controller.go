package admin_product_controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/service/admin_product_service"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_product_stocks_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_products_repository"
)

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

	newProduct, newProductStock := admin_product_service.Create(c, newProductRequest, newProductStockRequest)

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
	}

	c.JSON(http.StatusCreated, newProductResponse)
}

func Update(c *gin.Context) {
	body := UpdateRequestDto{}

	controller.ReadBody(c, &body)

	id := controller.GetId(c)

	adminUser := controller.GetClaims(c)

	admin_product_service.Update(c, admin_products_repository.UpdateParams{
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
	})

	c.JSON(http.StatusOK, true)
}

func SoftDelete(c *gin.Context) {
	id := controller.GetId(c)

	adminUser := controller.GetClaims(c)

	admin_product_service.SoftDelete(c, admin_products_repository.SoftDeleteParams{
		ID:        id,
		UpdatedBy: &adminUser.Id,
	})

	c.JSON(http.StatusOK, true)
}

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

func GetAllWithRelations(c *gin.Context) {
	products := admin_product_service.GetAllWithRelations(c)

	productsResponse := []GetAllWithRelationsResponseDto{}

	for _, product := range products {
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
		})
	}
	c.JSON(http.StatusOK, productsResponse)
}

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
