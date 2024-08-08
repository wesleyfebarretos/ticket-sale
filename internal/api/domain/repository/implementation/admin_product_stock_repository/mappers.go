package admin_product_stock_repository

import "github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_product_stock_connection"

func (this *CreateParams) ToEntity() admin_product_stock_connection.CreateParams {
	return admin_product_stock_connection.CreateParams{
		ProductID: this.ProductID,
		Qty:       this.Qty,
		MinQty:    this.MinQty,
		CreatedBy: this.CreatedBy,
	}
}

func (this *CreateResponse) FromEntity(p admin_product_stock_connection.ProductStock) CreateResponse {
	return CreateResponse{
		ID:        p.ID,
		ProductID: p.ProductID,
		Qty:       p.Qty,
		MinQty:    p.MinQty,
		CreatedBy: p.CreatedBy,
		UpdatedBy: p.UpdatedBy,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (this *UpdateParams) ToEntity() admin_product_stock_connection.UpdateParams {
	return admin_product_stock_connection.UpdateParams{
		Qty:       this.Qty,
		MinQty:    this.MinQty,
		UpdatedBy: this.UpdatedBy,
		UpdatedAt: this.UpdatedAt,
		ID:        this.ID,
	}
}
