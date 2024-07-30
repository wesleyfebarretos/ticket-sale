package admin_gateway_repository

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_gateway_connection"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

//  TODO: Thinking in some way to make repository structs works well with
//  IO objects in a natural way

type AdminProductRepository struct {
	queries *admin_gateway_connection.Queries
}

var (
	once       sync.Once
	repository *AdminProductRepository
)

func New() *AdminProductRepository {
	once.Do(func() {
		repository = &AdminProductRepository{
			queries: admin_gateway_connection.New(db.Conn),
		}
	})
	return repository
}

func (r *AdminProductRepository) WithTx(tx pgx.Tx) *AdminProductRepository {
	return &AdminProductRepository{
		queries: r.queries.WithTx(tx),
	}
}

func (r *AdminProductRepository) Create(ctx context.Context, createParams createParams) createResponse {
	newGateway, err := r.queries.Create(ctx, admin_gateway_connection.CreateParams{
		Name:             createParams.Name,
		Description:      createParams.Description,
		ClientID:         createParams.ClientID,
		ClientSecret:     createParams.ClientSecret,
		Order:            createParams.Order,
		Active:           createParams.Active,
		TestEnvironment:  createParams.TestEnvironment,
		NotifUser:        createParams.NotifUser,
		NotifPassword:    createParams.NotifPassword,
		SoftDescriptor:   createParams.SoftDescriptor,
		GatewayProcessID: createParams.GatewayProcessID,
		WebhookUrl:       createParams.WebhookUrl,
		Url:              createParams.Url,
		AuthType:         admin_gateway_connection.GatewayAuthType(createParams.AuthType),
		Use3ds:           createParams.Use3ds,
		AdqCode3ds:       createParams.AdqCode3ds,
		DefaultAdqCode:   createParams.DefaultAdqCode,
		UseAntifraud:     createParams.UseAntifraud,
		CreatedBy:        createParams.CreatedBy,
		UpdatedBy:        createParams.UpdatedBy,
	})
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return createResponse{
		ID:               newGateway.ID,
		Uuid:             newGateway.Uuid,
		Name:             newGateway.Name,
		Description:      newGateway.Description,
		ClientID:         newGateway.ClientID,
		ClientSecret:     newGateway.ClientSecret,
		Order:            newGateway.Order,
		Active:           newGateway.Active,
		IsDeleted:        newGateway.IsDeleted,
		TestEnvironment:  newGateway.TestEnvironment,
		NotifUser:        newGateway.NotifUser,
		NotifPassword:    newGateway.NotifPassword,
		SoftDescriptor:   newGateway.SoftDescriptor,
		GatewayProcessID: newGateway.GatewayProcessID,
		WebhookUrl:       newGateway.WebhookUrl,
		Url:              newGateway.Url,
		AuthType:         string(newGateway.AuthType),
		Use3ds:           newGateway.Use3ds,
		AdqCode3ds:       newGateway.AdqCode3ds,
		DefaultAdqCode:   newGateway.DefaultAdqCode,
		UseAntifraud:     newGateway.UseAntifraud,
		CreatedBy:        newGateway.CreatedBy,
		UpdatedBy:        newGateway.UpdatedBy,
		CreatedAt:        newGateway.CreatedAt,
		UpdatedAt:        newGateway.UpdatedAt,
	}
}

func (r *AdminProductRepository) Update(ctx context.Context, updateParams updateParams) bool {
	err := r.queries.Update(ctx, admin_gateway_connection.UpdateParams{
		ID:               updateParams.ID,
		Name:             updateParams.Name,
		Description:      updateParams.Description,
		ClientID:         updateParams.ClientID,
		ClientSecret:     updateParams.ClientSecret,
		Order:            updateParams.Order,
		Active:           updateParams.Active,
		TestEnvironment:  updateParams.TestEnvironment,
		NotifUser:        updateParams.NotifUser,
		NotifPassword:    updateParams.NotifPassword,
		SoftDescriptor:   updateParams.SoftDescriptor,
		GatewayProcessID: updateParams.GatewayProcessID,
		WebhookUrl:       updateParams.WebhookUrl,
		Url:              updateParams.Url,
		AuthType:         admin_gateway_connection.GatewayAuthType(updateParams.AuthType),
		Use3ds:           updateParams.Use3ds,
		AdqCode3ds:       updateParams.AdqCode3ds,
		DefaultAdqCode:   updateParams.DefaultAdqCode,
		UseAntifraud:     updateParams.UseAntifraud,
		UpdatedBy:        updateParams.UpdatedBy,
	})
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return true
}

