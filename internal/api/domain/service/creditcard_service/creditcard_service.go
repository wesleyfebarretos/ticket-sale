package creditcard_service

import (
	"regexp"
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

func Create(
	c *gin.Context,
	newCreditcard creditcard_repository.CreateParams,
) creditcard_repository.CreateResponse {
	return utils.WithTransaction(c, func(tx pgx.Tx) creditcard_repository.CreateResponse {
		regex := regexp.MustCompile("[^0-9]")

		newCreditcard.Number = regex.ReplaceAllString(newCreditcard.Number, "")

		newCreditcard.Number = utils.MaskCreditcardNumber(newCreditcard.Number)

		creditcardRepository := creditcard_repository.New().WithTx(tx)

		creditcard := creditcardRepository.Create(c, newCreditcard)

		_, err := gateway_service.FindOrCreateCustomer(c, newCreditcard.UserID)
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
