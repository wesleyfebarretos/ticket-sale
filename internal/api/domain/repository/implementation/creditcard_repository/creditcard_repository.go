package creditcard_repository

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/creditcard_connection"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

type CreditcardRepository struct {
	queries *creditcard_connection.Queries
}

var (
	once       sync.Once
	repository *CreditcardRepository
)

func New() *CreditcardRepository {
	once.Do(func() {
		repository = &CreditcardRepository{
			queries: creditcard_connection.New(db.Conn),
		}
	})
	return repository
}

func (r *CreditcardRepository) WithTx(tx pgx.Tx) *CreditcardRepository {
	return &CreditcardRepository{
		queries: r.queries.WithTx(tx),
	}
}

func (r *CreditcardRepository) Create(ctx context.Context, params CreateParams) CreateResponse {
	entityParams := params.ToEntity()
	createdCreditcard, err := r.queries.Create(ctx, entityParams)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	response := CreateResponse{}
	return response.FromEntity(createdCreditcard)
}

func (r *CreditcardRepository) GetAllUserCreditcards(ctx context.Context, userID int32) []GetAllUserCreditcardsResponse {
	userCreditcards, err := r.queries.GetAllUserCreditcards(ctx, userID)
	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	response := GetAllUserCreditcardsResponse{}
	return response.FromEntity(userCreditcards)
}

func (r *CreditcardRepository) SoftDelete(ctx context.Context, params SoftDeleteParams) {
	err := r.queries.SoftDelete(ctx, creditcard_connection.SoftDeleteParams{
		Uuid:      params.Uuid,
		UpdatedAt: params.UpdatedAt,
	})

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}
}

func (r *CreditcardRepository) Update(ctx context.Context, params UpdateParams) {
	err := r.queries.Update(ctx, params.ToEntity())
	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}
}
