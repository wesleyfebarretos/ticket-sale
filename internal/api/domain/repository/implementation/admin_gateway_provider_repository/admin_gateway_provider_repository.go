package admin_gateway_provider_repository

import (
	"sync"

	"github.com/jackc/pgx/v4"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_gateway_provider_connection"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

type AdminGatewayProviderRepository struct {
	queries *admin_gateway_provider_connection.Queries
}

var (
	once       sync.Once
	repository *AdminGatewayProviderRepository
)

func New() *AdminGatewayProviderRepository {
	once.Do(func() {
		repository = &AdminGatewayProviderRepository{
			queries: admin_gateway_provider_connection.New(db.Conn),
		}
	})
	return repository
}

func (this *AdminGatewayProviderRepository) WithTx(tx pgx.Tx) *AdminGatewayProviderRepository {
	return &AdminGatewayProviderRepository{
		queries: this.queries.WithTx(tx),
	}
}
