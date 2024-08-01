package admin_gateway_controller

import "github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_gateway_repository"

func createDtoToDomain(body CreateReqDTO) admin_gateway_repository.CreateParams {
	return admin_gateway_repository.CreateParams{
		Name:             body.Name,
		Description:      body.Description,
		ClientID:         body.ClientID,
		ClientSecret:     body.ClientSecret,
		Order:            body.Order,
		Active:           body.Active,
		TestEnvironment:  body.TestEnvironment,
		NotifUser:        body.NotifUser,
		NotifPassword:    body.NotifPassword,
		SoftDescriptor:   body.SoftDescriptor,
		GatewayProcessID: body.GatewayProcessID,
		WebhookUrl:       body.WebhookUrl,
		Url:              body.Url,
		AuthType:         body.AuthType,
		Use3ds:           body.Use3ds,
		AdqCode3ds:       body.AdqCode3ds,
		DefaultAdqCode:   body.DefaultAdqCode,
		UseAntifraud:     body.UseAntifraud,
		CreatedBy:        body.CreatedBy,
		UpdatedBy:        body.UpdatedBy,
	}
}

func createDomainToDto(body admin_gateway_repository.CreateResponse) CreateResDTO {
	return CreateResDTO{
		ID:               body.ID,
		Uuid:             body.Uuid,
		Name:             body.Name,
		Description:      body.Description,
		ClientID:         body.ClientID,
		ClientSecret:     body.ClientSecret,
		Order:            body.Order,
		Active:           body.Active,
		IsDeleted:        body.IsDeleted,
		TestEnvironment:  body.TestEnvironment,
		NotifUser:        body.NotifUser,
		NotifPassword:    body.NotifPassword,
		SoftDescriptor:   body.SoftDescriptor,
		GatewayProcessID: body.GatewayProcessID,
		WebhookUrl:       body.WebhookUrl,
		Url:              body.Url,
		AuthType:         body.AuthType,
		Use3ds:           body.Use3ds,
		AdqCode3ds:       body.AdqCode3ds,
		DefaultAdqCode:   body.DefaultAdqCode,
		UseAntifraud:     body.UseAntifraud,
		CreatedBy:        body.CreatedBy,
		UpdatedBy:        body.UpdatedBy,
		CreatedAt:        body.CreatedAt,
		UpdatedAt:        body.UpdatedAt,
	}
}

func updateDtoToDomain(body UpdateReqDTO) admin_gateway_repository.UpdateParams {
	return admin_gateway_repository.UpdateParams{
		ID:               body.ID,
		Name:             body.Name,
		Description:      body.Description,
		ClientID:         body.ClientID,
		ClientSecret:     body.ClientSecret,
		Order:            body.Order,
		Active:           body.Active,
		TestEnvironment:  body.TestEnvironment,
		NotifUser:        body.NotifUser,
		NotifPassword:    body.NotifPassword,
		SoftDescriptor:   body.SoftDescriptor,
		GatewayProcessID: body.GatewayProcessID,
		WebhookUrl:       body.WebhookUrl,
		Url:              body.Url,
		AuthType:         body.AuthType,
		Use3ds:           body.Use3ds,
		AdqCode3ds:       body.AdqCode3ds,
		DefaultAdqCode:   body.DefaultAdqCode,
		UseAntifraud:     body.UseAntifraud,
		UpdatedBy:        body.UpdatedBy,
	}
}

func getAllDomainToDto(p []admin_gateway_repository.GetAllResponse) []GetAllResDTO {
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

func getOneByIdDomainToDto(p *admin_gateway_repository.GetOneByIdResponse) *GetOneByIdResDTO {
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

func SoftDeleteDtoToDomain(p SoftDeleteDTO) admin_gateway_repository.SoftDeleteParams {
	return admin_gateway_repository.SoftDeleteParams{
		ID:        p.ID,
		UpdatedBy: &p.UpdatedBy,
	}
}
