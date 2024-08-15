package admin_gateway_controller

import (
	"time"

	"github.com/google/uuid"
)

type CreateReqDTO struct {
	Name              string  `json:"name" binding:"required,min=3" example:"Stripe"`
	Description       *string `json:"description" example:"Payment gateway for Stripe"`
	ClientID          *string `json:"clientId" example:"ZYDPLLBWSK3MVQJSIYHB1OR2JXCY0X2C5UJ2QAR2MAAIT5Q"`
	ClientSecret      *string `json:"clientSecret" example:"52caa7c6-d107-42bd-afac-0226a501e66e"`
	Order             int32   `json:"order" binding:"required,min=0" example:"1"`
	Active            bool    `json:"active" example:"true"`
	TestEnvironment   bool    `json:"testEnvironment" example:"false"`
	NotifUser         *string `json:"notifUser" example:"notification_user"`
	NotifPassword     *string `json:"notifPassword" example:"notification_password"`
	SoftDescriptor    *string `json:"softDescriptor" example:"Company Name"`
	GatewayProcessID  int32   `json:"gatewayProcessId" binding:"required,min=1" example:"1"`
	WebhookUrl        *string `json:"webhookUrl" example:"https://example.com/webhook"`
	Url               *string `json:"url" example:"https://api.stripe.com"`
	AuthType          string  `json:"authType" binding:"required,oneof=bearer basic" enums:"bearer,basic" example:"bearer"`
	Use3ds            bool    `json:"use3ds" example:"false"`
	AdqCode3ds        *string `json:"adqCode3ds" example:"123456"`
	DefaultAdqCode    *string `json:"defaultAdqCode" example:"654321"`
	UseAntifraud      bool    `json:"useAntifraud" example:"true"`
	CreatedBy         int32   `json:"-"`
	UpdatedBy         *int32  `json:"-"`
	PaymentTypes      []int32 `json:"paymentTypes" binding:"required,min=1,dive" example:"1,2"`
	GatewayProviderID int32   `json:"gatewayProviderId" binding:"required,min=1" example:"1"`
}

type CreatePaymentTypeResDTO struct {
	ID                   int32     `json:"id" example:"1"`
	GatewayID            int32     `json:"gatewayId" example:"1"`
	GatewayPaymentTypeID int32     `json:"gatewayPaymentTypeId" example:"1"`
	CreatedBy            int32     `json:"-"`
	UpdatedBy            *int32    `json:"-"`
	CreatedAt            time.Time `json:"-"`
	UpdatedAt            time.Time `json:"-"`
}

type CreateResDTO struct {
	ID                int32                     `json:"id" example:"1"`
	Uuid              uuid.UUID                 `json:"uuid" example:"78e5259f-1b4e-460c-90d2-83e640a7d024"`
	Name              string                    `json:"name" example:"Stripe"`
	Description       *string                   `json:"description" example:"Payment gateway for Stripe"`
	ClientID          *string                   `json:"clientId" example:"ZYDPLLBWSK3MVQJSIYHB1OR2JXCY0X2C5UJ2QAR2MAAIT5Q"`
	ClientSecret      *string                   `json:"clientSecret" example:"52caa7c6-d107-42bd-afac-0226a501e66e"`
	Order             int32                     `json:"order" example:"1"`
	Active            bool                      `json:"active" example:"true"`
	TestEnvironment   bool                      `json:"testEnvironment" example:"false"`
	NotifUser         *string                   `json:"notifUser" example:"notification_user"`
	NotifPassword     *string                   `json:"notifPassword" example:"notification_password"`
	SoftDescriptor    *string                   `json:"softDescriptor" example:"Company Name"`
	GatewayProcessID  int32                     `json:"gatewayProcessId" example:"1"`
	WebhookUrl        *string                   `json:"webhookUrl" example:"https://example.com/webhook"`
	Url               *string                   `json:"url" example:"https://api.stripe.com"`
	AuthType          string                    `json:"authType" example:"bearer"`
	Use3ds            bool                      `json:"use3ds" example:"false"`
	AdqCode3ds        *string                   `json:"adqCode3ds" example:"123456"`
	DefaultAdqCode    *string                   `json:"defaultAdqCode" example:"654321"`
	UseAntifraud      bool                      `json:"useAntifraud" example:"true"`
	IsDeleted         bool                      `json:"isDeleted" example:"false"`
	CreatedBy         int32                     `json:"-"`
	UpdatedBy         *int32                    `json:"-"`
	CreatedAt         time.Time                 `json:"createdAt" example:"2023-01-01T00:00:00Z"`
	UpdatedAt         time.Time                 `json:"updatedAt" example:"2023-01-01T00:00:00Z"`
	PaymentTypes      []CreatePaymentTypeResDTO `json:"paymentTypes"`
	GatewayProviderID int32                     `json:"gatewayProviderId" example:"1"`
}

