package main

import (
	"fmt"
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
		log.Fatalln("Couldn't start the web server on port 8080. Error was:", wErr.Error())
	}

}

// Home handles the home page URL or /
func Home(w http.ResponseWriter, r *http.Request) {
	nBytes, err := w.Write([]byte("<html>Hello World!, You are looking at the home page. <a href='/about'>About >></a></html>"))
	if err != nil {
		log.Println("response writer error:", err.Error())
		return
	} else {
		log.Println("Bytes written to response:", nBytes)
	}
}

// About handles the about page with URL /about
func About(w http.ResponseWriter, r *http.Request) {
	res := sum(2, 4)
	fmt.Fprintf(w, "<html>Now we are in the About page.<a href='/'>Home >></a>")
	fmt.Fprintf(w, "\nResult of addition of 2 numbers is:%d</html>", res)
}

// sum adds two integers and returns the results
func sum(x, y int) int {
	return x + y
}
