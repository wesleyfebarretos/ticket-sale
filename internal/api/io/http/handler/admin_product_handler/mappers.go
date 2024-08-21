package admin_product_handler

import (
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_product_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_product_stock_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/admin_product_service"
)

func (s *CreateRequestDto) ToDomain(userID int32) admin_product_service.CreateParams {

	newProductRequest := admin_product_repository.CreateParams{
		Name:           s.Name,
		Description:    s.Description,
		Price:          s.Price,
		DiscountPrice:  s.DiscountPrice,
		Active:         s.Active,
		Image:          s.Image,
		ImageMobile:    s.ImageMobile,
		ImageThumbnail: s.ImageThumbnail,
		CategoryID:     s.CategoryID,
		CreatedBy:      userID,
	}

	newProductStockRequest := admin_product_stock_repository.CreateParams{
		Qty:       s.Stock.Qty,
		MinQty:    s.Stock.MinQty,
		CreatedBy: userID,
	}

	newProductInstallments := []admin_product_repository.CreateInstallmentsParams{}

	for _, installment := range s.Installments {
		newInstallment := admin_product_repository.CreateInstallmentsParams{
			PaymentTypeID:     installment.PaymentTypeID,
			InstallmentTimeID: installment.ID,
			Fee:               0,
			Tariff:            0,
			CreatedBy:         userID,
		}

		if installment.Fee != nil {
			newInstallment.Fee = *installment.Fee
		}

		if installment.Tariff != nil {
			newInstallment.Tariff = *installment.Tariff
		}

		newProductInstallments = append(newProductInstallments, newInstallment)
	}

	return admin_product_service.CreateParams{
		Product:      newProductRequest,
		Stock:        newProductStockRequest,
		Installments: newProductInstallments,
	}
}

func (s *CreateResponseDto) FromDomain(p admin_product_service.CreateResponse) CreateResponseDto {
	newInstallments := []CreateInstallmentsResponseDto{}

	for _, i := range p.Installments {
		newInstallments = append(newInstallments, CreateInstallmentsResponseDto{
			ID:            i.ID,
			PaymentTypeID: i.PaymentTypeID,
			Fee:           i.Fee,
			Tariff:        i.Tariff,
			InstallmentID: i.InstallmentTimeID,
		})
	}

	return CreateResponseDto{
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
		Stock: CreateStockResponseDto{
			ID:        p.Stock.ID,
			ProductID: p.Stock.ProductID,
			Qty:       p.Stock.Qty,
			MinQty:    p.Stock.MinQty,
		},
		Installments: newInstallments,
	}
}

func (s *UpdateRequestDto) ToDomain(userID int32, productID int32) admin_product_service.UpdateParams {

	updateInstallments := []admin_product_repository.CreateInstallmentsParams{}

	for _, installment := range s.Installments {
		updatedInstallment := admin_product_repository.CreateInstallmentsParams{
			PaymentTypeID:     installment.PaymentTypeID,
			InstallmentTimeID: installment.ID,
			Fee:               0,
			Tariff:            0,
			CreatedBy:         userID,
			ProductID:         productID,
		}
		if installment.Fee != nil {
			updatedInstallment.Fee = *installment.Fee
		}

		if installment.Tariff != nil {
			updatedInstallment.Tariff = *installment.Tariff
		}

		updateInstallments = append(updateInstallments, updatedInstallment)
	}

	return admin_product_service.UpdateParams{
		Product: admin_product_repository.UpdateParams{
			Name:           s.Name,
			Description:    s.Description,
			Price:          s.Price,
			DiscountPrice:  s.DiscountPrice,
			Active:         s.Active,
			Image:          s.Image,
			ImageMobile:    s.ImageMobile,
			ImageThumbnail: s.ImageThumbnail,
			CategoryID:     s.CategoryID,
			UpdatedBy:      &userID,
			ID:             productID,
		},
		Installments: updateInstallments,
	}
}

func (s *SoftDeleteRequestDto) ToDomain(productID int32, userID int32) admin_product_repository.SoftDeleteParams {
	return admin_product_repository.SoftDeleteParams{
		ID:        productID,
		UpdatedBy: &userID,
	}
}

func (s *GetOneByIdResponseDto) FromDomain(p admin_product_repository.GetOneByIdResponse) GetOneByIdResponseDto {
	stock := &StockResponseDto{}

	if p.Stock != nil {
		stock.MinQty = p.Stock.MinQty
		stock.ID = p.Stock.ID
		stock.ProductID = p.Stock.ProductID
		stock.Qty = p.Stock.Qty
	} else {
		stock = nil
	}

	return GetOneByIdResponseDto{
		ID:             p.ID,
		Name:           p.Name,
		Description:    p.Description,
		Uuid:           p.Uuid,
		Price:          p.Price,
		DiscountPrice:  p.DiscountPrice,
		Active:         p.Active,
		IsDeleted:      p.IsDeleted,
		Image:          p.Image,
		ImageMobile:    p.ImageMobile,
		ImageThumbnail: p.ImageThumbnail,
		CategoryID:     p.CategoryID,
		CreatedBy:      p.CreatedBy,
		UpdatedBy:      p.UpdatedBy,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
		Stock:          stock,
		Category: CategoryResponseDto{
			Description: p.Category.Description,
			Name:        p.Category.Name,
			ID:          p.Category.ID,
		},
	}
}

