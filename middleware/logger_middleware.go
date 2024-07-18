package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/logger"
)

func Logger(c *gin.Context) {
	start := time.Now()

	l := logger.Get()

	c.Next()

	l.Info().
		Str("method", c.Request.Method).
		Str("url", c.Request.URL.String()).
		Str("user_agent", c.Request.UserAgent()).
		Dur("elapsed_ms", time.Since(start)).
		Int("status_code", c.Writer.Status()).
		Msg("incoming request")
}
