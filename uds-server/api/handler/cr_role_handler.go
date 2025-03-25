package handler

import (
	"myapp/dto"
	"myapp/pkg/service"
	"myapp/utils"

	"github.com/gofiber/fiber/v2"
)

func RoleIndex(service service.RoleService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(utils.RequestPaginate)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseQuery); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		currentPage, lastPage, totalData, result, err := service.GetRoles(req.Page, req.Limit)
		if err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		responses := make([]dto.RoleResponse, len(result))
		for i, p := range result {
			responses[i] = p.ToResponse()
		}

		return utils.ResponseOKWithPages(ctx, currentPage, lastPage, totalData, responses)
	}
}

func RoleShow(service service.RoleService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(utils.RequestID)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseParam); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		role, err := service.GetRoleById(req.ID)
		if err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseOK(ctx, role.ToResponse())
	}
}

func RoleInsert(service service.RoleService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(dto.RoleInsertRequest)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseBody); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		username, err := utils.Me(ctx)
		if err != nil {
			return utils.ResponseUnauthorized(ctx)
		}

		err = service.CreateRole(req.ToEntity(), username)
		if err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseCreated(ctx, 0)
	}
}

func RoleUpdate(service service.RoleService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(dto.RoleUpdateRequest)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseBody); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		username, err := utils.Me(ctx)
		if err != nil {
			return utils.ResponseUnauthorized(ctx)
		}

		err = service.UpdateRole(req.ToEntity(), username)
		if err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseUpdated(ctx)
	}
}

func RoleDelete(service service.RoleService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(utils.RequestID)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseParam); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		username, err := utils.Me(ctx)
		if err != nil {
			return utils.ResponseUnauthorized(ctx)
		}

		if err := service.DeleteRole(req.ID, username); err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseDeleted(ctx)
	}
}
