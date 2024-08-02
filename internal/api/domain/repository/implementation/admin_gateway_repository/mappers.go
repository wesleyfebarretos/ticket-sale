package admin_gateway_repository

import (
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_gateway_connection"
)

func (s *CreateResponse) FromEntity(p admin_gateway_connection.FinGateway) CreateResponse {
	return CreateResponse{
		ID:               p.ID,
		Uuid:             p.Uuid,
		Name:             p.Name,
		Description:      p.Description,
		ClientID:         p.ClientID,
		ClientSecret:     p.ClientSecret,
		Order:            p.Order,
		Active:           p.Active,
		IsDeleted:        p.IsDeleted,
		TestEnvironment:  p.TestEnvironment,
		NotifUser:        p.NotifUser,
		NotifPassword:    p.NotifPassword,
		SoftDescriptor:   p.SoftDescriptor,
		GatewayProcessID: p.GatewayProcessID,
		WebhookUrl:       p.WebhookUrl,
		Url:              p.Url,
		AuthType:         string(p.AuthType),
		Use3ds:           p.Use3ds,
		AdqCode3ds:       p.AdqCode3ds,
		DefaultAdqCode:   p.DefaultAdqCode,
		UseAntifraud:     p.UseAntifraud,
		CreatedBy:        p.CreatedBy,
		UpdatedBy:        p.UpdatedBy,
		CreatedAt:        p.CreatedAt,
		UpdatedAt:        p.UpdatedAt,
	}
}

func (s *CreateParams) ToEntity() admin_gateway_connection.CreateParams {
	return admin_gateway_connection.CreateParams{
		Name:             s.Name,
		Description:      s.Description,
		ClientID:         s.ClientID,
		ClientSecret:     s.ClientSecret,
		Order:            s.Order,
		Active:           s.Active,
		TestEnvironment:  s.TestEnvironment,
		NotifUser:        s.NotifUser,
		NotifPassword:    s.NotifPassword,
		SoftDescriptor:   s.SoftDescriptor,
		GatewayProcessID: s.GatewayProcessID,
		WebhookUrl:       s.WebhookUrl,
		Url:              s.Url,
		AuthType:         admin_gateway_connection.GatewayAuthType(s.AuthType),
		Use3ds:           s.Use3ds,
		AdqCode3ds:       s.AdqCode3ds,
		DefaultAdqCode:   s.DefaultAdqCode,
		UseAntifraud:     s.UseAntifraud,
		CreatedBy:        s.CreatedBy,
		UpdatedBy:        s.UpdatedBy,
	}
}

func (s *UpdateParams) ToEntity() admin_gateway_connection.UpdateParams {
	return admin_gateway_connection.UpdateParams{
		ID:               s.ID,
		Name:             s.Name,
		Description:      s.Description,
		ClientID:         s.ClientID,
		ClientSecret:     s.ClientSecret,
		Order:            s.Order,
		Active:           s.Active,
		TestEnvironment:  s.TestEnvironment,
		NotifUser:        s.NotifUser,
		NotifPassword:    s.NotifPassword,
		SoftDescriptor:   s.SoftDescriptor,
		GatewayProcessID: s.GatewayProcessID,
		WebhookUrl:       s.WebhookUrl,
		Url:              s.Url,
		AuthType:         admin_gateway_connection.GatewayAuthType(s.AuthType),
		Use3ds:           s.Use3ds,
		AdqCode3ds:       s.AdqCode3ds,
		DefaultAdqCode:   s.DefaultAdqCode,
		UseAntifraud:     s.UseAntifraud,
		UpdatedBy:        s.UpdatedBy,
	}
}

