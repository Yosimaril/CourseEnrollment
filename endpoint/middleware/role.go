package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"yosimaril/CourseEnrollment/constants"
)

func RequireRole(requiredRole constants.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleValue, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		role, ok := roleValue.(string)
		if !ok || constants.UserRole(role) != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
