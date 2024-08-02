package test_data

import (
	"context"
	"testing"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_gateway_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/admin_gateway_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/controller/admin_gateway_controller"
)

func NewGateway(t *testing.T, userID int32) admin_gateway_controller.CreateResDTO {
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

	gateway := admin_gateway_repository.CreateParams{
		Name:             "Testing",
		Description:      &description,
		ClientID:         &clientId,
		ClientSecret:     &clientSecret,
		Order:            1,
		Active:           true,
		TestEnvironment:  false,
		NotifUser:        &notifUser,
		NotifPassword:    &notifPassword,
		SoftDescriptor:   &softDescriptor,
		GatewayProcessID: 1,
		WebhookUrl:       &webhookUrl,
		Url:              &url,
		AuthType:         "bearer",
		Use3ds:           false,
		AdqCode3ds:       &adqCode3ds,
		DefaultAdqCode:   &defaultAdqCode,
		UseAntifraud:     false,
		CreatedBy:        userID,
		UpdatedBy:        &userID,
	}

	repository := admin_gateway_repository.New()
	ctx := context.Background()

	newGateway := repository.Create(ctx, gateway)

	paymentTypes := []admin_gateway_repository.CreatePaymentTypesParams{}

	paymentTypes = append(paymentTypes, admin_gateway_repository.CreatePaymentTypesParams{
		GatewayID:            newGateway.ID,
		GatewayPaymentTypeID: 1,
		CreatedBy:            userID,
		UpdatedBy:            &userID,
	})

	newPaymentTypes := repository.CreatePaymentTypes(ctx, paymentTypes)

	response := admin_gateway_controller.CreateResDTO{}

	return response.FromDomain(admin_gateway_service.CreateRes{
		Gateway:      newGateway,
		PaymentTypes: newPaymentTypes,
	})
}
