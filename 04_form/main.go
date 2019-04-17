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

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
	EnumO      string
	courses    []string
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

	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribe") == "on"
	e := req.FormValue("enum")

	err := tpl.ExecuteTemplate(w, "index.gohtml", person{f, l, s, e})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
	fmt.Println(f, l, s, e)
}
