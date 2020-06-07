package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type test int

func (t test) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		Submissions url.Values
	}{
		request.Method,
		request.Form,
	}
	tpl.ExecuteTemplate(writer, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var t test
	http.ListenAndServe("localhost:8080", t)
}
