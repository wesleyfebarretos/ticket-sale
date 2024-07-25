package admin_event_controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/api/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/api/internal/service/admin_event_service"
	"github.com/wesleyfebarretos/ticket-sale/api/io/http/controller"
	"github.com/wesleyfebarretos/ticket-sale/api/io/http/controller/admin_product_controller"
	"github.com/wesleyfebarretos/ticket-sale/api/repository/sqlc/admin_events_repository"
	"github.com/wesleyfebarretos/ticket-sale/api/repository/sqlc/admin_product_stocks_repository"
	"github.com/wesleyfebarretos/ticket-sale/api/repository/sqlc/admin_products_repository"
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
	adminUserClaims := controller.GetClaims(c)

	body := CreateRequestDto{}

	controller.ReadBody(c, &body)

	newEventReq := admin_events_repository.CreateParams{
		StartAt:   body.StartAt,
		EndAt:     body.EndAt,
		City:      body.City,
		State:     body.State,
		Location:  body.Location,
		CreatedBy: adminUserClaims.Id,
	}

	newProductReq := admin_products_repository.CreateParams{
		Name:           body.Product.Name,
		Description:    body.Product.Description,
		Price:          body.Product.Price,
		DiscountPrice:  body.Product.DiscountPrice,
		Active:         body.Product.Active,
		Image:          body.Product.Image,
		ImageMobile:    body.Product.ImageMobile,
		ImageThumbnail: body.Product.ImageThumbnail,
		CategoryID:     body.Product.CategoryID,
		CreatedBy:      adminUserClaims.Id,
	}

	newStockReq := admin_product_stocks_repository.CreateParams{
		Qty:       body.Product.Stock.Qty,
		MinQty:    body.Product.Stock.MinQty,
		CreatedBy: adminUserClaims.Id,
	}

	newProductInstallments := []admin_products_repository.CreateInstallmentsParams{}

	for _, installment := range body.Product.Installments {
		newProductInstallments = append(newProductInstallments, admin_products_repository.CreateInstallmentsParams{
			PaymentTypeID:     installment.PaymentTypeID,
			InstallmentTimeID: installment.ID,
			Fee:               *installment.Fee,
			Tariff:            *installment.Tariff,
			CreatedBy:         adminUserClaims.Id,
		})
	}

	res := admin_event_service.Create(c, newEventReq, newProductReq, newStockReq, newProductInstallments)

	installmentsResponse := []admin_product_controller.CreateInstallmentsResponseDto{}

	for _, newInstallment := range res.ProductInstallments {
		installmentsResponse = append(installmentsResponse, admin_product_controller.CreateInstallmentsResponseDto{
			ID:            newInstallment.ID,
			PaymentTypeID: newInstallment.PaymentTypeID,
			InstallmentID: newInstallment.InstallmentTimeID,
			Fee:           newInstallment.Fee,
			Tariff:        newInstallment.Tariff,
		})
	}

	newEventRes := CreateResponseDto{
		ID:        res.Event.ID,
		ProductID: res.Event.ProductID,
		StartAt:   res.Event.StartAt,
		EndAt:     res.Event.EndAt,
		City:      res.Event.City,
		State:     res.Event.State,
		Location:  res.Event.Location,
		Product: admin_product_controller.CreateResponseDto{
			ID:             res.Product.ID,
			Name:           res.Product.Name,
			Description:    res.Product.Description,
			Price:          res.Product.Price,
			DiscountPrice:  res.Product.DiscountPrice,
			Active:         res.Product.Active,
			Image:          res.Product.Image,
			ImageMobile:    res.Product.ImageMobile,
			ImageThumbnail: res.Product.ImageThumbnail,
			CategoryID:     res.Product.CategoryID,
			CreatedBy:      res.Product.CreatedBy,
			Uuid:           res.Product.Uuid,
			IsDeleted:      res.Product.IsDeleted,
			UpdatedBy:      res.Product.UpdatedBy,
			CreatedAt:      res.Product.CreatedAt,
			UpdatedAt:      res.Product.UpdatedAt,
			Stock: admin_product_controller.CreateStockResponseDto{
				ID:        res.ProductStock.ID,
				ProductID: res.ProductStock.ProductID,
				Qty:       res.ProductStock.Qty,
				MinQty:    res.ProductStock.MinQty,
			},
			Installments: installmentsResponse,
		},
	}

	c.JSON(http.StatusCreated, newEventRes)
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
	id := controller.GetId(c)
	adminUserClaims := controller.GetClaims(c)

	body := UpdateRequestDto{}

	controller.ReadBody(c, &body)

	updateEventReq := admin_events_repository.UpdateParams{
		ID:        id,
		StartAt:   body.StartAt,
		EndAt:     body.EndAt,
		City:      body.City,
		State:     body.State,
		Location:  body.Location,
		UpdatedAt: time.Now().UTC(),
	}

	updateProductReq := admin_products_repository.UpdateParams{
		Name:           body.Product.Name,
		Description:    body.Product.Description,
		Price:          body.Product.Price,
		DiscountPrice:  body.Product.DiscountPrice,
		Active:         body.Product.Active,
		Image:          body.Product.Image,
		ImageMobile:    body.Product.ImageMobile,
		ImageThumbnail: body.Product.ImageThumbnail,
		CategoryID:     body.Product.CategoryID,
		UpdatedBy:      &adminUserClaims.Id,
	}

	updateProductInstallments := []admin_products_repository.CreateInstallmentsParams{}

	for _, installment := range body.Product.Installments {
		updateProductInstallments = append(updateProductInstallments, admin_products_repository.CreateInstallmentsParams{
			PaymentTypeID:     installment.PaymentTypeID,
			InstallmentTimeID: installment.ID,
			Fee:               *installment.Fee,
			Tariff:            *installment.Tariff,
			CreatedBy:         adminUserClaims.Id,
		})
	}

	admin_event_service.Update(c, updateEventReq, updateProductReq, updateProductInstallments)

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
	id := controller.GetId(c)

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

	eventsResponse := []GetAllResponseDto{}

	bEvents, err := json.Marshal(events)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	if err := json.Unmarshal(bEvents, &eventsResponse); err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	c.JSON(http.StatusOK, eventsResponse)
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
	id := controller.GetId(c)

	event := admin_event_service.GetOneById(c, id)

	eventResponse := GetOneByIdResponseDto{}

	bEvent, err := json.Marshal(event)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	if err := json.Unmarshal(bEvent, &eventResponse); err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	c.JSON(http.StatusOK, eventResponse)
}
