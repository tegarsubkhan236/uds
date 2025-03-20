package middleware

import (
	"myapp/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Gateway(permission string) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user")
		if user == nil {
			return utils.ResponseUnauthorized(ctx)
		}

		token, ok := user.(*jwt.Token)
		if !ok {
			return utils.ResponseUnauthorized(ctx, "Invalid Token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return utils.ResponseUnauthorized(ctx, "Invalid Claims")
		}

		role, ok := claims["role"].(map[string]interface{})
		if !ok {
			return utils.ResponseUnauthorized(ctx, "Invalid Role")
		}

		if role["name"] == "SUPER_ADMIN" {
			return ctx.Next()
		}

		permissions, ok := claims["permissions"].([]interface{})
		if !ok {
			return utils.ResponseUnauthorized(ctx, "Invalid Permission")
		}

		// Has Permission
		for _, p := range permissions {
			if p == permission {
				return ctx.Next()
			}
		}

		return utils.ResponseForbidden(ctx, "You have no access")
	}
}
