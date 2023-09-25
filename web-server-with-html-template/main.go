package main

import (
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println("Starting web server on port:", portNumber)
	wErr := http.ListenAndServe(":8080", nil)
	if wErr != nil {
		log.Fatalln("couldn't start the web server on port 8080. Error was:", wErr.Error())
	}

}
