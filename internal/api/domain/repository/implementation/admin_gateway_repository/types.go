package admin_gateway_repository

import (
	"time"

	"github.com/google/uuid"
)

type CreateParams struct {
	Name              string  `json:"name"`
	Description       *string `json:"description"`
	ClientID          *string `json:"clientId"`
	ClientSecret      *string `json:"clientSecret"`
	Order             int32   `json:"order"`
	Active            bool    `json:"active"`
	TestEnvironment   bool    `json:"testEnvironment"`
	NotifUser         *string `json:"notifUser"`
	NotifPassword     *string `json:"notifPassword"`
	SoftDescriptor    *string `json:"softDescriptor"`
	GatewayProcessID  int32   `json:"gatewayProcessId"`
	WebhookUrl        *string `json:"webhookUrl"`
	Url               *string `json:"url"`
	AuthType          string  `json:"authType"`
	Use3ds            bool    `json:"use3ds"`
	AdqCode3ds        *string `json:"adqCode3ds"`
	DefaultAdqCode    *string `json:"defaultAdqCode"`
	UseAntifraud      bool    `json:"useAntifraud"`
	CreatedBy         int32   `json:"createdBy"`
	UpdatedBy         *int32  `json:"updatedBy"`
	GatewayProviderID int32   `json:"gatewayProviderId"`
}

