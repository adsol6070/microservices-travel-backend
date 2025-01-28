package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Secret Key (Should be stored in environment variables)
var jwtSecret = []byte("your-secret-key")

// GenerateJWT generates a new JWT token
func GenerateJWT(serviceName string) (string, error) {
	claims := jwt.MapClaims{
		"service": serviceName,
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateJWT validates a JWT token and extracts claims
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
