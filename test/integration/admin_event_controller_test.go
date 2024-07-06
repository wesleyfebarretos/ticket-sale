package integration_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum/product_categories_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller/admin_event_controller"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller/admin_product_controller"
	"github.com/wesleyfebarretos/ticket-sale/repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_events_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_products_repository"
	"github.com/wesleyfebarretos/ticket-sale/test/test_utils"
)

func TestAdminEventController(t *testing.T) {
	t.Run("it should create an event", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		startAt, err := time.Parse(time.RFC3339, "2024-01-01T00:00:00Z")
		if err != nil {
			t.Errorf("error on parse time %v", err)
		}
		endAt, err := time.Parse(time.RFC3339, "2024-01-10T00:00:00Z")
		if err != nil {
			t.Errorf("error on parse time %v", err)
		}

		newEvent := admin_event_controller.CreateRequestDto{
			StartAt:  &startAt,
			EndAt:    &endAt,
			City:     TPointer("Orlando"),
			State:    TPointer("FL"),
			Location: TPointer("Disney"),
			Product: admin_product_controller.CreateRequestDto{
				Name:           "Red Hot Chilly Peppers",
				Description:    TPointer("Fresh and fiery red hot chilly peppers, perfect for adding a spicy kick to your dishes."),
				Price:          5.99,
				DiscountPrice:  TPointer(4.99),
				Active:         false,
				Image:          TPointer("https://example.com/images/red-hot-chilly-peppers.jpg"),
				ImageMobile:    TPointer("https://example.com/images/red-hot-chilly-peppers-mobile.jpg"),
				ImageThumbnail: TPointer("https://example.com/images/red-hot-chilly-peppers-thumbnail.jpg"),
				CategoryID:     3,
				Stock: admin_product_controller.CreateStockRequestDto{
					Qty:    100,
					MinQty: TPointer(int32(50)),
				},
			},
		}

		res := TMakeRequest(t, http.MethodPost, "admin/events", newEvent)

		newEventResponse := admin_event_controller.CreateResponseDto{}

		test_utils.Decode(t, res.Body, &newEventResponse)

		expectedEvent := admin_event_controller.CreateResponseDto{
			ID:        newEventResponse.ID,
			ProductID: newEventResponse.ProductID,
			City:      newEvent.City,
			State:     newEvent.State,
			Location:  newEvent.Location,
			EndAt:     newEvent.EndAt,
			StartAt:   newEvent.StartAt,
			Product: admin_product_controller.CreateResponseDto{
				ID:             newEventResponse.Product.ID,
				Uuid:           newEventResponse.Product.Uuid,
				Name:           newEvent.Product.Name,
				Description:    newEvent.Product.Description,
				Price:          newEvent.Product.Price,
				DiscountPrice:  newEvent.Product.DiscountPrice,
				Active:         newEvent.Product.Active,
				Image:          newEvent.Product.Image,
				ImageMobile:    newEvent.Product.ImageMobile,
				ImageThumbnail: newEvent.Product.ImageThumbnail,
				CategoryID:     newEvent.Product.CategoryID,
				CreatedBy:      newEventResponse.Product.CreatedBy,
				IsDeleted:      newEventResponse.Product.IsDeleted,
				UpdatedBy:      newEventResponse.Product.UpdatedBy,
				CreatedAt:      newEventResponse.Product.CreatedAt,
				UpdatedAt:      newEventResponse.Product.UpdatedAt,
				Stock: admin_product_controller.CreateStockResponseDto{
					ID:        newEventResponse.Product.Stock.ID,
					ProductID: newEventResponse.Product.Stock.ProductID,
					Qty:       newEvent.Product.Stock.Qty,
					MinQty:    newEvent.Product.Stock.MinQty,
				},
			},
		}
		assert.Equal(t, http.StatusCreated, res.StatusCode)
		assert.Equal(t, expectedEvent, newEventResponse)
		assert.GreaterOrEqual(t, newEventResponse.ID, int32(1))
		assert.GreaterOrEqual(t, newEventResponse.Product.ID, int32(1))
		assert.GreaterOrEqual(t, newEventResponse.Product.Stock.ID, int32(1))
	}))

	t.Run("it should update an event", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newEvent := newEvent(t, adminUser.ID)

		updateEvent := admin_event_controller.UpdateRequestDto{
			StartAt:  nil,
			EndAt:    nil,
			City:     nil,
			State:    nil,
			Location: nil,
			Product: admin_product_controller.UpdateRequestDto{
				Name:           "update",
				Description:    TPointer("updated"),
				Price:          1,
				DiscountPrice:  nil,
				Active:         false,
				Image:          TPointer("updated"),
				ImageMobile:    TPointer("updated"),
				ImageThumbnail: TPointer("updated"),
				CategoryID:     product_categories_enum.PHYSICAL,
			},
		}

		res := TMakeRequest(t, http.MethodPut, fmt.Sprintf("admin/events/%d", newEvent.ID), updateEvent)

		updatedEventResponse := false

		event, err := repository.AdminEvents.GetOneById(context.Background(), newEvent.ID)
		if err != nil {
			t.Fatalf("error on try to get updated event %v", err)
		}

		updatedEvent := admin_event_controller.UpdateRequestDto{}

		bEvent, err := json.Marshal(event)
		if err != nil {
			t.Fatalf("could not marshal event json %v", err)
		}

		if err := json.Unmarshal(bEvent, &updatedEvent); err != nil {
			t.Fatalf("could not unmarshal event json %v", err)
		}

		test_utils.Decode(t, res.Body, &updatedEventResponse)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, true, updatedEventResponse)
		assert.Equal(t, updateEvent, updatedEvent)
	}))

	t.Run("it should delete an event", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newEvent := newEvent(t, adminUser.ID)

		res := TMakeRequest(t, http.MethodDelete, fmt.Sprintf("admin/events/%d", newEvent.ID), nil)

		deleteEventResponse := false

		_, err := repository.AdminEvents.GetOneById(context.Background(), newEvent.ID)

		test_utils.Decode(t, res.Body, &deleteEventResponse)

		assert.IsType(t, pgx.ErrNoRows, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, true, deleteEventResponse)
	}))

	t.Run("it should get all events", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newEvent(t, adminUser.ID)

		res := TMakeRequest(t, http.MethodGet, "admin/events", nil)

		getAllResponse := []admin_event_controller.GetAllResponseDto{}

		test_utils.Decode(t, res.Body, &getAllResponse)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, 1, len(getAllResponse))
		assert.IsType(t, admin_event_controller.GetAllResponseDto{}, getAllResponse[0])
	}))

	t.Run("it should get event by id", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newEvent := newEvent(t, adminUser.ID)

		res := TMakeRequest(t, http.MethodGet, fmt.Sprintf("admin/events/%d", newEvent.ID), nil)

		getOneByIdResponse := admin_event_controller.GetOneByIdResponseDto{}

		test_utils.Decode(t, res.Body, &getOneByIdResponse)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, newEvent.ID, getOneByIdResponse.ID)
		assert.IsType(t, admin_event_controller.GetOneByIdResponseDto{}, getOneByIdResponse)
	}))

	t.Run("it should not found an event by id", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		res := TMakeRequest(t, http.MethodGet, fmt.Sprintf("admin/events/%d", 1), nil)

		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	}))

	t.Run("it should validate required fields on create", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newEventRequest := admin_event_controller.CreateRequestDto{
			StartAt:  nil,
			EndAt:    nil,
			City:     nil,
			State:    nil,
			Location: nil,
			Product: admin_product_controller.CreateRequestDto{
				Name:           "test",
				Description:    nil,
				Price:          1,
				DiscountPrice:  nil,
				Active:         false,
				Image:          TPointer("test"),
				ImageMobile:    TPointer("test"),
				ImageThumbnail: TPointer("test"),
				CategoryID:     product_categories_enum.EVENT,
				Stock: admin_product_controller.CreateStockRequestDto{
					Qty:    1,
					MinQty: nil,
				},
			},
		}

		bStruct, err := json.Marshal(newEventRequest)
		if err != nil {
			t.Fatalf("error on marshal json :%v", err)
		}

		requiredFields := []string{
			"product.name",
			"product.price",
			"product.image",
			"product.imageMobile",
			"product.imageThumbnail",
			"product.categoryId",
			"product.stock.qty",
		}

		for _, k := range requiredFields {
			structMap := make(map[string]interface{})

			json.Unmarshal(bStruct, &structMap)

			test_utils.DeleteRequiredField(t, k, structMap)

			res := TMakeRequest(t, http.MethodPost, "admin/events", structMap)

			assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		}
	}))

	t.Run("it should validate required fields on update", TRun(func(t *testing.T) {
		adminUser := test_utils.CreateUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newEvent := newEvent(t, adminUser.ID)

		updateEventRequest := admin_event_controller.UpdateRequestDto{
			StartAt:  nil,
			EndAt:    nil,
			City:     nil,
			State:    nil,
			Location: nil,
			Product: admin_product_controller.UpdateRequestDto{
				Name:           "test",
				Description:    nil,
				Price:          1,
				DiscountPrice:  nil,
				Active:         false,
				Image:          TPointer("test"),
				ImageMobile:    TPointer("test"),
				ImageThumbnail: TPointer("test"),
				CategoryID:     product_categories_enum.EVENT,
			},
		}

		bStruct, err := json.Marshal(updateEventRequest)
		if err != nil {
			t.Fatalf("error on marshal json :%v", err)
		}

		requiredFields := []string{
			"product.name",
			"product.price",
			"product.image",
			"product.imageMobile",
			"product.imageThumbnail",
			"product.categoryId",
		}

		for _, k := range requiredFields {
			structMap := make(map[string]interface{})

			json.Unmarshal(bStruct, &structMap)

			test_utils.DeleteRequiredField(t, k, structMap)

			res := TMakeRequest(t, http.MethodPut, fmt.Sprintf("admin/products/%d", newEvent.ID), structMap)

			assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		}
	}))

	t.Run("it should make sure that only an admin can access this routes", TRun(func(t *testing.T) {
		user := test_utils.CreateUser(roles_enum.USER)
		TSetCookieWithUser(t, user)

		methods := []string{
			http.MethodGet,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		}

		routes := []string{
			"admin/events",
			"admin/events/1",
			"admin/events",
			"admin/events/1",
			"admin/events/1",
		}

		for i, route := range routes {
			res := TMakeRequest(t, methods[i], route, nil)
			assert.Equal(t, http.StatusForbidden, res.StatusCode)
		}
	}))
}

func newEvent(t *testing.T, userID int32) admin_event_controller.CreateResponseDto {
	uuid, err := uuid.NewV7()
	if err != nil {
		t.Errorf("error on creating UUID: %v", err)
	}
	ctx := context.Background()
	newProduct, err := repository.AdminProducts.Create(ctx, admin_products_repository.CreateParams{
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

	startAt, err := time.Parse(time.RFC3339, "2024-01-01T00:00:00Z")
	if err != nil {
		t.Errorf("error on parse time %v", err)
	}

	endAt, err := time.Parse(time.RFC3339, "2024-01-10T00:00:00Z")
	if err != nil {
		t.Errorf("error on parse time %v", err)
	}

	newEvent, err := repository.AdminEvents.Create(ctx, admin_events_repository.CreateParams{
		ProductID: newProduct.ID,
		StartAt:   &startAt,
		EndAt:     &endAt,
		City:      TPointer("Orlando"),
		State:     TPointer("FL"),
		Location:  TPointer("Disney"),
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