package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/nguyenduclam1711/go-templ/views"
)

func main() {
	app := fiber.New()

	// compress config
	app.Use(compress.New())

	// static files
	app.Static("/", "./public")

	// try out templ
	app.Get("/", func(c fiber.Ctx) error {
		return views.RenderHtml(c, views.Hello("lam nguyen"))
	})

	log.Fatal(app.Listen(":3000"))
}
