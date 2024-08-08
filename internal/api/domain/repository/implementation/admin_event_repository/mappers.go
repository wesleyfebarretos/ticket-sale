package admin_event_repository

import (
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_event_connection"
)

func (this *CreateParams) ToEntity() admin_event_connection.CreateParams {
	return admin_event_connection.CreateParams{
		ProductID: this.ProductID,
		StartAt:   this.StartAt,
		EndAt:     this.EndAt,
		City:      this.City,
		State:     this.State,
		Location:  this.Location,
		CreatedBy: this.CreatedBy,
	}
}

func (this *CreateResponse) FromEntity(p admin_event_connection.Event) CreateResponse {
	return CreateResponse{
		ID:        p.ID,
		ProductID: p.ProductID,
		StartAt:   p.StartAt,
		EndAt:     p.EndAt,
		City:      p.City,
		State:     p.State,
		Location:  p.Location,
		CreatedBy: p.CreatedBy,
		UpdatedBy: p.UpdatedBy,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (this *UpdateParams) ToEntity() admin_event_connection.UpdateParams {
	return admin_event_connection.UpdateParams{
		ID:        this.ID,
		StartAt:   this.StartAt,
		EndAt:     this.EndAt,
		City:      this.City,
		State:     this.State,
		Location:  this.Location,
		UpdatedAt: this.UpdatedAt,
	}
}

func (this *GetAllResponse) FromEntity(p []admin_event_connection.EventsGetAll) []GetAllResponse {
	res := []GetAllResponse{}

	for _, v := range p {
		res = append(res, GetAllResponse{
			ID:        v.ID,
			ProductID: v.ProductID,
			StartAt:   v.StartAt,
			EndAt:     v.EndAt,
			City:      v.City,
			State:     v.State,
			Location:  v.Location,
			CreatedBy: v.CreatedBy,
			UpdatedBy: v.UpdatedBy,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Product: EventProductResponse{
				ID:             v.Product.ID,
				Name:           v.Product.Name,
				Description:    v.Product.Description,
				Uuid:           v.Product.Uuid,
				Price:          v.Product.Price,
				DiscountPrice:  v.Product.DiscountPrice,
				Active:         v.Product.Active,
				IsDeleted:      v.Product.IsDeleted,
				Image:          v.Product.Image,
				ImageThumbnail: v.Product.ImageThumbnail,
				ImageMobile:    v.Product.ImageMobile,
				CategoryId:     v.Product.CategoryId,
				Category: ProductCategoryResponse{
					Description: v.Product.Category.Description,
					Name:        v.Product.Category.Name,
					ID:          v.Product.Category.ID,
				},
			},
		})
	}

	return res
}

func (this *GetOneByIdResponse) FromEntity(p admin_event_connection.EventsDetail) *GetOneByIdResponse {
	creditcard := []PaymentTypeInstallmentResponse{}
	paymentSlip := []PaymentTypeInstallmentResponse{}
	pix := []PaymentTypeInstallmentResponse{}

	for _, v := range p.Product.Installments.Creditcard {
		creditcard = append(creditcard, PaymentTypeInstallmentResponse{
			InstallmentTimeID:   v.InstallmentTimeID,
			InstallmentTimeName: v.InstallmentTimeName,
			Fee:                 v.Fee,
			Tariff:              v.Tariff,
		})
	}

	for _, v := range p.Product.Installments.PaymentSlip {
		paymentSlip = append(paymentSlip, PaymentTypeInstallmentResponse{
			InstallmentTimeID:   v.InstallmentTimeID,
			InstallmentTimeName: v.InstallmentTimeName,
			Fee:                 v.Fee,
			Tariff:              v.Tariff,
		})
	}
	for _, v := range p.Product.Installments.Pix {
		pix = append(pix, PaymentTypeInstallmentResponse{
			InstallmentTimeID:   v.InstallmentTimeID,
			InstallmentTimeName: v.InstallmentTimeName,
			Fee:                 v.Fee,
			Tariff:              v.Tariff,
		})
	}

	stock := &StockResponse{}

	if p.Product.Stock != nil {
		stock.MinQty = p.Product.Stock.MinQty
		stock.ID = p.Product.Stock.ID
		stock.ProductID = p.Product.Stock.ProductID
		stock.Qty = p.Product.Stock.Qty
	} else {
		stock = nil
	}

	return &GetOneByIdResponse{
		ID:        p.ID,
		ProductID: p.ProductID,
		StartAt:   p.StartAt,
		EndAt:     p.EndAt,
		City:      p.City,
		State:     p.State,
		Location:  p.Location,
		CreatedBy: p.CreatedBy,
		UpdatedBy: p.UpdatedBy,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Product: EventDetailsResponse{
			ID:             p.Product.ID,
			Name:           p.Product.Name,
			Description:    p.Product.Description,
			Uuid:           p.Product.Uuid,
			Price:          p.Product.Price,
			DiscountPrice:  p.Product.DiscountPrice,
			Active:         p.Product.Active,
			IsDeleted:      p.Product.IsDeleted,
			Image:          p.Product.Image,
			ImageMobile:    p.Product.ImageMobile,
			ImageThumbnail: p.Product.ImageThumbnail,
			CategoryID:     p.Product.CategoryID,
			CreatedBy:      p.Product.CreatedBy,
			UpdatedBy:      p.Product.UpdatedBy,
			CreatedAt:      p.Product.CreatedAt,
			UpdatedAt:      p.Product.UpdatedAt,
			Stock:          stock,
			Category: CategoryResponse{
				Description: p.Product.Category.Description,
				Name:        p.Product.Category.Name,
				ID:          p.Product.Category.ID,
			},
			Installments: InstallmentsResponse{
				Creditcard:  creditcard,
				PaymentSlip: paymentSlip,
				Pix:         pix,
			},
		},
	}
}
