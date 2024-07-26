package override

type ProductStock struct {
	MinQty    *int32 `json:"minQty" example:"50"`
	ID        int32  `json:"id" example:"1"`
	ProductID int32  `json:"productId" example:"1"`
	Qty       int32  `json:"qty" binding:"required,min=1" example:"100"`
}
