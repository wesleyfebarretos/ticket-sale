package test_data

import (
	"context"
	"testing"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/product_categories_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_products_repository"
)

func NewProduct(t *testing.T, userID int32) admin_products_repository.Product {
	description := "Fresh and fiery red hot chilly peppers, perfect for adding a spicy kick to your dishes."
	DiscountPrice := 4.99
	image := "https://example.com/images/red-hot-chilly-peppers.jpg"
	imageMobile := "https://example.com/images/red-hot-chilly-peppers-mobile.jpg"
	ImageThumbnail := "https://example.com/images/red-hot-chilly-peppers-thumbnail.jpg"

	newProduct, err := repository.AdminProducts.Create(context.Background(), admin_products_repository.CreateParams{
		Name:           "Red Hot Chilly Peppers",
		Description:    &description,
		Price:          5.99,
		DiscountPrice:  &DiscountPrice,
		Active:         true,
		Image:          &image,
		ImageMobile:    &imageMobile,
		ImageThumbnail: &ImageThumbnail,
		CategoryID:     product_categories_enum.EVENT,
		CreatedBy:      userID,
	})
	if err != nil {
		t.Fatalf("error on creating product: %v", err)
	}
	return newProduct
}
