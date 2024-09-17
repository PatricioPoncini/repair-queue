// Package auth handles authentication-related functionality,
package auth

import (
	"repair-queue/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CreateJWT generates a JSON Web Token (JWT) for a user with a specified user ID.
func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
		"role":      "ADMIN",
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
