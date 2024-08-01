package admin_gateway_repository

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_gateway_connection"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

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

func (r *AdminProductRepository) Create(ctx context.Context, createParams CreateParams) CreateResponse {
	newGateway, err := r.queries.Create(ctx, createToEntity(createParams))
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return createToDomain(newGateway)
}

func (r *AdminProductRepository) Update(ctx context.Context, updateParams UpdateParams) bool {
	err := r.queries.Update(ctx, updateToEntity(updateParams))
	if err == pgx.ErrNoRows {
		return false
	}
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return true
}

func (r *AdminProductRepository) SoftDelete(ctx context.Context, softDeleteParams SoftDeleteParams) bool {
	err := r.queries.SoftDelete(ctx, SoftDeleteToEntity(softDeleteParams))
	if err == pgx.ErrNoRows {
		return false
	}
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return true
}

func (r *AdminProductRepository) GetOneById(ctx context.Context, id int32) *GetOneByIdResponse {
	gateway, err := r.queries.GetOneById(ctx, id)

	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	return getOneByIdToDomain(gateway)
}

func (r *AdminProductRepository) GetAll(ctx context.Context) []GetAllResponse {
	gateways, err := r.queries.GetAll(ctx)
	if err == pgx.ErrNoRows {
		return []GetAllResponse{}
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return getAlltoDomain(gateways)
}
