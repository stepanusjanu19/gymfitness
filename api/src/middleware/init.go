package middleware

import (
	"api/lib/helpers"
	"api/lib/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
)

func AuthMiddleware(db *gorm.DB, requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.JSON(c.Writer, http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		if strings.HasPrefix(token, "Bearer ") {
			token = strings.TrimPrefix(token, "Bearer ")
		}
		claims, err := helpers.JWTValidate(token)
		if err != nil {
			response.JSON(c.Writer, http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			c.Abort()
			return
		}
		isAuthorized := false
		for _, role := range requiredRoles {
			if claims["role"] == role {
				isAuthorized = true
				break
			}
		}

		if !isAuthorized {
			response.JSON(c.Writer, http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		c.Set("userID", claims["user_id"])
		c.Set("role", claims["role"])
		c.Next()
	}
}
