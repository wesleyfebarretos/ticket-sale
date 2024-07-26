package creditcard_service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/sqlc/creditcard_repository"
)

func GetAllUserCreditcards(c *gin.Context, userID int32) []creditcard_repository.UserCreditcard {
	creditcards, err := repository.Creditcard.GetAllUserCreditcards(c, userID)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return creditcards
}

func Create(
	c *gin.Context,
	newCreditcard creditcard_repository.CreateParams,
) creditcard_repository.FinCreditcard {
	creditcard, err := repository.Creditcard.Create(c, newCreditcard)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return creditcard
}

func Update(
	c *gin.Context,
	updatedCreditcard creditcard_repository.UpdateParams,
) bool {
	err := repository.Creditcard.Update(c, updatedCreditcard)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return true
}

func SoftDelete(
	c *gin.Context,
	deleteParams creditcard_repository.SoftDeleteParams,
) bool {
	deleteParams.UpdatedAt = time.Now().UTC()

	err := repository.Creditcard.SoftDelete(c, deleteParams)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	return true
}
