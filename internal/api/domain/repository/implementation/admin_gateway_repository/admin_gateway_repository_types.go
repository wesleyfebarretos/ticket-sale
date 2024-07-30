package admin_gateway_repository

import (
	"time"

	"github.com/google/uuid"
)

type types struct {
	CreateParams       createParams
	CreateResponse     createResponse
	UpdateParams       updateParams
	GetAllResponse     getAllResponse
	GetOneByIdResponse getOneByIdResponse
}

type createParams struct {
	Name             string  `json:"name"`
	Description      *string `json:"description"`
	ClientID         *string `json:"clientId"`
	ClientSecret     *string `json:"clientSecret"`
	Order            int32   `json:"order"`
	Active           bool    `json:"active"`
	TestEnvironment  bool    `json:"testEnvironment"`
	NotifUser        *string `json:"notifUser"`
	NotifPassword    *string `json:"notifPassword"`
	SoftDescriptor   *string `json:"softDescriptor"`
	GatewayProcessID int32   `json:"gatewayProcessId"`
	WebhookUrl       *string `json:"webhookUrl"`
	Url              *string `json:"url"`
	AuthType         string  `json:"authType"`
	Use3ds           bool    `json:"use3ds"`
	AdqCode3ds       *string `json:"adqCode3ds"`
	DefaultAdqCode   *string `json:"defaultAdqCode"`
	UseAntifraud     bool    `json:"useAntifraud"`
	CreatedBy        int32   `json:"createdBy"`
	UpdatedBy        *int32  `json:"updatedBy"`
}

type createResponse struct {
	ID               int32     `json:"id"`
	Uuid             uuid.UUID `json:"uuid"`
	Name             string    `json:"name"`
	Description      *string   `json:"description"`
	ClientID         *string   `json:"clientId"`
	ClientSecret     *string   `json:"clientSecret"`
	Order            int32     `json:"order"`
	Active           bool      `json:"active"`
	IsDeleted        bool      `json:"isDeleted"`
	TestEnvironment  bool      `json:"testEnvironment"`
	NotifUser        *string   `json:"notifUser"`
	NotifPassword    *string   `json:"notifPassword"`
	SoftDescriptor   *string   `json:"softDescriptor"`
	GatewayProcessID int32     `json:"gatewayProcessId"`
	WebhookUrl       *string   `json:"webhookUrl"`
	Url              *string   `json:"url"`
	AuthType         string    `json:"authType"`
	Use3ds           bool      `json:"use3ds"`
	AdqCode3ds       *string   `json:"adqCode3ds"`
	DefaultAdqCode   *string   `json:"defaultAdqCode"`
	UseAntifraud     bool      `json:"useAntifraud"`
	CreatedBy        int32     `json:"createdBy"`
	UpdatedBy        *int32    `json:"updatedBy"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type updateParams struct {
	ID               int32   `json:"id"`
	Name             string  `json:"name"`
	Description      *string `json:"description"`
	ClientID         *string `json:"clientId"`
	ClientSecret     *string `json:"clientSecret"`
	Order            int32   `json:"order"`
	Active           bool    `json:"active"`
	TestEnvironment  bool    `json:"testEnvironment"`
	NotifUser        *string `json:"notifUser"`
	NotifPassword    *string `json:"notifPassword"`
	SoftDescriptor   *string `json:"softDescriptor"`
	GatewayProcessID int32   `json:"gatewayProcessId"`
	WebhookUrl       *string `json:"webhookUrl"`
	Url              *string `json:"url"`
	AuthType         string  `json:"authType"`
	Use3ds           bool    `json:"use3ds"`
	AdqCode3ds       *string `json:"adqCode3ds"`
	DefaultAdqCode   *string `json:"defaultAdqCode"`
	UseAntifraud     bool    `json:"useAntifraud"`
	UpdatedBy        *int32  `json:"updatedBy"`
}

type getAllResponse struct {
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

type getOneByIdResponse struct {
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

type gatewayProcess struct {
	Name string `json:"name"`
	ID   int32  `json:"id"`
}

type gatewayPaymentTypes struct {
	Name string `json:"name"`
	ID   int32  `json:"id"`
}

var Types types
