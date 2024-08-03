package admin_event_repository

import (
	"time"

	"github.com/google/uuid"
)

type CreateParams struct {
	ProductID int32      `json:"productId"`
	StartAt   *time.Time `json:"startAt"`
	EndAt     *time.Time `json:"endAt"`
	City      *string    `json:"city"`
	State     *string    `json:"state"`
	Location  *string    `json:"location"`
	CreatedBy int32      `json:"createdBy"`
}

type CreateResponse struct {
	ID        int32      `json:"id"`
	ProductID int32      `json:"productId"`
	StartAt   *time.Time `json:"startAt"`
	EndAt     *time.Time `json:"endAt"`
	City      *string    `json:"city"`
	State     *string    `json:"state"`
	Location  *string    `json:"location"`
	CreatedBy int32      `json:"createdBy"`
	UpdatedBy *int32     `json:"updatedBy"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type UpdateParams struct {
	ID        int32      `json:"id"`
	StartAt   *time.Time `json:"startAt"`
	EndAt     *time.Time `json:"endAt"`
	City      *string    `json:"city"`
	State     *string    `json:"state"`
	Location  *string    `json:"location"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type GetAllResponse struct {
	ID        int32                `json:"id"`
	ProductID int32                `json:"productId"`
	StartAt   *time.Time           `json:"startAt"`
	EndAt     *time.Time           `json:"endAt"`
	City      *string              `json:"city"`
	State     *string              `json:"state"`
	Location  *string              `json:"location"`
	CreatedBy int32                `json:"createdBy"`
	UpdatedBy *int32               `json:"updatedBy"`
	CreatedAt time.Time            `json:"createdAt"`
	UpdatedAt time.Time            `json:"updatedAt"`
	Product   EventProductResponse `json:"product"`
}

type EventProductResponse struct {
	ID             int32                   `json:"id"`
	Name           string                  `json:"name"`
	Description    *string                 `json:"description"`
	Uuid           uuid.UUID               `json:"uuid"`
	Price          float64                 `json:"price"`
	DiscountPrice  *float64                `json:"discountPrice"`
	Active         bool                    `json:"active"`
	IsDeleted      bool                    `json:"isDeleted"`
	Image          *string                 `json:"image"`
	ImageThumbnail *string                 `json:"imageThumbnail"`
	ImageMobile    *string                 `json:"imageMobile"`
	CategoryId     int32                   `json:"categoryId"`
	Category       ProductCategoryResponse `json:"category"`
}

type ProductCategoryResponse struct {
	Description *string `json:"description" example:"EVENT"`
	Name        string  `json:"name" example:"event"`
	ID          int32   `json:"id" example:"3"`
}

type GetOneByIdResponse struct {
	ID        int32                `json:"id"`
	ProductID int32                `json:"productId"`
	StartAt   *time.Time           `json:"startAt"`
	EndAt     *time.Time           `json:"endAt"`
	City      *string              `json:"city"`
	State     *string              `json:"state"`
	Location  *string              `json:"location"`
	CreatedBy int32                `json:"createdBy"`
	UpdatedBy *int32               `json:"updatedBy"`
	CreatedAt time.Time            `json:"createdAt"`
	UpdatedAt time.Time            `json:"updatedAt"`
	Product   EventDetailsResponse `json:"product"`
}

type EventDetailsResponse struct {
	ID             int32                `json:"id"`
	Name           string               `json:"name"`
	Description    *string              `json:"description"`
	Uuid           uuid.UUID            `json:"uuid"`
	Price          float64              `json:"price"`
	DiscountPrice  *float64             `json:"discountPrice"`
	Active         bool                 `json:"active"`
	IsDeleted      bool                 `json:"is_deleted"`
	Image          *string              `json:"image"`
	ImageMobile    *string              `json:"imageMobile"`
	ImageThumbnail *string              `json:"imageThumbnail"`
	CategoryID     int32                `json:"categoryId"`
	CreatedBy      int32                `json:"createdBy"`
	UpdatedBy      *int32               `json:"updatedBy"`
	CreatedAt      time.Time            `json:"createdAt"`
	UpdatedAt      time.Time            `json:"updatedAt"`
	Stock          *StockResponse       `json:"stock"`
	Category       CategoryResponse     `json:"category"`
	Installments   InstallmentsResponse `json:"installments"`
}

type InstallmentsResponse struct {
	Creditcard  []PaymentTypeInstallmentResponse `json:"creditcard"`
	PaymentSlip []PaymentTypeInstallmentResponse `json:"paymentSlip"`
	Pix         []PaymentTypeInstallmentResponse `json:"pix"`
}

type StockResponse struct {
	MinQty    *int32 `json:"minQty"`
	ID        int32  `json:"id"`
	ProductID int32  `json:"productId"`
	Qty       int32  `json:"qty"`
}

type CategoryResponse struct {
	Description *string `json:"description"`
	Name        string  `json:"name"`
	ID          int32   `json:"id"`
}

type PaymentTypeInstallmentResponse struct {
	InstallmentTimeID   int32   `json:"installmentTimeId" example:"1"`
	InstallmentTimeName string  `json:"installmentTimeName" example:"1x"`
	Fee                 float64 `json:"fee" example:"3.22"`
	Tariff              float64 `json:"tariff" example:"7.00"`
}
