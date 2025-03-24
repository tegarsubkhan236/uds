package handler

import (
	"fmt"
	"myapp/dto"
	"myapp/pkg/service"
	"myapp/utils"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func HandleFetchAllMovie(service service.MovieService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(utils.RequestPaginate)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseQuery); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		currentPage, lastPage, totalData, result, err := service.GetMovies(req.Page, req.Limit)
		if err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseOKWithPages(ctx, currentPage, lastPage, totalData, result)
	}
}

func HandleFetchDetailMovie(service service.MovieService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(utils.RequestID)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseParam); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		result, err := service.GetMovieById(req.ID)
		if err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseOK(ctx, result)
	}
}

func HandleCreateMovie(service service.MovieService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(dto.MovieRequest)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseBody); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		videoFile, err := ctx.FormFile("video_file")
		if err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		posterFile, err := ctx.FormFile("poster_file")
		if err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		//username, err := utils.Me(ctx)
		//if err != nil {
		//	return utils.ResponseUnauthorized(ctx)
		//}

		if err := service.CreateMovie(req.ToEntity(), videoFile, posterFile, "TEST CREATE USER"); err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseCreated(ctx)
	}
}

func HandleUpdateMovie(service service.MovieService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		reqID := new(utils.RequestID)
		if err := utils.ParseAndValidate(ctx, reqID, utils.ParseParam); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		req := new(dto.MovieRequest)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseBody); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		videoFile, err := ctx.FormFile("video_file")
		if err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		posterFile, err := ctx.FormFile("poster_file")
		if err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		//username, err := utils.Me(ctx)
		//if err != nil {
		//	return utils.ResponseUnauthorized(ctx)
		//}

		if err := service.UpdateMovie(reqID.ID, req.ToEntity(), videoFile, posterFile, "TEST UPDATE USER"); err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseUpdated(ctx)
	}
}

func HandleDeleteMovie(service service.MovieService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(utils.RequestID)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseParam); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		//username, err := utils.Me(ctx)
		//if err != nil {
		//	return utils.ResponseUnauthorized(ctx)
		//}

		if err := service.DeleteMovie(req.ID, "TEST DELETE USER"); err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		return utils.ResponseDeleted(ctx)
	}
}

func HandleStreamMovie(service service.MovieService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(utils.RequestID)
		if err := utils.ParseAndValidate(ctx, req, utils.ParseParam); err != nil {
			return utils.ResponseBadRequest(ctx, err.Error())
		}

		result, err := service.GetMovieById(req.ID)
		if err != nil {
			return utils.ResponseInternalServerError(ctx, err.Error())
		}

		fmt.Println("Path to video:", result.VideoUrl)

		// Handle file .ts
		if strings.HasSuffix(result.VideoUrl, ".ts") {
			fmt.Println("Streaming:", result.VideoUrl)

			// Ensure the .ts file exists before sending it
			if _, err := os.Stat(result.VideoUrl); os.IsNotExist(err) {
				return utils.ResponseNotFound(ctx, "TS file not found")
			}

			return ctx.SendFile(result.VideoUrl)
		}

		// Handle file .m3u8
		if strings.HasSuffix(result.VideoUrl, ".m3u8") {
			fmt.Println("Serving playlist:", result.VideoUrl)

			// Ensure the .m3u8 file exists before sending it
			if _, err := os.Stat(result.VideoUrl); os.IsNotExist(err) {
				return utils.ResponseNotFound(ctx, "M3U8 file not found")
			}

			return ctx.SendFile(result.VideoUrl)
		}

		// Default handling (non-TS, non-M3U8)
		// Ensure the default file exists before sending it
		if _, err := os.Stat(result.VideoUrl); os.IsNotExist(err) {
			return utils.ResponseNotFound(ctx, "File not found")
		}

		return ctx.Status(fiber.StatusOK).SendFile(result.VideoUrl, true)
	}
}
