package creditcard_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/creditcard_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/controller"
)

// GetAllUserCreditcards godoc
//
//	@Tags			Creditcard
//	@Summary		Get all user creditcards
//	@Description	Get all user creditcards
//	@Produce		json
//	@Success		200	{object}	[]GetAllUserCreditcardsResponseDto
//	@Failure		500	{object}	exception.HttpException
//	@Failure		400	{object}	exception.HttpException
//	@Failure		403	{object}	middleware.RolePermissionError
//	@Failure		401	{object}	middleware.AuthenticationError
//	@Router			/creditcard/user [get]
func GetAllUserCreditcards(c *gin.Context) {
	user := controller.GetClaims(c)

	creditcards := creditcard_service.GetAllUserCreditcards(c, user.Id)

	creditcardsResponse := GetAllUserCreditcardsResponseDto{}

	c.JSON(http.StatusOK, creditcardsResponse.FromDomain(creditcards))
}

// CreateCreditCard godoc
//
//	@Tags			Creditcard
//	@Summary		Create a Creditcard
//	@Description	Create a Creditcard
//	@Produce		json
//	@Param			Creditcard	body		CreateRequestDto	true	"New Creditcard"
//	@Success		201			{object}	CreateResponseDto
//	@Failure		500			{object}	exception.HttpException
//	@Failure		400			{object}	exception.HttpException
//	@Failure		403			{object}	middleware.RolePermissionError
//	@Failure		401			{object}	middleware.AuthenticationError
//	@Router			/creditcard [post]
func Create(c *gin.Context) {
	user := controller.GetClaims(c)

	body := CreateRequestDto{}

	controller.ReadBody(c, &body)

	newCreditcard := creditcard_service.Create(c, body.ToDomain(user.Id))

	res := CreateResponseDto{}

	c.JSON(http.StatusCreated, res.FromDomain(newCreditcard))
}

// UpdateProduct godoc
//
//	@Tags			Creditcard
//	@Summary		Update a creditcard
//	@Description	Update a creditcard
//	@Produce		json
//	@Param			uuid		path		string				true	"Creditcard UUID"
//	@Param			Creditcard	body		UpdateRequestDto	true	"Update Creditcard"
//	@Success		200			{object}	bool
//	@Failure		500			{object}	exception.HttpException
//	@Failure		404			{object}	exception.HttpException
//	@Failure		400			{object}	exception.HttpException
//	@Failure		403			{object}	middleware.RolePermissionError
//	@Failure		401			{object}	middleware.AuthenticationError
//	@Router			/creditcard/{uuid} [put]
func Update(c *gin.Context) {
	user := controller.GetClaims(c)

	uuid := controller.GetUuid(c)

	body := UpdateRequestDto{}

	controller.ReadBody(c, &body)

	creditcard_service.Update(c, body.ToDomain(uuid, user.Id))

	c.JSON(http.StatusOK, true)
}

// SoftDeleteProduct godoc
//
//	@Tags			Creditcard
//	@Summary		Soft Delete a creditcard
//	@Description	Soft Delete a creditcard
//	@Produce		json
//	@Param			uuid	path		string	true	"Creditcard UUID"
//	@Success		200		{object}	bool
//	@Failure		500		{object}	exception.HttpException
//	@Failure		400		{object}	exception.HttpException
//	@Failure		403		{object}	middleware.RolePermissionError
//	@Failure		401		{object}	middleware.AuthenticationError
//	@Router			/creditcard/{uuid} [delete]
func SoftDelete(c *gin.Context) {
	uuid := controller.GetUuid(c)

	params := SoftDeleteRequestDto{}

	creditcard_service.SoftDelete(c, params.ToDomain(uuid))

	c.JSON(http.StatusOK, true)
}
