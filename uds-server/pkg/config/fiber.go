package config

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/spf13/viper"
)

func NewFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		ServerHeader:      viper.GetString(AppName),
		AppName:           viper.GetString(AppName),
		Prefork:           viper.GetBool(AppPrefork),
		IdleTimeout:       5 * time.Second,
		BodyLimit:         100 * 1024 * 1024,
		StreamRequestBody: true,
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Use(filesystem.New(filesystem.Config{
		Root: http.Dir("/home/fauzi/www/github/uds/upload"),
	}))

	app.Static("/videos", "/home/fauzi/www/github/uds/upload/videos")
	app.Static("/images", "/home/fauzi/www/github/uds/upload")

	return app
}
