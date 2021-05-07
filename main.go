package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", frontPage)
	http.HandleFunc("/skillSet", skillSet)
	http.HandleFunc("/workExperience", workExperience)
	http.HandleFunc("/contactMe", contactMe)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("public")))) //initialize and load css.
	http.Handle("/favicon.ico", http.NotFoundHandler())                                       //keep browser from complaining about favicon missing
	http.ListenAndServe(":8080", nil)
}

func frontPage(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "frontPage.gohtml", req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

// var nv = template.FuncMap{
// 	"nv": navBar,
// }

func navBar(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "navBar.gohtml", req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

func skillSet(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "skillSet.gohtml", req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

func workExperience(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "workExperience.gohtml", req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

func contactMe(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "contactMe.gohtml", req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
