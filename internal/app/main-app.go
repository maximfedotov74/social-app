package app

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/maximfedotov74/social-app/internal/config"
)

func Start() {

	appConfig := config.MustLoadConfig()

	fiberApp := fiber.New()

	router := fiberApp.Group("")

	//TODO: Использовать иначе, вынести только для html роутинга
	router.Use(func(c *fiber.Ctx) error {
		encoding := c.Get("Accept-Encoding")
		if strings.Contains(encoding, "gzip") {
			c.Request().Header.Set("Accept-Encoding", "gzip")
		}
		return c.Next()
	})

	//TODO: Использовать иначе, вынести только для html роутинга
	router.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	dir, err := os.Getwd()

	static := path.Join(dir, "web", "public")

	router.Static("/public", static)

	router.Get("/test", func(ctx *fiber.Ctx) error {

		return ctx.JSON(fiber.Map{"hello": "world"})
	})

	err = fiberApp.Listen(fmt.Sprintf(":%d", appConfig.HTTP.Port))

	if err != nil {
		log.Fatal(err)
	}
}
