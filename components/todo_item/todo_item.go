package components_todo_item

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
)

type TodoItem struct {
	Content    string
	RenderItem func(item TodoItem) string
}

func TodoItemView(item TodoItem) string {
	wd, wdErr := os.Getwd()
	if wdErr != nil {
		log.Fatalln(wdErr)
	}
	var res bytes.Buffer
	tmpl := template.Must(template.ParseFiles(wd + "/components/todo_item/todo_item.html"))
	err := tmpl.Execute(&res, item)
	if err != nil {
		return fmt.Sprint("Todo item view err: ", err)
	}
	return res.String()
}
