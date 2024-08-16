package creditcard_service

import (
	"fmt"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"

	stripe_provider "github.com/wesleyfebarretos/ticket-sale/external/providers/gateways/stripe"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/roles_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/creditcard_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/gateway_customer_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/user_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
)

func GetAllUserCreditcards(c *gin.Context, userID int32) []creditcard_repository.GetAllUserCreditcardsResponse {
	return creditcard_repository.New().GetAllUserCreditcards(c, userID)
}

func Create(
	c *gin.Context,
	newCreditcard creditcard_repository.CreateParams,
	tokenize bool,
) creditcard_repository.CreateResponse {
	return utils.WithTransaction(c, func(tx pgx.Tx) creditcard_repository.CreateResponse {
		regex := regexp.MustCompile("[^0-9]")

		newCreditcard.Number = regex.ReplaceAllString(newCreditcard.Number, "")

		newCreditcard.Number = utils.MaskCreditcardNumber(newCreditcard.Number)

		creditcardRepository := creditcard_repository.New().WithTx(tx)

		creditcard := creditcardRepository.Create(c, newCreditcard)

		userRepository := user_repository.New().WithTx(tx)

		user := userRepository.GetOneById(c, user_repository.GetOneByIdParams{
			ID:   newCreditcard.UserID,
			Role: roles_enum.USER,
		})

		//  TODO: Make a function to take active gateway and check if customer exists
		//  if not create

		gatewayCustomerRepository := gateway_customer_repository.New().WithTx(tx)

		persistedCustomer := gatewayCustomerRepository.FindOneByGatewayAndUserIdResponse(c, gateway_customer_repository.FindOneByGatewayAndUserIdParams{
			UserID:    user.ID,
			GatewayID: 1,
		})

		if persistedCustomer == nil {
			gatewayCustomer, err := stripe_provider.CreateCustomer(&stripe_provider.CreateCustomerDTO{
				Name:  fmt.Sprintf("%s %s", user.FirstName, user.LastName),
				Email: user.Email,
			})
			if err != nil {
				panic(exception.InternalServerException(err.Error()))
			}

			//  TODO: Unmock gateway ID and create a functionality to take active gateway
			//  with a helper function to create a customer in this gateway and return
			gatewayCustomerRepository.Create(c, gateway_customer_repository.CreateParams{
				UserID:            user.ID,
				GatewayID:         1,
				GatewayCustomerID: gatewayCustomer.ID,
			})

			//  TODO: Tokenize card in gateway
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
