package override

import "github.com/google/uuid"

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