type UpdateReqDTO struct {
	Name              string  `json:"name" binding:"required,min=3" example:"Stripe"`
	Description       *string `json:"description" example:"Payment gateway for Stripe"`
	ClientID          *string `json:"clientId" example:"ZYDPLLBWSK3MVQJSIYHB1OR2JXCY0X2C5UJ2QAR2MAAIT5Q"`
	ClientSecret      *string `json:"clientSecret" example:"52caa7c6-d107-42bd-afac-0226a501e66e"`
	Order             int32   `json:"order" binding:"required,min=0" example:"1"`
	Active            bool    `json:"active" example:"true"`
	TestEnvironment   bool    `json:"testEnvironment" example:"false"`
	NotifUser         *string `json:"notifUser" example:"notification_user"`
	NotifPassword     *string `json:"notifPassword" example:"notification_password"`
	SoftDescriptor    *string `json:"softDescriptor" example:"Company Name"`
	GatewayProcessID  int32   `json:"gatewayProcessId" binding:"required,min=1" example:"1"`
	WebhookUrl        *string `json:"webhookUrl" example:"https://example.com/webhook"`
	Url               *string `json:"url" example:"https://api.stripe.com"`
	AuthType          string  `json:"authType" binding:"required,oneof=bearer basic" enums:"bearer,basic" example:"bearer"`
	Use3ds            bool    `json:"use3ds" example:"false"`
	AdqCode3ds        *string `json:"adqCode3ds" example:"123456"`
	DefaultAdqCode    *string `json:"defaultAdqCode" example:"654321"`
	UseAntifraud      bool    `json:"useAntifraud" example:"true"`
	GatewayProviderID int32   `json:"gatewayProviderId" binding:"required,min=1" example:"1"`
}

type GetAllResDTO struct {
	ID                  int32                    `json:"id" example:"1"`
	Uuid                uuid.UUID                `json:"uuid" example:"78e5259f-1b4e-460c-90d2-83e640a7d024"`
	Name                string                   `json:"name" example:"Stripe"`
	Description         *string                  `json:"description" example:"Payment gateway for Stripe"`
	ClientID            *string                  `json:"clientId" example:"ZYDPLLBWSK3MVQJSIYHB1OR2JXCY0X2C5UJ2QAR2MAAIT5Q"`
	ClientSecret        *string                  `json:"clientSecret" example:"52caa7c6-d107-42bd-afac-0226a501e66e"`
	Order               int32                    `json:"order" example:"1"`
	Active              bool                     `json:"active" example:"true"`
	TestEnvironment     bool                     `json:"testEnvironment" example:"false"`
	NotifUser           *string                  `json:"notifUser" example:"notification_user"`
	NotifPassword       *string                  `json:"notifPassword" example:"notification_password"`
	SoftDescriptor      *string                  `json:"softDescriptor" example:"Company Name"`
	GatewayProcessID    int32                    `json:"gatewayProcessId" example:"1"`
	WebhookUrl          *string                  `json:"webhookUrl" example:"https://example.com/webhook"`
	Url                 *string                  `json:"url" example:"https://api.stripe.com"`
	AuthType            string                   `json:"authType" example:"bearer"`
	Use3ds              bool                     `json:"use3ds" example:"false"`
	AdqCode3ds          *string                  `json:"adqCode3ds" example:"123456"`
	DefaultAdqCode      *string                  `json:"defaultAdqCode" example:"654321"`
	UseAntifraud        bool                     `json:"useAntifraud" example:"true"`
	CreatedBy           int32                    `json:"createdBy" example:"1"`
	UpdatedBy           *int32                   `json:"updatedBy" example:"1"`
	CreatedAt           time.Time                `json:"createdAt" example:"2023-01-01T00:00:00Z"`
	UpdatedAt           time.Time                `json:"updatedAt" example:"2023-01-01T00:00:00Z"`
	GatewayProcess      GatewayProcessDTO        `json:"gatewayProcess"`
	GatewayPaymentTypes []GatewayPaymentTypesDTO `json:"gatewayPaymentTypes"`
}

