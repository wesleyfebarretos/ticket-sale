package admin_gateway_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_gateway_repository"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/utils"
)

type CreateRes struct {
	Gateway      admin_gateway_repository.CreateResponse
	PaymentTypes []admin_gateway_repository.CreatePaymentTypesResponse
}

type CreateReq struct {
	Gateway      admin_gateway_repository.CreateParams
	PaymentTypes []admin_gateway_repository.CreatePaymentTypesParams
}

func Create(c *gin.Context, body CreateReq, userID int32) CreateRes {
	body.Gateway.CreatedBy = userID
	body.Gateway.UpdatedBy = &userID

	return utils.WithTransaction(c, func(tx pgx.Tx) CreateRes {
		repository := admin_gateway_repository.New().WithTx(tx)

		gateway := repository.Create(c, body.Gateway)

		for i := range body.PaymentTypes {
			body.PaymentTypes[i].GatewayID = gateway.ID
			body.PaymentTypes[i].CreatedBy = userID
			body.PaymentTypes[i].UpdatedBy = &userID
		}

		paymentTypes := repository.CreatePaymentTypes(c, body.PaymentTypes)

		return CreateRes{
			Gateway:      gateway,
			PaymentTypes: paymentTypes,
		}
	})
}

func Update(c *gin.Context, body admin_gateway_repository.UpdateParams, userID int32) bool {
	repository := admin_gateway_repository.New()

	body.UpdatedBy = &userID

	res := repository.Update(c, body)

	return res
}

func GetAll(c *gin.Context) []admin_gateway_repository.GetAllResponse {
	repository := admin_gateway_repository.New()

	gateways := repository.GetAll(c)

	return gateways
}

func GetOneById(c *gin.Context, id int32) *admin_gateway_repository.GetOneByIdResponse {
	repository := admin_gateway_repository.New()

	gateway := repository.GetOneById(c, id)

	if gateway == nil {
		panic(exception.NotFoundException(fmt.Sprintf("gateway of id %d not found", id)))
	}

	return gateway
}

func SoftDelete(c *gin.Context, softDeleteParams admin_gateway_repository.SoftDeleteParams) bool {
	repository := admin_gateway_repository.New()

	gateway := repository.SoftDelete(c, softDeleteParams)

	if !gateway {
		panic(exception.NotFoundException(fmt.Sprintf("gateway of id %d not found", softDeleteParams.ID)))
	}

	return gateway
}
