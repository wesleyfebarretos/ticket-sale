package integration_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

	t.Run("it should get all products", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newProduct(t, adminUser.ID)

		res := TMakeRequest(t, http.MethodGet, "admin/products", nil)

		getAllResponse := []admin_product_controller.GetAllResponseDto{}

		test_utils.Decode(t, res.Body, &getAllResponse)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, 1, len(getAllResponse))
		assert.IsType(t, admin_product_controller.GetAllResponseDto{}, getAllResponse[0])
	}))

	t.Run("it should get all products with relations", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newProduct(t, adminUser.ID)

		res := TMakeRequest(t, http.MethodGet, "admin/products/details", nil)

		getAllWithRleationsResponse := []admin_product_controller.GetAllWithRelationsResponseDto{}

		test_utils.Decode(t, res.Body, &getAllWithRleationsResponse)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, 1, len(getAllWithRleationsResponse))
		assert.IsType(t, admin_product_controller.GetAllWithRelationsResponseDto{}, getAllWithRleationsResponse[0])
	}))

	t.Run("it should get product by id", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newProduct := newProduct(t, adminUser.ID)

		res := TMakeRequest(t, http.MethodGet, fmt.Sprintf("admin/products/%d", newProduct.ID), nil)

		getOneByIdResponse := admin_product_controller.GetOneByIdResponseDto{}

		test_utils.Decode(t, res.Body, &getOneByIdResponse)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, newProduct.ID, getOneByIdResponse.ID)
		assert.IsType(t, admin_product_controller.GetOneByIdResponseDto{}, getOneByIdResponse)
	}))

	t.Run("it should get product by uuid", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newProduct := newProduct(t, adminUser.ID)

		res := TMakeRequest(t, http.MethodGet, fmt.Sprintf("admin/products/uuid/%s", newProduct.Uuid), nil)

		getOneByUuidResponse := admin_product_controller.GetOneByUuidResponseDto{}

		test_utils.Decode(t, res.Body, &getOneByUuidResponse)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, newProduct.Uuid, getOneByUuidResponse.Uuid)
		assert.IsType(t, admin_product_controller.GetOneByUuidResponseDto{}, getOneByUuidResponse)
	}))

	t.Run("it should validate required fields on create", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newProductRequest := admin_product_controller.CreateRequestDto{
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

		bStruct, err := json.Marshal(newProductRequest)
		if err != nil {
			t.Fatalf("error on marshal json :%v", err)
		}

		requiredFields := []string{
			"name",
			"price",
			"image",
			"imageMobile",
			"imageThumbnail",
			"categoryId",
			"stock.qty",
		}

		for _, k := range requiredFields {
			structMap := make(map[string]interface{})

			json.Unmarshal(bStruct, &structMap)

			if strings.Contains(k, ".") {
				keys := strings.Split(k, ".")
				delete(structMap[keys[0]].(map[string]any), keys[1])
			} else {
				delete(structMap, k)
			}

			res := TMakeRequest(t, http.MethodPost, "admin/products", structMap)

			assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		}
	}))

	t.Run("it should validate required fields on update", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newProduct := newProduct(t, adminUser.ID)

		updateProductRequest := admin_product_controller.UpdateRequestDto{
			Name:           "Red Hot Chilly Peppers",
			Description:    TPointer("Fresh and fiery red hot chilly peppers, perfect for adding a spicy kick to your dishes."),
			Price:          5.99,
			DiscountPrice:  TPointer(4.99),
			Active:         true,
			Image:          TPointer("https://example.com/images/red-hot-chilly-peppers.jpg"),
			ImageMobile:    TPointer("https://example.com/images/red-hot-chilly-peppers-mobile.jpg"),
			ImageThumbnail: TPointer("https://example.com/images/red-hot-chilly-peppers-thumbnail.jpg"),
			CategoryID:     product_categories_enum.EVENT,
		}

		bStruct, err := json.Marshal(updateProductRequest)
		if err != nil {
			t.Fatalf("error on marshal json :%v", err)
		}

		requiredFields := []string{
			"name",
			"price",
			"image",
			"imageMobile",
			"imageThumbnail",
			"categoryId",
		}

		for _, k := range requiredFields {
			structMap := make(map[string]interface{})

			json.Unmarshal(bStruct, &structMap)

			if strings.Contains(k, ".") {
				keys := strings.Split(k, ".")
				delete(structMap[keys[0]].(map[string]any), keys[1])
			} else {
				delete(structMap, k)
			}

			res := TMakeRequest(t, http.MethodPut, fmt.Sprintf("admin/products/%d", newProduct.ID), structMap)

			assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		}

		t.Run("it should make sure that only an admin can access this routes", TRun(func(t *testing.T) {
			user := test_utils.CreateUser(roles_enum.USER)
			TSetCookieWithUser(t, user)

			methods := []string{
				http.MethodGet,
				http.MethodGet,
				http.MethodGet,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
			}

			routes := []string{
				"admin/products",
				"admin/products/details",
				"admin/products/1",
				"admin/products/uuid/1",
				"admin/products",
				"admin/products/1",
				"admin/products/1",
			}

			for i, route := range routes {
				res := TMakeRequest(t, methods[i], route, nil)
				assert.Equal(t, http.StatusForbidden, res.StatusCode)
			}
		}))
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
