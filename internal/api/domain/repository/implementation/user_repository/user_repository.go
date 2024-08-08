package user_repository

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/user_connection"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

type UserRepository struct {
	queries *user_connection.Queries
}

var (
	once       sync.Once
	repository *UserRepository
)

func New() *UserRepository {
	once.Do(func() {
		repository = &UserRepository{
			queries: user_connection.New(db.Conn),
		}
	})
	return repository
}

func (this *UserRepository) WithTx(tx pgx.Tx) *UserRepository {
	return &UserRepository{
		queries: this.queries.WithTx(tx),
	}
}

func (this *UserRepository) Create(ctx context.Context, p CreateParams) CreateResponse {
	p.Role = roles_enum.USER
	user, err := this.queries.Create(ctx, p.ToEntity())
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := CreateResponse{}

	return res.FromEntity(user)
}

func (this *UserRepository) Update(ctx context.Context, p UpdateParams) {
	err := this.queries.Update(ctx, p.ToEntity())

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}
}

func (this *UserRepository) CheckIfEmailExists(
	ctx context.Context,
	p CheckIfEmailExistsParams,
) *CheckIfEmailExistsResponse {
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

func (this *UserRepository) GetOneWithPasswordByEmail(ctx context.Context, email string) *GetOneWithPasswordByEmailResponse {
	user, err := this.queries.GetOneWithPasswordByEmail(ctx, email)
	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetOneWithPasswordByEmailResponse{}

	return res.FromEntity(user)
}

func (this *UserRepository) GetOneByEmailAndRole(ctx context.Context, p GetOneByEmailAndRoleParams) *GetOneByEmailAndRoleResponse {
	user, err := this.queries.GetOneByEmailAndRole(ctx, p.ToEntity())
	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetOneByEmailAndRoleResponse{}

	return res.FromEntity(user)
}

func (this *UserRepository) GetProfile(ctx context.Context, id int32) *GetProfileResponse {
	user, err := this.queries.GetProfile(ctx, id)
	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetProfileResponse{}

	return res.FromEntity(user)
}

func (this *UserRepository) GetOneByEmail(ctx context.Context, email string) *GetOneByEmailResponse {
	user, err := this.queries.GetOneByEmail(ctx, email)

	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetOneByEmailResponse{}

	return res.FromEntity(user)
}

func (this *UserRepository) GetOneById(ctx context.Context, p GetOneByIdParams) *GetOneByIdResponse {
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

func (this *UserRepository) GetAll(ctx context.Context) []GetAllResponse {
	users, err := this.queries.GetAll(ctx, roles_enum.USER)

	if err != nil && err != pgx.ErrNoRows {
		panic(exception.InternalServerException(err.Error()))
	}

	res := GetAllResponse{}

	return res.FromEntity(users)
}
