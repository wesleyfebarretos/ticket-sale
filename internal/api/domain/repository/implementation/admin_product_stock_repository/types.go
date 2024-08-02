package admin_product_stock_repository

import "time"

type CreateParams struct {
	ProductID int32  `json:"productId"`
	Qty       int32  `json:"qty"`
	MinQty    *int32 `json:"minQty"`
	CreatedBy int32  `json:"createdBy"`
}

type CreateResponse struct {
	ID        int32     `json:"id"`
	ProductID int32     `json:"productId"`
	Qty       int32     `json:"qty"`
	MinQty    *int32    `json:"minQty"`
	CreatedBy int32     `json:"createdBy"`
	UpdatedBy *int32    `json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UpdateParams struct {
	Qty       int32     `json:"qty"`
	MinQty    *int32    `json:"minQty"`
	UpdatedBy *int32    `json:"updatedBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	ID        int32     `json:"id"`
}
