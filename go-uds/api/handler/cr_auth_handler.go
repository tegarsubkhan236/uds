package handler

import (
	"github.com/gofiber/fiber/v2"
	"myapp/dto"
	"myapp/pkg/service"
	"myapp/utils"
)

func Login(service service.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(dto.AuthRequest)
		if err := utils.ParseAndValidate(ctx, req); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		token, err := service.AuthenticateUser(req.Identity, req.Password)
		if err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseOK(ctx, token)
	}
}

func Me() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user")
		if user == nil {
			return utils.ResponseUnauthorized(ctx)
		}

		return utils.ResponseOK(ctx, user)
	}
}

func Logout(service service.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user")
		if user == nil {
			return utils.ResponseUnauthorized(ctx)
		}

		err := service.UnAuthenticateUser(user)
		if err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseOK(ctx, "Logged out successfully")
	}
}
