package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorization(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if claims, ok := c.Get(IDENTITY_KEY); !ok || claims.(*UserClaims).Role != role {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "permission denied.",
			})
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
