package main

import (
	"log"
	"net/http"
)

// Home handles the home page URL or /
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html") //Using .html notation to use Emmet abbreviations in VS Code
	log.Println("Rendered Home")
}

// About handles the about page with URL /about
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html") //Using .html notation to use Emmet abbreviations in VS Code
	log.Println("Rendered About")
}
