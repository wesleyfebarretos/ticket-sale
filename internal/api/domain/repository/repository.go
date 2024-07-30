package repository

import (
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_events_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_gateway_connection"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_product_stocks_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_products_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_users_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/creditcard_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/users_addresses_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/users_phones_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/users_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/infra/db"
)

var (
	AdminUsers         = &admin_users_repository.Queries{}
	AdminProducts      = &admin_products_repository.Queries{}
	AdminProductStocks = &admin_product_stocks_repository.Queries{}
	AdminEvents        = &admin_events_repository.Queries{}
	Users              = &users_repository.Queries{}
	UsersAdresses      = &users_addresses_repository.Queries{}
	UsersPhones        = &users_phones_repository.Queries{}
	Creditcard         = &creditcard_repository.Queries{}
	AdminGateway       = &admin_gateway_connection.Queries{}
)

func Bind() {
	AdminUsers = admin_users_repository.New(db.Conn)
	AdminProducts = admin_products_repository.New(db.Conn)
	AdminProductStocks = admin_product_stocks_repository.New(db.Conn)
	AdminEvents = admin_events_repository.New(db.Conn)
	Users = users_repository.New(db.Conn)
	UsersAdresses = users_addresses_repository.New(db.Conn)
	UsersPhones = users_phones_repository.New(db.Conn)
	Creditcard = creditcard_repository.New(db.Conn)
	//  TODO: Maybe i'll make my own repository
	AdminGateway = admin_gateway_connection.New(db.Conn)
}
