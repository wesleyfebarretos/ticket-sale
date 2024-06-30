package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/ticket-sale/internal/enum"
)

type RolePermissionError struct {
	StatusCode int    `json:"statusCode" example:"403"`
	Message    string `json:"message" example:"permission denied."`
}

func Authorization(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if claims, ok := c.Get(IDENTITY_KEY); !ok || claims.(*UserClaims).Role != role && claims.(*UserClaims).Role != enum.SUPER_ADMIN {
			c.JSON(http.StatusForbidden, RolePermissionError{
				StatusCode: http.StatusForbidden,
				Message:    "permission denied.",
			})
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