type GetOneByIdResDTO struct {
	ID                  int32                    `json:"id" example:"1"`
	Uuid                uuid.UUID                `json:"uuid" example:"78e5259f-1b4e-460c-90d2-83e640a7d024"`
	Name                string                   `json:"name" example:"Stripe"`
	Description         *string                  `json:"description" example:"Payment gateway for Stripe"`
	ClientID            *string                  `json:"clientId" example:"ZYDPLLBWSK3MVQJSIYHB1OR2JXCY0X2C5UJ2QAR2MAAIT5Q"`
	ClientSecret        *string                  `json:"clientSecret" example:"52caa7c6-d107-42bd-afac-0226a501e66e"`
	Order               int32                    `json:"order" example:"1"`
	Active              bool                     `json:"active" example:"true"`
	TestEnvironment     bool                     `json:"testEnvironment" example:"false"`
	NotifUser           *string                  `json:"notifUser" example:"notification_user"`
	NotifPassword       *string                  `json:"notifPassword" example:"notification_password"`
	SoftDescriptor      *string                  `json:"softDescriptor" example:"Company Name"`
	GatewayProcessID    int32                    `json:"gatewayProcessId" example:"1"`
	WebhookUrl          *string                  `json:"webhookUrl" example:"https://example.com/webhook"`
	Url                 *string                  `json:"url" example:"https://api.stripe.com"`
	AuthType            string                   `json:"authType" example:"bearer"`
	Use3ds              bool                     `json:"use3ds" example:"false"`
	AdqCode3ds          *string                  `json:"adqCode3ds" example:"123456"`
	DefaultAdqCode      *string                  `json:"defaultAdqCode" example:"654321"`
	UseAntifraud        bool                     `json:"useAntifraud" example:"true"`
	CreatedBy           int32                    `json:"createdBy" example:"1"`
	UpdatedBy           *int32                   `json:"updatedBy" example:"1"`
	CreatedAt           time.Time                `json:"createdAt" example:"2023-01-01T00:00:00Z"`
	UpdatedAt           time.Time                `json:"updatedAt" example:"2023-01-01T00:00:00Z"`
	GatewayProcess      GatewayProcessDTO        `json:"gatewayProcess"`
	GatewayPaymentTypes []GatewayPaymentTypesDTO `json:"gatewayPaymentTypes"`
}

type SoftDeleteDTO struct {
	ID        int32 `json:"id"`
	UpdatedBy int32 `json:"updatedBy"`
}

type GatewayProcessDTO struct {
	Name string `json:"name" example:"Authorize Only"`
	ID   int32  `json:"id" example:"2"`
}

type GatewayPaymentTypesDTO struct {
	Name string `json:"name" example:"Installment"`
	ID   int32  `json:"id" example:"1"`
}
