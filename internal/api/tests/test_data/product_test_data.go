package test_data

import (
	"context"
	"testing"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/product_categories_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_product_repository"
)

func NewProduct(t *testing.T, userID int32) admin_product_repository.CreateResponse {
	description := "Fresh and fiery red hot chilly peppers, perfect for adding a spicy kick to your dishes."
	DiscountPrice := 4.99
	image := "https://example.com/images/red-hot-chilly-peppers.jpg"
	imageMobile := "https://example.com/images/red-hot-chilly-peppers-mobile.jpg"
	ImageThumbnail := "https://example.com/images/red-hot-chilly-peppers-thumbnail.jpg"

	repository := admin_product_repository.New()

	newProduct := repository.Create(context.Background(), admin_product_repository.CreateParams{
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

	return newProduct
}
