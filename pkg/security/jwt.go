package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	SecretKey               = "your-secret-key"
	TokenExpirationDuration = time.Hour * 24 // 24 hours
)

type JWTClaims struct {
	CustomClaims map[string]interface{} `json:"custom_claims"`
	jwt.StandardClaims
}

func GenerateJWT(customClaims map[string]interface{}) (string, error) {
	claims := JWTClaims{
		CustomClaims: customClaims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpirationDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateJWT(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims.CustomClaims, nil
	}

	return nil, errors.New("invalid token")
}
