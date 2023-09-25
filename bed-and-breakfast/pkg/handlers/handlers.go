package handler

import (
	"log"
	"net/http"

	"github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/render"
)

const (
	homeTpl  = "home.page.tmpl"
	aboutTpl = "about.page.tmpl"
)

// Home handles the home page URL or /
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, homeTpl) //Using .html notation to use Emmet abbreviations in VS Code
	log.Println("Rendered Home")
}

// About handles the about page with URL /about
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, aboutTpl) //Using .html notation to use Emmet abbreviations in VS Code
	log.Println("Rendered About")
}
