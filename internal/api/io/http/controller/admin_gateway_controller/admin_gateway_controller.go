package admin_gateway_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/admin_gateway_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/controller"
)

func Create(c *gin.Context) {
	adminUser := controller.GetClaims(c)

	body := CreateReqDTO{}

	controller.ReadBody(c, &body)

	gateway := admin_gateway_service.Create(c, body.ToDomain(), adminUser.Id)

	res := CreateResDTO{}

	c.JSON(http.StatusCreated, res.FromDomain(gateway))
}

func Update(c *gin.Context) {
	adminUser := controller.GetClaims(c)

	body := UpdateReqDTO{}

	controller.ReadBody(c, &body)

	res := admin_gateway_service.Update(c, body.ToDomain(), adminUser.Id)

	c.JSON(http.StatusOK, res)
}

func SoftDelete(c *gin.Context) {
	adminUser := controller.GetClaims(c)

	id := controller.GetId(c)

	params := SoftDeleteDTO{
		ID:        id,
		UpdatedBy: adminUser.Id,
	}

	res := admin_gateway_service.SoftDelete(c, params.ToDomain())

	c.JSON(http.StatusOK, res)
}

func GetAll(c *gin.Context) {

	gateways := admin_gateway_service.GetAll(c)

	res := GetAllResDTO{}

	c.JSON(http.StatusOK, res.FromDomain(gateways))
}

func GetOneById(c *gin.Context) {
	id := controller.GetId(c)

	gateway := admin_gateway_service.GetOneById(c, id)

	res := GetOneByIdResDTO{}

	c.JSON(http.StatusOK, res.FromDomain(gateway))
}
