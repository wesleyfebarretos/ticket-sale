package user_phone_repository

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/user_phone_connection"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

type UserPhoneRepository struct {
	queries *user_phone_connection.Queries
}

var (
	once       sync.Once
	repository *UserPhoneRepository
)

func New() *UserPhoneRepository {
	once.Do(func() {
		repository = &UserPhoneRepository{
			queries: user_phone_connection.New(db.Conn),
		}
	})
	return repository
}

func (this *UserPhoneRepository) WithTx(tx pgx.Tx) *UserPhoneRepository {
	return &UserPhoneRepository{
		queries: this.queries.WithTx(tx),
	}
}

func (this *UserPhoneRepository) Create(ctx context.Context, p CreateParams) CreateResponse {
	user, err := this.queries.Create(ctx, p.ToEntity())
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := CreateResponse{}

	return res.FromEntity(user)
}
