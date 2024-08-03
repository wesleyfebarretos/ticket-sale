package repository

import (
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_users_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/users_addresses_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/users_phones_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/users_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

var (
	AdminUsers    = &admin_users_repository.Queries{}
	Users         = &users_repository.Queries{}
	UsersAdresses = &users_addresses_repository.Queries{}
	UsersPhones   = &users_phones_repository.Queries{}
)

func Bind() {
	AdminUsers = admin_users_repository.New(db.Conn)
	Users = users_repository.New(db.Conn)
	UsersAdresses = users_addresses_repository.New(db.Conn)
	UsersPhones = users_phones_repository.New(db.Conn)
}
