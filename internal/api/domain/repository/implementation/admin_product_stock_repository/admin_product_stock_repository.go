package admin_product_stock_repository

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_product_stock_connection"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

type AdminProductStockRepository struct {
	queries *admin_product_stock_connection.Queries
}

var (
	once       sync.Once
	repository *AdminProductStockRepository
)

func New() *AdminProductStockRepository {
	once.Do(func() {
		repository = &AdminProductStockRepository{
			queries: admin_product_stock_connection.New(db.Conn),
		}
	})
	return repository
}

func (r *AdminProductStockRepository) WithTx(tx pgx.Tx) *AdminProductStockRepository {
	return &AdminProductStockRepository{
		queries: r.queries.WithTx(tx),
	}
}

func (r *AdminProductStockRepository) Create(c context.Context, createParams CreateParams) CreateResponse {
	stock, err := r.queries.Create(c, createParams.ToEntity())

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := CreateResponse{}

	return res.FromEntity(stock)
}

func (r *AdminProductStockRepository) Update(c context.Context, updateParams UpdateParams) {
	err := r.queries.Update(c, updateParams.ToEntity())

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}
}
