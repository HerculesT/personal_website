package main

import (
	"log"
	"net/http"
	"os"
	"text/template"

	_ "github.com/heroku/x/hmetrics/onload"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	http.HandleFunc("/", frontPage)
	http.HandleFunc("/skillSet", skillSet)
	http.HandleFunc("/workExperience", workExperience)
	http.HandleFunc("/certificates", certificates)
	http.HandleFunc("/contactMe", contactMe)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("public")))) //initialize and load css.
	http.Handle("/favicon.ico", http.NotFoundHandler())                                       //keep browser from complaining about favicon missing
	// http.ListenAndServe(":8080", nil)

	router.Run(":" + port)

}

func frontPage(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "frontPage.gohtml", req)
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

func certificates(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "certificates.gohtml", req)
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
