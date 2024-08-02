package admin_gateway_repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_gateway_connection"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

type AdminGatewayRepository struct {
	queries *admin_gateway_connection.Queries
}

var (
	once       sync.Once
	repository *AdminGatewayRepository
)

func New() *AdminGatewayRepository {
	once.Do(func() {
		repository = &AdminGatewayRepository{
			queries: admin_gateway_connection.New(db.Conn),
		}
	})
	return repository
}

func (r *AdminGatewayRepository) WithTx(tx pgx.Tx) *AdminGatewayRepository {
	return &AdminGatewayRepository{
		queries: r.queries.WithTx(tx),
	}
}

func (r *AdminGatewayRepository) Create(ctx context.Context, createParams CreateParams) CreateResponse {
	newGateway, err := r.queries.Create(ctx, createParams.ToEntity())
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := CreateResponse{}

	return res.FromEntity(newGateway)
}

func (r *AdminGatewayRepository) CreatePaymentTypes(ctx context.Context, createParams []CreatePaymentTypesParams) []CreatePaymentTypesResponse {
	paymentTypes := []admin_gateway_connection.CreatePaymentTypesParams{}
	for _, v := range createParams {
		paymentTypes = append(paymentTypes, v.ToEntity())
	}

	paymentTypesBatch := r.queries.CreatePaymentTypes(ctx, paymentTypes)

	newPaymentTypes := []admin_gateway_connection.FinGatewayPaymentTypeAssociation{}

	paymentTypesBatch.QueryRow(func(i int, fgpta admin_gateway_connection.FinGatewayPaymentTypeAssociation, err error) {
		if err != nil {
			panic(exception.InternalServerException(fmt.Sprintf("query of index %d failed: %v", i, err)))
		}

		newPaymentTypes = append(newPaymentTypes, fgpta)
	})

	res := []CreatePaymentTypesResponse{}

	for _, v := range newPaymentTypes {
		domainModel := CreatePaymentTypesResponse{}
		res = append(res, domainModel.FromEntity(v))
	}

	return res
}

func (r *AdminGatewayRepository) Update(ctx context.Context, updateParams UpdateParams) bool {
	err := r.queries.Update(ctx, updateParams.ToEntity())
	if err == pgx.ErrNoRows {
		return false
	}
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return true
}

func (r *AdminGatewayRepository) SoftDelete(ctx context.Context, softDeleteParams SoftDeleteParams) bool {
	err := r.queries.SoftDelete(ctx, softDeleteParams.ToEntity())
	if err == pgx.ErrNoRows {
		return false
	}
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return true
}

func (r *AdminGatewayRepository) GetOneById(ctx context.Context, id int32) *GetOneByIdResponse {
	gateway, err := r.queries.GetOneById(ctx, id)

	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetOneByIdResponse{}

	return res.FromEntity(gateway)
}

func (r *AdminGatewayRepository) GetAll(ctx context.Context) []GetAllResponse {
	gateways, err := r.queries.GetAll(ctx)
	if err == pgx.ErrNoRows {
		return []GetAllResponse{}
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetAllResponse{}

	return res.FromEntity(gateways)
}
