package integration_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/product_categories_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_event_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/handler/admin_event_handler"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/handler/admin_product_handler"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/tests/test_data"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/tests/test_utils"
)

func TestAdminEventHandler(t *testing.T) {
	t.Run("it should create an event", TRun(func(t *testing.T) {
		adminUser := test_data.NewUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		startAt, err := time.Parse(time.RFC3339, "2024-01-01T00:00:00Z")
		if err != nil {
			t.Errorf("error on parse time %v", err)
		}
		endAt, err := time.Parse(time.RFC3339, "2024-01-10T00:00:00Z")
		if err != nil {
			t.Errorf("error on parse time %v", err)
		}

		installments := []admin_product_handler.CreateInstallmentsRequestDto{
			{
				ID:            1,
				PaymentTypeID: 1,
				Fee:           TPointer(3.22),
				Tariff:        TPointer(4.22),
			},
		}

		newEvent := admin_event_handler.CreateRequestDto{
			StartAt:  &startAt,
			EndAt:    &endAt,
			City:     TPointer("Orlando"),
			State:    TPointer("FL"),
			Location: TPointer("Disney"),
			Product: admin_product_handler.CreateRequestDto{
				Name:           "Red Hot Chilly Peppers",
				Description:    TPointer("Fresh and fiery red hot chilly peppers, perfect for adding a spicy kick to your dishes."),
				Price:          5.99,
				DiscountPrice:  TPointer(4.99),
				Active:         false,
				Image:          TPointer("https://example.com/images/red-hot-chilly-peppers.jpg"),
				ImageMobile:    TPointer("https://example.com/images/red-hot-chilly-peppers-mobile.jpg"),
				ImageThumbnail: TPointer("https://example.com/images/red-hot-chilly-peppers-thumbnail.jpg"),
				CategoryID:     3,
				Stock: admin_product_handler.CreateStockRequestDto{
					Qty:    100,
					MinQty: TPointer(int32(50)),
				},
				Installments: installments,
			},
		}

		res := TMakeRequest(t, http.MethodPost, "admin/events", newEvent)

		newEventResponse := admin_event_handler.CreateResponseDto{}

		test_utils.Decode(t, res.Body, &newEventResponse)

		expectedEvent := admin_event_handler.CreateResponseDto{
			ID:        newEventResponse.ID,
			ProductID: newEventResponse.ProductID,
			City:      newEvent.City,
			State:     newEvent.State,
			Location:  newEvent.Location,
			EndAt:     newEvent.EndAt,
			StartAt:   newEvent.StartAt,
			Product: admin_product_handler.CreateResponseDto{
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
				Stock: admin_product_handler.CreateStockResponseDto{
					ID:        newEventResponse.Product.Stock.ID,
					ProductID: newEventResponse.Product.Stock.ProductID,
					Qty:       newEvent.Product.Stock.Qty,
					MinQty:    newEvent.Product.Stock.MinQty,
				},
				Installments: newEventResponse.Product.Installments,
			},
		}
		assert.Equal(t, http.StatusCreated, res.StatusCode)
		assert.Equal(t, expectedEvent, newEventResponse)
		assert.GreaterOrEqual(t, newEventResponse.ID, int32(1))
		assert.GreaterOrEqual(t, newEventResponse.Product.ID, int32(1))
		assert.GreaterOrEqual(t, newEventResponse.Product.Stock.ID, int32(1))
	}))

	t.Run("it should update an event", TRun(func(t *testing.T) {
		adminUser := test_data.NewUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newEvent := test_data.NewEvent(t, adminUser.ID)

		updateEvent := admin_event_handler.UpdateRequestDto{
			StartAt:  nil,
			EndAt:    nil,
			City:     nil,
			State:    nil,
			Location: nil,
			Product: admin_product_handler.UpdateRequestDto{
				Name:           "update",
				Description:    TPointer("updated"),
				Price:          1,
				DiscountPrice:  nil,
				Active:         false,
				Image:          TPointer("updated"),
				ImageMobile:    TPointer("updated"),
				ImageThumbnail: TPointer("updated"),
				CategoryID:     product_categories_enum.PHYSICAL,
				Installments: []admin_product_handler.UpdateInstallmentsRequestDto{
					{
						ID:            1,
						PaymentTypeID: 1,
						Fee:           TPointer(7.1),
						Tariff:        TPointer(7.1),
					},
				},
			},
		}

		res := TMakeRequest(t, http.MethodPut, fmt.Sprintf("admin/events/%d", newEvent.ID), updateEvent)

		updatedEventResponse := false

		event := admin_event_repository.New().GetOneById(context.Background(), newEvent.ID)
		if event == nil {
			t.Fatal("error on try to get updated event")
		}

		test_utils.Decode(t, res.Body, &updatedEventResponse)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, true, updatedEventResponse)
		assert.Equal(t, len(event.Product.Installments.Creditcard), len(updateEvent.Product.Installments))
		assert.Equal(t, event.Product.Installments.Creditcard[0].Fee, *updateEvent.Product.Installments[0].Fee)
		assert.Equal(t, event.Product.Installments.Creditcard[0].Tariff, *updateEvent.Product.Installments[0].Tariff)
		assert.Equal(t, event.Product.Installments.Creditcard[0].InstallmentTimeID, updateEvent.Product.Installments[0].ID)
	}))

	t.Run("it should delete an event", TRun(func(t *testing.T) {
		adminUser := test_data.NewUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newEvent := test_data.NewEvent(t, adminUser.ID)

		res := TMakeRequest(t, http.MethodDelete, fmt.Sprintf("admin/events/%d", newEvent.ID), nil)

		deleteEventResponse := false

		event := admin_event_repository.New().GetOneById(context.Background(), newEvent.ID)

		test_utils.Decode(t, res.Body, &deleteEventResponse)

		assert.Nil(t, event)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, true, deleteEventResponse)
	}))

	t.Run("it should get all events", TRun(func(t *testing.T) {
		adminUser := test_data.NewUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		test_data.NewEvent(t, adminUser.ID)

		res := TMakeRequest(t, http.MethodGet, "admin/events", nil)

		getAllResponse := []admin_event_handler.GetAllResponseDto{}

		test_utils.Decode(t, res.Body, &getAllResponse)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, 1, len(getAllResponse))
		assert.IsType(t, admin_event_handler.GetAllResponseDto{}, getAllResponse[0])
	}))

	t.Run("it should get event by id", TRun(func(t *testing.T) {
		adminUser := test_data.NewUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newEvent := test_data.NewEvent(t, adminUser.ID)

		res := TMakeRequest(t, http.MethodGet, fmt.Sprintf("admin/events/%d", newEvent.ID), nil)

		getOneByIdResponse := admin_event_handler.GetOneByIdResponseDto{}

		test_utils.Decode(t, res.Body, &getOneByIdResponse)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, newEvent.ID, getOneByIdResponse.ID)
		assert.IsType(t, admin_event_handler.GetOneByIdResponseDto{}, getOneByIdResponse)
	}))

	t.Run("it should not found an event by id", TRun(func(t *testing.T) {
		adminUser := test_data.NewUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		res := TMakeRequest(t, http.MethodGet, fmt.Sprintf("admin/events/%d", 1), nil)

		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	}))

	t.Run("it should validate required fields on create", TRun(func(t *testing.T) {
		adminUser := test_data.NewUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newEventRequest := admin_event_handler.CreateRequestDto{
			StartAt:  nil,
			EndAt:    nil,
			City:     nil,
			State:    nil,
			Location: nil,
			Product: admin_product_handler.CreateRequestDto{
				Name:           "test",
				Description:    nil,
				Price:          1,
				DiscountPrice:  nil,
				Active:         false,
				Image:          TPointer("test"),
				ImageMobile:    TPointer("test"),
				ImageThumbnail: TPointer("test"),
				CategoryID:     product_categories_enum.EVENT,
				Stock: admin_product_handler.CreateStockRequestDto{
					Qty:    1,
					MinQty: nil,
				},
				Installments: []admin_product_handler.CreateInstallmentsRequestDto{
					{
						ID:            1,
						PaymentTypeID: 1,
						Fee:           TPointer(7.1),
						Tariff:        TPointer(7.1),
					},
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
			"product.installments",
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
		adminUser := test_data.NewUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newEvent := test_data.NewEvent(t, adminUser.ID)

		updateEventRequest := admin_event_handler.UpdateRequestDto{
			StartAt:  nil,
			EndAt:    nil,
			City:     nil,
			State:    nil,
			Location: nil,
			Product: admin_product_handler.UpdateRequestDto{
				Name:           "test",
				Description:    nil,
				Price:          1,
				DiscountPrice:  nil,
				Active:         false,
				Image:          TPointer("test"),
				ImageMobile:    TPointer("test"),
				ImageThumbnail: TPointer("test"),
				CategoryID:     product_categories_enum.EVENT,
				Installments: []admin_product_handler.UpdateInstallmentsRequestDto{
					{
						ID:            1,
						PaymentTypeID: 1,
						Fee:           TPointer(7.1),
						Tariff:        TPointer(7.1),
					},
				},
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
			"product.installments",
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
		user := test_data.NewUser(roles_enum.USER)
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
