package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

// JWTMiddleware checks for a valid JWT token in the Authorization header
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		// Extract service name from claims and log it (for debugging)
		if service, ok := claims["service"].(string); ok {
			fmt.Printf("Authenticated request from service: %s\n", service)
		}

		next.ServeHTTP(w, r)
	})
}
