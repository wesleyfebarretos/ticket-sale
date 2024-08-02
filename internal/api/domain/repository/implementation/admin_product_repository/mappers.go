package admin_product_repository

import (
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/admin_product_connection"
)

func (s *CreateResponse) FromEntity(p admin_product_connection.Product) CreateResponse {
	return CreateResponse{
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
	}
}

func (s *CreateParams) ToEntity() admin_product_connection.CreateParams {
	return admin_product_connection.CreateParams{
		Name:           s.Name,
		Description:    s.Description,
		Price:          s.Price,
		DiscountPrice:  s.DiscountPrice,
		Active:         s.Active,
		Image:          s.Image,
		ImageMobile:    s.ImageMobile,
		ImageThumbnail: s.ImageThumbnail,
		CategoryID:     s.CategoryID,
		CreatedBy:      s.CreatedBy,
	}
}

func (s *UpdateParams) ToEntity() admin_product_connection.UpdateParams {
	return admin_product_connection.UpdateParams{
		Name:           s.Name,
		Description:    s.Description,
		Price:          s.Price,
		DiscountPrice:  s.DiscountPrice,
		Active:         s.Active,
		Image:          s.Image,
		ImageMobile:    s.ImageMobile,
		ImageThumbnail: s.ImageThumbnail,
		CategoryID:     s.CategoryID,
		UpdatedBy:      s.UpdatedBy,
		ID:             s.ID,
	}
}

func (_ *GetAllResponse) FromEntity(p []admin_product_connection.Product) []GetAllResponse {
	r := []GetAllResponse{}

	for _, v := range p {
		r = append(r, GetAllResponse{
			ID:             v.ID,
			Name:           v.Name,
			Description:    v.Description,
			Uuid:           v.Uuid,
			Price:          v.Price,
			DiscountPrice:  v.DiscountPrice,
			Active:         v.Active,
			IsDeleted:      v.IsDeleted,
			Image:          v.Image,
			ImageMobile:    v.ImageMobile,
			ImageThumbnail: v.ImageThumbnail,
			CategoryID:     v.CategoryID,
			CreatedBy:      v.CreatedBy,
			UpdatedBy:      v.UpdatedBy,
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
		})

	}

	return r
}

func (s *SoftDeleteParams) ToEntity() admin_product_connection.SoftDeleteParams {
	return admin_product_connection.SoftDeleteParams{
		ID:        s.ID,
		UpdatedBy: s.UpdatedBy,
	}
}

func (_ *GetOneByIdResponse) FromEntity(p admin_product_connection.ProductsDetail) *GetOneByIdResponse {
	creditcard := []PaymentTypeInstallment{}
	paymentSlip := []PaymentTypeInstallment{}
	pix := []PaymentTypeInstallment{}

	for _, v := range p.Installments.Creditcard {
		creditcard = append(creditcard, PaymentTypeInstallment{
			InstallmentTimeID:   v.InstallmentTimeID,
			InstallmentTimeName: v.InstallmentTimeName,
			Fee:                 v.Fee,
			Tariff:              v.Tariff,
		})

	}

	for _, v := range p.Installments.PaymentSlip {
		paymentSlip = append(paymentSlip, PaymentTypeInstallment{
			InstallmentTimeID:   v.InstallmentTimeID,
			InstallmentTimeName: v.InstallmentTimeName,
			Fee:                 v.Fee,
			Tariff:              v.Tariff,
		})

	}
	for _, v := range p.Installments.Pix {
		pix = append(pix, PaymentTypeInstallment{
			InstallmentTimeID:   v.InstallmentTimeID,
			InstallmentTimeName: v.InstallmentTimeName,
			Fee:                 v.Fee,
			Tariff:              v.Tariff,
		})

	}

	stock := &ProductStockResponse{}

	if p.Stock != nil {
		stock.MinQty = p.Stock.MinQty
		stock.ID = p.Stock.ID
		stock.ProductID = p.Stock.ProductID
		stock.Qty = p.Stock.Qty
	} else {
		stock = nil
	}

	return &GetOneByIdResponse{
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
		Category: ProductCategoryResponse{
			Description: p.Category.Description,
			Name:        p.Category.Name,
			ID:          p.Category.ID,
		},
		Installments: ProductInstallmentsResponse{
			Creditcard:  creditcard,
			PaymentSlip: paymentSlip,
			Pix:         pix,
		},
	}
}

func (s *GetOneByUuidResponse) FromEntity(p admin_product_connection.ProductsDetail) *GetOneByUuidResponse {
	creditcard := []PaymentTypeInstallment{}
	paymentSlip := []PaymentTypeInstallment{}
	pix := []PaymentTypeInstallment{}

	for _, v := range p.Installments.Creditcard {
		creditcard = append(creditcard, PaymentTypeInstallment{
			InstallmentTimeID:   v.InstallmentTimeID,
			InstallmentTimeName: v.InstallmentTimeName,
			Fee:                 v.Fee,
			Tariff:              v.Tariff,
		})

	}

	for _, v := range p.Installments.PaymentSlip {
		paymentSlip = append(paymentSlip, PaymentTypeInstallment{
			InstallmentTimeID:   v.InstallmentTimeID,
			InstallmentTimeName: v.InstallmentTimeName,
			Fee:                 v.Fee,
			Tariff:              v.Tariff,
		})

	}
	for _, v := range p.Installments.Pix {
		pix = append(pix, PaymentTypeInstallment{
			InstallmentTimeID:   v.InstallmentTimeID,
			InstallmentTimeName: v.InstallmentTimeName,
			Fee:                 v.Fee,
			Tariff:              v.Tariff,
		})

	}

	stock := &ProductStockResponse{}

	if p.Stock != nil {
		stock.MinQty = p.Stock.MinQty
		stock.ID = p.Stock.ID
		stock.ProductID = p.Stock.ProductID
		stock.Qty = p.Stock.Qty
	} else {
		stock = nil
	}

	return &GetOneByUuidResponse{
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
		Category: ProductCategoryResponse{
			Description: p.Category.Description,
			Name:        p.Category.Name,
			ID:          p.Category.ID,
		},
		Installments: ProductInstallmentsResponse{
			Creditcard:  creditcard,
			PaymentSlip: paymentSlip,
			Pix:         pix,
		},
	}
}

