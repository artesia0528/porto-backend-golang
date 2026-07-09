package middleware

import (
	"net/http"
	"strings"

	"portfolio-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthRequired membuat middleware yang memvalidasi JWT token
// dan menyimpan user_id ke Gin context.
func AuthRequired(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization") // format: "Bearer <token>"
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Token tidak ditemukan",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := utils.ValidateToken(tokenString, jwtSecret)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Token tidak valid",
			})
			return
		}

		// Extract user_id dari claims dan simpan ke context
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Claims tidak valid",
			})
			return
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "User ID tidak ditemukan di token",
			})
			return
		}

		// Set userID ke context agar bisa diakses handler di belakang middleware
		c.Set("userID", uint(userID))
		c.Next()
	}
}