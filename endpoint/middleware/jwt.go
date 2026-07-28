package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"yosimaril/CourseEnrollment/utils/token"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authorizationHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		claims, err := token.ValidateToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Set("user_id", claims.UserID)
		c.Set("user_role", string(claims.UserRole))
		c.Next()
	}
}
