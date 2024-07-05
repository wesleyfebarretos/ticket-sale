package admin_event_controller

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller/admin_product_controller"
)

type CreateRequestDto struct {
	StartAt  *time.Time                                `json:"startAt" example:"2024-01-01T00:00:00Z"`
	EndAt    *time.Time                                `json:"endAt" example:"2024-01-10T00:00:00Z"`
	City     *string                                   `json:"city" example:"Orlando"`
	State    *string                                   `json:"state" example:"FL"`
	Location *string                                   `json:"location" example:"Disney"`
	Product  admin_product_controller.CreateRequestDto `json:"product"`
}

type CreateResponseDto struct {
	ID        int32                                      `json:"id" example:"1"`
	ProductID int32                                      `json:"productId" example:"1"`
	City      *string                                    `json:"city" example:"Orlando"`
	State     *string                                    `json:"state" example:"FL"`
	Location  *string                                    `json:"location" example:"Disney"`
	EndAt     *time.Time                                 `json:"endAt" example:"2024-01-10T00:00:00Z"`
	StartAt   *time.Time                                 `json:"startAt" example:"2024-01-01T00:00:00Z"`
	Product   admin_product_controller.CreateResponseDto `json:"product"`
}

type UpdateRequestDto struct {
	StartAt  *time.Time                                `json:"startAt" example:"2024-01-01T00:00:00Z"`
	EndAt    *time.Time                                `json:"endAt" example:"2024-01-10T00:00:00Z"`
	City     *string                                   `json:"city" example:"Orlando"`
	State    *string                                   `json:"state" example:"FL"`
	Location *string                                   `json:"location" example:"Disney"`
	Product  admin_product_controller.UpdateRequestDto `json:"product"`
}

type GetAllResponseDto struct {
	ID        int32            `json:"id" example:"1"`
	ProductID int32            `json:"productId" example:"1"`
	City      *string          `json:"city" example:"Orlando"`
	State     *string          `json:"state" example:"FL"`
	Location  *string          `json:"location" example:"Disney"`
	EndAt     *time.Time       `json:"endAt" example:"2024-01-10T00:00:00Z"`
	StartAt   *time.Time       `json:"startAt" example:"2024-01-01T00:00:00Z"`
	Product   GetAllProductDto `json:"product"`
}

type GetAllProductDto struct {
	ID             int32       `json:"id" example:"1"`
	Name           string      `json:"name" example:"Red Hot Chilly Peppers"`
	Description    *string     `json:"description" example:"Fresh and fiery red hot chilly peppers, perfect for adding a spicy kick to your dishes."`
	Price          float64     `json:"price" example:"5.99"`
	DiscountPrice  *float64    `json:"discountPrice" example:"4.99"`
	Active         bool        `json:"active" example:"true"`
	Image          *string     `json:"image" example:"https://example.com/images/red-hot-chilly-peppers.jpg"`
	ImageMobile    *string     `json:"imageMobile" example:"https://example.com/images/red-hot-chilly-peppers-mobile.jpg"`
	ImageThumbnail *string     `json:"imageThumbnail" example:"https://example.com/images/red-hot-chilly-peppers-thumbnail.jpg"`
	CategoryID     int32       `json:"categoryId" example:"3"`
	Uuid           uuid.UUID   `json:"uuid" example:"998f91f3-4dd7-419d-a543-0d26a0e945ec"`
	IsDeleted      bool        `json:"isDeleted" example:"false"`
	CategoryId     int32       `json:"categoryId" example:"3"`
	Category       CategoryDto `json:"category"`
}

type CategoryDto struct {
	Description *string `json:"description" example:"EVENT"`
	Name        string  `json:"name" example:"event"`
	ID          int32   `json:"id" example:"3"`
}

type GetOneByIdResponseDto struct {
	ID        int32                `json:"id" example:"1"`
	ProductID int32                `json:"productId" example:"1"`
	City      *string              `json:"city" example:"Orlando"`
	State     *string              `json:"state" example:"FL"`
	Location  *string              `json:"location" example:"Disney"`
	EndAt     *time.Time           `json:"endAt" example:"2024-01-10T00:00:00Z"`
	StartAt   *time.Time           `json:"startAt" example:"2024-01-01T00:00:00Z"`
	Product   GetOneByIdProductDto `json:"product"`
}
type GetOneByIdProductDto struct {
	ID             int32       `json:"id" example:"1"`
	Name           string      `json:"name" example:"Red Hot Chilly Peppers"`
	Description    *string     `json:"description" example:"Fresh and fiery red hot chilly peppers, perfect for adding a spicy kick to your dishes."`
	Price          float64     `json:"price" example:"5.99"`
	DiscountPrice  *float64    `json:"discountPrice" example:"4.99"`
	Active         bool        `json:"active" example:"true"`
	Image          *string     `json:"image" example:"https://example.com/images/red-hot-chilly-peppers.jpg"`
	ImageMobile    *string     `json:"imageMobile" example:"https://example.com/images/red-hot-chilly-peppers-mobile.jpg"`
	ImageThumbnail *string     `json:"imageThumbnail" example:"https://example.com/images/red-hot-chilly-peppers-thumbnail.jpg"`
	CategoryID     int32       `json:"categoryId" example:"3"`
	Uuid           uuid.UUID   `json:"uuid" example:"998f91f3-4dd7-419d-a543-0d26a0e945ec"`
	IsDeleted      bool        `json:"isDeleted" example:"false"`
	Category       CategoryDto `json:"category"`
	Stock          StockDto    `json:"stock"`
}

type StockDto struct {
	Qty       int32  `json:"qty" binding:"required,min=1" example:"100"`
	MinQty    *int32 `json:"minQty" example:"50"`
	ID        int32  `json:"id" example:"1"`
	ProductID int32  `json:"productId" example:"1"`
}
