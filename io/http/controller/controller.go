package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
)

type BaseController struct{}

func (bs *BaseController) GetId(c *gin.Context) int32 {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		panic(exception.InternalServerException(fmt.Sprintf("invalid input value for id %s", id)))
	}

	return int32(intId)
}

func (bs *BaseController) ReadBody(c *gin.Context, body any) {
	err := c.ShouldBindJSON(&body)
	if err != nil {
		panic(exception.BadRequestException(err.Error()))
	}
}
