package admin_product_repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_product_connection"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

type AdminProductRepository struct {
	queries *admin_product_connection.Queries
}

var (
	once       sync.Once
	repository *AdminProductRepository
)

func New() *AdminProductRepository {
	once.Do(func() {
		repository = &AdminProductRepository{
			queries: admin_product_connection.New(db.Conn),
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
	newProduct, err := r.queries.Create(ctx, createParams.ToEntity())
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := CreateResponse{}

	return res.FromEntity(newProduct)
}

func (r *AdminProductRepository) Update(ctx context.Context, updateParams UpdateParams) bool {
	err := r.queries.Update(ctx, updateParams.ToEntity())
	if err == pgx.ErrNoRows {
		return false
	}
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return true
}

func (r *AdminProductRepository) SoftDelete(ctx context.Context, softDeleteParams SoftDeleteParams) {
	err := r.queries.SoftDelete(ctx, softDeleteParams.ToEntity())

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}
}

func (r *AdminProductRepository) GetOneById(ctx context.Context, id int32) *GetOneByIdResponse {
	product, err := r.queries.GetOneById(ctx, id)

	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetOneByIdResponse{}

	return res.FromEntity(product)
}

func (r *AdminProductRepository) GetOneByUuid(ctx context.Context, uuid uuid.UUID) *GetOneByUuidResponse {
	product, err := r.queries.GetOneByUuid(ctx, uuid)

	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetOneByUuidResponse{}

	return res.FromEntity(product)
}

func (r *AdminProductRepository) GetAll(ctx context.Context) []GetAllResponse {
	products, err := r.queries.GetAll(ctx)

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetAllResponse{}

	return res.FromEntity(products)
}

func (r *AdminProductRepository) GetAllInstallmentTimes(ctx context.Context, productID int32) []GetAllInstallmentTimeResponse {
	product, err := r.queries.GetAllProductInstallmentTimes(ctx, productID)

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetAllInstallmentTimeResponse{}

	return res.FromEntity(product)
}

func (r *AdminProductRepository) DeleteAllInstallmentTimes(ctx context.Context, productID int32) {
	err := r.queries.DeleteAllProductInstallmentTimes(ctx, productID)

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}
}

func (r *AdminProductRepository) GetAllWithDetails(ctx context.Context) []GetAllWithDetailsResponse {
	products, err := r.queries.GetAllProductsDetails(ctx)

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetAllWithDetailsResponse{}

	return res.FromEntity(products)
}

func (r *AdminProductRepository) CreateInstallments(ctx context.Context, installments []CreateInstallmentsParams) []CreateInstallmentsResponse {
	createInstallments := []admin_product_connection.CreateInstallmentsParams{}

	for _, v := range installments {
		createInstallments = append(createInstallments, v.ToEntity())
	}

	batch := r.queries.CreateInstallments(ctx, createInstallments)

	newInstallments := []admin_product_connection.FinProductPaymentTypeInstallmentTime{}

	batch.QueryRow(func(i int, installment admin_product_connection.FinProductPaymentTypeInstallmentTime, err error) {
		if err != nil {
			panic(exception.InternalServerException(fmt.Sprintf("query of index %d failed: %v", i, err)))
		}

		newInstallments = append(newInstallments, installment)
	})

	res := CreateInstallmentsResponse{}

	return res.FromEntity(newInstallments)
}
