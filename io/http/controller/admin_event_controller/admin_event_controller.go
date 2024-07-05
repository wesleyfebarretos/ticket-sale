package admin_event_controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/service/admin_event_service"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller"
	"github.com/wesleyfebarretos/ticket-sale/io/http/controller/admin_product_controller"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_events_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_product_stocks_repository"
	"github.com/wesleyfebarretos/ticket-sale/repository/admin_products_repository"
)

func Create(c *gin.Context) {
	adminUserClaims := controller.GetClaims(c)

	body := CreateRequestDto{}

	controller.ReadBody(c, &body)

	newEventReq := admin_events_repository.CreateParams{
		StartAt:  body.StartAt,
		EndAt:    body.EndAt,
		City:     body.City,
		State:    body.State,
		Location: body.Location,
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

	res := admin_event_service.Create(c, newEventReq, newProductReq, newStockReq)

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
		},
	}

	c.JSON(http.StatusCreated, newEventRes)
}

func Update(c *gin.Context) {
	id := controller.GetId(c)
	adminUserClaims := controller.GetClaims(c)

	body := UpdateRequestDto{}

	controller.ReadBody(c, &body)

	updateEventReq := admin_events_repository.UpdateParams{
		ID:       id,
		StartAt:  body.StartAt,
		EndAt:    body.EndAt,
		City:     body.City,
		State:    body.State,
		Location: body.Location,
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

	admin_event_service.Update(c, updateEventReq, updateProductReq)

	c.JSON(http.StatusOK, true)
}

func SoftDelete(c *gin.Context) {
	id := controller.GetId(c)

	admin_event_service.SoftDelete(c, id)

	c.JSON(http.StatusOK, true)
}

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