func (s *GetOneByUuidResponseDto) FromDomain(p admin_product_repository.GetOneByUuidResponse) GetOneByUuidResponseDto {

	stock := &StockResponseDto{}

	if p.Stock != nil {
		stock.MinQty = p.Stock.MinQty
		stock.ID = p.Stock.ID
		stock.ProductID = p.Stock.ProductID
		stock.Qty = p.Stock.Qty
	} else {
		stock = nil
	}

	return GetOneByUuidResponseDto{
		ID:             p.ID,
		Name:           p.Name,
		Description:    p.Description,
		Uuid:           p.Uuid,
		Price:          p.Price,
		DiscountPrice:  p.DiscountPrice,
		Active:         p.Active,
		IsDeleted:      p.IsDeleted,
		Image:          p.Image,
		ImageMobile:    p.ImageMobile,
		ImageThumbnail: p.ImageThumbnail,
		CategoryID:     p.CategoryID,
		CreatedBy:      p.CreatedBy,
		UpdatedBy:      p.UpdatedBy,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
		Stock:          stock,
		Category: CategoryResponseDto{
			Description: p.Category.Description,
			Name:        p.Category.Name,
			ID:          p.Category.ID,
		},
	}

}

func (s *GetAllResponseDto) FromDomain(p []admin_product_repository.GetAllResponse) []GetAllResponseDto {
	productsResponse := []GetAllResponseDto{}

	for _, product := range p {
		productsResponse = append(productsResponse, GetAllResponseDto{
			ID:             product.ID,
			Name:           product.Name,
			Description:    product.Description,
			Uuid:           product.Uuid,
			Price:          product.Price,
			DiscountPrice:  product.DiscountPrice,
			Active:         product.Active,
			IsDeleted:      product.IsDeleted,
			Image:          product.Image,
			ImageMobile:    product.ImageMobile,
			ImageThumbnail: product.ImageThumbnail,
			CategoryID:     product.CategoryID,
			CreatedBy:      product.CreatedBy,
			UpdatedBy:      product.UpdatedBy,
			CreatedAt:      product.CreatedAt,
			UpdatedAt:      product.UpdatedAt,
		})
	}

	return productsResponse
}

func (s *GetAllWithRelationsResponseDto) FromDomain(p []admin_product_repository.GetAllWithDetailsResponse) []GetAllWithRelationsResponseDto {
	productsResponse := []GetAllWithRelationsResponseDto{}

	for _, product := range p {
		stock := &StockResponseDto{}

		if product.Stock != nil {
			stock.MinQty = product.Stock.MinQty
			stock.ID = product.Stock.ID
			stock.ProductID = product.Stock.ProductID
			stock.Qty = product.Stock.Qty
		} else {
			stock = nil
		}

		creditcard := []PaymentTypeInstallment{}
		paymentSlip := []PaymentTypeInstallment{}
		pix := []PaymentTypeInstallment{}

		for _, v := range product.Installments.Creditcard {
			creditcard = append(creditcard, PaymentTypeInstallment{
				InstallmentTimeID:   v.InstallmentTimeID,
				InstallmentTimeName: v.InstallmentTimeName,
				Fee:                 v.Fee,
				Tariff:              v.Tariff,
			})

		}

		for _, v := range product.Installments.PaymentSlip {
			paymentSlip = append(paymentSlip, PaymentTypeInstallment{
				InstallmentTimeID:   v.InstallmentTimeID,
				InstallmentTimeName: v.InstallmentTimeName,
				Fee:                 v.Fee,
				Tariff:              v.Tariff,
			})

		}
		for _, v := range product.Installments.Pix {
			pix = append(pix, PaymentTypeInstallment{
				InstallmentTimeID:   v.InstallmentTimeID,
				InstallmentTimeName: v.InstallmentTimeName,
				Fee:                 v.Fee,
				Tariff:              v.Tariff,
			})

		}

		productsResponse = append(productsResponse, GetAllWithRelationsResponseDto{
			ID:             product.ID,
			Name:           product.Name,
			Description:    product.Description,
			Uuid:           product.Uuid,
			Price:          product.Price,
			DiscountPrice:  product.DiscountPrice,
			Active:         product.Active,
			IsDeleted:      product.IsDeleted,
			Image:          product.Image,
			ImageMobile:    product.ImageMobile,
			ImageThumbnail: product.ImageThumbnail,
			CategoryID:     product.CategoryID,
			CreatedBy:      product.CreatedBy,
			UpdatedBy:      product.UpdatedBy,
			CreatedAt:      product.CreatedAt,
			UpdatedAt:      product.UpdatedAt,
			Stock:          stock,
			Category: CategoryResponseDto{
				Description: product.Category.Description,
				Name:        product.Category.Name,
				ID:          product.Category.ID,
			},
			Installments: InstallmentsResponseDto{
				Creditcard:  creditcard,
				PaymentSlip: paymentSlip,
				Pix:         pix,
			},
		})
	}

	return productsResponse
}
