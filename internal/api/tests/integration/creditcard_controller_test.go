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
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/creditcard_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/controller/creditcard_controller"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/tests/test_utils"
)

func TestCreditcardController(t *testing.T) {
	t.Run("it should create a creditcard", TRun(func(t *testing.T) {
		user := test_utils.CreateUser(roles_enum.USER)

		TSetCookieWithUser(t, user)

		expiration := time.Now().UTC().AddDate(3, 0, 0)

		newCreditcard := creditcard_controller.CreateRequestDto{
			Name:             "Testing",
			Number:           "5574723384289379",
			Expiration:       expiration,
			Priority:         1,
			NotifyExpiration: true,
			CreditcardTypeID: 1,
			CreditcardFlagID: 1,
		}

		req := TMakeRequest(t, http.MethodPost, "creditcard", newCreditcard)

		// test, _ := io.ReadAll(req.Body)
		//
		// fmt.Println(string(test))
		// os.Exit(1)

		newCreditcardResponse := creditcard_controller.CreateResponseDto{}

		test_utils.Decode(t, req.Body, &newCreditcardResponse)

		expectedDate := expiration.Format(time.DateOnly)
		currentDate := newCreditcardResponse.Expiration.Format(time.DateOnly)

		assert.Equal(t, http.StatusCreated, req.StatusCode)
		assert.Equal(t, expectedDate, currentDate)
		assert.Equal(t, newCreditcard.Name, newCreditcardResponse.Name)
		assert.Equal(t, newCreditcard.Number, newCreditcardResponse.Number)
		assert.Equal(t, newCreditcard.Priority, newCreditcardResponse.Priority)
		assert.Equal(t, newCreditcard.NotifyExpiration, newCreditcardResponse.NotifyExpiration)
		assert.Equal(t, newCreditcard.CreditcardTypeID, newCreditcardResponse.CreditcardTypeID)
		assert.Equal(t, newCreditcard.CreditcardFlagID, newCreditcardResponse.CreditcardFlagID)
		assert.NotEqual(t, uuid.Nil, newCreditcardResponse.Uuid)
	}))

	t.Run("it should update a creditcard", TRun(func(t *testing.T) {
		user := test_utils.CreateUser(roles_enum.USER)

		TSetCookieWithUser(t, user)

		newCreditcard := createCreditCard(t, user.ID)

		expiration := time.Now().AddDate(3, 0, 0).UTC()

		updateCreditcard := creditcard_controller.UpdateRequestDto{
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
		user := test_utils.CreateUser(roles_enum.USER)

		TSetCookieWithUser(t, user)

		newCreditcard := createCreditCard(t, user.ID)

		req := TMakeRequest(t, http.MethodDelete, fmt.Sprintf("creditcard/%s", newCreditcard.Uuid), nil)

		var deleteRequestResponse bool

		test_utils.Decode(t, req.Body, &deleteRequestResponse)

		assert.Equal(t, http.StatusOK, req.StatusCode)
		assert.True(t, deleteRequestResponse)
	}))

	t.Run("it should get all user creditcards", TRun(func(t *testing.T) {
		user := test_utils.CreateUser(roles_enum.USER)

		TSetCookieWithUser(t, user)

		newCcQty := 5

		for i := 1; i <= newCcQty; i++ {
			createCreditCard(t, user.ID)
		}

		req := TMakeRequest(t, http.MethodGet, "creditcard/user", nil)

		response := []creditcard_controller.GetAllUserCreditcardsResponseDto{}

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
		user := test_utils.CreateUser(roles_enum.USER)
		TSetCookieWithUser(t, user)

		newCreditcard := createCreditCard(t, user.ID)

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
		user := test_utils.CreateUser(roles_enum.USER)
		TSetCookieWithUser(t, user)

		newCreditcard := createCreditCard(t, user.ID)

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

func createCreditCard(t *testing.T, userID int32) creditcard_repository.FinCreditcard {
	creditcard, err := repository.Creditcard.Create(context.Background(), creditcard_repository.CreateParams{
		Name:             "Testing",
		Number:           "5574723384289379",
		Expiration:       time.Now().AddDate(3, 0, 0).UTC(),
		Priority:         1,
		NotifyExpiration: true,
		CreditcardTypeID: 1,
		CreditcardFlagID: 1,
		UserID:           userID,
	})
	if err != nil {
		t.Errorf("error on create creditcard: %v", err)
	}

	return creditcard
}
