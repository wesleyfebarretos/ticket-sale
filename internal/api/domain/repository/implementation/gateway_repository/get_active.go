package gateway_repository

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/gateway_connection"
)

type GetActiveResponse struct {
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	AdqCode3ds        *string   `json:"adqCode3ds"`
	Url               *string   `json:"url"`
	ClientID          *string   `json:"clientId"`
	ClientSecret      *string   `json:"clientSecret"`
	SoftDescriptor    *string   `json:"softDescriptor"`
	UpdatedBy         *int32    `json:"updatedBy"`
	DefaultAdqCode    *string   `json:"defaultAdqCode"`
	Description       *string   `json:"description"`
	NotifUser         *string   `json:"notifUser"`
	NotifPassword     *string   `json:"notifPassword"`
	WebhookUrl        *string   `json:"webhookUrl"`
	Name              string    `json:"name"`
	AuthType          string    `json:"authType"`
	ID                int32     `json:"id"`
	GatewayProviderID int32     `json:"gatewayProviderId"`
	CreatedBy         int32     `json:"createdBy"`
	GatewayProcessID  int32     `json:"gatewayProcessId"`
	Order             int32     `json:"order"`
	Uuid              uuid.UUID `json:"uuid"`
	Use3ds            bool      `json:"use3ds"`
	TestEnvironment   bool      `json:"testEnvironment"`
	IsDeleted         bool      `json:"isDeleted"`
	UseAntifraud      bool      `json:"useAntifraud"`
	Active            bool      `json:"active"`
}

func (this *GetActiveResponse) FromEntity(p gateway_connection.FinGateway) *GetActiveResponse {
	if p.ID == 0 {
		return nil
	}

	return &GetActiveResponse{
		ID:                p.ID,
		Uuid:              uuid.UUID(p.Uuid),
		Name:              p.Name,
		Description:       p.Description,
		ClientID:          p.ClientID,
		ClientSecret:      p.ClientSecret,
		Order:             p.Order,
		Active:            p.Active,
		IsDeleted:         p.IsDeleted,
		TestEnvironment:   p.TestEnvironment,
		NotifUser:         p.NotifUser,
		NotifPassword:     p.NotifPassword,
		SoftDescriptor:    p.SoftDescriptor,
		GatewayProcessID:  p.GatewayProcessID,
		WebhookUrl:        p.WebhookUrl,
		Url:               p.Url,
		AuthType:          string(p.AuthType),
		Use3ds:            p.Use3ds,
		AdqCode3ds:        p.AdqCode3ds,
		DefaultAdqCode:    p.DefaultAdqCode,
		UseAntifraud:      p.UseAntifraud,
		CreatedBy:         p.CreatedBy,
		UpdatedBy:         p.UpdatedBy,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
		GatewayProviderID: p.GatewayProviderID,
	}
}

func (this *GatewayRepository) GetActive(ctx context.Context) *GetActiveResponse {
	gateway, err := this.queries.GetActive(ctx)
	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetActiveResponse{}

	return res.FromEntity(gateway)
}
