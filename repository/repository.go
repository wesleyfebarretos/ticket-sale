package repository

import (
	"github.com/wesleyfebarretos/ticket-sale/infra/db"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_users_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/users_addresses_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/users_repository"
)

var (
	User       = &users_repository.Queries{}
	UserAdress = &users_addresses_repository.Queries{}
	AdminUser  = &admin_users_repository.Queries{}
)

func Bind() {
	User = users_repository.New(db.Conn)
	UserAdress = users_addresses_repository.New(db.Conn)
	AdminUser = admin_users_repository.New(db.Conn)
}
