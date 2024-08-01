package admin_gateway_controller

import (
	"time"

	"github.com/google/uuid"
)

type CreateReqDTO struct {
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

type CreateResDTO struct {
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

type UpdateReqDTO struct {
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

type GetAllResDTO struct {
	ID                  int32                    `json:"id"`
	Uuid                uuid.UUID                `json:"uuid"`
	Name                string                   `json:"name"`
	Description         *string                  `json:"description"`
	ClientID            *string                  `json:"clientId"`
	ClientSecret        *string                  `json:"clientSecret"`
	Order               int32                    `json:"order"`
	Active              bool                     `json:"active"`
	TestEnvironment     bool                     `json:"testEnvironment"`
	NotifUser           *string                  `json:"notifUser"`
	NotifPassword       *string                  `json:"notifPassword"`
	SoftDescriptor      *string                  `json:"softDescriptor"`
	GatewayProcessID    int32                    `json:"gatewayProcessId"`
	WebhookUrl          *string                  `json:"webhookUrl"`
	Url                 *string                  `json:"url"`
	AuthType            string                   `json:"authType"`
	Use3ds              bool                     `json:"use3ds"`
	AdqCode3ds          *string                  `json:"adqCode3ds"`
	DefaultAdqCode      *string                  `json:"defaultAdqCode"`
	UseAntifraud        bool                     `json:"useAntifraud"`
	CreatedBy           int32                    `json:"createdBy"`
	UpdatedBy           *int32                   `json:"updatedBy"`
	CreatedAt           time.Time                `json:"createdAt"`
	UpdatedAt           time.Time                `json:"updatedAt"`
	GatewayProcess      GatewayProcessDTO        `json:"gatewayProcess"`
	GatewayPaymentTypes []GatewayPaymentTypesDTO `json:"gatewayPaymentTypes"`
}

type GetOneByIdResDTO struct {
	ID                  int32                    `json:"id"`
	Uuid                uuid.UUID                `json:"uuid"`
	Name                string                   `json:"name"`
	Description         *string                  `json:"description"`
	ClientID            *string                  `json:"clientId"`
	ClientSecret        *string                  `json:"clientSecret"`
	Order               int32                    `json:"order"`
	Active              bool                     `json:"active"`
	TestEnvironment     bool                     `json:"testEnvironment"`
	NotifUser           *string                  `json:"notifUser"`
	NotifPassword       *string                  `json:"notifPassword"`
	SoftDescriptor      *string                  `json:"softDescriptor"`
	GatewayProcessID    int32                    `json:"gatewayProcessId"`
	WebhookUrl          *string                  `json:"webhookUrl"`
	Url                 *string                  `json:"url"`
	AuthType            string                   `json:"authType"`
	Use3ds              bool                     `json:"use3ds"`
	AdqCode3ds          *string                  `json:"adqCode3ds"`
	DefaultAdqCode      *string                  `json:"defaultAdqCode"`
	UseAntifraud        bool                     `json:"useAntifraud"`
	CreatedBy           int32                    `json:"createdBy"`
	UpdatedBy           *int32                   `json:"updatedBy"`
	CreatedAt           time.Time                `json:"createdAt"`
	UpdatedAt           time.Time                `json:"updatedAt"`
	GatewayProcess      GatewayProcessDTO        `json:"gatewayProcess"`
	GatewayPaymentTypes []GatewayPaymentTypesDTO `json:"gatewayPaymentTypes"`
}

type SoftDeleteDTO struct {
	ID        int32
	UpdatedBy int32
}

type GatewayProcessDTO struct {
	Name string `json:"name"`
	ID   int32  `json:"id"`
}

type GatewayPaymentTypesDTO struct {
	Name string `json:"name"`
	ID   int32  `json:"id"`
}
