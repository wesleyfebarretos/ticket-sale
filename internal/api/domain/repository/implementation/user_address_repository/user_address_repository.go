package user_address_repository

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/user_address_connection"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

type UserAddressRepository struct {
	queries *user_address_connection.Queries
}

var (
	once       sync.Once
	repository *UserAddressRepository
)

func New() *UserAddressRepository {
	once.Do(func() {
		repository = &UserAddressRepository{
			queries: user_address_connection.New(db.Conn),
		}
	})
	return repository
}

func (this *UserAddressRepository) WithTx(tx pgx.Tx) *UserAddressRepository {
	return &UserAddressRepository{
		queries: this.queries.WithTx(tx),
	}
}

func (this *UserAddressRepository) Create(ctx context.Context, p CreateParams) CreateResponse {
	user, err := this.queries.Create(ctx, p.ToEntity())
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := CreateResponse{}

	return res.FromEntity(user)
}

func (this *UserAddressRepository) Update(ctx context.Context, p UpdateParams) {
	err := this.queries.Update(ctx, p.ToEntity())
	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}
}
