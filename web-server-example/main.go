package main

import (
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		nBytes, err := w.Write([]byte("Hello World!"))
		if err != nil {
			log.Println("response writer error:", err.Error())
		} else {
			log.Println("Bytes written to response:", nBytes)
		}
	})

	wErr := http.ListenAndServe(":8080", nil)
	if wErr != nil {
		log.Fatalln("Couldn't start the web server on port 8080. Error was:", wErr.Error())
	} else {
		log.Println("Web server started on port:", portNumber)
	}

}
