package admin_gateway_controller

import (
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_gateway_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/admin_gateway_service"
)

func (s *CreateReqDTO) ToDomain() admin_gateway_service.CreateReq {
	paymentTypes := []admin_gateway_repository.CreatePaymentTypesParams{}

	for _, v := range s.PaymentTypes {
		paymentTypes = append(paymentTypes, admin_gateway_repository.CreatePaymentTypesParams{
			GatewayPaymentTypeID: v,
		})
	}
	return admin_gateway_service.CreateReq{
		Gateway: admin_gateway_repository.CreateParams{
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
			AuthType:         s.AuthType,
			Use3ds:           s.Use3ds,
			AdqCode3ds:       s.AdqCode3ds,
			DefaultAdqCode:   s.DefaultAdqCode,
			UseAntifraud:     s.UseAntifraud,
			CreatedBy:        s.CreatedBy,
			UpdatedBy:        s.UpdatedBy,
		},
		PaymentTypes: paymentTypes,
	}
}

func (_ *CreateResDTO) FromDomain(p admin_gateway_service.CreateRes) CreateResDTO {
	paymentTypes := []CreatePaymentTypeResDTO{}
	for _, v := range p.PaymentTypes {
		paymentTypes = append(paymentTypes, CreatePaymentTypeResDTO{
			ID:                   v.ID,
			GatewayID:            v.GatewayID,
			GatewayPaymentTypeID: v.GatewayPaymentTypeID,
			CreatedBy:            v.CreatedBy,
			UpdatedBy:            v.UpdatedBy,
			CreatedAt:            v.CreatedAt,
			UpdatedAt:            v.UpdatedAt,
		})
	}
	return CreateResDTO{
		ID:               p.Gateway.ID,
		Uuid:             p.Gateway.Uuid,
		Name:             p.Gateway.Name,
		Description:      p.Gateway.Description,
		ClientID:         p.Gateway.ClientID,
		ClientSecret:     p.Gateway.ClientSecret,
		Order:            p.Gateway.Order,
		Active:           p.Gateway.Active,
		IsDeleted:        p.Gateway.IsDeleted,
		TestEnvironment:  p.Gateway.TestEnvironment,
		NotifUser:        p.Gateway.NotifUser,
		NotifPassword:    p.Gateway.NotifPassword,
		SoftDescriptor:   p.Gateway.SoftDescriptor,
		GatewayProcessID: p.Gateway.GatewayProcessID,
		WebhookUrl:       p.Gateway.WebhookUrl,
		Url:              p.Gateway.Url,
		AuthType:         p.Gateway.AuthType,
		Use3ds:           p.Gateway.Use3ds,
		AdqCode3ds:       p.Gateway.AdqCode3ds,
		DefaultAdqCode:   p.Gateway.DefaultAdqCode,
		UseAntifraud:     p.Gateway.UseAntifraud,
		CreatedBy:        p.Gateway.CreatedBy,
		UpdatedBy:        p.Gateway.UpdatedBy,
		CreatedAt:        p.Gateway.CreatedAt,
		UpdatedAt:        p.Gateway.UpdatedAt,
		PaymentTypes:     paymentTypes,
	}
}

func (s *UpdateReqDTO) ToDomain() admin_gateway_repository.UpdateParams {
	return admin_gateway_repository.UpdateParams{
		Name:         s.Name,
		Description:  s.Description,
		ClientID:     s.ClientID,
		ClientSecret: s.ClientSecret,
		Order:        s.Order,
		Active:       s.Active,

		TestEnvironment:  s.TestEnvironment,
		NotifUser:        s.NotifUser,
		NotifPassword:    s.NotifPassword,
		SoftDescriptor:   s.SoftDescriptor,
		GatewayProcessID: s.GatewayProcessID,
		WebhookUrl:       s.WebhookUrl,
		Url:              s.Url,
		AuthType:         s.AuthType,
		Use3ds:           s.Use3ds,
		AdqCode3ds:       s.AdqCode3ds,
		DefaultAdqCode:   s.DefaultAdqCode,
		UseAntifraud:     s.UseAntifraud,
		UpdatedBy:        s.UpdatedBy,
	}
}

func (_ *GetAllResDTO) FromDomain(p []admin_gateway_repository.GetAllResponse) []GetAllResDTO {
	r := []GetAllResDTO{}

	for _, v := range p {
		paymentTypes := []GatewayPaymentTypesDTO{}

		for _, pt := range v.GatewayPaymentTypes {
			paymentTypes = append(paymentTypes, GatewayPaymentTypesDTO{
				Name: pt.Name,
				ID:   pt.ID,
			})
		}

		r = append(r, GetAllResDTO{
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
			GatewayProcess: GatewayProcessDTO{
				Name: v.GatewayProcess.Name,
				ID:   v.GatewayProcessID,
			},
			GatewayPaymentTypes: paymentTypes,
		})

	}

	return r
}

func (s *GetOneByIdResDTO) FromDomain(p *admin_gateway_repository.GetOneByIdResponse) *GetOneByIdResDTO {
	paymentTypes := []GatewayPaymentTypesDTO{}

	for _, v := range p.GatewayPaymentTypes {
		paymentTypes = append(paymentTypes, GatewayPaymentTypesDTO{
			Name: v.Name,
			ID:   v.ID,
		})
	}

	return &GetOneByIdResDTO{
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
		AuthType:         p.AuthType,
		Use3ds:           p.Use3ds,
		AdqCode3ds:       p.AdqCode3ds,
		DefaultAdqCode:   p.DefaultAdqCode,
		UseAntifraud:     p.UseAntifraud,
		CreatedBy:        p.CreatedBy,
		UpdatedBy:        p.UpdatedBy,
		CreatedAt:        p.CreatedAt,
		UpdatedAt:        p.UpdatedAt,
		GatewayProcess: GatewayProcessDTO{
			Name: p.GatewayProcess.Name,
			ID:   p.GatewayProcessID,
		},
		GatewayPaymentTypes: paymentTypes,
	}
}

func (s *SoftDeleteDTO) ToDomain() admin_gateway_repository.SoftDeleteParams {
	return admin_gateway_repository.SoftDeleteParams{
		ID:        s.ID,
		UpdatedBy: &s.UpdatedBy,
	}
}

func (_ *CreatePaymentTypeResDTO) FromDomain(p []admin_gateway_repository.CreatePaymentTypesResponse) []CreatePaymentTypeResDTO {
	paymentTypes := []CreatePaymentTypeResDTO{}

	for _, v := range p {
		paymentTypes = append(paymentTypes, CreatePaymentTypeResDTO{
			ID:                   v.ID,
			GatewayID:            v.GatewayID,
			GatewayPaymentTypeID: v.GatewayPaymentTypeID,
			CreatedBy:            v.CreatedBy,
			UpdatedBy:            v.UpdatedBy,
			CreatedAt:            v.CreatedAt,
			UpdatedAt:            v.UpdatedAt,
		})
	}

	return paymentTypes
}
