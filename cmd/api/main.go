package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"navicstein/private-gpt/internal/api"
	"navicstein/private-gpt/internal/config"
	"navicstein/private-gpt/internal/database"
	"os"
)

func main() {

	// 1. load config
	if err := config.Setup(); err != nil {
		panic(err)
	}

	// 3. init logger with a sentry writer
	multi := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stderr})
	zerologLogger := zerolog.New(multi).With().Caller().Timestamp().Logger()
	log.Logger = zerologLogger
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// 4. init fiber
	engine := html.New("./views", ".html")
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024,
	})

	app.Static("/", "./views")

	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	apis := api.NewAPI(app, db)
	apis.RegisterRoutes()

	err = app.Listen(":1337")
	if err != nil {
		panic(err)
	}
}
