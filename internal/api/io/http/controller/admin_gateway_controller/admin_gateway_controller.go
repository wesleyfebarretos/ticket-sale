package admin_gateway_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/admin_gateway_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/controller"
)

// CreateGateway godoc
//
//	@Tags			Admin Gateway
//	@Summary		Create a gateway
//	@Description	Create a gateway
//	@Produce		json
//	@Param			Gateway	body		CreateReqDTO	true	"New Gateway"
//	@Success		201		{object}	CreateResDTO
//	@Failure		500		{object}	exception.HttpException
//	@Failure		400		{object}	exception.HttpException
//	@Failure		403		{object}	middleware.RolePermissionError
//	@Failure		401		{object}	middleware.AuthenticationError
//	@Router			/admin/gateway [post]
func Create(c *gin.Context) {
	adminUser := controller.GetClaims(c)

	body := CreateReqDTO{}

	controller.ReadBody(c, &body)

	gateway := admin_gateway_service.Create(c, body.ToDomain(), adminUser.Id)

	res := CreateResDTO{}

	c.JSON(http.StatusCreated, res.FromDomain(gateway))
}

// UpdateGateway godoc
//
//	@Tags			Admin Gateway
//	@Summary		Update a gateway
//	@Description	Update a gateway
//	@Produce		json
//	@Param			id		path		int				true	"Gateway ID"
//	@Param			Gateway	body		UpdateReqDTO	true	"Update gateway"
//	@Success		200		{object}	bool
//	@Failure		500		{object}	exception.HttpException
//	@Failure		404		{object}	exception.HttpException
//	@Failure		400		{object}	exception.HttpException
//	@Failure		403		{object}	middleware.RolePermissionError
//	@Failure		401		{object}	middleware.AuthenticationError
//	@Router			/admin/gateway/{id} [put]
func Update(c *gin.Context) {
	adminUser := controller.GetClaims(c)

	body := UpdateReqDTO{}

	controller.ReadBody(c, &body)

	res := admin_gateway_service.Update(c, body.ToDomain(), adminUser.Id)

	c.JSON(http.StatusOK, res)
}

// SoftDeleteGateway godoc
//
//	@Tags			Admin Gateway
//	@Summary		Soft Delete a gateway
//	@Description	Soft Delete a gateway
//	@Produce		json
//	@Param			id	path		int	true	"Gateway ID"
//	@Success		200	{object}	bool
//	@Failure		500	{object}	exception.HttpException
//	@Failure		400	{object}	exception.HttpException
//	@Failure		403	{object}	middleware.RolePermissionError
//	@Failure		401	{object}	middleware.AuthenticationError
//	@Router			/admin/gateway/{id} [delete]
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

// GetAllGateways godoc
//
//	@Tags			Admin Gateway
//	@Summary		Get all gateway
//	@Description	Get all gateway
//	@Produce		json
//	@Success		200	{object}	[]GetAllResDTO
//	@Failure		500	{object}	exception.HttpException
//	@Failure		400	{object}	exception.HttpException
//	@Failure		403	{object}	middleware.RolePermissionError
//	@Failure		401	{object}	middleware.AuthenticationError
//	@Router			/admin/gateway [get]
func GetAll(c *gin.Context) {

	gateways := admin_gateway_service.GetAll(c)

	res := GetAllResDTO{}

	c.JSON(http.StatusOK, res.FromDomain(gateways))
}

// GetOneById godoc
//
//	@Tags			Admin Gateway
//	@Summary		Get One By Id
//	@Description	Get One By Id
//	@Produce		json
//	@Param			id	path		int	true	"Gateway ID"
//	@Success		200	{object}	GetOneByIdResDTO
//	@Failure		500	{object}	exception.HttpException
//	@Failure		404	{object}	exception.HttpException
//	@Failure		400	{object}	exception.HttpException
//	@Failure		403	{object}	middleware.RolePermissionError
//	@Failure		401	{object}	middleware.AuthenticationError
//	@Router			/admin/gateway/{id} [get]
func GetOneById(c *gin.Context) {
	id := controller.GetId(c)

	gateway := admin_gateway_service.GetOneById(c, id)

	res := GetOneByIdResDTO{}

	c.JSON(http.StatusOK, res.FromDomain(gateway))
}
