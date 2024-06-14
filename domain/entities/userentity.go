package entities

import "github.com/wesleyfebarretos/ticket-sale/domain/enums"

type UserEntity struct {
	FirstName string
	LastName  string
	email     string
	password  string
	Role      enums.Roles
	Id        int
}

func (u *UserEntity) TableName() string {
	return "user"
}
