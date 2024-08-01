package admin_gateway_service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/admin_gateway_repository"
)

func Create(c *gin.Context, body admin_gateway_repository.CreateParams) admin_gateway_repository.CreateResponse {
	repository := admin_gateway_repository.New()

	gateway := repository.Create(c, body)

	return gateway
}

func Update(c *gin.Context, body admin_gateway_repository.UpdateParams) bool {
	repository := admin_gateway_repository.New()

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
