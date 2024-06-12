package entities

import "github.com/wesleyfebarretos/ticket-sale/domain/enums"

type User struct {
	FirstName string
	LastName  string
	email     string
	password  string
	Role      enums.Roles
	Id        int
}

func (u *User) TableName() string {
	return "user"
}
