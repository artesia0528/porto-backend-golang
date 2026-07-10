package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken membuat JWT token baru dengan user_id sebagai claim.
func GenerateToken(userID string, secret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidateToken memvalidasi JWT token dan memastikan signing method yang benar.
func ValidateToken(tokenString string, secret string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Pastikan signing method adalah HMAC (HS256/HS384/HS512)
		// Ini mencegah "alg: none" attack dan signing method mismatch
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}