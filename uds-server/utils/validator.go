package utils

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ent "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
	"log"
	"reflect"
)

var (
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
)

const (
	ParseBody  = "body"
	ParseQuery = "query"
	ParseParam = "param"
)

func init() {
	validate = validator.New()

	uni = ut.New(en.New())
	trans, _ = uni.GetTranslator("en")

	if err := ent.RegisterDefaultTranslations(validate, trans); err != nil && !fiber.IsChild() {
		log.Println(err)
	}
}

func ParseAndValidate(ctx *fiber.Ctx, body any, source string) error {
	v := reflect.ValueOf(body)
	if v.Kind() != reflect.Pointer {
		return ResponseBadRequest(ctx, "Body must be pointer")
	}

	switch source {
	case ParseBody:
		if err := parseBody(ctx, body); err != nil {
			return err
		}
		return validateStruct(body)
	case ParseQuery:
		if err := parseQuery(ctx, body); err != nil {
			return err
		}
		return validateStruct(body)
	case ParseParam:
		if err := parseParam(ctx, body); err != nil {
			return err
		}
		return validateStruct(body)
	default:
		return ResponseBadRequest(ctx, "Invalid parse source")
	}
}

func parseBody(ctx *fiber.Ctx, body any) error {
	if err := ctx.BodyParser(body); err != nil {
		return err
	}
	return nil
}

func parseQuery(ctx *fiber.Ctx, body any) error {
	if err := ctx.QueryParser(body); err != nil {
		return err
	}
	return nil
}

func parseParam(ctx *fiber.Ctx, body any) error {
	if err := ctx.ParamsParser(body); err != nil {
		return err
	}
	return nil
}

func validateStruct(input any) error {
	return validate.Struct(input)
}
