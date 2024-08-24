package payment_order_repository

import (
	"sync"

	"github.com/jackc/pgx/v4"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/payment_order_connection"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

type PaymentOrderRepository struct {
	queries *payment_order_connection.Queries
}

var (
	once       sync.Once
	repository *PaymentOrderRepository
)

func New() *PaymentOrderRepository {
	once.Do(func() {
		repository = &PaymentOrderRepository{
			queries: payment_order_connection.New(db.Conn),
		}
	})
	return repository
}

func (this *PaymentOrderRepository) WithTx(tx pgx.Tx) *PaymentOrderRepository {
	return &PaymentOrderRepository{
		queries: this.queries.WithTx(tx),
	}
}
