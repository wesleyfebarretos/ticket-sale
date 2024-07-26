package creditcard_controller

import (
	"time"

	"github.com/google/uuid"
)

type CreditcardFlag struct {
	Id          int32   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Regex       string  `json:"regex"`
}

type CreditcardType struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type GetAllUserCreditcardsResponseDto struct {
	Uuid           uuid.UUID      `json:"uuid"`
	Name           string         `json:"name"`
	Number         string         `json:"number"`
	Expiration     time.Time      `json:"expiration"`
	UserID         int32          `json:"userId"`
	CreatedAt      time.Time      `json:"createdAt"`
	CreditcardFlag CreditcardFlag `json:"creditcardFlag"`
	CreditcardType CreditcardType `json:"creditcardType"`
}

type CreateRequestDto struct {
	Name             string    `json:"name"`
	Number           string    `json:"number"`
	Expiration       time.Time `json:"expiration"`
	Priority         int32     `json:"priority"`
	NotifyExpiration bool      `json:"notifyExpiration"`
	UserID           int32     `json:"userId"`
	CreditcardTypeID int32     `json:"creditcardTypeId"`
	CreditcardFlagID int32     `json:"creditcardFlagId"`
}

type CreateResponseDto struct {
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

type UpdateRequestDto struct {
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
