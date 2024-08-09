package test_data

import (
	"context"
	"log"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_user_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/user_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
)

const UserTestPassword = "123"

func NewUser(role string) *user_repository.GetOneWithPasswordByEmailResponse {
	password, err := utils.HashPassword(UserTestPassword)
	if err != nil {
		log.Fatalf("could not hash password: %v", err)
	}

	newUser := user_repository.CreateParams{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoetest@gmail.com",
		Password:  password,
	}

	_repository := user_repository.New()

	ctx := context.Background()

	user := _repository.Create(ctx, newUser)

	if role == roles_enum.ADMIN {
		admin_user_repository.New().Update(ctx, admin_user_repository.UpdateParams{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      roles_enum.ADMIN,
		})
	}

	nUser := _repository.GetOneWithPasswordByEmail(ctx, user.Email)

	return nUser
}