func (s *GetAllInstallmentTimeResponse) FromEntity(p []admin_product_connection.GetAllProductInstallmentTimesRow) []GetAllInstallmentTimeResponse {
	res := []GetAllInstallmentTimeResponse{}
	for _, v := range p {
		res = append(res, GetAllInstallmentTimeResponse{
			ID:                v.ID,
			Fee:               v.Fee,
			Tariff:            v.Tariff,
			PaymentTypeID:     v.PaymentTypeID,
			InstallmentTimeID: v.InstallmentTimeID,
		})
	}

	return res
}

func (s *GetAllWithDetailsResponse) FromEntity(p []admin_product_connection.ProductsDetail) []GetAllWithDetailsResponse {
	res := []GetAllWithDetailsResponse{}

	for _, pd := range p {
		creditcard := []PaymentTypeInstallment{}
		paymentSlip := []PaymentTypeInstallment{}
		pix := []PaymentTypeInstallment{}

		for _, v := range pd.Installments.Creditcard {
			creditcard = append(creditcard, PaymentTypeInstallment{
				InstallmentTimeID:   v.InstallmentTimeID,
				InstallmentTimeName: v.InstallmentTimeName,
				Fee:                 v.Fee,
				Tariff:              v.Tariff,
			})

		}

		for _, v := range pd.Installments.PaymentSlip {
			paymentSlip = append(paymentSlip, PaymentTypeInstallment{
				InstallmentTimeID:   v.InstallmentTimeID,
				InstallmentTimeName: v.InstallmentTimeName,
				Fee:                 v.Fee,
				Tariff:              v.Tariff,
			})

		}
		for _, v := range pd.Installments.Pix {
			pix = append(pix, PaymentTypeInstallment{
				InstallmentTimeID:   v.InstallmentTimeID,
				InstallmentTimeName: v.InstallmentTimeName,
				Fee:                 v.Fee,
				Tariff:              v.Tariff,
			})

		}

		stock := &ProductStockResponse{}

		if pd.Stock != nil {
			stock.MinQty = pd.Stock.MinQty
			stock.ID = pd.Stock.ID
			stock.ProductID = pd.Stock.ProductID
			stock.Qty = pd.Stock.Qty
		} else {
			stock = nil
		}

		res = append(res, GetAllWithDetailsResponse{
			ID:             pd.ID,
			Name:           pd.Name,
			Description:    pd.Description,
			Uuid:           pd.Uuid,
			Price:          pd.Price,
			DiscountPrice:  pd.DiscountPrice,
			Active:         pd.Active,
			IsDeleted:      pd.IsDeleted,
			Image:          pd.Image,
			ImageMobile:    pd.ImageMobile,
			ImageThumbnail: pd.ImageThumbnail,
			CategoryID:     pd.CategoryID,
			CreatedBy:      pd.CreatedBy,
			UpdatedBy:      pd.UpdatedBy,
			CreatedAt:      pd.CreatedAt,
			UpdatedAt:      pd.UpdatedAt,
			Stock:          stock,
			Category: ProductCategoryResponse{
				Description: pd.Category.Description,
				Name:        pd.Category.Name,
				ID:          pd.Category.ID,
			},
			Installments: ProductInstallmentsResponse{
				Creditcard:  creditcard,
				PaymentSlip: paymentSlip,
				Pix:         pix,
			},
		})
	}

	return res
}

func (s *CreateInstallmentsParams) ToEntity() admin_product_connection.CreateInstallmentsParams {
	return admin_product_connection.CreateInstallmentsParams{
		ProductID:         s.ProductID,
		PaymentTypeID:     s.PaymentTypeID,
		InstallmentTimeID: s.InstallmentTimeID,
		Fee:               s.Fee,
		Tariff:            s.Tariff,
		CreatedBy:         s.CreatedBy,
	}
}

func (s *CreateInstallmentsResponse) FromEntity(p []admin_product_connection.FinProductPaymentTypeInstallmentTime) []CreateInstallmentsResponse {
	res := []CreateInstallmentsResponse{}

	for _, v := range p {
		res = append(res, CreateInstallmentsResponse{
			ID:                v.ID,
			ProductID:         v.ProductID,
			PaymentTypeID:     v.PaymentTypeID,
			InstallmentTimeID: v.InstallmentTimeID,
			Fee:               v.Fee,
			Tariff:            v.Tariff,
			CreatedBy:         v.CreatedBy,
			UpdatedBy:         v.UpdatedBy,
			CreatedAt:         v.CreatedAt,
			UpdatedAt:         v.UpdatedAt,
		})
	}

	return res
}
