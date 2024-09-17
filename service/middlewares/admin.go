// Package middlewares defines the middlewares for the endpoints.
package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"repair-queue/config"
	"repair-queue/utils"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type userIDKey struct{}

// JWTAuthMiddleware is a middleware that validates JWT tokens and checks user roles.
func JWTAuthMiddleware() func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			secret := []byte(config.Envs.JWTSecret)
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Authorization header is required"))
				return
			}

			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) != 2 {
				utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("Invalid token format"))
				return
			}

			token, err := jwt.Parse(bearerToken[1], func(_ *jwt.Token) (interface{}, error) {
				return secret, nil
			})
			if err != nil {
				utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("Invalid token"))
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				if float64(time.Now().Unix()) > claims["expiredAt"].(float64) {
					utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("Token expired"))
					return
				}

				role, ok := claims["role"].(string)
				if !ok || role != "ADMIN" {
					utils.WriteError(w, http.StatusForbidden, fmt.Errorf("Insufficient permissions"))
					return
				}

				r = r.WithContext(context.WithValue(r.Context(), userIDKey{}, claims["userID"]))
			} else {
				utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("Invalid token claims"))
				return
			}

			next.ServeHTTP(w, r)
		}
	}
}
