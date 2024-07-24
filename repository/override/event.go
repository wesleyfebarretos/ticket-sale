package override

import (
	"time"

	"github.com/google/uuid"
)

type EventProduct struct {
	ID             int32           `json:"id"`
	Name           string          `json:"name"`
	Description    *string         `json:"description"`
	Uuid           uuid.UUID       `json:"uuid"`
	Price          float64         `json:"price"`
	DiscountPrice  *float64        `json:"discountPrice"`
	Active         bool            `json:"active"`
	IsDeleted      bool            `json:"isDeleted"`
	Image          *string         `json:"image"`
	ImageThumbnail *string         `json:"imageThumbnail"`
	ImageMobile    *string         `json:"imageMobile"`
	CategoryId     int32           `json:"categoryId"`
	Category       ProductCategory `json:"category"`
}

type EventDetails struct {
	ID             int32                   `json:"id"`
	Name           string                  `json:"name"`
	Description    *string                 `json:"description"`
	Uuid           uuid.UUID               `json:"uuid"`
	Price          float64                 `json:"price"`
	DiscountPrice  *float64                `json:"discountPrice"`
	Active         bool                    `json:"active"`
	IsDeleted      bool                    `json:"is_deleted"`
	Image          *string                 `json:"image"`
	ImageMobile    *string                 `json:"imageMobile"`
	ImageThumbnail *string                 `json:"imageThumbnail"`
	CategoryID     int32                   `json:"categoryId"`
	CreatedBy      int32                   `json:"createdBy"`
	UpdatedBy      *int32                  `json:"updatedBy"`
	CreatedAt      time.Time               `json:"createdAt"`
	UpdatedAt      time.Time               `json:"updatedAt"`
	Stock          *StockResponseDto       `json:"stock"`
	Category       *CategoryResponseDto    `json:"category"`
	Installments   InstallmentsResponseDto `json:"installments"`
}

type InstallmentsResponseDto struct {
	Creditcard  []PaymentTypeInstallment `json:"creditcard"`
	PaymentSlip []PaymentTypeInstallment `json:"paymentSlip"`
	Pix         []PaymentTypeInstallment `json:"pix"`
}

type StockResponseDto struct {
	MinQty    *int32 `json:"minQty"`
	ID        int32  `json:"id"`
	ProductID int32  `json:"productId"`
	Qty       int32  `json:"qty"`
}

type CategoryResponseDto struct {
	Description *string `json:"description"`
	Name        string  `json:"name"`
	ID          int32   `json:"id"`
}

type EventProductWithRelations struct {
	ID             int32           `json:"id"`
	Name           string          `json:"name"`
	Description    *string         `json:"description"`
	Uuid           uuid.UUID       `json:"uuid"`
	Price          float64         `json:"price"`
	DiscountPrice  *float64        `json:"discountPrice"`
	Active         bool            `json:"active"`
	IsDeleted      bool            `json:"isDeleted"`
	Image          *string         `json:"image"`
	ImageThumbnail *string         `json:"imageThumbnail"`
	ImageMobile    *string         `json:"imageMobile"`
	CategoryId     int32           `json:"categoryId"`
	Category       ProductCategory `json:"category"`
	Stock          ProductStock    `json:"stock"`
}
