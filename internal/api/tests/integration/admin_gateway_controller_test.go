package integration_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/controller/admin_gateway_controller"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/tests/test_data"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/tests/test_utils"
)

func TestAdminGatewayController(t *testing.T) {
	t.Run("it should create a gateway", TRun(func(t *testing.T) {
		adminUser := test_data.NewUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		description := "testing"
		clientId := "testing"
		clientSecret := "testing"
		notifUser := "testing"
		notifPassword := "testing"
		webhookUrl := "testing"
		softDescriptor := "testing"
		url := "testing"
		adqCode3ds := "testing"
		defaultAdqCode := "testing"

		newGateway := admin_gateway_controller.CreateReqDTO{
			Name:              "Testing",
			Description:       &description,
			ClientID:          &clientId,
			ClientSecret:      &clientSecret,
			Order:             1,
			Active:            true,
			TestEnvironment:   false,
			NotifUser:         &notifUser,
			NotifPassword:     &notifPassword,
			SoftDescriptor:    &softDescriptor,
			GatewayProcessID:  1,
			WebhookUrl:        &webhookUrl,
			Url:               &url,
			AuthType:          "bearer",
			Use3ds:            false,
			AdqCode3ds:        &adqCode3ds,
			DefaultAdqCode:    &defaultAdqCode,
			UseAntifraud:      false,
			PaymentTypes:      []int32{1, 2},
			GatewayProviderID: 1,
		}

		res := TMakeRequest(t, http.MethodPost, "admin/gateway", newGateway)

		gatewayResponse := admin_gateway_controller.CreateResDTO{}

		test_utils.Decode(t, res.Body, &gatewayResponse)

		assert.Equal(t, http.StatusCreated, res.StatusCode)
		assert.GreaterOrEqual(t, gatewayResponse.ID, int32(1))
		assert.Equal(t, newGateway.Name, gatewayResponse.Name)
		assert.Equal(t, newGateway.Description, gatewayResponse.Description)
		assert.Equal(t, newGateway.ClientID, gatewayResponse.ClientID)
		assert.Equal(t, newGateway.ClientSecret, gatewayResponse.ClientSecret)
		assert.Equal(t, newGateway.Order, gatewayResponse.Order)
		assert.Equal(t, newGateway.Active, gatewayResponse.Active)
		assert.Equal(t, newGateway.TestEnvironment, gatewayResponse.TestEnvironment)
		assert.Equal(t, newGateway.NotifUser, gatewayResponse.NotifUser)
		assert.Equal(t, newGateway.NotifPassword, gatewayResponse.NotifPassword)
		assert.Equal(t, newGateway.SoftDescriptor, gatewayResponse.SoftDescriptor)
		assert.Equal(t, newGateway.GatewayProcessID, gatewayResponse.GatewayProcessID)
		assert.Equal(t, newGateway.WebhookUrl, gatewayResponse.WebhookUrl)
		assert.Equal(t, newGateway.Url, gatewayResponse.Url)
		assert.Equal(t, newGateway.AuthType, gatewayResponse.AuthType)
		assert.Equal(t, newGateway.Use3ds, gatewayResponse.Use3ds)
		assert.Equal(t, newGateway.AdqCode3ds, gatewayResponse.AdqCode3ds)
		assert.Equal(t, newGateway.DefaultAdqCode, gatewayResponse.DefaultAdqCode)
		assert.Equal(t, newGateway.UseAntifraud, gatewayResponse.UseAntifraud)
		assert.Equal(t, len(newGateway.PaymentTypes), len(gatewayResponse.PaymentTypes))
		assert.GreaterOrEqual(t, gatewayResponse.PaymentTypes[0].ID, int32(1))
	}))

	t.Run("it should update a gateway", TRun(func(t *testing.T) {
		adminUser := test_data.NewUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newGateway := test_data.NewGateway(t, adminUser.ID)

		updateParams := admin_gateway_controller.UpdateReqDTO{
			Name:              "Updated Gateway Name",
			Description:       TPointer("Updated Description"),
			ClientID:          TPointer("new-client-id"),
			ClientSecret:      TPointer("new-client-secret"),
			Order:             2,
			Active:            true,
			TestEnvironment:   true,
			NotifUser:         TPointer("new-notif-user"),
			NotifPassword:     TPointer("new-notif-password"),
			SoftDescriptor:    TPointer("new-soft-descriptor"),
			GatewayProcessID:  2,
			WebhookUrl:        TPointer("https://new-webhook-url.com"),
			Url:               TPointer("https://new-url.com"),
			AuthType:          "basic",
			Use3ds:            true,
			AdqCode3ds:        TPointer("new-adq-code-3ds"),
			DefaultAdqCode:    TPointer("new-default-adq-code"),
			UseAntifraud:      true,
			GatewayProviderID: 1,
		}

		res := TMakeRequest(t, http.MethodPut, fmt.Sprintf("admin/gateway/%d", newGateway.ID), updateParams)

		updateResponse := false

		test_utils.Decode(t, res.Body, &updateResponse)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, true, updateResponse)
	}))

	t.Run("it should delete a gateway", TRun(func(t *testing.T) {
		adminUser := test_data.NewUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newGateway := test_data.NewGateway(t, adminUser.ID)

		res := TMakeRequest(t, http.MethodDelete, fmt.Sprintf("admin/gateway/%d", newGateway.ID), nil)

		deleteResponse := false

		test_utils.Decode(t, res.Body, &deleteResponse)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, true, deleteResponse)
	}))

	t.Run("it should get all gateways", TRun(func(t *testing.T) {
		adminUser := test_data.NewUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		test_data.NewGateway(t, adminUser.ID)

		res := TMakeRequest(t, http.MethodGet, "admin/gateway", nil)

		getAllResponse := []admin_gateway_controller.GetAllResDTO{}

		test_utils.Decode(t, res.Body, &getAllResponse)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, 1, len(getAllResponse))
		assert.IsType(t, admin_gateway_controller.GetAllResDTO{}, getAllResponse[0])
	}))

	t.Run("it should get gateway by id", TRun(func(t *testing.T) {
		adminUser := test_data.NewUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newGateway := test_data.NewGateway(t, adminUser.ID)

		res := TMakeRequest(t, http.MethodGet, fmt.Sprintf("admin/gateway/%d", newGateway.ID), nil)

		getOneByIdResponse := admin_gateway_controller.GetOneByIdResDTO{}

		test_utils.Decode(t, res.Body, &getOneByIdResponse)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, newGateway.ID, getOneByIdResponse.ID)
	}))

	t.Run("it should validate required fields on create", TRun(func(t *testing.T) {
		adminUser := test_data.NewUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		newGateway := test_data.NewGateway(t, adminUser.ID)

		bStruct, err := json.Marshal(newGateway)
		if err != nil {
			t.Fatalf("error on marshal json :%v", err)
		}

		requiredFields := []string{
			"name",
			"order",
			"gatewayProcessId",
			"authType",
			"paymentTypes",
		}

		for _, k := range requiredFields {
			structMap := make(map[string]interface{})

			json.Unmarshal(bStruct, &structMap)

			test_utils.DeleteRequiredField(t, k, structMap)

			res := TMakeRequest(t, http.MethodPost, "admin/gateway", structMap)

			assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		}
	}))

	t.Run("it should validate required fields on update", TRun(func(t *testing.T) {
		adminUser := test_data.NewUser(roles_enum.ADMIN)
		TSetCookieWithUser(t, adminUser)

		updateParams := admin_gateway_controller.UpdateReqDTO{
			Name:             "Updated Gateway Name",
			Description:      TPointer("Updated Description"),
			ClientID:         TPointer("new-client-id"),
			ClientSecret:     TPointer("new-client-secret"),
			Order:            2,
			Active:           true,
			TestEnvironment:  true,
			NotifUser:        TPointer("new-notif-user"),
			NotifPassword:    TPointer("new-notif-password"),
			SoftDescriptor:   TPointer("new-soft-descriptor"),
			GatewayProcessID: 2,
			WebhookUrl:       TPointer("https://new-webhook-url.com"),
			Url:              TPointer("https://new-url.com"),
			AuthType:         "basic",
			Use3ds:           true,
			AdqCode3ds:       TPointer("new-adq-code-3ds"),
			DefaultAdqCode:   TPointer("new-default-adq-code"),
			UseAntifraud:     true,
		}

		bStruct, err := json.Marshal(updateParams)
		if err != nil {
			t.Fatalf("error on marshal json :%v", err)
		}

		requiredFields := []string{
			"name",
			"order",
			"gatewayProcessId",
			"authType",
		}

		newGateway := test_data.NewGateway(t, adminUser.ID)

		for _, k := range requiredFields {
			structMap := make(map[string]interface{})

			json.Unmarshal(bStruct, &structMap)

			test_utils.DeleteRequiredField(t, k, structMap)

			res := TMakeRequest(t, http.MethodPut, fmt.Sprintf("admin/gateway/%d", newGateway.ID), structMap)

			assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		}
	}))

	t.Run("it should make sure that only an admin can access this routes", TRun(func(t *testing.T) {
		user := test_data.NewUser(roles_enum.USER)
		TSetCookieWithUser(t, user)

		methods := []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodGet,
			http.MethodPut,
			http.MethodDelete,
		}

		routes := []string{
			"admin/gateway",
			"admin/gateway",
			"admin/gateway/1",
			"admin/gateway/1",
			"admin/gateway/1",
		}

		for i, route := range routes {
			res := TMakeRequest(t, methods[i], route, nil)
			assert.Equal(t, http.StatusForbidden, res.StatusCode)
		}
	}))
}
