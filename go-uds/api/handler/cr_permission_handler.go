package handler

import (
	"github.com/gofiber/fiber/v2"
	"myapp/dto"
	"myapp/pkg/service"
	"myapp/utils"
)

func PermissionIndex(service service.PermissionService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(utils.RequestPaginate)
		if err := utils.ParseAndValidate(ctx, req); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		currentPage, lastPage, totalData, result, err := service.GetPermissions(req.Page, req.Limit)
		if err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		responses := make([]dto.PermissionResponse, len(result))
		for i, p := range result {
			responses[i] = p.ToResponse()
		}

		return utils.ResponseOKWithPages(ctx, currentPage, lastPage, totalData, responses)
	}
}

func PermissionShow(service service.PermissionService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(utils.RequestID)
		if err := utils.ParseAndValidate(ctx, req); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		result, err := service.GetPermissionById(req.ID)
		if err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseOK(ctx, result.ToResponse())
	}
}

func PermissionInsert(service service.PermissionService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(dto.PermissionInsertRequest)
		if err := utils.ParseAndValidate(ctx, req); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		username, err := utils.Me(ctx)
		if err != nil {
			return utils.ResponseUnauthorized(ctx)
		}

		if err := service.CreatePermission(req.ToEntity(), username); err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseCreated(ctx)
	}
}

func PermissionUpdate(service service.PermissionService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(dto.PermissionUpdateRequest)
		if err := utils.ParseAndValidate(ctx, req); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		username, err := utils.Me(ctx)
		if err != nil {
			return utils.ResponseUnauthorized(ctx)
		}

		if err := service.UpdatePermission(req.ToEntity(), username); err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseUpdated(ctx)
	}
}

func PermissionDelete(service service.PermissionService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(utils.RequestID)
		if err := utils.ParseAndValidate(ctx, req); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		username, err := utils.Me(ctx)
		if err != nil {
			return utils.ResponseUnauthorized(ctx)
		}

		if err := service.DeletePermission(req.ID, username); err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseDeleted(ctx)
	}
}
