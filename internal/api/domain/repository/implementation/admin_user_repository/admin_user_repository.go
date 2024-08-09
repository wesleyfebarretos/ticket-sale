package admin_user_repository

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_user_connection"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

type AdminUserRepository struct {
	queries *admin_user_connection.Queries
}

var (
	once       sync.Once
	repository *AdminUserRepository
)

func New() *AdminUserRepository {
	once.Do(func() {
		repository = &AdminUserRepository{
			queries: admin_user_connection.New(db.Conn),
		}
	})
	return repository
}

func (this *AdminUserRepository) WithTx(tx pgx.Tx) *AdminUserRepository {
	return &AdminUserRepository{
		queries: this.queries.WithTx(tx),
	}
}

func (this *AdminUserRepository) Create(ctx context.Context, p CreateParams) CreateResponse {
	user, err := this.queries.Create(ctx, p.ToEntity())
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := CreateResponse{}

	return res.FromEntity(user)
}

func (this *AdminUserRepository) Update(ctx context.Context, p UpdateParams) {
	err := this.queries.Update(ctx, p.ToEntity())

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}
}

func (this *AdminUserRepository) GetOneByEmail(ctx context.Context, p GetOneByEmailParams) *GetOneByEmailResponse {
	user, err := this.queries.GetOneByEmail(ctx, p.ToEntity())

	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetOneByEmailResponse{}

	return res.FromEntity(user)
}

func (this *AdminUserRepository) GetOneById(ctx context.Context, p GetOneByIdParams) *GetOneByIdResponse {
	user, err := this.queries.GetOneById(ctx, p.ToEntity())

	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetOneByIdResponse{}

	return res.FromEntity(user)
}

func (this *AdminUserRepository) GetAll(ctx context.Context) []GetAllResponse {
	users, err := this.queries.GetAll(ctx, roles_enum.ADMIN)

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetAllResponse{}

	return res.FromEntity(users)
}

func (this *AdminUserRepository) Delete(ctx context.Context, id int32) {
	err := this.queries.Delete(ctx, id)

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}
}

func (this *AdminUserRepository) CheckIfEmailExists(ctx context.Context, p CheckIfEmailExistsParams) *CheckIfEmailExistsResponse {
	user, err := this.queries.CheckIfEmailExists(ctx, p.ToEntity())

	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := CheckIfEmailExistsResponse{}

	return res.FromEntity(user)
}

func (this *AdminUserRepository) GetOneByEmailAndRoles(ctx context.Context, p GetOneByEmailAndRolesParams) *GetOneByEmailAndRolesResponse {
	user, err := this.queries.GetOneByEmailAndRoles(ctx, p.ToEntity())

	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetOneByEmailAndRolesResponse{}

	return res.FromEntity(user)
}
