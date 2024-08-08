package admin_event_repository

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_event_connection"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

type AdminEventRepository struct {
	queries *admin_event_connection.Queries
}

var (
	once       sync.Once
	repository *AdminEventRepository
)

func New() *AdminEventRepository {
	once.Do(func() {
		repository = &AdminEventRepository{
			queries: admin_event_connection.New(db.Conn),
		}
	})
	return repository
}

func (this *AdminEventRepository) WithTx(tx pgx.Tx) *AdminEventRepository {
	return &AdminEventRepository{
		queries: this.queries.WithTx(tx),
	}
}

func (this *AdminEventRepository) Create(c context.Context, createParams CreateParams) CreateResponse {
	event, err := this.queries.Create(c, createParams.ToEntity())
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}
	res := CreateResponse{}

	return res.FromEntity(event)
}

func (this *AdminEventRepository) Update(c context.Context, updateParams UpdateParams) int32 {
	productID, err := this.queries.Update(c, updateParams.ToEntity())

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	return productID
}

func (this *AdminEventRepository) GetAll(c context.Context) []GetAllResponse {
	events, err := this.queries.GetAll(c)
	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetAllResponse{}

	return res.FromEntity(events)
}

func (this *AdminEventRepository) GetOneById(c context.Context, eventID int32) *GetOneByIdResponse {
	event, err := this.queries.GetOneById(c, eventID)
	if err == pgx.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := &GetOneByIdResponse{}

	return res.FromEntity(event)
}

func (this *AdminEventRepository) SoftDelete(c context.Context, eventID int32) {
	err := this.queries.SoftDelete(c, eventID)
	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}
}
