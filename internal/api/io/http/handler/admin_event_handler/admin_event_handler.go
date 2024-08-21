package admin_event_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/admin_event_service"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/io/http/handler"
)

// CreateEvent godoc
//
//	@Tags			Admin Event
//	@Summary		Create a event
//	@Description	Create a event
//	@Produce		json
//	@Param			Event	body		CreateRequestDto	true	"New Event"
//	@Success		201		{object}	CreateResponseDto
//	@Failure		500		{object}	exception.HttpException
//	@Failure		400		{object}	exception.HttpException
//	@Failure		403		{object}	middleware.RolePermissionError
//	@Failure		401		{object}	middleware.AuthenticationError
//	@Router			/admin/events [post]
func Create(c *gin.Context) {
	adminUserClaims := handler.GetClaims(c)

	body := CreateRequestDto{}

	handler.ReadBody(c, &body)

	domainObj := body.ToDomain(adminUserClaims.Id)

	newEvent := admin_event_service.Create(c, domainObj.NewEvent, domainObj.NewProduct, domainObj.NewStock, domainObj.NewProductInstallments)

	res := CreateResponseDto{}

	c.JSON(http.StatusCreated, res.FromDomain(newEvent))
}

// UpdateEvent godoc
//
//	@Tags			Admin Event
//	@Summary		Update a event
//	@Description	Update a event
//	@Produce		json
//	@Param			id		path		int					true	"Event ID"
//	@Param			Event	body		UpdateRequestDto	true	"Update Event"
//	@Success		200		{object}	bool
//	@Failure		500		{object}	exception.HttpException
//	@Failure		404		{object}	exception.HttpException
//	@Failure		400		{object}	exception.HttpException
//	@Failure		403		{object}	middleware.RolePermissionError
//	@Failure		401		{object}	middleware.AuthenticationError
//	@Router			/admin/events/{id} [put]
func Update(c *gin.Context) {
	id := handler.GetId(c)
	adminUserClaims := handler.GetClaims(c)

	body := UpdateRequestDto{}

	handler.ReadBody(c, &body)

	domainObj := body.ToDomain(id, adminUserClaims.Id)

	admin_event_service.Update(c, domainObj.UpdateEvent, domainObj.UpdateProduct, domainObj.UpdateProductInstallments)

	c.JSON(http.StatusOK, true)
}

// SoftDeleteEvent godoc
//
//	@Tags			Admin Event
//	@Summary		Soft Delete a event
//	@Description	Soft Delete a event
//	@Produce		json
//	@Param			id	path		int	true	"Event ID"
//	@Success		200	{object}	bool
//	@Failure		500	{object}	exception.HttpException
//	@Failure		400	{object}	exception.HttpException
//	@Failure		403	{object}	middleware.RolePermissionError
//	@Failure		401	{object}	middleware.AuthenticationError
//	@Router			/admin/events/{id} [delete]k
func SoftDelete(c *gin.Context) {
	id := handler.GetId(c)

	admin_event_service.SoftDelete(c, id)

	c.JSON(http.StatusOK, true)
}

// GetAllEvents godoc
//
//	@Tags			Admin Event
//	@Summary		Get all events
//	@Description	Get all events
//	@Produce		json
//	@Success		200	{object}	[]GetAllResponseDto
//	@Failure		500	{object}	exception.HttpException
//	@Failure		400	{object}	exception.HttpException
//	@Failure		403	{object}	middleware.RolePermissionError
//	@Failure		401	{object}	middleware.AuthenticationError
//	@Router			/admin/events [get]
func GetAll(c *gin.Context) {
	events := admin_event_service.GetAll(c)

	eventsResponse := GetAllResponseDto{}

	c.JSON(http.StatusOK, eventsResponse.FromDomain(events))
}

// GetOneById godoc
//
//	@Tags			Admin Event
//	@Summary		Get One By Id
//	@Description	Get One By Id
//	@Produce		json
//	@Param			id	path		int	true	"Event ID"
//	@Success		200	{object}	GetOneByIdResponseDto
//	@Failure		500	{object}	exception.HttpException
//	@Failure		404	{object}	exception.HttpException
//	@Failure		400	{object}	exception.HttpException
//	@Failure		403	{object}	middleware.RolePermissionError
//	@Failure		401	{object}	middleware.AuthenticationError
//	@Router			/admin/events/{id} [get]
func GetOneById(c *gin.Context) {
	id := handler.GetId(c)

	event := admin_event_service.GetOneById(c, id)

	eventResponse := GetOneByIdResponseDto{}

	c.JSON(http.StatusOK, eventResponse.FromDomain(event))
}
