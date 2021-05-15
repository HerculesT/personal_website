package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/spf13/viper"
	"gopkg.in/mail.v2"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("../templates/*.gohtml"))
}

func main() {

	// port := os.Getenv("PORT")

	// if port == "" {
	// 	log.Fatal("$PORT must be set")
	// }

	http.HandleFunc("/", frontPage)
	http.HandleFunc("/skillSet", skillSet)
	http.HandleFunc("/workExperience", workExperience)
	http.HandleFunc("/certificates", certificates)
	http.HandleFunc("/contactMe", contactMe)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("../assets")))) //initialize and load css.
	http.Handle("/favicon.ico", http.NotFoundHandler())                                          //keep browser from complaining about favicon missing
	http.ListenAndServe(":8080", nil)

	// http.ListenAndServe(":"+port, nil)

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

func viperEnvVariable(key string) string {

	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory
	viper.SetConfigFile("../config.env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(key).(string)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

func contactMe(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		err := tpl.ExecuteTemplate(w, "contactMe.gohtml", nil) //execute the template with bool false shows  the form.
		if err != nil {
			http.Error(w, err.Error(), 500)
			log.Fatalln(err)
		}
	} else {
		req.ParseForm()
		tpl.ExecuteTemplate(w, "contactMe.gohtml", struct{ Success bool }{true}) //execute the template with bool true shows thank you msg.
		user := viperEnvVariable("GMAIL_USER")
		pass := viperEnvVariable("GMAIL_PASS")
		d := mail.NewDialer("smtp.gmail.com", 587, user, pass)
		m := mail.NewMessage()

		nm := req.FormValue("name")
		em := req.FormValue("email")
		su := req.FormValue("subject")
		msg := req.FormValue("message")

		m.SetHeader("From", em)
		m.SetHeader("To", "sko.hercules@hotmail.com")
		m.SetHeader("Subject", su)
		m.SetBody("text/plain", "Mail sent From: "+em+"\n\n"+"Name: "+nm+"\n\n"+msg)
		if err := d.DialAndSend(m); err != nil {
			fmt.Println("Failed sending mail")
			panic(err)
		}
		fmt.Println("Mail sent without error")
	}
}
