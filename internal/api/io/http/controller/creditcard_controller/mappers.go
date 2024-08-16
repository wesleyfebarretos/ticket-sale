package creditcard_controller

import (
	"github.com/google/uuid"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/creditcard_repository"
)

func (s *CreateRequestDto) ToDomain(userID int32) creditcard_repository.CreateParams {
	return creditcard_repository.CreateParams{
		Name:             s.Name,
		Number:           s.Number,
		Expiration:       s.Expiration,
		Priority:         s.Priority,
		NotifyExpiration: s.NotifyExpiration,
		UserID:           userID,
		CreditcardTypeID: s.CreditcardTypeID,
		CreditcardFlagID: s.CreditcardFlagID,
	}
}

func (s *CreateResponseDto) FromDomain(p creditcard_repository.CreateResponse) CreateResponseDto {
	return CreateResponseDto{
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

func (s *GetAllUserCreditcardsResponseDto) FromDomain(p []creditcard_repository.GetAllUserCreditcardsResponse) []GetAllUserCreditcardsResponseDto {
	res := []GetAllUserCreditcardsResponseDto{}
	for _, v := range p {
		res = append(res, GetAllUserCreditcardsResponseDto{
			Uuid:       v.Uuid,
			Name:       v.Name,
			Number:     v.Number,
			Expiration: v.Expiration,
			UserID:     v.UserID,
			CreatedAt:  v.CreatedAt,
			CreditcardFlag: CreditcardFlag{
				Id:          v.CreditcardFlag.ID,
				Name:        v.CreditcardFlag.Name,
				Description: v.CreditcardFlag.Description,
				Regex:       v.CreditcardFlag.Regex,
			},
			CreditcardType: CreditcardType{
				Id:   v.CreditcardType.ID,
				Name: v.CreditcardType.Name,
			},
		})
	}
	return res
}

func (s *UpdateRequestDto) ToDomain(cardUUID uuid.UUID, userID int32) creditcard_repository.UpdateParams {
	return creditcard_repository.UpdateParams{
		Name:             s.Name,
		Number:           s.Number,
		Expiration:       s.Expiration,
		Priority:         s.Priority,
		NotifyExpiration: s.NotifyExpiration,
		UserID:           userID,
		CreditcardTypeID: s.CreditcardTypeID,
		CreditcardFlagID: s.CreditcardFlagID,
		Uuid:             cardUUID,
	}
}

func (s *SoftDeleteRequestDto) ToDomain(cardUUID uuid.UUID) creditcard_repository.SoftDeleteParams {
	return creditcard_repository.SoftDeleteParams{
		Uuid: cardUUID,
	}
}
