package handler

import (
	"log"
	"net/http"

	"github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/config"
	"github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/models"
	"github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/render"
)

const (
	homeTpl  = "home.page.tmpl"
	aboutTpl = "about.page.tmpl"
)

// Repo stores the pointer to a Repository
var Repo *Repository

// Repository defines the structure to store the app config
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository with the app config and returns a pointer to the new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{a}
}

// SetRepo initializes the reference to a Repository
func SetRepo(r *Repository) {
	Repo = r
}

// Home handles the home page URL or /
func (rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	if !isURLPathHome(r) {
		http.NotFound(w, r)
		return
	}
	err := render.RenderTemplate(w, homeTpl, &models.TemplateData{}) //Using .html notation to use Emmet abbreviations in VS Code
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
func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {

	td := dynContentAbout()
	err := render.RenderTemplate(w, aboutTpl, td) //Using .html notation to use Emmet abbreviations in VS Code
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		log.Println("Rendered About")
	}
}

func dynContentAbout() *models.TemplateData {
	sMap := make(map[string]string)
	sMap["About me"] = "Software Engineering Manager from Germany"
	td := models.TemplateData{StringMap: sMap}
	return &td
}
