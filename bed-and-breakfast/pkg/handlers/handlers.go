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
	if !isURLPathHome(r) {
		http.NotFound(w, r)
		return
	}
	err := render.RenderTemplate(w, homeTpl) //Using .html notation to use Emmet abbreviations in VS Code
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		log.Println("Rendered Home")
	}
}

// isURLPathHome checks for the path '/' in the URL and returns true if path is '/'
// This is used to avoid duplicate requests because of favicon
func isURLPathHome(r *http.Request) bool {
	return r.URL.Path == "/"
}

// About handles the about page with URL /about
func About(w http.ResponseWriter, r *http.Request) {
	err := render.RenderTemplate(w, aboutTpl) //Using .html notation to use Emmet abbreviations in VS Code
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		log.Println("Rendered About")
	}
}
