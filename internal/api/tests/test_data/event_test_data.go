package test_data

import (
	"context"
	"testing"
	"time"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/product_categories_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_product_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_events_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/controller/admin_event_controller"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/controller/admin_product_controller"
)

func NewEvent(t *testing.T, userID int32) admin_event_controller.CreateResponseDto {
	ctx := context.Background()

	description := "Fresh and fiery red hot chilly peppers, perfect for adding a spicy kick to your dishes."
	discountPrice := 4.99
	image := "https://example.com/images/red-hot-chilly-peppers.jpg"
	imageMobile := "https://example.com/images/red-hot-chilly-peppers-mobile.jpg"
	ImageThumbnail := "https://example.com/images/red-hot-chilly-peppers-thumbnail.jpg"

	adminProductRepository := admin_product_repository.New()

	newProduct := adminProductRepository.Create(ctx, admin_product_repository.CreateParams{
		Name:           "Red Hot Chilly Peppers",
		Description:    &description,
		Price:          5.99,
		DiscountPrice:  &discountPrice,
		Active:         true,
		Image:          &image,
		ImageMobile:    &imageMobile,
		ImageThumbnail: &ImageThumbnail,
		CategoryID:     product_categories_enum.EVENT,
		CreatedBy:      userID,
	})

	startAt, err := time.Parse(time.RFC3339, "2024-01-01T00:00:00Z")
	if err != nil {
		t.Errorf("error on parse time %v", err)
	}

	endAt, err := time.Parse(time.RFC3339, "2024-01-10T00:00:00Z")
	if err != nil {
		t.Errorf("error on parse time %v", err)
	}

	city := "Orlando"
	state := "FL"
	location := "Disney"

	newEvent, err := repository.AdminEvents.Create(ctx, admin_events_repository.CreateParams{
		ProductID: newProduct.ID,
		StartAt:   &startAt,
		EndAt:     &endAt,
		City:      &city,
		State:     &state,
		Location:  &location,
		CreatedBy: userID,
	})
	if err != nil {
		t.Fatalf("error on creating event: %v", err)
	}

	return admin_event_controller.CreateResponseDto{
		ID:        newEvent.ID,
		ProductID: newEvent.ProductID,
		City:      newEvent.City,
		State:     newEvent.State,
		Location:  newEvent.Location,
		EndAt:     newEvent.EndAt,
		StartAt:   newEvent.StartAt,
		Product: admin_product_controller.CreateResponseDto{
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
		},
	}
}
