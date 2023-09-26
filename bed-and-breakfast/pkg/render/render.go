package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const tmplPath = "./templates/"

var tc = make(map[string]*template.Template, 3)

// RenderTemplate takes the name of the template and renders it
func RenderTemplate(w http.ResponseWriter, tmplName string) error {

	if _, tExists := tc[tmplName]; !tExists {
		//cache miss, create cache
		err := createTmplCache(tmplName)
		if err != nil {
			log.Println("couldn't create cache:", err.Error())
			return err
		}
	}
	t := tc[tmplName]
	err := t.Execute(w, nil)
	if err != nil {
		log.Println("could not render HTML template ", tmplName, ". Received error:", err.Error())
		return err
	} else {
		log.Println("Rendered template:", tmplName)
	}
	return nil
}

// createTmplCache creates the cache from disk when the template is not yet in cache
func createTmplCache(tn string) error {
	templates := []string{
		fmt.Sprintf("%s%s", tmplPath, tn),
		fmt.Sprintf("%s%s", tmplPath, "base.layout.tmpl"),
	}

	//parse and return err
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	//create cache
	tc[tn] = tmpl
	return nil
}
