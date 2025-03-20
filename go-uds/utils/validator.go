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

func init() {
	validate = validator.New()

	uni = ut.New(en.New())
	trans, _ = uni.GetTranslator("en")

	if err := ent.RegisterDefaultTranslations(validate, trans); err != nil && !fiber.IsChild() {
		log.Println(err)
	}
}

func ParseAndValidate(c *fiber.Ctx, body any) error {
	v := reflect.ValueOf(body)

	switch v.Kind() {
	case reflect.Ptr:
		switch c.Method() {
		case fiber.MethodPost, fiber.MethodPut:
			err := parseBody(c, body)
			if err != nil {
				return err
			}
		case fiber.MethodGet, fiber.MethodDelete:
			err := parseQuery(c, body)
			if err != nil {
				return err
			}
		}
		return validateStruct(v.Elem().Interface())
	case reflect.Struct:
		switch c.Method() {
		case fiber.MethodPost, fiber.MethodPut:
			err := parseBody(c, &body)
			if err != nil {
				return err
			}
		case fiber.MethodGet, fiber.MethodDelete:
			err := parseQuery(c, &body)
			if err != nil {
				return err
			}
		}
		return validateStruct(v)
	default:
		return nil
	}
}

func parseBody(c *fiber.Ctx, body any) error {
	if err := c.BodyParser(body); err != nil {
		return err
	}
	return nil
}

func parseQuery(c *fiber.Ctx, body any) error {
	if err := c.QueryParser(body); err != nil {
		return err
	}
	return nil
}

func validateStruct(input any) error {
	return validate.Struct(input)
}
