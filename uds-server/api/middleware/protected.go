package middleware

import (
	"myapp/pkg/config"
	"myapp/utils"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func Protected() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.ResponseUnauthorized(c, "Authorization header is missing")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 && parts[0] != "Bearer" {
			return utils.ResponseUnauthorized(c, "Invalid authorization header format")
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString(config.JwtSecret)), nil
		})
		if err != nil || !token.Valid {
			return utils.ResponseUnauthorized(c, "Invalid token")
		}

		claims := token.Claims.(jwt.MapClaims)
		expTime := time.Unix(int64(claims["exp"].(float64)), 0)
		currentTime := time.Now()
		if currentTime.After(expTime) {
			return utils.ResponseUnauthorized(c, "Token has expired")
		}

		c.Locals("user", token)
		return c.Next()
	}
}
