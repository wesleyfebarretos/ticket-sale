package creditcard_repository

import "github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/creditcard_connection"

func (this *CreateParams) ToEntity() creditcard_connection.CreateParams {
	return creditcard_connection.CreateParams{
		Name:             this.Name,
		Number:           this.Number,
		Expiration:       this.Expiration,
		Priority:         this.Priority,
		NotifyExpiration: this.NotifyExpiration,
		UserID:           this.UserID,
		CreditcardTypeID: this.CreditcardTypeID,
		CreditcardFlagID: this.CreditcardFlagID,
	}
}

func (this *CreateResponse) FromEntity(p creditcard_connection.FinCreditcard) CreateResponse {
	return CreateResponse{
		ID:               p.ID,
		Uuid:             p.Uuid,
		Name:             p.Name,
		Number:           p.Number,
		Expiration:       p.Expiration,
		Priority:         p.Priority,
		NotifyExpiration: p.NotifyExpiration,
		UserID:           p.UserID,
		CreditcardTypeID: p.CreditcardTypeID,
		CreditcardFlagID: p.CreditcardFlagID,
		IsDeleted:        p.IsDeleted,
		CreatedAt:        p.CreatedAt,
		UpdatedAt:        p.UpdatedAt,
	}
}

func (this *UpdateParams) ToEntity() creditcard_connection.UpdateParams {
	return creditcard_connection.UpdateParams{
		Name:             this.Name,
		Number:           this.Number,
		Expiration:       this.Expiration,
		Priority:         this.Priority,
		NotifyExpiration: this.NotifyExpiration,
		UserID:           this.UserID,
		CreditcardTypeID: this.CreditcardTypeID,
		CreditcardFlagID: this.CreditcardFlagID,
		UpdatedAt:        this.UpdatedAt,
		Uuid:             this.Uuid,
	}
}

func (this *GetAllUserCreditcardsResponse) FromEntity(p []creditcard_connection.UserCreditcard) []GetAllUserCreditcardsResponse {
	res := []GetAllUserCreditcardsResponse{}

	for _, v := range p {
		res = append(res, GetAllUserCreditcardsResponse{
			Uuid:       v.Uuid,
			Name:       v.Name,
			Number:     v.Number,
			Expiration: v.Expiration,
			UserID:     v.UserID,
			CreatedAt:  v.CreatedAt,
			CreditcardFlag: CreditcardFlagResponse{
				ID:          v.CreditcardFlag.Id,
				Name:        v.CreditcardFlag.Name,
				Description: v.CreditcardFlag.Description,
				Regex:       v.CreditcardFlag.Regex,
			},
			CreditcardType: CreditcardTypeResponse{
				ID:   v.CreditcardType.Id,
				Name: v.CreditcardType.Name,
			},
		})
	}

	return res
}
