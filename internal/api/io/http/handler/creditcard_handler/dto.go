package creditcard_handler

import (
	"time"

	"github.com/google/uuid"
)

type CreditcardFlag struct {
	Id          int32   `json:"id" example:"2"`
	Name        string  `json:"name" example:"visa"`
	Description *string `json:"description" example:"VISA"`
	Regex       string  `json:"regex" example:"'^4d{5}'"`
}

type CreditcardType struct {
	Id   int32  `json:"id" example:"1"`
	Name string `json:"name" example:"Credit"`
}

type GetAllUserCreditcardsResponseDto struct {
	Uuid           uuid.UUID      `json:"uuid" example:"767c1587-1ade-4e6a-a9fe-f0e07c83bda9"`
	Name           string         `json:"name" binding:"required" example:"Testing"`
	Number         string         `json:"number" binding:"required,min=8" example:"4242********4242"`
	Expiration     time.Time      `json:"expiration" binding:"required" example:"2025-01-01T00:00:00Z"`
	UserID         int32          `json:"userId" example:"1"`
	CreatedAt      time.Time      `json:"createdAt" example:"2024-01-01T00:00:00Z"`
	CreditcardFlag CreditcardFlag `json:"creditcardFlag"`
	CreditcardType CreditcardType `json:"creditcardType"`
}

type CreateRequestDto struct {
	Name             string    `json:"name" binding:"required" example:"Testing"`
	Number           string    `json:"number" binding:"required,min=8" example:"4242424242424242"`
	Expiration       time.Time `json:"expiration" binding:"required" example:"2025-01-01"`
	Priority         int32     `json:"priority" example:"1"`
	NotifyExpiration bool      `json:"notifyExpiration" example:"true"`
	CreditcardTypeID int32     `json:"creditcardTypeId" binding:"required,min=1" example:"1"`
	CreditcardFlagID int32     `json:"creditcardFlagId" binding:"required,min=1" example:"1"`
	Tokenize         bool      `json:"tokenize" binding:"required" example:"true"`
}

type CreateResponseDto struct {
	ID               int32     `json:"id" example:"1"`
	Uuid             uuid.UUID `json:"uuid" example:"767c1587-1ade-4e6a-a9fe-f0e07c83bda9"`
	Name             string    `json:"name" binding:"required" example:"Testing"`
	Number           string    `json:"number" binding:"required,min=8" example:"4242********4242"`
	Expiration       time.Time `json:"expiration" binding:"required" example:"2025-01-01"`
	Priority         int32     `json:"priority" example:"1"`
	NotifyExpiration bool      `json:"notifyExpiration" example:"true"`
	UserID           int32     `json:"userId" example:"1"`
	CreditcardTypeID int32     `json:"creditcardTypeId" binding:"required,min=1" example:"1"`
	CreditcardFlagID int32     `json:"creditcardFlagId" binding:"required,min=1" example:"1"`
	IsDeleted        bool      `json:"isDeleted" example:"false"`
	CreatedAt        time.Time `json:"createdAt" example:"2024-01-01T00:00:00Z"`
	UpdatedAt        time.Time `json:"updatedAt" example:"2024-01-01T00:00:00Z"`
}

type UpdateRequestDto struct {
	Name             string    `json:"name" binding:"required" example:"Testing Update"`
	Number           string    `json:"number" binding:"required,min=8" example:"6643213384289379"`
	Expiration       time.Time `json:"expiration" binding:"required" example:"2025-01-01"`
	Priority         int32     `json:"priority" example:"2"`
	NotifyExpiration bool      `json:"notifyExpiration" example:"true"`
	CreditcardTypeID int32     `json:"creditcardTypeId" binding:"required,min=1" example:"2"`
	CreditcardFlagID int32     `json:"creditcardFlagId" binding:"required,min=1" example:"2"`
}

type SoftDeleteRequestDto struct {
	Uuid      uuid.UUID `json:"uuid"`
	UpdatedAt time.Time `json:"updatedAt"`
}
