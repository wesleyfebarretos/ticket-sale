package integration_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/gateway_customer_card_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/gateway_customer_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/gateway_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/handler/creditcard_handler"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/tests/test_data"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/tests/test_utils"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
)

func TestCreditcardHandler(t *testing.T) {
	t.Run("it should create a creditcard", TRun(func(t *testing.T) {
		user := test_data.NewUser(roles_enum.USER)

		TSetCookieWithUser(t, user)

		expiration := time.Now().UTC().AddDate(3, 0, 0).Format(time.DateOnly)

		newCreditcard := creditcard_handler.CreateRequestDto{
			Name:             "Testing",
			Number:           "5574723384289379",
			Expiration:       expiration,
			Priority:         1,
			NotifyExpiration: true,
			CreditcardTypeID: 1,
			CreditcardFlagID: 1,
			CVC:              "434",
		}

		req := TMakeRequest(t, http.MethodPost, "creditcard", newCreditcard)

		newCreditcardResponse := creditcard_handler.CreateResponseDto{}

		test_utils.Decode(t, req.Body, &newCreditcardResponse)

		currentDate := newCreditcardResponse.Expiration.Format(time.DateOnly)

		assert.Equal(t, http.StatusCreated, req.StatusCode)
		assert.Equal(t, expiration, currentDate)
		assert.Equal(t, newCreditcard.Name, newCreditcardResponse.Name)
		assert.Equal(t, utils.MaskCreditcardNumber(newCreditcard.Number), newCreditcardResponse.Number)
		assert.Equal(t, newCreditcard.Priority, newCreditcardResponse.Priority)
		assert.Equal(t, newCreditcard.NotifyExpiration, newCreditcardResponse.NotifyExpiration)
		assert.Equal(t, newCreditcard.CreditcardTypeID, newCreditcardResponse.CreditcardTypeID)
		assert.Equal(t, newCreditcard.CreditcardFlagID, newCreditcardResponse.CreditcardFlagID)
		assert.NotEqual(t, uuid.Nil, newCreditcardResponse.Uuid)

		ctx := context.Background()

		gateway := gateway_repository.New().GetActive(ctx)

		assert.NotNil(t, gateway)

		customer := gateway_customer_repository.New().FindOneByGatewayAndUserId(ctx, gateway_customer_repository.FindOneByGatewayAndUserIdParams{
			UserID:    user.ID,
			GatewayID: gateway.ID,
		})

		assert.NotNil(t, customer)

		customerCard := gateway_customer_card_repository.New().GetByCardAndGatewayId(ctx, gateway.ID, newCreditcardResponse.ID)

		assert.NotNil(t, customerCard)
	}))

	t.Run("it should update a creditcard", TRun(func(t *testing.T) {
		user := test_data.NewUser(roles_enum.USER)

		TSetCookieWithUser(t, user)

		newCreditcard := test_data.NewCreditCard(t, user.ID)

		expiration := time.Now().AddDate(3, 0, 0).UTC()

		updateCreditcard := creditcard_handler.UpdateRequestDto{
			Name:             "Testing update",
			Number:           "5574723384289371",
			Expiration:       expiration,
			Priority:         2,
			NotifyExpiration: false,
			CreditcardTypeID: 2,
			CreditcardFlagID: 2,
		}

		req := TMakeRequest(t, http.MethodPut, fmt.Sprintf("creditcard/%s", newCreditcard.Uuid), updateCreditcard)

		var updateRequestResponse bool

		test_utils.Decode(t, req.Body, &updateRequestResponse)

		assert.Equal(t, http.StatusOK, req.StatusCode)
		assert.True(t, updateRequestResponse)
	}))

	t.Run("it should delete a creditcard", TRun(func(t *testing.T) {
		user := test_data.NewUser(roles_enum.USER)

		TSetCookieWithUser(t, user)

		newCreditcard := test_data.NewCreditCard(t, user.ID)

		req := TMakeRequest(t, http.MethodDelete, fmt.Sprintf("creditcard/%s", newCreditcard.Uuid), nil)

		var deleteRequestResponse bool

		test_utils.Decode(t, req.Body, &deleteRequestResponse)

		assert.Equal(t, http.StatusOK, req.StatusCode)
		assert.True(t, deleteRequestResponse)
	}))

	t.Run("it should get all user creditcards", TRun(func(t *testing.T) {
		user := test_data.NewUser(roles_enum.USER)

		TSetCookieWithUser(t, user)

		newCcQty := 5

		for i := 1; i <= newCcQty; i++ {
			test_data.NewCreditCard(t, user.ID)
		}

		req := TMakeRequest(t, http.MethodGet, "creditcard/user", nil)

		response := []creditcard_handler.GetAllUserCreditcardsResponseDto{}

		test_utils.Decode(t, req.Body, &response)

		assert.Equal(t, http.StatusOK, req.StatusCode)
		assert.Equal(t, len(response), newCcQty)
	}))

	t.Run("it should validate login to access routes", TRun(func(t *testing.T) {
		methods := []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
			http.MethodDelete,
		}

		routes := []string{
			"creditcard",
			"creditcard/user",
			"creditcard/4eeaaeac-2531-493f-845a-9590d0d2640e",
			"creditcard/4eeaaeac-2531-493f-845a-9590d0d2640e",
		}

		for i, route := range routes {
			res := TMakeRequest(t, methods[i], route, nil)
			assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
		}
	}))

	t.Run("it should validate required fields on create", TRun(func(t *testing.T) {
		user := test_data.NewUser(roles_enum.USER)
		TSetCookieWithUser(t, user)

		newCreditcard := test_data.NewCreditCard(t, user.ID)

		bStruct, err := json.Marshal(newCreditcard)
		if err != nil {
			t.Fatalf("error on marshal json :%v", err)
		}

		requiredFields := []string{
			"name",
			"number",
			"expiration",
			"creditcardTypeId",
			"creditcardFlagId",
		}

		for _, k := range requiredFields {
			structMap := make(map[string]interface{})

			json.Unmarshal(bStruct, &structMap)

			test_utils.DeleteRequiredField(t, k, structMap)

			res := TMakeRequest(t, http.MethodPost, "creditcard", structMap)

			assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		}
	}))

	t.Run("it should validate required fields on update", TRun(func(t *testing.T) {
		user := test_data.NewUser(roles_enum.USER)
		TSetCookieWithUser(t, user)

		newCreditcard := test_data.NewCreditCard(t, user.ID)

		bStruct, err := json.Marshal(newCreditcard)
		if err != nil {
			t.Fatalf("error on marshal json :%v", err)
		}

		requiredFields := []string{
			"name",
			"number",
			"expiration",
			"creditcardTypeId",
			"creditcardFlagId",
		}

		for _, k := range requiredFields {
			structMap := make(map[string]interface{})

			json.Unmarshal(bStruct, &structMap)

			test_utils.DeleteRequiredField(t, k, structMap)

			res := TMakeRequest(t, http.MethodPut, fmt.Sprintf("creditcard/%s", newCreditcard.Uuid), structMap)

			assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		}
	}))
}
