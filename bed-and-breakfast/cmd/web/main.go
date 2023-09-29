package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/config"
	handler "github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/handlers"
	"github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/render"
)

const portNumber string = ":8080"

// declare app config
var app config.AppConfig

func main() {

	s := CreateSession()
	SetSession(&app, s)

	//create the Template Cache
	tc, err := render.GetTemplateCache()
	if err != nil {
		log.Fatalln("cannot load template cache. Exiting app:", err.Error())
	}

	// Set the template cache to app config and render package
	SetTmplCache(&app, tc)
	SetUseCache(&app, false)
	render.SetAppConfig(&app)
	// Init a handler repository with app config
	repo := handler.NewRepo(&app)
	handler.SetRepo(repo)

	// Setup server and run with pat
	srv := http.Server{
		Addr:    portNumber,
		Handler: chiRoutes(&app),
	}
	log.Println("Starting web server on port:", portNumber)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln("couldn't start the web server on port "+portNumber+
			" Error was:", err.Error())
	}

	// //Set handlers for the apis
	// http.HandleFunc("/", handler.Repo.Home)
	// http.HandleFunc("/about", handler.Repo.About)

	// //Start the web server
	// log.Println("Starting web server on port:", portNumber)
	// wErr := http.ListenAndServe(portNumber, nil)
	// if wErr != nil {
	// 	log.Fatalln("couldn't start the web server on port 8080. Error was:", wErr.Error())
	// }

}

func CreateSession() *scs.SessionManager {
	s := scs.New()
	s.Lifetime = 24 * time.Hour
	s.Cookie.Persist = true
	s.Cookie.SameSite = http.SameSiteLaxMode
	s.Cookie.Secure = app.InProduction
	return s
}

func SetSession(app *config.AppConfig, s *scs.SessionManager) {
	app.Session = s
}

func SetUseCache(app *config.AppConfig, uc bool) {
	app.UseCache = uc
}

func SetTmplCache(app *config.AppConfig, t map[string]*template.Template) {
	app.TemplateCache = t
}
