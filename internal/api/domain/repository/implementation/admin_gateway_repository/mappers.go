package admin_gateway_repository

import (
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_gateway_connection"
)

func (this *CreateResponse) FromEntity(p admin_gateway_connection.FinGateway) CreateResponse {
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

func (this *CreateParams) ToEntity() admin_gateway_connection.CreateParams {
	return admin_gateway_connection.CreateParams{
		Name:             this.Name,
		Description:      this.Description,
		ClientID:         this.ClientID,
		ClientSecret:     this.ClientSecret,
		Order:            this.Order,
		Active:           this.Active,
		TestEnvironment:  this.TestEnvironment,
		NotifUser:        this.NotifUser,
		NotifPassword:    this.NotifPassword,
		SoftDescriptor:   this.SoftDescriptor,
		GatewayProcessID: this.GatewayProcessID,
		WebhookUrl:       this.WebhookUrl,
		Url:              this.Url,
		AuthType:         admin_gateway_connection.GatewayAuthType(this.AuthType),
		Use3ds:           this.Use3ds,
		AdqCode3ds:       this.AdqCode3ds,
		DefaultAdqCode:   this.DefaultAdqCode,
		UseAntifraud:     this.UseAntifraud,
		CreatedBy:        this.CreatedBy,
		UpdatedBy:        this.UpdatedBy,
	}
}

func (this *UpdateParams) ToEntity() admin_gateway_connection.UpdateParams {
	return admin_gateway_connection.UpdateParams{
		ID:               this.ID,
		Name:             this.Name,
		Description:      this.Description,
		ClientID:         this.ClientID,
		ClientSecret:     this.ClientSecret,
		Order:            this.Order,
		Active:           this.Active,
		TestEnvironment:  this.TestEnvironment,
		NotifUser:        this.NotifUser,
		NotifPassword:    this.NotifPassword,
		SoftDescriptor:   this.SoftDescriptor,
		GatewayProcessID: this.GatewayProcessID,
		WebhookUrl:       this.WebhookUrl,
		Url:              this.Url,
		AuthType:         admin_gateway_connection.GatewayAuthType(this.AuthType),
		Use3ds:           this.Use3ds,
		AdqCode3ds:       this.AdqCode3ds,
		DefaultAdqCode:   this.DefaultAdqCode,
		UseAntifraud:     this.UseAntifraud,
		UpdatedBy:        this.UpdatedBy,
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

func (this *SoftDeleteParams) ToEntity() admin_gateway_connection.SoftDeleteParams {
	return admin_gateway_connection.SoftDeleteParams{
		ID:        this.ID,
		UpdatedBy: this.UpdatedBy,
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

func (this *CreatePaymentTypesParams) ToEntity() admin_gateway_connection.CreatePaymentTypesParams {
	return admin_gateway_connection.CreatePaymentTypesParams{
		GatewayID:            this.GatewayID,
		GatewayPaymentTypeID: this.GatewayPaymentTypeID,
		CreatedBy:            this.CreatedBy,
		UpdatedBy:            this.UpdatedBy,
	}
}

func (this *CreatePaymentTypesResponse) FromEntity(p admin_gateway_connection.FinGatewayPaymentTypeAssociation) CreatePaymentTypesResponse {
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
