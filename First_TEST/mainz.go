package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", header)
	tpl = template.Must(template.ParseGlob("styles/*"))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func header(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "header.gohtml", req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
