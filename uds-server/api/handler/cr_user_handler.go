package handler

import (
	"github.com/gofiber/fiber/v2"
	"myapp/dto"
	"myapp/pkg/service"
	"myapp/utils"
)

func UserIndex(service service.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(utils.RequestPaginate)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseQuery); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		currentPage, lastPage, totalData, result, err := service.GetUsers(req.Page, req.Limit)
		if err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		responses := make([]dto.UserResponse, len(result))
		for i, p := range result {
			responses[i] = p.ToResponse()
		}

		return utils.ResponseOKWithPages(ctx, currentPage, lastPage, totalData, responses)
	}
}

func UserShow(service service.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(utils.RequestID)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseParam); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		user, err := service.GetUserById(req.ID)
		if err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseOK(ctx, user.ToResponse())
	}
}

func UserInsert(service service.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(dto.UserInsertRequest)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseBody); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		username, err := utils.Me(ctx)
		if err != nil {
			return utils.ResponseUnauthorized(ctx)
		}

		err = service.CreateUser(req.ToEntity(), username)
		if err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseCreated(ctx)
	}
}

func UserUpdate(service service.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(dto.UserUpdateRequest)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseBody); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		username, err := utils.Me(ctx)
		if err != nil {
			return utils.ResponseUnauthorized(ctx)
		}

		err = service.UpdateUser(req.ToEntity(), username)
		if err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseUpdated(ctx)
	}
}

func UserDelete(service service.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(utils.RequestID)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseParam); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		username, err := utils.Me(ctx)
		if err != nil {
			return utils.ResponseUnauthorized(ctx)
		}

		if err := service.DeleteUser(req.ID, username); err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseDeleted(ctx)
	}
}
