package controller

import (
	"fmt"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/wesleyfebarretos/ticket-sale/internal/exception"
	"github.com/wesleyfebarretos/ticket-sale/repository/sqlc"
)

type BaseController struct {
	conn *pgx.Conn
}

func (bc *BaseController) GetId(c *gin.Context) int32 {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		panic(exception.InternalServerException(fmt.Sprintf("invalid id parameter %s", id)))
	}

	return int32(intId)
}

func (bc *BaseController) ReadBody(c *gin.Context, body any) {
	err := c.ShouldBindJSON(&body)
	if err == io.EOF {
		panic(exception.BadRequestException("empty request body"))
	}
	if err != nil {
		panic(exception.BadRequestException(err.Error()))
	}
}

func (bc *BaseController) NewConnection() *sqlc.Queries {
	return sqlc.New(bc.conn)
}
