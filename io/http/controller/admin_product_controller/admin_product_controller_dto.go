package admin_product_controller

import (
	"time"

	"github.com/google/uuid"
)

// REQUESTS
type CreateRequestDto struct {
	Name           string   `json:"name" binding:"required,max=255" example:"Red Hot Chilly Peppers"`
	Description    *string  `json:"description" binding:"max=2000" example:"Fresh and fiery red hot chilly peppers, perfect for adding a spicy kick to your dishes."`
	Price          float64  `json:"price" binding:"required" example:"5.99"`
	DiscountPrice  *float64 `json:"discountPrice" example:"4.99"`
	Active         bool     `json:"active" example:"true"`
	Image          *string  `json:"image" binding:"required,max=2000" example:"https://example.com/images/red-hot-chilly-peppers.jpg"`
	ImageMobile    *string  `json:"imageMobile" binding:"required,max=2000" example:"https://example.com/images/red-hot-chilly-peppers-mobile.jpg"`
	ImageThumbnail *string  `json:"imageThumbnail" binding:"required,max=2000" example:"https://example.com/images/red-hot-chilly-peppers-thumbnail.jpg"`
	CategoryID     int32    `json:"categoryId" binding:"required,min=1" example:"3"`
}

type UpdateRequestDto struct {
	Name           string   `json:"name" binding:"required,max=255" example:"Update Red Hot Chilly Peppers"`
	Description    *string  `json:"description" binding:"max=2000" example:"Update Fresh and fiery red hot chilly peppers, perfect for adding a spicy kick to your dishes."`
	Price          float64  `json:"price" binding:"required" example:"6.11"`
	DiscountPrice  *float64 `json:"discountPrice" example:"5.11"`
	Active         bool     `json:"active" example:"false"`
	Image          *string  `json:"image" binding:"required,max=2000" example:"https://example.com/images/red-hot-chilly-peppers.png"`
	ImageMobile    *string  `json:"imageMobile" binding:"required,max=2000" example:"https://example.com/images/red-hot-chilly-peppers-mobile.png"`
	ImageThumbnail *string  `json:"imageThumbnail" binding:"required,max=2000" example:"https://example.com/images/red-hot-chilly-peppers-thumbnail.png"`
	CategoryID     int32    `json:"categoryId" binding:"required,min=1" example:"1"`
}

// RESPONSES
type CreateResponseDto struct {
	Name           string    `json:"name" example:"Red Hot Chilly Peppers"`
	Description    *string   `json:"description" example:"Fresh and fiery red hot chilly peppers, perfect for adding a spicy kick to your dishes."`
	Price          float64   `json:"price" example:"5.99"`
	DiscountPrice  *float64  `json:"discountPrice" example:"4.99"`
	Active         bool      `json:"active" example:"true"`
	Image          *string   `json:"image" example:"https://example.com/images/red-hot-chilly-peppers.jpg"`
	ImageMobile    *string   `json:"imageMobile" example:"https://example.com/images/red-hot-chilly-peppers-mobile.jpg"`
	ImageThumbnail *string   `json:"imageThumbnail" example:"https://example.com/images/red-hot-chilly-peppers-thumbnail.jpg"`
	CategoryID     int32     `json:"categoryId" example:"3"`
	CreatedBy      int32     `json:"createdBy" example:"1"`
	Uuid           uuid.UUID `json:"uuid" example:"998f91f3-4dd7-419d-a543-0d26a0e945ec"`
}

type GetAllResponseDto struct {
	ID             int32      `json:"id" example:"1"`
	Name           string     `json:"name" example:"Red Hot Chilly Peppers"`
	Description    *string    `json:"description" example:"Fresh and fiery red hot chilly peppers, perfect for adding a spicy kick to your dishes."`
	Uuid           uuid.UUID  `json:"uuid" example:"998f91f3-4dd7-419d-a543-0d26a0e945ec"`
	Price          float64    `json:"price" example:"5.99"`
	DiscountPrice  *float64   `json:"discountPrice" example:"4.99"`
	Active         bool       `json:"active" example:"true"`
	IsDeleted      bool       `json:"isDeleted" example:"false"`
	Image          *string    `json:"image" example:"https://example.com/images/red-hot-chilly-peppers.jpg"`
	ImageMobile    *string    `json:"imageMobile" example:"https://example.com/images/red-hot-chilly-peppers-mobile.jpg"`
	ImageThumbnail *string    `json:"imageThumbnail" example:"https://example.com/images/red-hot-chilly-peppers-thumbnail.jpg"`
	CategoryID     int32      `json:"categoryId" example:"3"`
	CreatedBy      int32      `json:"createdBy" example:"1"`
	UpdatedBy      *int32     `json:"updatedBy" example:"1"`
	CreatedAt      time.Time  `json:"createdAt" example:"2023-01-01T00:00:00Z"`
	UpdatedAt      *time.Time `json:"updatedAt" example:"2023-01-01T00:00:00Z"`
}

