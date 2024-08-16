package gateway_customer_repository

import (
	"sync"

	"github.com/jackc/pgx/v4"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/gateway_customer_connection"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

type GatewayCustomerRepository struct {
	queries *gateway_customer_connection.Queries
}

var (
	once       sync.Once
	repository *GatewayCustomerRepository
)

func New() *GatewayCustomerRepository {
	once.Do(func() {
		repository = &GatewayCustomerRepository{
			queries: gateway_customer_connection.New(db.Conn),
		}
	})
	return repository
}

func (this *GatewayCustomerRepository) WithTx(tx pgx.Tx) *GatewayCustomerRepository {
	return &GatewayCustomerRepository{
		queries: this.queries.WithTx(tx),
	}
}
