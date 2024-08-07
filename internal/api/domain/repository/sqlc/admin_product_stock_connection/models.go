// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package admin_product_stock_connection

import (
	"time"
)

type ProductStock struct {
	ID        int32     `json:"id"`
	ProductID int32     `json:"productId"`
	Qty       int32     `json:"qty"`
	MinQty    *int32    `json:"minQty"`
	CreatedBy int32     `json:"createdBy"`
	UpdatedBy *int32    `json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
