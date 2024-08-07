package controller

import (
	"fmt"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/middleware"
)

func GetId(c *gin.Context) int32 {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		panic(exception.BadRequestException(fmt.Sprintf("invalid id parameter %s", id)))
	}

	return int32(intId)
}

func GetUuid(c *gin.Context) uuid.UUID {
	uuidRequest := c.Param("uuid")

	parsedUuid, err := uuid.Parse(uuidRequest)
	if err != nil {
		panic(exception.BadRequestException(fmt.Sprintf("invalid uuid parameter %s", uuidRequest)))
	}

	return parsedUuid
}

func ReadBody[B any](c *gin.Context, body *B) {
	err := c.ShouldBindJSON(body)
	if err == io.EOF {
		panic(exception.BadRequestException("empty request body"))
	}
	if err != nil {
		panic(exception.BadRequestException(err.Error()))
	}
}

func GetClaims(c *gin.Context) *middleware.UserClaims {
	if claims, ok := c.Get(middleware.IDENTITY_KEY); ok {
		return claims.(*middleware.UserClaims)
	}
	panic(exception.InternalServerException("JwtError: Fail on get claims"))
}
