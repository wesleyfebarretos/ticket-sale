package admin_product_stock_repository

import "github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_product_stock_connection"

func (s *CreateParams) ToEntity() admin_product_stock_connection.CreateParams {
	return admin_product_stock_connection.CreateParams{
		ProductID: s.ProductID,
		Qty:       s.Qty,
		MinQty:    s.MinQty,
		CreatedBy: s.CreatedBy,
	}
}

func (s *CreateResponse) FromEntity(p admin_product_stock_connection.ProductStock) CreateResponse {
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

func (s *UpdateParams) ToEntity() admin_product_stock_connection.UpdateParams {
	return admin_product_stock_connection.UpdateParams{
		Qty:       s.Qty,
		MinQty:    s.MinQty,
		UpdatedBy: s.UpdatedBy,
		UpdatedAt: s.UpdatedAt,
		ID:        s.ID,
	}
}
