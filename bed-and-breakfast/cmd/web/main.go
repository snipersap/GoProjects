package main

import (
	"log"
	"net/http"

	"github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/config"
	handler "github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/handlers"
	"github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/render"
)

const portNumber = ":8080"

func main() {

	// declare app config
	var app config.AppConfig

	//create the Template Cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("cannot load template cache. Exiting app:", err.Error())
	}
	app.TemplateCache = tc
	render.SetAppConfig(&app)

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	log.Println("Starting web server on port:", portNumber)
	wErr := http.ListenAndServe(":8080", nil)
	if wErr != nil {
		log.Fatalln("couldn't start the web server on port 8080. Error was:", wErr.Error())
	}

}
