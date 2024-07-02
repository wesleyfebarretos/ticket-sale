package repository

import (
	"github.com/wesleyfebarretos/ticket-sale/infra/db"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_products_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_users_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/users_addresses_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/users_phones_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/users_repository"
)

var (
	AdminUsers    = &admin_users_repository.Queries{}
	AdminProducts = &admin_products_repository.Queries{}
	Users         = &users_repository.Queries{}
	UsersAdresses = &users_addresses_repository.Queries{}
	UsersPhones   = &users_phones_repository.Queries{}
)

func Bind() {
	AdminUsers = admin_users_repository.New(db.Conn)
	AdminProducts = admin_products_repository.New(db.Conn)
	Users = users_repository.New(db.Conn)
	UsersAdresses = users_addresses_repository.New(db.Conn)
	UsersPhones = users_phones_repository.New(db.Conn)
}
