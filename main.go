package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	components_todo_item "github.com/nguyenduclam1711/go-templ/components/todo_item"
	components_todo_list "github.com/nguyenduclam1711/go-templ/components/todo_list"
)

func main() {
	app := fiber.New()

	// compress config
	app.Use(compress.New())

	// static files
	app.Static("/", "./public")

	// try out templ
	app.Get("/", func(c fiber.Ctx) error {
		list := components_todo_list.TodoList{}
		for i := 0; i < 5; i++ {
			list = append(list, components_todo_item.TodoItem{
				Content: fmt.Sprint("Content", i+1),
				RenderItem: func(item components_todo_item.TodoItem) string {
					return components_todo_item.TodoItemView(item)
				},
			})
		}
		res := components_todo_list.TodoListView(list)

		c.Set("Content-Type", "text/html")
		return c.SendString(res)
	})

	log.Fatal(app.Listen(":3000"))
}
