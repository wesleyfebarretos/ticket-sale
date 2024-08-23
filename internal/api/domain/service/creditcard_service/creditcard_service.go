package creditcard_service

import (
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/creditcard_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/gateway_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
)

func GetAllUserCreditcards(c *gin.Context, userID int32) []creditcard_repository.GetAllUserCreditcardsResponse {
	return creditcard_repository.New().GetAllUserCreditcards(c, userID)
}

type CreateParamsDTO struct {
	Name             string    `json:"name"`
	Number           string    `json:"number"`
	Expiration       time.Time `json:"expiration"`
	Priority         int32     `json:"priority"`
	NotifyExpiration bool      `json:"notifyExpiration"`
	UserID           int32     `json:"userId"`
	CreditcardTypeID int32     `json:"creditcardTypeId"`
	CreditcardFlagID int32     `json:"creditcardFlagId"`
	CVC              string    `json:"cvc"`
}

func Create(
	c *gin.Context,
	newCreditcard CreateParamsDTO,
) creditcard_repository.CreateResponse {
	return utils.WithTransaction(c, func(tx pgx.Tx) creditcard_repository.CreateResponse {
		regex := regexp.MustCompile("[^0-9]")

		newCreditcard.Number = regex.ReplaceAllString(newCreditcard.Number, "")

		nonMaskedCardNumber := newCreditcard.Number

		newCreditcard.Number = utils.MaskCreditcardNumber(newCreditcard.Number)

		creditcardRepository := creditcard_repository.New().WithTx(tx)

		creditcard := creditcardRepository.Create(c, creditcard_repository.CreateParams{
			Name:             newCreditcard.Name,
			Number:           newCreditcard.Number,
			Expiration:       newCreditcard.Expiration,
			Priority:         newCreditcard.Priority,
			NotifyExpiration: newCreditcard.NotifyExpiration,
			UserID:           newCreditcard.UserID,
			CreditcardTypeID: newCreditcard.CreditcardTypeID,
			CreditcardFlagID: newCreditcard.CreditcardFlagID,
		})

		_, err := gateway_service.FindOrCreateCustomer(c, newCreditcard.UserID)
		if err != nil {
			panic(exception.InternalServerException(err.Error()))
		}

		expYear := strconv.Itoa(newCreditcard.Expiration.Year())
		expMonth := strconv.Itoa(int(newCreditcard.Expiration.Month()))

		_, err = gateway_service.CreateCard(c, gateway_service.CreateCardDTO{
			Number:   nonMaskedCardNumber,
			ExpMonth: expMonth,
			CVC:      newCreditcard.CVC,
			ExpYear:  expYear,
			CardID:   creditcard.ID,
			UserID:   newCreditcard.UserID,
		}, &tx)
		if err != nil {
			panic(exception.InternalServerException(err.Error()))
		}

		return creditcard
	})
}

func Update(
	c *gin.Context,
	updatedCreditcard creditcard_repository.UpdateParams,
) bool {
	regex := regexp.MustCompile("[^0-9]")

	updatedCreditcard.Number = regex.ReplaceAllString(updatedCreditcard.Number, "")

	creditcard_repository.New().Update(c, updatedCreditcard)

	return true
}

func SoftDelete(
	c *gin.Context,
	deleteParams creditcard_repository.SoftDeleteParams,
) bool {
	deleteParams.UpdatedAt = time.Now().UTC()

	creditcard_repository.New().SoftDelete(c, deleteParams)

	return true
}
