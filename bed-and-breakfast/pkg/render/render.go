package render

import (
	"html/template"
	"log"
	"net/http"
)

const tmplPath = "./templates/"
const baseTmpl = tmplPath + "base.layout.tmpl"

// RenderTemplate takes the name of the template and renders it
func RenderTemplate(w http.ResponseWriter, tmplName string) {
	t, err := template.ParseFiles(tmplPath+tmplName, baseTmpl)
	if err != nil {
		log.Println("could not read HTML template ", tmplName, ". Received error:", err.Error())
		return
	} else {
		pErr := t.Execute(w, nil)
		if pErr != nil {
			log.Println("could not parse HTML template ", tmplName, ". Received error:", err.Error())
			return
		} else {
			log.Println("Rendered template:", tmplName)
		}
	}
}
