package admin_gateway_controller

import (
	"time"

	"github.com/google/uuid"
)

type CreateReqDTO struct {
	Name             string  `json:"name" binding:"required,min=3" example:"Stripe"`
	Description      *string `json:"description" example:"Payment gateway for Stripe"`
	ClientID         *string `json:"clientId" example:"ZYDPLLBWSK3MVQJSIYHB1OR2JXCY0X2C5UJ2QAR2MAAIT5Q"`
	ClientSecret     *string `json:"clientSecret" example:"52caa7c6-d107-42bd-afac-0226a501e66e"`
	Order            int32   `json:"order" binding:"required,min=0" example:"1"`
	Active           bool    `json:"active" example:"true"`
	TestEnvironment  bool    `json:"testEnvironment" example:"false"`
	NotifUser        *string `json:"notifUser" example:"notification_user"`
	NotifPassword    *string `json:"notifPassword" example:"notification_password"`
	SoftDescriptor   *string `json:"softDescriptor" example:"Company Name"`
	GatewayProcessID int32   `json:"gatewayProcessId" binding:"required,min=1" example:"1"`
	WebhookUrl       *string `json:"webhookUrl" example:"https://example.com/webhook"`
	Url              *string `json:"url" example:"https://api.stripe.com"`
	AuthType         string  `json:"authType" binding:"required,oneof=bearer basic" enums:"bearer,basic" example:"bearer"`
	Use3ds           bool    `json:"use3ds" example:"false"`
	AdqCode3ds       *string `json:"adqCode3ds" example:"123456"`
	DefaultAdqCode   *string `json:"defaultAdqCode" example:"654321"`
	UseAntifraud     bool    `json:"useAntifraud" example:"true"`
	CreatedBy        int32   `json:"-"`
	UpdatedBy        *int32  `json:"-"`
	PaymentTypes     []int32 `json:"paymentTypes" binding:"required,min=1,dive" example:"[1]"`
}

type CreatePaymentTypeResDTO struct {
	ID                   int32     `json:"id"`
	GatewayID            int32     `json:"gatewayId"`
	GatewayPaymentTypeID int32     `json:"gatewayPaymentTypeId"`
	CreatedBy            int32     `json:"-"`
	UpdatedBy            *int32    `json:"-"`
	CreatedAt            time.Time `json:"-"`
	UpdatedAt            time.Time `json:"-"`
}

type CreateResDTO struct {
	ID               int32                     `json:"id"`
	Uuid             uuid.UUID                 `json:"uuid"`
	Name             string                    `json:"name"`
	Description      *string                   `json:"description"`
	ClientID         *string                   `json:"clientId"`
	ClientSecret     *string                   `json:"clientSecret"`
	Order            int32                     `json:"order"`
	Active           bool                      `json:"active"`
	IsDeleted        bool                      `json:"isDeleted"`
	TestEnvironment  bool                      `json:"testEnvironment"`
	NotifUser        *string                   `json:"notifUser"`
	NotifPassword    *string                   `json:"notifPassword"`
	SoftDescriptor   *string                   `json:"softDescriptor"`
	GatewayProcessID int32                     `json:"gatewayProcessId"`
	WebhookUrl       *string                   `json:"webhookUrl"`
	Url              *string                   `json:"url"`
	AuthType         string                    `json:"authType"`
	Use3ds           bool                      `json:"use3ds"`
	AdqCode3ds       *string                   `json:"adqCode3ds"`
	DefaultAdqCode   *string                   `json:"defaultAdqCode"`
	UseAntifraud     bool                      `json:"useAntifraud"`
	CreatedBy        int32                     `json:"-"`
	UpdatedBy        *int32                    `json:"-"`
	CreatedAt        time.Time                 `json:"createdAt"`
	UpdatedAt        time.Time                 `json:"updatedAt"`
	PaymentTypes     []CreatePaymentTypeResDTO `json:"paymentTypes"`
}

type UpdateReqDTO struct {
	Name             string  `json:"name" binding:"required,min=3" example:"Stripe"`
	Description      *string `json:"description" example:"Payment gateway for Stripe"`
	ClientID         *string `json:"clientId" example:"ZYDPLLBWSK3MVQJSIYHB1OR2JXCY0X2C5UJ2QAR2MAAIT5Q"`
	ClientSecret     *string `json:"clientSecret" example:"52caa7c6-d107-42bd-afac-0226a501e66e"`
	Order            int32   `json:"order" binding:"required,min=0" example:"1"`
	Active           bool    `json:"active" example:"true"`
	TestEnvironment  bool    `json:"testEnvironment" example:"false"`
	NotifUser        *string `json:"notifUser" example:"notification_user"`
	NotifPassword    *string `json:"notifPassword" example:"notification_password"`
	SoftDescriptor   *string `json:"softDescriptor" example:"Company Name"`
	GatewayProcessID int32   `json:"gatewayProcessId" binding:"required,min=1" example:"1"`
	WebhookUrl       *string `json:"webhookUrl" example:"https://example.com/webhook"`
	Url              *string `json:"url" example:"https://api.stripe.com"`
	AuthType         string  `json:"authType" binding:"required,oneof=bearer basic" enums:"bearer,basic" example:"bearer"`
	Use3ds           bool    `json:"use3ds" example:"false"`
	AdqCode3ds       *string `json:"adqCode3ds" example:"123456"`
	DefaultAdqCode   *string `json:"defaultAdqCode" example:"654321"`
	UseAntifraud     bool    `json:"useAntifraud" example:"true"`
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
	ID        int32 `json:"id"`
	UpdatedBy int32 `json:"updatedBy"`
}

type GatewayProcessDTO struct {
	Name string `json:"name"`
	ID   int32  `json:"id"`
}

type GatewayPaymentTypesDTO struct {
	Name string `json:"name"`
	ID   int32  `json:"id"`
}
