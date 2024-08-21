package admin_event_handler

import (
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_event_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_product_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_product_stock_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/admin_event_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/handler/admin_product_handler"
)

func (s *CreateRequestDto) ToDomain(userID int32) admin_event_service.CreateParams {
	installments := []admin_product_repository.CreateInstallmentsParams{}

	for _, v := range s.Product.Installments {
		installments = append(installments, admin_product_repository.CreateInstallmentsParams{
			PaymentTypeID:     v.PaymentTypeID,
			InstallmentTimeID: v.ID,
			Fee:               *v.Fee,
			Tariff:            *v.Tariff,
			CreatedBy:         userID,
		})
	}

	return admin_event_service.CreateParams{
		NewEvent: admin_event_repository.CreateParams{
			StartAt:   s.StartAt,
			EndAt:     s.EndAt,
			City:      s.City,
			State:     s.State,
			Location:  s.Location,
			CreatedBy: userID,
		},
		NewProduct: admin_product_repository.CreateParams{
			Name:           s.Product.Name,
			Description:    s.Product.Description,
			Price:          s.Product.Price,
			DiscountPrice:  s.Product.DiscountPrice,
			Active:         s.Product.Active,
			Image:          s.Product.Image,
			ImageMobile:    s.Product.ImageMobile,
			ImageThumbnail: s.Product.ImageThumbnail,
			CategoryID:     s.Product.CategoryID,
			CreatedBy:      userID,
		},
		NewStock: admin_product_stock_repository.CreateParams{
			Qty:       s.Product.Stock.Qty,
			MinQty:    s.Product.Stock.MinQty,
			CreatedBy: userID,
		},
		NewProductInstallments: installments,
	}

}

func (s *CreateResponseDto) FromDomain(p admin_event_service.CreateResponse) CreateResponseDto {
	installments := []admin_product_handler.CreateInstallmentsResponseDto{}

	for _, v := range p.ProductInstallments {
		installments = append(installments, admin_product_handler.CreateInstallmentsResponseDto{
			ID:            v.ID,
			PaymentTypeID: v.PaymentTypeID,
			InstallmentID: v.InstallmentTimeID,
			Fee:           v.Fee,
			Tariff:        v.Tariff,
		})
	}
	return CreateResponseDto{
		ID:        p.Event.ID,
		ProductID: p.Event.ProductID,
		City:      p.Event.City,
		State:     p.Event.State,
		Location:  p.Event.Location,
		EndAt:     p.Event.EndAt,
		StartAt:   p.Event.StartAt,
		Product: admin_product_handler.CreateResponseDto{
			ID:             p.Product.ID,
			Name:           p.Product.Name,
			Description:    p.Product.Description,
			Price:          p.Product.Price,
			DiscountPrice:  p.Product.DiscountPrice,
			Active:         p.Product.Active,
			Image:          p.Product.Image,
			ImageMobile:    p.Product.ImageMobile,
			ImageThumbnail: p.Product.ImageThumbnail,
			CategoryID:     p.Product.CategoryID,
			CreatedBy:      p.Product.CreatedBy,
			Uuid:           p.Product.Uuid,
			IsDeleted:      p.Product.IsDeleted,
			UpdatedBy:      p.Product.UpdatedBy,
			CreatedAt:      p.Product.CreatedAt,
			UpdatedAt:      p.Product.UpdatedAt,
			Stock: admin_product_handler.CreateStockResponseDto{
				ID:        p.ProductStock.ID,
				ProductID: p.ProductStock.ProductID,
				Qty:       p.ProductStock.Qty,
				MinQty:    p.ProductStock.MinQty,
			},
			Installments: installments,
		},
	}
}

func (s *UpdateRequestDto) ToDomain(eventID int32, userID int32) admin_event_service.UpdateParams {
	installments := []admin_product_repository.CreateInstallmentsParams{}

	for _, v := range s.Product.Installments {
		installments = append(installments, admin_product_repository.CreateInstallmentsParams{
			PaymentTypeID:     v.PaymentTypeID,
			InstallmentTimeID: v.ID,
			Fee:               *v.Fee,
			Tariff:            *v.Tariff,
			CreatedBy:         userID,
		})
	}

	return admin_event_service.UpdateParams{
		UpdateEvent: admin_event_repository.UpdateParams{
			ID:       eventID,
			StartAt:  s.StartAt,
			EndAt:    s.EndAt,
			City:     s.City,
			State:    s.State,
			Location: s.Location,
		},
		UpdateProduct: admin_product_repository.UpdateParams{
			Name:           s.Product.Name,
			Description:    s.Product.Description,
			Price:          s.Product.Price,
			DiscountPrice:  s.Product.DiscountPrice,
			Active:         s.Product.Active,
			Image:          s.Product.Image,
			ImageMobile:    s.Product.ImageMobile,
			ImageThumbnail: s.Product.ImageThumbnail,
			CategoryID:     s.Product.CategoryID,
			UpdatedBy:      &userID,
		},
		UpdateProductInstallments: installments,
	}

}

