package user_repository

import (
	"time"
)

type CreateParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

type CreateResponse struct {
	ID        int32     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UpdateParams struct {
	ID        int32  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

type CheckIfEmailExistsParams struct {
	Email string `json:"email"`
	ID    int32  `json:"id"`
}

type CheckIfEmailExistsResponse struct {
	ID        int32     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetOneWithPasswordByEmailResponse struct {
	ID        int32     `json:"id"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetOneByEmailAndRoleParams struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type GetOneByEmailAndRoleResponse struct {
	ID        int32     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetProfileResponse struct {
	ID        int32                        `json:"id"`
	FirstName string                       `json:"firstName"`
	LastName  string                       `json:"lastName"`
	Email     string                       `json:"email"`
	Role      string                       `json:"role"`
	CreatedAt time.Time                    `json:"createdAt"`
	UpdatedAt time.Time                    `json:"updatedAt"`
	Addresses []UserProfileAddressResponse `json:"addresses"`
	Phones    []UserProfilePhoneResponse   `json:"phones"`
}

type UserProfilePhoneResponse struct {
	ID     int32  `json:"id"`
	UserID int32  `json:"userId"`
	DDD    string `json:"ddd"`
	Number string `json:"number"`
	Type   string `json:"type"`
}

type UserProfileAddressResponse struct {
	ID            int32  `json:"id"`
	UserID        int32  `json:"userId"`
	StreetAddress string `json:"streetAddress"`
	City          string `json:"city"`
	Complement    string `json:"complement"`
	State         string `json:"state"`
	PostalCode    string `json:"postalCode"`
	Country       string `json:"country"`
	AddressType   string `json:"addressType"`
	Favorite      bool   `json:"favorite"`
}

type GetOneByEmailResponse struct {
	ID        int32     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetOneByIdParams struct {
	ID   int32  `json:"id"`
	Role string `json:"role"`
}

type GetOneByIdResponse struct {
	ID        int32     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetAllResponse struct {
	ID        int32     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
