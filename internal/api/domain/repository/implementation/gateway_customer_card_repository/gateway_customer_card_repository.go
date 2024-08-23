package gateway_customer_card_repository

import (
	"sync"

	"github.com/jackc/pgx/v4"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/gateway_customer_card_connection"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

type GatewayCustomerCardRepository struct {
	queries *gateway_customer_card_connection.Queries
}

var (
	once       sync.Once
	repository *GatewayCustomerCardRepository
)

func New() *GatewayCustomerCardRepository {
	once.Do(func() {
		repository = &GatewayCustomerCardRepository{
			queries: gateway_customer_card_connection.New(db.Conn),
		}
	})
	return repository
}

func (this *GatewayCustomerCardRepository) WithTx(tx pgx.Tx) *GatewayCustomerCardRepository {
	return &GatewayCustomerCardRepository{
		queries: this.queries.WithTx(tx),
	}
}
