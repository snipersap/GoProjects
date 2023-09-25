package main

import (
	"log"
	"net/http"

	handler "github.com/snipersap/GoProjects/tree/main/bed-and-breakfast/pkg/handlers"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	log.Println("Starting web server on port:", portNumber)
	wErr := http.ListenAndServe(":8080", nil)
	if wErr != nil {
		log.Fatalln("couldn't start the web server on port 8080. Error was:", wErr.Error())
	}

}
