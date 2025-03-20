package utils

import (
	"github.com/gofiber/fiber/v2"
)

type meta struct {
	CurrentPage int `json:"current_page,omitempty"`
	LastPage    int `json:"last_page,omitempty"`
	TotalData   int `json:"total_data,omitempty"`
}

type errorResponse struct {
	Code     int    `json:"code"`
	Status   string `json:"status"`
	Messages string `json:"message"`
	Data     any    `json:"data,omitempty"`
}

type successResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Meta    meta   `json:"meta,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// ResponseOKWithPages responds with paginated data
func ResponseOKWithPages(ctx *fiber.Ctx, currentPage, lastPage, totalData int, data any) error {
	return ctx.Status(fiber.StatusOK).JSON(&successResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Meta: meta{
			CurrentPage: currentPage,
			LastPage:    lastPage,
			TotalData:   totalData,
		},
		Data: data,
	})
}

// ResponseOK responds with a success message and data
func ResponseOK(ctx *fiber.Ctx, data any) error {
	return ctx.Status(fiber.StatusOK).JSON(&successResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   data,
	})
}

// ResponseCreated responds with success for created data
func ResponseCreated(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusCreated).JSON(&successResponse{
		Code:    fiber.StatusCreated,
		Status:  "Created",
		Message: "Data created successfully",
	})
}

// ResponseUpdated responds with success for updated data
func ResponseUpdated(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&successResponse{
		Code:    fiber.StatusOK,
		Status:  "Updated",
		Message: "Data updated successfully",
	})
}

// ResponseDeleted responds with success for deleted data
func ResponseDeleted(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&successResponse{
		Code:    fiber.StatusOK,
		Status:  "Deleted",
		Message: "Data deleted successfully",
	})
}

// ResponseBadRequest responds with an error for bad requests
func ResponseBadRequest(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(&errorResponse{
		Code:     fiber.StatusBadRequest,
		Status:   "Bad Request",
		Messages: message,
	})
}

// ResponseUnauthorized responds with an error for unauthorized requests
func ResponseUnauthorized(ctx *fiber.Ctx, messages ...string) error {
	message := "You are not authorized to access this resource"
	if len(messages) > 0 {
		message = messages[0]
	}

	return ctx.Status(fiber.StatusUnauthorized).JSON(&errorResponse{
		Code:     fiber.StatusUnauthorized,
		Status:   "Unauthorized",
		Messages: message,
	})
}

// ResponseForbidden responds with an error for forbidden requests
func ResponseForbidden(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusForbidden).JSON(&errorResponse{
		Code:     fiber.StatusForbidden,
		Status:   "Forbidden",
		Messages: message,
	})
}

// ResponseNotFound responds with an error when the requested resource is not found
func ResponseNotFound(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusNotFound).JSON(&errorResponse{
		Code:     fiber.StatusNotFound,
		Status:   "Not Found",
		Messages: message,
	})
}

// ResponseInternalServerError responds with an error for internal server errors
func ResponseInternalServerError(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(&errorResponse{
		Code:     fiber.StatusInternalServerError,
		Status:   "Internal Server Error",
		Messages: message,
	})
}
