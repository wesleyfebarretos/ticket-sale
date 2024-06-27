package repository

import (
	"github.com/wesleyfebarretos/ticket-sale/infra/db"
	"github.com/wesleyfebarretos/ticket-sale/repository/user_address_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/user_repository"
)

var (
	User       = &user_repository.Queries{}
	UserAdress = &user_address_repository.Queries{}
)

func BindAll() {
	User = user_repository.New(db.Conn)
	UserAdress = user_address_repository.New(db.Conn)
}
