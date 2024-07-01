// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package users_repository

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type PhoneTypes string

const (
	PhoneTypesPhone     PhoneTypes = "phone"
	PhoneTypesTellphone PhoneTypes = "tellphone"
)

func (e *PhoneTypes) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PhoneTypes(s)
	case string:
		*e = PhoneTypes(s)
	default:
		return fmt.Errorf("unsupported scan type for PhoneTypes: %T", src)
	}
	return nil
}

type NullPhoneTypes struct {
	PhoneTypes PhoneTypes `json:"phoneTypes"`
	Valid      bool       `json:"valid"` // Valid is true if PhoneTypes is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPhoneTypes) Scan(value interface{}) error {
	if value == nil {
		ns.PhoneTypes, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PhoneTypes.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPhoneTypes) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PhoneTypes), nil
}

func (e PhoneTypes) Valid() bool {
	switch e {
	case PhoneTypesPhone,
		PhoneTypesTellphone:
		return true
	}
	return false
}

func AllPhoneTypesValues() []PhoneTypes {
	return []PhoneTypes{
		PhoneTypesPhone,
		PhoneTypesTellphone,
	}
}

type Roles string

const (
	RolesAdmin      Roles = "admin"
	RolesUser       Roles = "user"
	RolesWebservice Roles = "webservice"
	RolesSuperadmin Roles = "super admin"
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

func (e Roles) Valid() bool {
	switch e {
	case RolesAdmin,
		RolesUser,
		RolesWebservice,
		RolesSuperadmin:
		return true
	}
	return false
}

func AllRolesValues() []Roles {
	return []Roles{
		RolesAdmin,
		RolesUser,
		RolesWebservice,
		RolesSuperadmin,
	}
}

type Event struct {
	ID        int32      `json:"id"`
	ProductID int32      `json:"productId"`
	StartAt   *time.Time `json:"startAt"`
	EndAt     *time.Time `json:"endAt"`
	City      *string    `json:"city"`
	State     *string    `json:"state"`
	Location  *string    `json:"location"`
}

type Product struct {
	ID             int32      `json:"id"`
	Name           string     `json:"name"`
	Description    *string    `json:"description"`
	Uuid           uuid.UUID  `json:"uuid"`
	Price          float64    `json:"price"`
	DiscountPrice  *float64   `json:"discountPrice"`
	Active         bool       `json:"active"`
	IsDeleted      bool       `json:"isDeleted"`
	Image          *string    `json:"image"`
	ImageMobile    *string    `json:"imageMobile"`
	ImageThumbnail *string    `json:"imageThumbnail"`
	CategoryID     int32      `json:"categoryId"`
	CreatedBy      int32      `json:"createdBy"`
	UpdatedBy      *int32     `json:"updatedBy"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt"`
}

type ProductCategory struct {
	ID          int32      `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	CreatedBy   int32      `json:"createdBy"`
	UpdatedBy   *int32     `json:"updatedBy"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

type ProductStock struct {
	ID        int32      `json:"id"`
	ProductID int32      `json:"productId"`
	Qty       int32      `json:"qty"`
	MinQty    *int32     `json:"minQty"`
	CreatedBy int32      `json:"createdBy"`
	UpdatedBy *int32     `json:"updatedBy"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
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

type UsersPhone struct {
	ID     int32      `json:"id"`
	UserID int32      `json:"userId"`
	Ddd    string     `json:"ddd"`
	Number string     `json:"number"`
	Type   PhoneTypes `json:"type"`
}
