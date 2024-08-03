package creditcard_repository

import (
	"time"

	"github.com/google/uuid"
)

type CreateParams struct {
	Name             string    `json:"name"`
	Number           string    `json:"number"`
	Expiration       time.Time `json:"expiration"`
	Priority         int32     `json:"priority"`
	NotifyExpiration bool      `json:"notifyExpiration"`
	UserID           int32     `json:"userId"`
	CreditcardTypeID int32     `json:"creditcardTypeId"`
	CreditcardFlagID int32     `json:"creditcardFlagId"`
}

type CreateResponse struct {
	ID               int32     `json:"id"`
	Uuid             uuid.UUID `json:"uuid"`
	Name             string    `json:"name"`
	Number           string    `json:"number"`
	Expiration       time.Time `json:"expiration"`
	Priority         int32     `json:"priority"`
	NotifyExpiration bool      `json:"notifyExpiration"`
	UserID           int32     `json:"userId"`
	CreditcardTypeID int32     `json:"creditcardTypeId"`
	CreditcardFlagID int32     `json:"creditcardFlagId"`
	IsDeleted        bool      `json:"isDeleted"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type UpdateParams struct {
	Name             string    `json:"name"`
	Number           string    `json:"number"`
	Expiration       time.Time `json:"expiration"`
	Priority         int32     `json:"priority"`
	NotifyExpiration bool      `json:"notifyExpiration"`
	UserID           int32     `json:"userId"`
	CreditcardTypeID int32     `json:"creditcardTypeId"`
	CreditcardFlagID int32     `json:"creditcardFlagId"`
	UpdatedAt        time.Time `json:"updatedAt"`
	Uuid             uuid.UUID `json:"uuid"`
}

type GetAllUserCreditcardsResponse struct {
	Uuid           uuid.UUID              `json:"uuid"`
	Name           string                 `json:"name"`
	Number         string                 `json:"number"`
	Expiration     time.Time              `json:"expiration"`
	UserID         int32                  `json:"userId"`
	CreatedAt      time.Time              `json:"createdAt"`
	CreditcardFlag CreditcardFlagResponse `json:"creditcardFlag"`
	CreditcardType CreditcardTypeResponse `json:"creditcardType"`
}

type CreditcardFlagResponse struct {
	ID          int32
	Name        string
	Description *string
	Regex       string
}

type CreditcardTypeResponse struct {
	ID   int32
	Name string
}

type SoftDeleteParams struct {
	Uuid      uuid.UUID `json:"uuid"`
	UpdatedAt time.Time `json:"updatedAt"`
}