func (_ *GetAllResponse) FromEntity(p []admin_gateway_connection.GatewayDetail) []GetAllResponse {
	r := []GetAllResponse{}

	for _, v := range p {
		paymentTypes := []gatewayPaymentTypes{}

		for _, pt := range v.GatewayPaymentTypes {
			paymentTypes = append(paymentTypes, gatewayPaymentTypes{
				Name: pt.Name,
				ID:   pt.ID,
			})
		}

		r = append(r, GetAllResponse{
			ID:               v.ID,
			Uuid:             v.Uuid,
			Name:             v.Name,
			Description:      v.Description,
			ClientID:         v.ClientID,
			ClientSecret:     v.ClientSecret,
			Order:            v.Order,
			Active:           v.Active,
			TestEnvironment:  v.TestEnvironment,
			NotifUser:        v.NotifUser,
			NotifPassword:    v.NotifPassword,
			SoftDescriptor:   v.SoftDescriptor,
			GatewayProcessID: v.GatewayProcessID,
			WebhookUrl:       v.WebhookUrl,
			Url:              v.Url,
			AuthType:         string(v.AuthType),
			Use3ds:           v.Use3ds,
			AdqCode3ds:       v.AdqCode3ds,
			DefaultAdqCode:   v.DefaultAdqCode,
			UseAntifraud:     v.UseAntifraud,
			CreatedBy:        v.CreatedBy,
			UpdatedBy:        v.UpdatedBy,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
			GatewayProcess: gatewayProcess{
				Name: v.GatewayProcess.Name,
				ID:   v.GatewayProcessID,
			},
			GatewayPaymentTypes: paymentTypes,
		})

	}

	return r
}

func (s *SoftDeleteParams) ToEntity() admin_gateway_connection.SoftDeleteParams {
	return admin_gateway_connection.SoftDeleteParams{
		ID:        s.ID,
		UpdatedBy: s.UpdatedBy,
	}
}

func (_ *GetOneByIdResponse) FromEntity(p admin_gateway_connection.GatewayDetail) *GetOneByIdResponse {
	paymentTypes := []gatewayPaymentTypes{}

	for _, pt := range p.GatewayPaymentTypes {
		paymentTypes = append(paymentTypes, gatewayPaymentTypes{
			Name: pt.Name,
			ID:   pt.ID,
		})
	}

	return &GetOneByIdResponse{
		ID:               p.ID,
		Uuid:             p.Uuid,
		Name:             p.Name,
		Description:      p.Description,
		ClientID:         p.ClientID,
		ClientSecret:     p.ClientSecret,
		Order:            p.Order,
		Active:           p.Active,
		TestEnvironment:  p.TestEnvironment,
		NotifUser:        p.NotifUser,
		NotifPassword:    p.NotifPassword,
		SoftDescriptor:   p.SoftDescriptor,
		GatewayProcessID: p.GatewayProcessID,
		WebhookUrl:       p.WebhookUrl,
		Url:              p.Url,
		AuthType:         string(p.AuthType),
		Use3ds:           p.Use3ds,
		AdqCode3ds:       p.AdqCode3ds,
		DefaultAdqCode:   p.DefaultAdqCode,
		UseAntifraud:     p.UseAntifraud,
		CreatedBy:        p.CreatedBy,
		UpdatedBy:        p.UpdatedBy,
		CreatedAt:        p.CreatedAt,
		UpdatedAt:        p.UpdatedAt,
		GatewayProcess: gatewayProcess{
			Name: p.GatewayProcess.Name,
			ID:   p.GatewayProcessID,
		},
		GatewayPaymentTypes: paymentTypes,
	}
}

func (s *CreatePaymentTypesParams) ToEntity() admin_gateway_connection.CreatePaymentTypesParams {
	return admin_gateway_connection.CreatePaymentTypesParams{
		GatewayID:            s.GatewayID,
		GatewayPaymentTypeID: s.GatewayPaymentTypeID,
		CreatedBy:            s.CreatedBy,
		UpdatedBy:            s.UpdatedBy,
	}
}

func (s *CreatePaymentTypesResponse) FromEntity(p admin_gateway_connection.FinGatewayPaymentTypeAssociation) CreatePaymentTypesResponse {
	return CreatePaymentTypesResponse{
		ID:                   p.ID,
		GatewayID:            p.GatewayID,
		GatewayPaymentTypeID: p.GatewayPaymentTypeID,
		CreatedBy:            p.CreatedBy,
		UpdatedBy:            p.UpdatedBy,
		CreatedAt:            p.CreatedAt,
		UpdatedAt:            p.UpdatedAt,
	}
}
