package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"myapp/dto"
	"myapp/pkg/service"
	"myapp/utils"
	"time"
)

func GetAllCategoryHandler(service service.CategoryService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(utils.RequestPaginate)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseQuery); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		currentPage, lastPage, totalData, result, err := service.GetCategory(req.Page, req.Limit)

		log.Default().Println(result, "di handler")

		if err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		responses := make([]dto.CategoryResponse, len(result))
		for i, category := range result {
			responses[i] = dto.CategoryResponse{
				ID:        category.ID,
				Name:      category.Name,
				CreatedBy: category.CreatedBy,
				CreatedAt: category.CreatedAt.Format(time.RFC3339),
			}
		}

		return utils.ResponseOKWithPages(ctx, currentPage, lastPage, totalData, responses)
	}
}
func CategoryInsert(service service.CategoryService) fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		req := new(dto.CategoryInsertRequest)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseBody); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		// Debug: Cek apakah name terbaca
		fmt.Println("Nama category:", req.Name)

		playlist := req.ToEntity()
		fmt.Println("Entity category:", playlist.Name)

		// Simpan playlist tanpa autentikasi
		if err := service.CreateCategory(playlist, "system"); err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseCreated(ctx)
	}
}

func CategoryUpdate(service service.CategoryService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(dto.CategoryInsertRequest)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseBody); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		username, err := utils.Me(ctx)
		if err != nil {
			return utils.ResponseUnauthorized(ctx)
		}

		if err := service.UpdateCategory(req.ToEntity(), username); err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseUpdated(ctx)
	}
}
