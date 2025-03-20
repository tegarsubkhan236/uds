package main

import (
	"fmt"
	"log"
	"myapp/api/router"
	"myapp/pkg"
	"myapp/pkg/config"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func main() {
	if err := config.NewViper(); err != nil {
		log.Fatal(err)
	}

	db, err := config.NewMysqlDB()
	if err != nil {
		log.Fatal(err)
	}

	deps := pkg.NewDependencies(db)

	app := config.NewFiber()
	api := app.Group("/api", logger.New())

	router.SetupRouter(api, deps)

	handleShutdown(app)

	address := fmt.Sprintf("%s:%d", viper.GetString(config.AppHost), viper.GetInt(config.AppPort))
	if err := app.Listen(address); err != nil {
		log.Panic(err)
	}
}

func handleShutdown(app *fiber.App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()
}
