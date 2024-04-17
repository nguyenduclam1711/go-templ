package views

import (
	"context"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
)

func RenderHtml(c fiber.Ctx, component templ.Component) error {
	c.Set("Content-type", "text/html")
	return component.Render(context.Background(), c.Response().BodyWriter())
}
