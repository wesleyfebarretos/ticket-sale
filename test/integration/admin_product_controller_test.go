package integration_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum/product_categories_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller/admin_product_controller"
	"github.com/wesleyfebarretos/ticket-sale/repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_products_repository"
	"github.com/wesleyfebarretos/ticket-sale/test/test_utils"
)

func TestAdminProductController(t *testing.T) {
	t.Run("it should create a product", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newProduct := admin_product_controller.CreateRequestDto{
			Name:           "Red Hot Chilly Peppers",
			Description:    TPointer("Fresh and fiery red hot chilly peppers, perfect for adding a spicy kick to your dishes."),
			Price:          5.99,
			DiscountPrice:  TPointer(4.99),
			Active:         true,
			Image:          TPointer("https://example.com/images/red-hot-chilly-peppers.jpg"),
			ImageMobile:    TPointer("https://example.com/images/red-hot-chilly-peppers-mobile.jpg"),
			ImageThumbnail: TPointer("https://example.com/images/red-hot-chilly-peppers-thumbnail.jpg"),
			CategoryID:     product_categories_enum.EVENT,
			Stock: admin_product_controller.CreateStockRequestDto{
				Qty:    100,
				MinQty: TPointer(int32(50)),
			},
		}

		res := TMakeRequest(t, http.MethodPost, "admin/products", newProduct)

		newProductResponse := admin_product_controller.CreateResponseDto{}

		test_utils.Decode(t, res.Body, &newProductResponse)

		expectedProduct := admin_product_controller.CreateResponseDto{
			ID:             newProductResponse.ID,
			Name:           newProduct.Name,
			Description:    newProduct.Description,
			Price:          newProduct.Price,
			DiscountPrice:  newProduct.DiscountPrice,
			Active:         newProduct.Active,
			Image:          newProduct.Image,
			ImageMobile:    newProduct.ImageMobile,
			ImageThumbnail: newProduct.ImageThumbnail,
			CategoryID:     newProduct.CategoryID,
			CreatedBy:      adminUser.ID,
			Uuid:           newProductResponse.Uuid,
			IsDeleted:      false,
			UpdatedBy:      newProductResponse.UpdatedBy,
			CreatedAt:      newProductResponse.CreatedAt,
			UpdatedAt:      newProductResponse.UpdatedAt,
			Stock: admin_product_controller.CreateStockResponseDto{
				ID:        newProductResponse.Stock.ID,
				ProductID: newProductResponse.ID,
				Qty:       newProduct.Stock.Qty,
				MinQty:    newProduct.Stock.MinQty,
			},
		}

		assert.Equal(t, http.StatusCreated, res.StatusCode)
		assert.Equal(t, expectedProduct, newProductResponse)
		assert.GreaterOrEqual(t, newProductResponse.ID, int32(1))
		assert.GreaterOrEqual(t, newProductResponse.Stock.ID, int32(1))
	}))

	t.Run("it should update an product", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newProduct := newProduct(t, adminUser.ID)

		updateProduct := admin_product_controller.UpdateRequestDto{
			Name:           "update",
			Description:    TPointer("update"),
			Price:          5.99,
			DiscountPrice:  TPointer(4.99),
			Active:         true,
			Image:          TPointer("update"),
			ImageMobile:    TPointer("update"),
			ImageThumbnail: TPointer("update"),
			CategoryID:     product_categories_enum.DIGITAL,
		}

		res := TMakeRequest(t, http.MethodPut, fmt.Sprintf("admin/products/%d", newProduct.ID), updateProduct)

		updatedProductResponse := false

		test_utils.Decode(t, res.Body, &updatedProductResponse)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, true, updatedProductResponse)
	}))

	t.Run("it should delete a product", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newProduct := newProduct(t, adminUser.ID)

		res := TMakeRequest(t, http.MethodDelete, fmt.Sprintf("admin/products/%d", newProduct.ID), nil)

		deleteProductResponse := false

		test_utils.Decode(t, res.Body, &deleteProductResponse)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, true, deleteProductResponse)
	}))
}

func newProduct(t *testing.T, userID int32) admin_products_repository.Product {
	uuid, err := uuid.NewV7()
	if err != nil {
		t.Errorf("error on creating UUID: %v", err)
	}
	newProduct, err := repository.AdminProducts.Create(context.Background(), admin_products_repository.CreateParams{
		Name:           "Red Hot Chilly Peppers",
		Description:    TPointer("Fresh and fiery red hot chilly peppers, perfect for adding a spicy kick to your dishes."),
		Price:          5.99,
		DiscountPrice:  TPointer(4.99),
		Active:         true,
		Image:          TPointer("https://example.com/images/red-hot-chilly-peppers.jpg"),
		ImageMobile:    TPointer("https://example.com/images/red-hot-chilly-peppers-mobile.jpg"),
		ImageThumbnail: TPointer("https://example.com/images/red-hot-chilly-peppers-thumbnail.jpg"),
		Uuid:           uuid,
		CategoryID:     product_categories_enum.EVENT,
		CreatedBy:      userID,
	})
	if err != nil {
		t.Fatalf("error on creating product: %v", err)
	}
	return newProduct
}
