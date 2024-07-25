package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/api/internal/exception"
)

func ExceptionHandler(c *gin.Context, recovered any) {
	if exception, ok := recovered.(*exception.HttpException); ok {
		c.JSON(exception.StatusCode, gin.H{"code": exception.StatusCode, "message": exception.Message})
	} else {
		log.Printf("Exception not mapped: %s", recovered)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
