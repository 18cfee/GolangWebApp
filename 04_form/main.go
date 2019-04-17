package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type Todo struct {
	Title string
	Id    string
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
	Enum       string
	Todos      []Todo
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/exit", exit)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8085", nil)
}

func exit(w http.ResponseWriter, req *http.Request) {
	os.Exit(3)
}

func foo(w http.ResponseWriter, req *http.Request) {
	test := "enum"

	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribe") == "on"
	e := req.FormValue(test)

	data := person{
		FirstName:  f,
		LastName:   l,
		Subscribed: s,
		Enum:       test,
		Todos: []Todo{
			{Title: "Task 1", Id: test},
			{Title: "Task 2", Id: test},
			{Title: "Task 3", Id: test},
			{Title: "Task 4", Id: test},
		},
	}

	err := tpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
	fmt.Println(req.Form)
	fmt.Println(f, l, s, e)
}
