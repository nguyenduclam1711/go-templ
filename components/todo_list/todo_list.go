package components_todo_list

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"

	components_todo_item "github.com/nguyenduclam1711/go-templ/components/todo_item"
)

type TodoList = []components_todo_item.TodoItem

type TodoListProps struct {
	List TodoList
}

func TodoListView(list TodoList) string {
	props := TodoListProps{
		List: list,
	}
	var res bytes.Buffer
	wd, wdErr := os.Getwd()

	if wdErr != nil {
		log.Fatalln(wdErr)
	}
	newTmpl := template.New("todo_list").Funcs(template.FuncMap{
		"RenderItem": func(item components_todo_item.TodoItem) string {
			return components_todo_item.TodoItemView(item)
		},
	})
	tmpl := template.Must(newTmpl.ParseFiles(wd + "/components/todo_list/todo_list.html"))
	err := tmpl.Execute(&res, props)
	if err != nil {
		return fmt.Sprint("Todo list view err: ", err)
	}
	return res.String()
}
