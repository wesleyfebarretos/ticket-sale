// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package users_phones_repository

import (
	"database/sql/driver"
	"fmt"
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

type UsersPhone struct {
	ID     int32      `json:"id"`
	UserID int32      `json:"userId"`
	Ddd    string     `json:"ddd"`
	Number string     `json:"number"`
	Type   PhoneTypes `json:"type"`
}
