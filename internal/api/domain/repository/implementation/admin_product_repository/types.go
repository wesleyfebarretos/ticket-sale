package admin_product_repository

import (
	"time"

	"github.com/google/uuid"
)

type CreateParams struct {
	Name           string   `json:"name"`
	Description    *string  `json:"description"`
	Price          float64  `json:"price"`
	DiscountPrice  *float64 `json:"discountPrice"`
	Active         bool     `json:"active"`
	Image          *string  `json:"image"`
	ImageMobile    *string  `json:"imageMobile"`
	ImageThumbnail *string  `json:"imageThumbnail"`
	CategoryID     int32    `json:"categoryId"`
	CreatedBy      int32    `json:"createdBy"`
}

type CreateResponse struct {
	ID             int32     `json:"id"`
	Name           string    `json:"name"`
	Description    *string   `json:"description"`
	Uuid           uuid.UUID `json:"uuid"`
	Price          float64   `json:"price"`
	DiscountPrice  *float64  `json:"discountPrice"`
	Active         bool      `json:"active"`
	IsDeleted      bool      `json:"isDeleted"`
	Image          *string   `json:"image"`
	ImageMobile    *string   `json:"imageMobile"`
	ImageThumbnail *string   `json:"imageThumbnail"`
	CategoryID     int32     `json:"categoryId"`
	CreatedBy      int32     `json:"createdBy"`
	UpdatedBy      *int32    `json:"updatedBy"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type UpdateParams struct {
	Name           string   `json:"name"`
	Description    *string  `json:"description"`
	Price          float64  `json:"price"`
	DiscountPrice  *float64 `json:"discountPrice"`
	Active         bool     `json:"active"`
	Image          *string  `json:"image"`
	ImageMobile    *string  `json:"imageMobile"`
	ImageThumbnail *string  `json:"imageThumbnail"`
	CategoryID     int32    `json:"categoryId"`
	UpdatedBy      *int32   `json:"updatedBy"`
	ID             int32    `json:"id"`
}

type GetAllResponse struct {
	ID             int32     `json:"id"`
	Name           string    `json:"name"`
	Description    *string   `json:"description"`
	Uuid           uuid.UUID `json:"uuid"`
	Price          float64   `json:"price"`
	DiscountPrice  *float64  `json:"discountPrice"`
	Active         bool      `json:"active"`
	IsDeleted      bool      `json:"isDeleted"`
	Image          *string   `json:"image"`
	ImageMobile    *string   `json:"imageMobile"`
	ImageThumbnail *string   `json:"imageThumbnail"`
	CategoryID     int32     `json:"categoryId"`
	CreatedBy      int32     `json:"createdBy"`
	UpdatedBy      *int32    `json:"updatedBy"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type GetOneByIdResponse struct {
	ID             int32                       `json:"id"`
	Name           string                      `json:"name"`
	Description    *string                     `json:"description"`
	Uuid           uuid.UUID                   `json:"uuid"`
	Price          float64                     `json:"price"`
	DiscountPrice  *float64                    `json:"discountPrice"`
	Active         bool                        `json:"active"`
	IsDeleted      bool                        `json:"isDeleted"`
	Image          *string                     `json:"image"`
	ImageMobile    *string                     `json:"imageMobile"`
	ImageThumbnail *string                     `json:"imageThumbnail"`
	CategoryID     int32                       `json:"categoryId"`
	CreatedBy      int32                       `json:"createdBy"`
	UpdatedBy      *int32                      `json:"updatedBy"`
	CreatedAt      time.Time                   `json:"createdAt"`
	UpdatedAt      time.Time                   `json:"updatedAt"`
	Stock          *ProductStockResponse       `json:"stock"`
	Category       ProductCategoryResponse     `json:"category"`
	Installments   ProductInstallmentsResponse `json:"installments"`
}

type ProductStockResponse struct {
	MinQty    *int32 `json:"minQty" example:"50"`
	ID        int32  `json:"id" example:"1"`
	ProductID int32  `json:"productId" example:"1"`
	Qty       int32  `json:"qty" binding:"required,min=1" example:"100"`
}

type ProductCategoryResponse struct {
	Description *string `json:"description" example:"EVENT"`
	Name        string  `json:"name" example:"event"`
	ID          int32   `json:"id" example:"3"`
}
type ProductInstallmentsResponse struct {
	Creditcard  []PaymentTypeInstallment `json:"creditcard"`
	PaymentSlip []PaymentTypeInstallment `json:"paymentSlip"`
	Pix         []PaymentTypeInstallment `json:"pix"`
}

type PaymentTypeInstallment struct {
	InstallmentTimeID   int32   `json:"installmentTimeId" example:"1"`
	InstallmentTimeName string  `json:"installmentTimeName" example:"1x"`
	Fee                 float64 `json:"fee" example:"3.22"`
	Tariff              float64 `json:"tariff" example:"7.00"`
}

type SoftDeleteParams struct {
	ID        int32  `json:"id"`
	UpdatedBy *int32 `json:"updatedBy"`
}

type gatewayProcess struct {
	Name string `json:"name"`
	ID   int32  `json:"id"`
}

type gatewayPaymentTypes struct {
	Name string `json:"name"`
	ID   int32  `json:"id"`
}

type CreatePaymentTypesParams struct {
	GatewayID            int32  `json:"gatewayId"`
	GatewayPaymentTypeID int32  `json:"gatewayPaymentTypeId"`
	CreatedBy            int32  `json:"createdBy"`
	UpdatedBy            *int32 `json:"updatedBy"`
}

type CreatePaymentTypesResponse struct {
	ID                   int32     `json:"id"`
	GatewayID            int32     `json:"gatewayId"`
	GatewayPaymentTypeID int32     `json:"gatewayPaymentTypeId"`
	CreatedBy            int32     `json:"createdBy"`
	UpdatedBy            *int32    `json:"updatedBy"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}

type GetOneByUuidResponse struct {
	ID             int32                       `json:"id"`
	Name           string                      `json:"name"`
	Description    *string                     `json:"description"`
	Uuid           uuid.UUID                   `json:"uuid"`
	Price          float64                     `json:"price"`
	DiscountPrice  *float64                    `json:"discountPrice"`
	Active         bool                        `json:"active"`
	IsDeleted      bool                        `json:"isDeleted"`
	Image          *string                     `json:"image"`
	ImageMobile    *string                     `json:"imageMobile"`
	ImageThumbnail *string                     `json:"imageThumbnail"`
	CategoryID     int32                       `json:"categoryId"`
	CreatedBy      int32                       `json:"createdBy"`
	UpdatedBy      *int32                      `json:"updatedBy"`
	CreatedAt      time.Time                   `json:"createdAt"`
	UpdatedAt      time.Time                   `json:"updatedAt"`
	Stock          *ProductStockResponse       `json:"stock"`
	Category       ProductCategoryResponse     `json:"category"`
	Installments   ProductInstallmentsResponse `json:"installments"`
}

type GetAllInstallmentTimeResponse struct {
	ID                int32   `json:"id"`
	Fee               float64 `json:"fee"`
	Tariff            float64 `json:"tariff"`
	PaymentTypeID     int32   `json:"paymentTypeId"`
	InstallmentTimeID int32   `json:"installmentTimeId"`
}

type GetAllWithDetailsResponse struct {
	ID             int32                       `json:"id"`
	Name           string                      `json:"name"`
	Description    *string                     `json:"description"`
	Uuid           uuid.UUID                   `json:"uuid"`
	Price          float64                     `json:"price"`
	DiscountPrice  *float64                    `json:"discountPrice"`
	Active         bool                        `json:"active"`
	IsDeleted      bool                        `json:"isDeleted"`
	Image          *string                     `json:"image"`
	ImageMobile    *string                     `json:"imageMobile"`
	ImageThumbnail *string                     `json:"imageThumbnail"`
	CategoryID     int32                       `json:"categoryId"`
	CreatedBy      int32                       `json:"createdBy"`
	UpdatedBy      *int32                      `json:"updatedBy"`
	CreatedAt      time.Time                   `json:"createdAt"`
	UpdatedAt      time.Time                   `json:"updatedAt"`
	Stock          *ProductStockResponse       `json:"stock"`
	Category       ProductCategoryResponse     `json:"category"`
	Installments   ProductInstallmentsResponse `json:"installments"`
}

type CreateInstallmentsParams struct {
	ProductID         int32   `json:"productId"`
	PaymentTypeID     int32   `json:"paymentTypeId"`
	InstallmentTimeID int32   `json:"installmentTimeId"`
	Fee               float64 `json:"fee"`
	Tariff            float64 `json:"tariff"`
	CreatedBy         int32   `json:"createdBy"`
}

type CreateInstallmentsResponse struct {
	ID                int32     `json:"id"`
	ProductID         int32     `json:"productId"`
	PaymentTypeID     int32     `json:"paymentTypeId"`
	InstallmentTimeID int32     `json:"installmentTimeId"`
	Fee               float64   `json:"fee"`
	Tariff            float64   `json:"tariff"`
	CreatedBy         int32     `json:"createdBy"`
	UpdatedBy         *int32    `json:"updatedBy"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
