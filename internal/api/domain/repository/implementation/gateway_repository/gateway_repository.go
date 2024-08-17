package gateway_repository

import (
	"sync"

	"github.com/jackc/pgx/v4"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/gateway_connection"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

type GatewayRepository struct {
	queries *gateway_connection.Queries
}

var (
	once       sync.Once
	repository *GatewayRepository
)

func New() *GatewayRepository {
	once.Do(func() {
		repository = &GatewayRepository{
			queries: gateway_connection.New(db.Conn),
		}
	})
	return repository
}

func (this *GatewayRepository) WithTx(tx pgx.Tx) *GatewayRepository {
	return &GatewayRepository{
		queries: this.queries.WithTx(tx),
	}
}
