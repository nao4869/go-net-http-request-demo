package main

import (
	"html/template"
	"log"
	"net/http"
)

type test int

func (m test) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(writer, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var t test
	http.ListenAndServe("localhost:8080", t) // 基本的に第二引数はnil - ハンドラの引数にnilが渡された場合、デフォルトでDefaultServeMuxというServeMux型のハンドラが使用されます。
}