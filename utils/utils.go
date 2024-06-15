package utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetId(c *gin.Context) (int32, error) {
	id := c.Param("id")

	if id == "" {
		return 0, errors.New("expected id and receive nothing")
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return int32(intId), nil
}

func WriteError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{"error": err.Error()})
}