func (s *GetAllResponseDto) FromDomain(p []admin_event_repository.GetAllResponse) []GetAllResponseDto {
	res := []GetAllResponseDto{}
	for _, v := range p {
		res = append(res, GetAllResponseDto{
			ID:        v.ID,
			ProductID: v.ProductID,
			City:      v.City,
			State:     v.State,
			Location:  v.Location,
			EndAt:     v.EndAt,
			StartAt:   v.StartAt,
			Product: GetAllProductDto{
				ID:             v.Product.ID,
				Name:           v.Product.Name,
				Description:    v.Product.Description,
				Price:          v.Product.Price,
				DiscountPrice:  v.Product.DiscountPrice,
				Active:         v.Product.Active,
				Image:          v.Product.Image,
				ImageMobile:    v.Product.ImageMobile,
				ImageThumbnail: v.Product.ImageThumbnail,
				CategoryID:     v.Product.CategoryId,
				Uuid:           v.Product.Uuid,
				IsDeleted:      v.Product.IsDeleted,
				CategoryId:     v.Product.CategoryId,
				Category: CategoryDto{
					Description: v.Product.Category.Description,
					Name:        v.Product.Category.Name,
					ID:          v.Product.Category.ID,
				},
			},
		})
	}

	return res
}

func (s *GetOneByIdResponseDto) FromDomain(p *admin_event_repository.GetOneByIdResponse) GetOneByIdResponseDto {
	stock := &StockDto{}

	if p.Product.Stock != nil {
		stock.MinQty = p.Product.Stock.MinQty
		stock.ProductID = p.Product.Stock.ProductID
		stock.Qty = p.Product.Stock.Qty
	} else {
		stock = nil
	}

	creditcard := []admin_product_handler.PaymentTypeInstallment{}
	paymentSlip := []admin_product_handler.PaymentTypeInstallment{}
	pix := []admin_product_handler.PaymentTypeInstallment{}

	for _, v := range p.Product.Installments.Creditcard {
		creditcard = append(creditcard, admin_product_handler.PaymentTypeInstallment{
			InstallmentTimeID:   v.InstallmentTimeID,
			InstallmentTimeName: v.InstallmentTimeName,
			Fee:                 v.Fee,
			Tariff:              v.Tariff,
		})

	}

	for _, v := range p.Product.Installments.PaymentSlip {
		paymentSlip = append(paymentSlip, admin_product_handler.PaymentTypeInstallment{
			InstallmentTimeID:   v.InstallmentTimeID,
			InstallmentTimeName: v.InstallmentTimeName,
			Fee:                 v.Fee,
			Tariff:              v.Tariff,
		})

	}
	for _, v := range p.Product.Installments.Pix {
		pix = append(pix, admin_product_handler.PaymentTypeInstallment{
			InstallmentTimeID:   v.InstallmentTimeID,
			InstallmentTimeName: v.InstallmentTimeName,
			Fee:                 v.Fee,
			Tariff:              v.Tariff,
		})

	}

	return GetOneByIdResponseDto{
		ID:        p.ID,
		ProductID: p.ProductID,
		City:      p.City,
		State:     p.State,
		Location:  p.Location,
		EndAt:     p.EndAt,
		StartAt:   p.StartAt,
		Product: GetOneByIdProductDto{
			ID:             p.Product.ID,
			Name:           p.Product.Name,
			Description:    p.Product.Description,
			Price:          p.Product.Price,
			DiscountPrice:  p.Product.DiscountPrice,
			Active:         p.Product.Active,
			Image:          p.Product.Image,
			ImageMobile:    p.Product.ImageMobile,
			ImageThumbnail: p.Product.ImageThumbnail,
			CategoryID:     p.Product.CategoryID,
			Uuid:           p.Product.Uuid,
			IsDeleted:      p.Product.IsDeleted,
			Category: CategoryDto{
				Description: p.Product.Category.Description,
				Name:        p.Product.Category.Name,
				ID:          p.Product.Category.ID,
			},
			Stock: stock,
			Installments: admin_product_handler.InstallmentsResponseDto{
				Creditcard:  creditcard,
				PaymentSlip: paymentSlip,
				Pix:         pix,
			},
		},
	}
}
