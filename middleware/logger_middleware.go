package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/infra/logger"
)

func Logger(c *gin.Context) {
	start := time.Now()

	l := logger.Get()
	l.Info()

	l.Info().
		Str("method", c.Request.Method).
		Str("url", c.Request.URL.String()).
		Str("user_agent", c.Request.UserAgent()).
		Dur("elapsed_ms", time.Since(start)).
		Msg("incoming request")
}
