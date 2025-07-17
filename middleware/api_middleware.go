package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationKey := c.GetHeader("Authorization")

		if authorizationKey != secretKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		c.Next()
	}
}
