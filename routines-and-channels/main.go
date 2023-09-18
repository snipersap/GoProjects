package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://honeywell.com",
		"http://arunisgood.com",
	}
	// Create channel
	c := make(chan string)

	// Call routine and pass each link to it
	for _, link := range links {
		go checkLink(link, c)
	}

	// Wait for the completion of the routines of each of the links
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	// Wait using separate Println statements
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		println("Maybe", link, "is down!")
		c <- "vvv DOWN vvv"
	} else {
		println("Yay!!", link, "is up.")
		c <- "^^^ UP ^^^"
	}
}
