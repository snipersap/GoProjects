package main

import (
	"html/template"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println("Starting web server on port:", portNumber)
	wErr := http.ListenAndServe(":8080", nil)
	if wErr != nil {
		log.Fatalln("couldn't start the web server on port 8080. Error was:", wErr.Error())
	}

}

// Home handles the home page URL or /
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html") //Using .html notation to use Emmet abbreviations in VS Code
	log.Println("render from Home called")
}

// About handles the about page with URL /about
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html") //Using .html notation to use Emmet abbreviations in VS Code
	log.Println("render from About called")
}

func renderTemplate(w http.ResponseWriter, tmplName string) {
	t, err := template.ParseFiles("./templates/" + tmplName)
	if err != nil {
		log.Println("could not read HTML template ", tmplName, ". Received error:", err.Error())
		return
	} else {
		pErr := t.Execute(w, nil)
		if pErr != nil {
			log.Println("could not parse HTML template ", tmplName, ". Received error:", err.Error())
			return
		} else {
			log.Println("rT: Rendered template:", tmplName)
		}
	}
}
