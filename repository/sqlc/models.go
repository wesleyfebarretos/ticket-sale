// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Roles string

const (
	RolesAdmin      Roles = "admin"
	RolesUser       Roles = "user"
	RolesWebservice Roles = "webservice"
)

func (e *Roles) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Roles(s)
	case string:
		*e = Roles(s)
	default:
		return fmt.Errorf("unsupported scan type for Roles: %T", src)
	}
	return nil
}

type NullRoles struct {
	Roles Roles `json:"roles"`
	Valid bool  `json:"valid"` // Valid is true if Roles is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRoles) Scan(value interface{}) error {
	if value == nil {
		ns.Roles, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Roles.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRoles) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Roles), nil
}

type User struct {
	ID        int32      `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Role      Roles      `json:"role"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type UsersAddress struct {
	ID            int32      `json:"id"`
	UserID        int32      `json:"userId"`
	StreetAddress string     `json:"streetAddress"`
	City          string     `json:"city"`
	Complement    *string    `json:"complement"`
	State         string     `json:"state"`
	PostalCode    *string    `json:"postalCode"`
	Country       string     `json:"country"`
	AddressType   *string    `json:"addressType"`
	Favorite      *bool      `json:"favorite"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
}
