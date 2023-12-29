package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/bobykurniawan11/starter-go/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// GenerateToken generates a JWT token for the given user ID.
// It takes the user ID as input and returns the generated token as a string.
// If an error occurs during token generation, it is also returned.
func GenerateToken(id uuid.UUID) (string, error) {

	config := config.GetConfig()
	secret := config.GetString("api.secret")
	token_lifespan := config.GetInt("token.life")

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))

}

// ExtractToken extracts the token from the request context or query parameter.
// It first checks if the token is present in the query parameter "token".
// If not found, it then checks the "Authorization" header for a bearer token.
// If a valid bearer token is found, it returns the token.
// If no token is found, it returns an empty string.
func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// TokenValid checks if the token in the request context is valid.
// It extracts the token from the request, parses it, and verifies the signing method and secret.
// If the token is valid, it returns nil. Otherwise, it returns an error.
func TokenValid(c *gin.Context) error {

	config := config.GetConfig()
	secret := config.GetString("api.secret")

	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}
	return nil
}

// ExtractTokenID extracts the user ID from the JWT token in the given Gin context.
// It returns the user ID as a UUID and an error if any.
func ExtractTokenID(c *gin.Context) (uuid.UUID, error) {
	config := config.GetConfig()
	secret := config.GetString("api.secret")
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid := claims["user_id"]

		// Convert uid to string
		uidStr, ok := uid.(string)
		if !ok {
			return uuid.UUID{}, fmt.Errorf("failed to convert user_id to string")
		}

		// Parse string to uuid.UUID
		uidUUID, err := uuid.Parse(uidStr)
		if err != nil {
			return uuid.UUID{}, err
		}

		return uidUUID, nil
	}
	return uuid.UUID{}, nil
}
