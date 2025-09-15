package middleware

import (
	"bookstore-api/app/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "No token provided"})
			return
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "provided token is invalid"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Set("user_id", uint(claims["sub"].(float64)))
		c.Set("role", claims["role"].(string))
		c.Next()
	}
}

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "403 forbidden"})
			return
		}
		if r.(string) != role {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "403 forbidden"})
			return
		}
		c.Next()
	}
}
