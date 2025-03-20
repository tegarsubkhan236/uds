package utils

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Me(ctx *fiber.Ctx) (string, error) {
	user := ctx.Locals("user")
	if user == nil {
		return "", errors.New("unauthorized")
	}

	token, ok := user.(*jwt.Token)
	if !ok {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("invalid username")
	}

	return username, nil
}
