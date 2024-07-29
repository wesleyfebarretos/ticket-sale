package test_data

import (
	"context"
	"log"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/users_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
)

const UserTestPassword = "123"

func NewUser(role string) users_repository.GetOneWithPasswordByEmailRow {
	password, err := utils.HashPassword(UserTestPassword)
	if err != nil {
		log.Fatalf("could not hash password: %v", err)
	}

	newUser := users_repository.CreateParams{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoetest@gmail.com",
		Password:  password,
		Role:      users_repository.Roles(role),
	}

	user, _ := repository.Users.Create(context.Background(), newUser)

	nUser, _ := repository.Users.GetOneWithPasswordByEmail(context.Background(), user.Email)

	return nUser
}
