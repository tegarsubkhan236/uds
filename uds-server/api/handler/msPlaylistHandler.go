package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"myapp/dto"
	"myapp/pkg/service"
	"myapp/utils"
)

func GetAllPlaylistHandler(service service.PlaylistService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(utils.RequestPaginate)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseQuery); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		currentPage, lastPage, totalData, result, err := service.GetPlaylist(req.Page, req.Limit)

		log.Default().Println(result, "di handler")

		if err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		responses := make([]dto.CategoryResponse, len(result))
		for i, category := range result {
			responses[i] = dto.CategoryResponse{
				ID:        category.ID,
				Name:      category.Name,
				CreatedBy: "",
			}
		}

		return utils.ResponseOKWithPages(ctx, currentPage, lastPage, totalData, responses)
	}
}

func PlaylistInsert(service service.PlaylistService) fiber.Handler {
	//return func(ctx *fiber.Ctx) error {
	//	req := new(dto.PlaylistInsertRequest)
	//	if err := utils.ParseAndValidate(ctx, req, utils.ParseBody); err != nil {
	//		return utils.ResponseBadRequest(ctx, err.Error())
	//	}
	//
	//	// Tanpa autentikasi, langsung simpan playlist
	//	if err := service.CreatePlaylist(req.ToEntity(), "system"); err != nil {
	//		return utils.ResponseInternalServerError(ctx, err.Error())
	//	}
	//
	//	return utils.ResponseCreated(ctx)
	//}

	return func(ctx *fiber.Ctx) error {
		req := new(dto.PlaylistInsertRequest)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseBody); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		// Debug: Cek apakah name terbaca
		fmt.Println("Nama Playlist:", req.Name)

		playlist := req.ToEntity()
		fmt.Println("Entity Playlist:", playlist.Name)

		// Simpan playlist tanpa autentikasi
		if err := service.CreatePlaylist(playlist, "system"); err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseCreated(ctx)
	}
}
