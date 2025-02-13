package app

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/maximfedotov74/social-app-backend/internal/config"
)

func Start() {

	appConfig := config.MustLoadConfig()

	fiberApp := fiber.New()

	router := fiberApp.Group("")

	router.Get("/test", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"hello": "world"})
	})

	err := fiberApp.Listen(fmt.Sprintf(":%d", appConfig.HTTP.Port))

	if err != nil {
		log.Fatal(err)
	}

}