type GetAllWithRelationsResponseDto struct {
	ID             int32                `json:"id" example:"1"`
	Name           string               `json:"name" example:"Red Hot Chilly Peppers"`
	Description    *string              `json:"description" example:"Fresh and fiery red hot chilly peppers, perfect for adding a spicy kick to your dishes."`
	Uuid           uuid.UUID            `json:"uuid" example:"998f91f3-4dd7-419d-a543-0d26a0e945ec"`
	Price          float64              `json:"price" example:"5.99"`
	DiscountPrice  *float64             `json:"discountPrice" example:"4.99"`
	Active         bool                 `json:"active" example:"true"`
	IsDeleted      bool                 `json:"isDeleted" example:"false"`
	Image          *string              `json:"image" example:"https://example.com/images/red-hot-chilly-peppers.jpg"`
	ImageMobile    *string              `json:"imageMobile" example:"https://example.com/images/red-hot-chilly-peppers-mobile.jpg"`
	ImageThumbnail *string              `json:"imageThumbnail" example:"https://example.com/images/red-hot-chilly-peppers-thumbnail.jpg"`
	CategoryID     int32                `json:"categoryId" example:"3"`
	CreatedBy      int32                `json:"createdBy" example:"1"`
	UpdatedBy      *int32               `json:"updatedBy" example:"1"`
	CreatedAt      time.Time            `json:"createdAt" example:"2023-01-01T00:00:00Z"`
	UpdatedAt      *time.Time           `json:"updatedAt" example:"2023-01-01T00:00:00Z"`
	Stock          *StockResponseDto    `json:"stock"`
	Category       *CategoryResponseDto `json:"category"`
}

type StockResponseDto struct {
	ID        int32 `json:"id"`
	ProductID int32 `json:"productId"`
	Qty       int32 `json:"qty"`
	MinQty    int32 `json:"minQty"`
}

type CategoryResponseDto struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type GetOneByIdResponseDto struct {
	ID             int32               `json:"id" example:"1"`
	Name           string              `json:"name" example:"Red Hot Chilly Peppers"`
	Description    *string             `json:"description" example:"Fresh and fiery red hot chilly peppers, perfect for adding a spicy kick to your dishes."`
	Uuid           uuid.UUID           `json:"uuid" example:"998f91f3-4dd7-419d-a543-0d26a0e945ec"`
	Price          float64             `json:"price" example:"5.99"`
	DiscountPrice  *float64            `json:"discountPrice" example:"4.99"`
	Active         bool                `json:"active" example:"true"`
	IsDeleted      bool                `json:"isDeleted" example:"false"`
	Image          *string             `json:"image" example:"https://example.com/images/red-hot-chilly-peppers.jpg"`
	ImageMobile    *string             `json:"imageMobile" example:"https://example.com/images/red-hot-chilly-peppers-mobile.jpg"`
	ImageThumbnail *string             `json:"imageThumbnail" example:"https://example.com/images/red-hot-chilly-peppers-thumbnail.jpg"`
	CategoryID     int32               `json:"categoryId" example:"3"`
	CreatedBy      int32               `json:"createdBy" example:"1"`
	UpdatedBy      *int32              `json:"updatedBy" example:"1"`
	CreatedAt      time.Time           `json:"createdAt" example:"2023-01-01T00:00:00Z"`
	UpdatedAt      *time.Time          `json:"updatedAt" example:"2023-01-01T00:00:00Z"`
	Stock          StockResponseDto    `json:"stock"`
	Category       CategoryResponseDto `json:"category"`
}

type GetOneByUuidResponseDto struct {
	ID             int32               `json:"id" example:"1"`
	Name           string              `json:"name" example:"Red Hot Chilly Peppers"`
	Description    *string             `json:"description" example:"Fresh and fiery red hot chilly peppers, perfect for adding a spicy kick to your dishes."`
	Uuid           uuid.UUID           `json:"uuid" example:"998f91f3-4dd7-419d-a543-0d26a0e945ec"`
	Price          float64             `json:"price" example:"5.99"`
	DiscountPrice  *float64            `json:"discountPrice" example:"4.99"`
	Active         bool                `json:"active" example:"true"`
	IsDeleted      bool                `json:"isDeleted" example:"false"`
	Image          *string             `json:"image" example:"https://example.com/images/red-hot-chilly-peppers.jpg"`
	ImageMobile    *string             `json:"imageMobile" example:"https://example.com/images/red-hot-chilly-peppers-mobile.jpg"`
	ImageThumbnail *string             `json:"imageThumbnail" example:"https://example.com/images/red-hot-chilly-peppers-thumbnail.jpg"`
	CategoryID     int32               `json:"categoryId" example:"3"`
	CreatedBy      int32               `json:"createdBy" example:"1"`
	UpdatedBy      *int32              `json:"updatedBy" example:"1"`
	CreatedAt      time.Time           `json:"createdAt" example:"2023-01-01T00:00:00Z"`
	UpdatedAt      *time.Time          `json:"updatedAt" example:"2023-01-01T00:00:00Z"`
	Stock          StockResponseDto    `json:"stock"`
	Category       CategoryResponseDto `json:"category"`
}