func (r *AdminProductRepository) SoftDelete(ctx context.Context, id int32) (bool, error) {
	err := r.queries.SoftDelete(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *AdminProductRepository) GetOneById(ctx context.Context, id int32) *getOneByIdResponse {
	gateway, err := r.queries.GetOneById(ctx, id)

	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	paymentTypes := []gatewayPaymentTypes{}

	for _, v := range gateway.GatewayPaymentTypes {
		paymentTypes = append(paymentTypes, gatewayPaymentTypes(v))
	}

	return &getOneByIdResponse{
		ID:                  gateway.ID,
		Uuid:                gateway.Uuid,
		Name:                gateway.Name,
		Description:         gateway.Description,
		ClientID:            gateway.ClientID,
		ClientSecret:        gateway.ClientSecret,
		Order:               gateway.Order,
		Active:              gateway.Active,
		TestEnvironment:     gateway.TestEnvironment,
		NotifUser:           gateway.NotifUser,
		NotifPassword:       gateway.NotifPassword,
		SoftDescriptor:      gateway.SoftDescriptor,
		GatewayProcessID:    gateway.GatewayProcessID,
		WebhookUrl:          gateway.WebhookUrl,
		Url:                 gateway.Url,
		AuthType:            string(gateway.AuthType),
		Use3ds:              gateway.Use3ds,
		AdqCode3ds:          gateway.AdqCode3ds,
		DefaultAdqCode:      gateway.DefaultAdqCode,
		UseAntifraud:        gateway.UseAntifraud,
		CreatedBy:           gateway.CreatedBy,
		UpdatedBy:           gateway.UpdatedBy,
		CreatedAt:           gateway.CreatedAt,
		UpdatedAt:           gateway.UpdatedAt,
		GatewayProcess:      gatewayProcess(gateway.GatewayProcess),
		GatewayPaymentTypes: paymentTypes,
	}
}

func (r *AdminProductRepository) GetAll(ctx context.Context) []getAllResponse {
	gatewaysResponse := []getAllResponse{}

	gateways, err := r.queries.GetAll(ctx)
	if err == pgx.ErrNoRows {
		return gatewaysResponse
	}
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	for _, v := range gateways {
		paymentTypes := []gatewayPaymentTypes{}

		for _, y := range v.GatewayPaymentTypes {
			paymentTypes = append(paymentTypes, gatewayPaymentTypes{
				Name: y.Name,
				ID:   y.ID,
			})
		}

		gatewaysResponse = append(gatewaysResponse, getAllResponse{
			ID:                  v.ID,
			Uuid:                v.Uuid,
			Name:                v.Name,
			Description:         v.Description,
			ClientID:            v.ClientID,
			ClientSecret:        v.ClientSecret,
			Order:               v.Order,
			Active:              v.Active,
			TestEnvironment:     v.TestEnvironment,
			NotifUser:           v.NotifUser,
			NotifPassword:       v.NotifPassword,
			SoftDescriptor:      v.SoftDescriptor,
			GatewayProcessID:    v.GatewayProcessID,
			WebhookUrl:          v.WebhookUrl,
			Url:                 v.Url,
			AuthType:            string(v.AuthType),
			Use3ds:              v.Use3ds,
			AdqCode3ds:          v.AdqCode3ds,
			DefaultAdqCode:      v.DefaultAdqCode,
			UseAntifraud:        v.UseAntifraud,
			CreatedBy:           v.CreatedBy,
			UpdatedBy:           v.UpdatedBy,
			CreatedAt:           v.CreatedAt,
			UpdatedAt:           v.UpdatedAt,
			GatewayProcess:      gatewayProcess(v.GatewayProcess),
			GatewayPaymentTypes: paymentTypes,
		})
	}

	return gatewaysResponse
}
