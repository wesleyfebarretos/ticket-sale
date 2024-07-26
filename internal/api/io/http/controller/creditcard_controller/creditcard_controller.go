package creditcard_controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/creditcard_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/creditcard_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/controller"
)

// TODO: Implement swagger documentation config and examples in structs
func GetAllUserCreditcards(c *gin.Context) {
	user := controller.GetClaims(c)

	creditcards := creditcard_service.GetAllUserCreditcards(c, user.Id)

	creditcardsResponse := []GetAllUserCreditcardsResponseDto{}

	for _, cc := range creditcards {
		creditcardsResponse = append(creditcardsResponse, GetAllUserCreditcardsResponseDto{
			Uuid:       cc.Uuid,
			Name:       cc.Name,
			Number:     cc.Number,
			Expiration: cc.Expiration,
			UserID:     cc.UserID,
			CreatedAt:  cc.CreatedAt,
			CreditcardFlag: CreditcardFlag{
				Id:          cc.CreditcardFlag.Id,
				Name:        cc.CreditcardFlag.Name,
				Description: cc.CreditcardFlag.Description,
				Regex:       cc.CreditcardFlag.Regex,
			},
			CreditcardType: CreditcardType{
				Id:   cc.CreditcardType.Id,
				Name: cc.CreditcardType.Name,
			},
		})
	}

	c.JSON(http.StatusOK, creditcardsResponse)
}

func Create(c *gin.Context) {
	user := controller.GetClaims(c)

	body := CreateRequestDto{}

	controller.ReadBody(c, &body)

	newCreditcardRequest := creditcard_repository.CreateParams{
		Name:             body.Name,
		Number:           body.Number,
		Expiration:       body.Expiration,
		Priority:         body.Priority,
		NotifyExpiration: body.NotifyExpiration,
		UserID:           user.Id,
		CreditcardTypeID: body.CreditcardTypeID,
		CreditcardFlagID: body.CreditcardFlagID,
	}

	newCreditcard := creditcard_service.Create(c, newCreditcardRequest)

	newCreditcardResponse := CreateResponseDto{
		ID:               newCreditcard.ID,
		Uuid:             newCreditcard.Uuid,
		Name:             newCreditcard.Name,
		Number:           newCreditcard.Number,
		Expiration:       newCreditcard.Expiration,
		Priority:         newCreditcard.Priority,
		NotifyExpiration: newCreditcard.NotifyExpiration,
		UserID:           newCreditcard.UserID,
		CreditcardTypeID: newCreditcard.CreditcardTypeID,
		CreditcardFlagID: newCreditcard.CreditcardFlagID,
		IsDeleted:        newCreditcard.IsDeleted,
		CreatedAt:        newCreditcard.CreatedAt,
		UpdatedAt:        newCreditcard.UpdatedAt,
	}

	c.JSON(http.StatusCreated, newCreditcardResponse)
}

func Update(c *gin.Context) {
	user := controller.GetClaims(c)

	uuid := controller.GetUuid(c)

	body := UpdateRequestDto{}

	controller.ReadBody(c, &body)

	updateCreditcardRequest := creditcard_repository.UpdateParams{
		Name:             body.Name,
		Number:           body.Number,
		Expiration:       body.Expiration,
		Priority:         body.Priority,
		NotifyExpiration: body.NotifyExpiration,
		UserID:           user.Id,
		CreditcardTypeID: body.CreditcardTypeID,
		CreditcardFlagID: body.CreditcardFlagID,
		UpdatedAt:        time.Now().UTC(),
		Uuid:             uuid,
	}

	creditcard_service.Update(c, updateCreditcardRequest)

	c.JSON(http.StatusOK, true)
}

func SoftDelete(c *gin.Context) {
	uuid := controller.GetUuid(c)

	creditcard_service.SoftDelete(c, creditcard_repository.SoftDeleteParams{
		Uuid:      uuid,
		UpdatedAt: time.Now().UTC(),
	})

	c.JSON(http.StatusOK, true)
}
