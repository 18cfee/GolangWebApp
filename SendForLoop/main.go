package main

import (
	"html/template"
	"net/http"
	"os"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
				{Title: "Task 4", Done: false},
			},
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(3)
	})

	http.ListenAndServe(":8080", nil)
}
