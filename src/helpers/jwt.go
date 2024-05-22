package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(secretKey string, payload map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println("payload > ", payload)

	for key, value := range payload {
		claims[key] = value
	}

	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(secretKey string, payload map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	for key, value := range payload {
		claims[key] = value
	}

	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	refreshTokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}
