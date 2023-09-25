package main

import (
	"html/template"
	"log"
	"net/http"
)

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
			log.Println("Rendered template:", tmplName)
		}
	}
}
