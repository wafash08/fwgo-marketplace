package middlewares

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ExtractToken(c *fiber.Ctx) string {
	bearerToken := c.Get("Authorization")
	if strings.HasPrefix(bearerToken, "Bearer ") {
		return strings.TrimPrefix(bearerToken, "Bearer ")
	}
	return ""
}

func JwtMiddleware() fiber.Handler {
	secretKey := os.Getenv("SECRET_KEY")
	return func(c *fiber.Ctx) error {
		tokenString := ExtractToken(c)
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    fiber.StatusUnauthorized,
				"message": "Unauthorized",
			})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    fiber.StatusUnauthorized,
				"message": "Unauthorized",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			fmt.Println("claims > ", claims)
			c.Locals("user", claims)
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":     "unauthorized",
				"statusCode": 401,
				"message":    "Token claims unauthorized",
			})
		}

		return c.Next()
	}
}

func UserLocals(c *fiber.Ctx) jwt.MapClaims {
	user := c.Locals("user").(jwt.MapClaims)

	return user
}