type CreateResponse struct {
	ID                int32     `json:"id"`
	Uuid              uuid.UUID `json:"uuid"`
	Name              string    `json:"name"`
	Description       *string   `json:"description"`
	ClientID          *string   `json:"clientId"`
	ClientSecret      *string   `json:"clientSecret"`
	Order             int32     `json:"order"`
	Active            bool      `json:"active"`
	IsDeleted         bool      `json:"isDeleted"`
	TestEnvironment   bool      `json:"testEnvironment"`
	NotifUser         *string   `json:"notifUser"`
	NotifPassword     *string   `json:"notifPassword"`
	SoftDescriptor    *string   `json:"softDescriptor"`
	GatewayProcessID  int32     `json:"gatewayProcessId"`
	WebhookUrl        *string   `json:"webhookUrl"`
	Url               *string   `json:"url"`
	AuthType          string    `json:"authType"`
	Use3ds            bool      `json:"use3ds"`
	AdqCode3ds        *string   `json:"adqCode3ds"`
	DefaultAdqCode    *string   `json:"defaultAdqCode"`
	UseAntifraud      bool      `json:"useAntifraud"`
	CreatedBy         int32     `json:"createdBy"`
	UpdatedBy         *int32    `json:"updatedBy"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	GatewayProviderID int32     `json:"gatewayProviderId"`
}

type UpdateParams struct {
	ID                int32   `json:"id"`
	Name              string  `json:"name"`
	Description       *string `json:"description"`
	ClientID          *string `json:"clientId"`
	ClientSecret      *string `json:"clientSecret"`
	Order             int32   `json:"order"`
	Active            bool    `json:"active"`
	TestEnvironment   bool    `json:"testEnvironment"`
	NotifUser         *string `json:"notifUser"`
	NotifPassword     *string `json:"notifPassword"`
	SoftDescriptor    *string `json:"softDescriptor"`
	GatewayProcessID  int32   `json:"gatewayProcessId"`
	WebhookUrl        *string `json:"webhookUrl"`
	Url               *string `json:"url"`
	AuthType          string  `json:"authType"`
	Use3ds            bool    `json:"use3ds"`
	AdqCode3ds        *string `json:"adqCode3ds"`
	DefaultAdqCode    *string `json:"defaultAdqCode"`
	UseAntifraud      bool    `json:"useAntifraud"`
	UpdatedBy         *int32  `json:"updatedBy"`
	GatewayProviderID int32   `json:"gatewayProviderId"`
}

type GetAllResponse struct {
	ID                  int32                 `json:"id"`
	Uuid                uuid.UUID             `json:"uuid"`
	Name                string                `json:"name"`
	Description         *string               `json:"description"`
	ClientID            *string               `json:"clientId"`
	ClientSecret        *string               `json:"clientSecret"`
	Order               int32                 `json:"order"`
	Active              bool                  `json:"active"`
	TestEnvironment     bool                  `json:"testEnvironment"`
	NotifUser           *string               `json:"notifUser"`
	NotifPassword       *string               `json:"notifPassword"`
	SoftDescriptor      *string               `json:"softDescriptor"`
	GatewayProcessID    int32                 `json:"gatewayProcessId"`
	WebhookUrl          *string               `json:"webhookUrl"`
	Url                 *string               `json:"url"`
	AuthType            string                `json:"authType"`
	Use3ds              bool                  `json:"use3ds"`
	AdqCode3ds          *string               `json:"adqCode3ds"`
	DefaultAdqCode      *string               `json:"defaultAdqCode"`
	UseAntifraud        bool                  `json:"useAntifraud"`
	CreatedBy           int32                 `json:"createdBy"`
	UpdatedBy           *int32                `json:"updatedBy"`
	CreatedAt           time.Time             `json:"createdAt"`
	UpdatedAt           time.Time             `json:"updatedAt"`
	GatewayProcess      gatewayProcess        `json:"gatewayProcess"`
	GatewayPaymentTypes []gatewayPaymentTypes `json:"gatewayPaymentTypes"`
}

type GetOneByIdResponse struct {
	ID                  int32                 `json:"id"`
	Uuid                uuid.UUID             `json:"uuid"`
	Name                string                `json:"name"`
	Description         *string               `json:"description"`
	ClientID            *string               `json:"clientId"`
	ClientSecret        *string               `json:"clientSecret"`
	Order               int32                 `json:"order"`
	Active              bool                  `json:"active"`
	TestEnvironment     bool                  `json:"testEnvironment"`
	NotifUser           *string               `json:"notifUser"`
	NotifPassword       *string               `json:"notifPassword"`
	SoftDescriptor      *string               `json:"softDescriptor"`
	GatewayProcessID    int32                 `json:"gatewayProcessId"`
	WebhookUrl          *string               `json:"webhookUrl"`
	Url                 *string               `json:"url"`
	AuthType            string                `json:"authType"`
	Use3ds              bool                  `json:"use3ds"`
	AdqCode3ds          *string               `json:"adqCode3ds"`
	DefaultAdqCode      *string               `json:"defaultAdqCode"`
	UseAntifraud        bool                  `json:"useAntifraud"`
	CreatedBy           int32                 `json:"createdBy"`
	UpdatedBy           *int32                `json:"updatedBy"`
	CreatedAt           time.Time             `json:"createdAt"`
	UpdatedAt           time.Time             `json:"updatedAt"`
	GatewayProcess      gatewayProcess        `json:"gatewayProcess"`
	GatewayPaymentTypes []gatewayPaymentTypes `json:"gatewayPaymentTypes"`
}

type SoftDeleteParams struct {
	ID        int32  `json:"id"`
	UpdatedBy *int32 `json:"updatedBy"`
}

type gatewayProcess struct {
	Name string `json:"name"`
	ID   int32  `json:"id"`
}

type gatewayPaymentTypes struct {
	Name string `json:"name"`
	ID   int32  `json:"id"`
}

type CreatePaymentTypesParams struct {
	GatewayID            int32  `json:"gatewayId"`
	GatewayPaymentTypeID int32  `json:"gatewayPaymentTypeId"`
	CreatedBy            int32  `json:"createdBy"`
	UpdatedBy            *int32 `json:"updatedBy"`
}

type CreatePaymentTypesResponse struct {
	ID                   int32     `json:"id"`
	GatewayID            int32     `json:"gatewayId"`
	GatewayPaymentTypeID int32     `json:"gatewayPaymentTypeId"`
	CreatedBy            int32     `json:"createdBy"`
	UpdatedBy            *int32    `json:"updatedBy"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
